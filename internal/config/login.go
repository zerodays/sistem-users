package config

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
)

var ErrKeyMismatch = errors.New("private and public keys don't match")

// login section
type login struct {
	// Private and public keys used for signing tokens.
	SigningPrivateKey *rsa.PrivateKey
	SigningPublicKey  *rsa.PublicKey
}

func (l login) TokenTtl() int {
	val, _ := cfg.GetInt("login", "token_ttl")
	return val
}

func (l login) SigningPrivateKeyLocation() string {
	val, _ := cfg.GetString("login", "signing_private_key_location")
	return val
}

func (l login) SigningPublicKeyLocation() string {
	val, _ := cfg.GetString("login", "signing_public_key_location")
	return val
}

// LoadKeys loads keys used for signing jwt tokens.
func (l *login) LoadKeys() error {
	// Open file for private key.
	privFile, err := os.Open(l.SigningPrivateKeyLocation())
	if err != nil {
		return err
	}
	defer privFile.Close()

	// Read the file.
	privKey, err := ioutil.ReadAll(privFile)
	if err != nil {
		return err
	}

	// Parse private key.
	l.SigningPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		return err
	}

	// Open file for public key
	pubFile, err := os.Open(l.SigningPublicKeyLocation())
	if err != nil {
		return err
	}
	defer pubFile.Close()

	// Read the file.
	pubKey, err := ioutil.ReadAll(pubFile)
	if err != nil {
		return err
	}

	// Parse public key.
	l.SigningPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		return err
	}

	// Check if private key is valid.
	err = l.SigningPrivateKey.Validate()
	if err != nil {
		return err
	}

	// Check if private and public keys match.
	if !l.SigningPrivateKey.PublicKey.Equal(l.SigningPublicKey) {
		return ErrKeyMismatch
	}

	return nil
}
