package dataStore

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/dgraph-io/badger/v4"
	"log"
)

func ConnectDB() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions("/tmp/athghno.db"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetObject(db *badger.DB, key []byte) ([]byte, error) {
	if string(key) == "/PrivateKey" {
		return nil, errors.New("unauthorized access of server private key")
	}
	txn := db.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get(key)
	if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)
}

func PutObject(db *badger.DB, key []byte, value []byte) error {
	txn := db.NewTransaction(true)
	defer txn.Discard()
	err := txn.Set(key, value)
	if err != nil {
		return err
	}
	return txn.Commit()
}

func GetServerPrivateKey(db *badger.DB) (*rsa.PrivateKey, error) {
	txn := db.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte("/PrivateKey"))
	if err != nil {
		return nil, err
	}
	privKey, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privKey)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func SetServerPrivateKey(db *badger.DB, privateKey *rsa.PrivateKey) {
	privKeypcks1 := x509.MarshalPKCS1PrivateKey(privateKey)
	privKeyPem := string(pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privKeypcks1,
	}))
	privKeyEntry := badger.NewEntry([]byte("/PrivateKey"), []byte(privKeyPem))
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.SetEntry(privKeyEntry)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
