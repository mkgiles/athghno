package serverProto

import (
	"athghno/internal/dataStore"
	"athghno/internal/streams"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/dgraph-io/badger/v4"
	"net/http"
	"strings"
)

func GetObject(registry *streams.TypeRegistry, db *badger.DB, request http.Request) ([]byte, error) {
	signature := request.Header.Get("Signature")
	if signature == "" {
		return nil, errors.New("Missing signature")
	}
	valid, err := ValidateSignature(registry, db, signature)
	if !valid {
		return nil, errors.New("Invalid signature")
	}
	if err != nil {
		return nil, err
	}
	object, err := dataStore.GetObject(db, []byte(request.URL.String()))
	if err != nil {
		return nil, err
	}
	return object, nil
}

func ValidateSignature(registry *streams.TypeRegistry, db *badger.DB, signature string) (bool, error) {
	parts := strings.Split(signature, ",")
	partMap := map[string]string{}
	for _, part := range parts {
		keyAndValue := strings.Split(part, "=")
		if len(keyAndValue) != 2 {
			continue
		}
		partMap[keyAndValue[0]] = keyAndValue[1]
	}
	res, err := FetchActor(registry, partMap["keyId"], db)
	if err != nil {
		return false, err
	}
	var pubKey streams.PublicKey
	switch res.(type) {
	case streams.PersonAS2:
		person := res.(streams.PersonAS2)
		pubKey = person.PublicKey.Complex.(streams.PublicKey)
	case streams.ApplicationAS2:
		app := res.(streams.ApplicationAS2)
		pubKey = app.PublicKey.Complex.(streams.PublicKey)
	case streams.GroupAS2:
		group := res.(streams.GroupAS2)
		pubKey = group.PublicKey.Complex.(streams.PublicKey)
	case streams.ServiceAS2:
		service := res.(streams.ServiceAS2)
		pubKey = service.PublicKey.Complex.(streams.PublicKey)
	case streams.OrganizationAS2:
		org := res.(streams.OrganizationAS2)
		pubKey = org.PublicKey.Complex.(streams.PublicKey)
	case streams.ActorAS2:
		actor := res.(streams.ActorAS2)
		pubKey = actor.PublicKey.Complex.(streams.PublicKey)
	default:
		return false, errors.New("unknown type")
	}
	block, _ := pem.Decode([]byte(pubKey.PublicKeyPem))
	if block == nil {
		return false, errors.New("failed to decode public key")
	}
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, []byte(signature), []byte(partMap["signature"]))
	if err != nil {
		return false, err
	}
	return true, nil
}
