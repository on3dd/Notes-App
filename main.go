package main

import (
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"notes-app/api"
	dbpkg "notes-app/db"
)

const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath = "keys/app.rsa.pub"
)

func main() {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	singKey, verifyKey := initKeys()

	db := dbpkg.New()

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
	}

	apiHandler := api.New(db, singKey, verifyKey)
	http.Handle("/api/", apiHandler)

	log.Printf("Server successfully started at port %v\n", server.Addr)
	log.Println(server.ListenAndServe())
}

func initKeys() (signKey, verifyKey []byte) {
	var err error

	signKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Error reading private key, error: %v", err)
	}

	verifyKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("Error reading public key, error: %v", err)
	}

	return signKey, verifyKey
}