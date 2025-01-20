package main

import (
	"athghno/internal/dataStore"
	"athghno/internal/serverProto"
	"athghno/internal/streams"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"github.com/lpernett/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hostname := os.Getenv("ATHGHNO_HOSTNAME")
	port := os.Getenv("ATHGHNO_PORT")
	registry := streams.InitRegistry(hostname)
	db := dataStore.ConnectDB()
	defer db.Close()
	if _, err := dataStore.GetObject(db, []byte(fmt.Sprintf("%s/actor", hostname))); err != nil {
		fmt.Println("No server actor found. Generating new one")
		actor, err := createServerActor(db, registry)
		if err != nil {
			panic(err)
		} else {
			fmt.Println(actor)
		}
	}
	storedObj, err := dataStore.GetObject(db, []byte(fmt.Sprintf("%s/create/1", hostname)))
	if err != nil {
		if err.Error() != "Key not found" {
			panic(err)
		} else {
			jsonData, err := createObject(registry, db)
			if err != nil {
				panic(err)
			}
			storedObj = jsonData
			err = dataStore.PutObject(db, []byte(fmt.Sprintf("%s/create/1", hostname)), storedObj)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println(string(storedObj))
	as2, err := registry.UnmarshalIntoAS2Type(storedObj)
	if err != nil {
		panic(err)
	}
	fmt.Println(as2)
	server := http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}
	server.Handler = handler{
		db,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type handler struct {
	Db *badger.DB
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Scheme == "" {
		r.URL.Scheme = "http"
	}
	switch r.URL.Path {
	case "/":
		Index(w, r)
	case "/.well-known/webfinger":
		serverProto.WebFinger(h.Db, w, r)
	default:
		fullUrl := r.URL.Scheme + "://" + r.Host + r.URL.Path
		fmt.Println(fullUrl)
		data, err := dataStore.GetObject(h.Db, []byte(fullUrl))
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/activity+json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html><html><head><title>Athghno</title></head><body><h1>Athghno</h1></body></html>`))
}

func createServerActor(db *badger.DB, registry *streams.TypeRegistry) (streams.ApplicationAS2, error) {
	var app streams.ApplicationAS2
	hostname := registry.Hostname
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	id := hostname + "/actor"
	if err != nil {
		return app, err
	}
	pubKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})
	pubKeyObj := streams.PublicKey{
		Id:           id + "#main-key",
		Type:         "Key",
		Owner:        id,
		PublicKeyPem: string(pubKey),
	}
	app.LDContext.Compound = make([]interface{}, 0)
	app.LDContext.Compound = append(app.LDContext.Compound, "https://www.w3.org/ns/activitystreams")
	app.LDContext.Compound = append(app.LDContext.Compound, "https://w3id.org/security/v1")
	app.Id.Simple = id
	app.Type.Simple = "Application"
	now := time.Now().Format(time.RFC3339)
	app.Published = &streams.PropertyAS2{
		Simple: now,
	}
	app.Updated = &streams.PropertyAS2{
		Simple: now,
	}
	app.Inbox = &streams.PropertyAS2{
		Simple: id + "/inbox",
	}
	app.Outbox = &streams.PropertyAS2{
		Simple: id + "/outbox",
	}
	app.PreferredUsername = &streams.PropertyAS2{
		Simple: hostname,
	}
	app.PublicKey = &streams.PropertyAS2{
		Complex: pubKeyObj,
	}
	jsonData, err := json.Marshal(app)
	if err != nil {
		return app, err
	}
	err = dataStore.PutObject(db, []byte(app.Id.Simple), jsonData)
	if err != nil {
		return app, err
	}
	err = serverProto.CreateWF(db, "acct:"+registry.Hostname+"@"+registry.Hostname, app)
	dataStore.SetServerPrivateKey(db, privateKey)
	return app, err
}

func createObject(registry *streams.TypeRegistry, db *badger.DB) ([]byte, error) {
	var create streams.CreateAS2
	var note streams.NoteAS2
	cid := registry.Hostname + "/create/1"
	nid := registry.Hostname + "/note/1"
	actor := registry.Hostname + "/actor"
	LDContext := streams.PropertyAS2{
		Simple:   "https://w3.org/ns/activitystreams",
		Complex:  nil,
		Compound: nil,
	}
	create.LDContext = LDContext
	create.Id.Simple = cid
	create.Type.Simple = "Create"
	create.Actor = &streams.PropertyAS2{}
	create.Actor.Simple = actor
	note.LDContext = LDContext
	note.Id.Simple = nid
	note.Type.Simple = "Note"
	note.Published = &streams.PropertyAS2{}
	note.Published.Simple = time.Now().Format(time.RFC3339)
	note.AttributedTo = &streams.PropertyAS2{}
	note.AttributedTo.Simple = create.Actor.Simple
	note.Content = &streams.PropertyAS2{}
	note.Content.Simple = "Hello, world!"
	note.To = &streams.PropertyAS2{}
	note.To.Simple = "https://w3.org/ns/activitystreams#Public"
	create.Object = streams.PropertyAS2{}
	create.Object.Complex = note
	create.Map_ = map[string]*streams.PropertyAS2{}
	create.Map_["test"] = &streams.PropertyAS2{Simple: "test"}
	createJson, err := registry.MarshalFromAS2Type(create)
	noteJson, err := registry.MarshalFromAS2Type(note)
	dataStore.PutObject(db, []byte(note.Id.Simple), noteJson)
	dataStore.PutObject(db, []byte(create.Type.Simple), createJson)
	return createJson, err
}
