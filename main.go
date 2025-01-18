package main

import (
	"athghno/internal/dataStore"
	"athghno/internal/streams"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"
)

func main() {
	registry := streams.InitRegistry()
	db := dataStore.ConnectDB()
	defer db.Close()
	if _, err := dataStore.GetServerPublicKey(db); err != nil {
		fmt.Println("No server public key found. Generating new one")
		privKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			panic(err)
		}
		dataStore.SetServerPrivateKey(db, privKey)
	}
	storedObj, err := dataStore.GetObject(db, []byte("https://example.com/create/1"))
	if err != nil {
		if err.Error() != "Key not found" {
			panic(err)
		} else {
			jsonData, err := createObject(registry)
			if err != nil {
				panic(err)
			}
			storedObj = jsonData
			err = dataStore.PutObject(db, []byte("https://example.com/create/1"), storedObj)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println(string(storedObj))
	as2, err := registry.UnmarshalIntoAS2Type(storedObj)
	if err != nil {
		return
	}
	fmt.Println(as2)
}

func createObject(registry *streams.TypeRegistry) ([]byte, error) {
	var create streams.CreateAS2
	var note streams.NoteAS2
	LDContext := streams.PropertyAS2{
		Simple:   "https://w3.org/ns/activitystreams",
		Complex:  nil,
		Compound: nil,
	}
	create.LDContext = LDContext
	create.Id.Simple = "https://example.com/create/1"
	create.Type.Simple = "Create"
	create.Actor = &streams.PropertyAS2{}
	create.Actor.Simple = "https://example.com/actor"
	note.LDContext = LDContext
	note.Id.Simple = "https://example.com/note/1"
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
	jsonData, err := registry.MarshalFromAS2Type(create)
	return jsonData, err
}
