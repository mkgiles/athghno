package serverProto

import (
	"athghno/internal/streams"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/dgraph-io/badger/v4"
	"strings"
)

func ValidateSignature(registry *streams.TypeRegistry, db *badger.DB, signature string) (interface{}, error) {
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
		return nil, err
	}
	switch res.(type) {
	case streams.PersonAS2:
		person := res.(streams.PersonAS2)
		pubKey := person.PublicKey.Complex.(streams.PublicKey)
		block, _ := pem.Decode([]byte(pubKey.PublicKeyPem))
		if block == nil {
			return nil, errors.New("failed to decode public key")
		}
		key, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		err = rsa.VerifyPKCS1v15(key, crypto.SHA256, []byte(signature), []byte(partMap["signature"]))
		if err != nil {
			return nil, err
		}
		return person, nil
	default:
		return nil, errors.New("unknown type")
	}
}
