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

func GetServerPublicKey(db *badger.DB) ([]byte, error) {
	txn := db.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte("/PublicKey"))
	if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)
}

func GetServerPrivateKey(db *badger.DB) ([]byte, error) {
	txn := db.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte("/PrivateKey"))
	if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)
}

func SetServerPrivateKey(db *badger.DB, privateKey *rsa.PrivateKey) {
	pubKeyEntry := badger.NewEntry([]byte("/PublicKey"), pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}))
	privKeyEntry := badger.NewEntry([]byte("/PrivateKey"), pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}))
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.SetEntry(privKeyEntry)
		if err != nil {
			return err
		}
		err = txn.SetEntry(pubKeyEntry)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
