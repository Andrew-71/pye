package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"git.a71.su/Andrew71/pye/internal/config"
	"git.a71.su/Andrew71/pye/internal/models/user"
	"github.com/golang-jwt/jwt/v5"
)

var key *rsa.PrivateKey

// LoadKey attempts to load a private RS256 key from file.
// If the file does not exist, it generates a new key (and saves it)
func MustLoadKey() {
	// If the key doesn't exist, create it
	if _, err := os.Stat(config.Cfg.KeyFile); errors.Is(err, os.ErrNotExist) {
		key, err = rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			slog.Error("error generating key", "error", err)
			os.Exit(1)
		}

		// Save key to disk
		km := x509.MarshalPKCS1PrivateKey(key)
		block := pem.Block{Bytes: km, Type: "RSA PRIVATE KEY"}
		f, err := os.OpenFile(config.Cfg.KeyFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			slog.Error("error opening/creating file", "error", err)
			os.Exit(1)
		}
		f.Write(pem.EncodeToMemory(&block))
		if err := f.Close(); err != nil {
			slog.Error("error closing file", "error", err)
			os.Exit(1)
		}
		slog.Debug("generated new key", "file", config.Cfg.KeyFile)
	} else {
		km, err := os.ReadFile(config.Cfg.KeyFile)
		if err != nil {
			slog.Error("error reading key", "error", err)
			os.Exit(1)
		}
		key, err = jwt.ParseRSAPrivateKeyFromPEM(km)
		if err != nil {
			slog.Error("error parsing key", "error", err)
			os.Exit(1)
		}
		slog.Debug("loaded private key", "file", config.Cfg.KeyFile)
	}
}

// ServePublicKey returns our public key as PEM block over HTTP
func ServePublicKey(w http.ResponseWriter, r *http.Request) {
	key_marshalled := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	block := pem.Block{Bytes: key_marshalled, Type: "RSA PUBLIC KEY"}
	pem.Encode(w, &block)
}

// Create creates a JSON Web Token that expires after a week
func Create(user user.User) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"iss": "pye",
			"uid": user.Uuid,
			"sub": user.Email,
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		})
	token, err = t.SignedString(key)
	if err != nil {
		slog.Error("error creating JWT", "error", err)
		return "", err
	}
	return
}

// Verify receives a JWT and PEM-encoded public key,
// then returns whether the token is valid
func Verify(token string, publicKey []byte) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}
		return key, nil
	})
	return t, err
}

// VerifyLocal calls Verify with public key set to current local one
func VerifyLocal(token string) (*jwt.Token, error) {
	key_marshalled := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	block := pem.Block{Bytes: key_marshalled, Type: "RSA PUBLIC KEY"}
	return Verify(token, pem.EncodeToMemory(&block))
}
