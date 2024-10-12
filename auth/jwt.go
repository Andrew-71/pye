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

	"git.a71.su/Andrew71/pye/config"
	"git.a71.su/Andrew71/pye/storage"
	"github.com/golang-jwt/jwt/v5"
)

var (
	key     *rsa.PrivateKey
)

// LoadKey attempts to load a private key from KeyFile.
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
		slog.Info("generated new key", "file", config.Cfg.KeyFile)
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
		slog.Info("loaded private key", "file", config.Cfg.KeyFile)
	}
}

func init() {
	MustLoadKey()
}

// PublicKey returns our public key as PEM block over http
func PublicKey(w http.ResponseWriter, r *http.Request) {
	key_marshalled := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	block := pem.Block{Bytes: key_marshalled, Type: "RSA PUBLIC KEY"}
	pem.Encode(w, &block)
}

func CreateJWT(user storage.User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"iss": "pye",
			"uid": user.Uuid,
			"sub": user.Email,
			"iat": time.Now().Unix(),
			"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		})
	s, err := t.SignedString(key)
	if err != nil {
		slog.Error("error creating JWT", "error", err)
		return "", err
	}
	return s, nil
}

// VerifyToken receives a JWT and PEM-encoded public key,
// then returns whether the token is valid
func VerifyJWT(token string, publicKey []byte) (*jwt.Token, error) {
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