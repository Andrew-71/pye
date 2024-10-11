package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"log/slog"

	// "github.com/golang-jwt/jwt/v5"
)

// var (
// 	key *ecdsa.PrivateKey
// 	t   *jwt.Token
// 	s   string
// 	key string
// )

func CreateKey() {
	// TODO: Is this a secure key?
	k, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		slog.Error("Error generating key", "error", err)
	}
	km, _ := x509.MarshalECPrivateKey(k)
	slog.Info("Key", "key", km)
}

// func CreateJWT(usr User) string {

// 	t := jwt.NewWithClaims(jwt.SigningMethodES256,
// 		jwt.MapClaims{
// 			"iss": "my-auth-server",
// 			"sub": "john",
// 			"foo": 2,
// 		})
// 	s, err := t.SignedString(key)
// 	if err != nil {
// 		slog.Error("Error creating JWT", "error", err)
// 		// TODO: Something
// 	}
// 	return s
// }
