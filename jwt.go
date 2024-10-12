package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var KeyFile = "key"

var (
	key *ecdsa.PrivateKey
	// t   *jwt.Token
)

// LoadKey attempts to load a private key from KeyFile.
// If the file does not exist, it generates a new key (and saves it)
func LoadKey() {
	// If the key doesn't exist, create it
	if _, err := os.Stat(KeyFile); errors.Is(err, os.ErrNotExist) {
		key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			slog.Error("error generating key", "error", err)
			os.Exit(1)
		}
		km, err := x509.MarshalECPrivateKey(key) // Save private key to disk
		if err != nil {
			slog.Error("error marshalling key", "error", err)
			os.Exit(1)
		}
		os.WriteFile(KeyFile, km, 0644)
		slog.Info("generated new key")
	} else {
		km, err := os.ReadFile(KeyFile)
		if err != nil {
			slog.Error("error reading key", "error", err)
			os.Exit(1)
		}
		key, err = x509.ParseECPrivateKey(km)
		if err != nil {
			slog.Error("error parsing key", "error", err)
			os.Exit(1)
		}
		slog.Info("loaded private key")
	}
	slog.Debug("private key", "key", key)
}

// publicKey returns our public key in PKIX, ASN.1 DER form
func publicKey(w http.ResponseWriter, r *http.Request) {
	key_marshalled, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		slog.Error("error marshalling public key", "error", err)
		http.Error(w, "error marshalling public key", http.StatusInternalServerError)
		return
	}
	// w.Write(key_marshalled)
	block := pem.Block{Bytes: key_marshalled, Type: "ECDSA PUBLIC KEY"}
	// slog.Info("public key", "orig", key_marshalled, "block", block)
	pem.Encode(w, &block)
}

func init() {
	LoadKey()
}

func CreateJWT(usr User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"iss": "pye",
			"sub": "john",
			"foo": 2,
		})
	s, err := t.SignedString(key)
	if err != nil {
		slog.Error("Error creating JWT", "error", err)
		return "", err
	}
	return s, nil
}
