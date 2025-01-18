package serverProto

import (
	"athghno/internal/dataStore"
	"athghno/internal/streams"
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"io"
	"net/http"
	"net/url"
	"time"
)

func FetchActor(registry *streams.TypeRegistry, id string, db *badger.DB) (interface{}, error) {
	privKey, err := dataStore.GetServerPrivateKey(db)
	if err != nil {
		return nil, err
	}
	actorUrl, err := url.Parse(id)
	if err != nil {
		return nil, err
	}
	httpClient := http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("GET", actorUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/activity+json,application/ld+json,application/json")
	date := time.Now().UTC().Format(time.ANSIC)
	req.Header.Set("Date", date)
	cleartext := fmt.Sprintf("(request-target): get %s\nhost: %s\ndate: %s", actorUrl.Path, actorUrl.Host, date)
	hash := sha256.New()
	hash.Write([]byte(cleartext))
	signed, err := privKey.Sign(rand.Reader, hash.Sum(nil), crypto.SHA256)
	signature := fmt.Sprintf(`keyId="%s",headers="%s",signature="%s",algorithm="rsa-sha256"`,
		registry.Hostname+"/actor#main-key",
		"(request-target) host date",
		base64.URLEncoding.EncodeToString(signed))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Signature", signature)
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		resData, _ := io.ReadAll(resp.Body)
		fmt.Println(string(resData))
		return nil, err
	}
	defer resp.Body.Close()
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		fmt.Println(string(data))
		return nil, errors.New(resp.Status)
	}
	as2, err := registry.UnmarshalIntoAS2Type(jsonData)
	if err != nil {
		return nil, err
	}
	return as2, nil
}
