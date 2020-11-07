package config

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Load loads config and secret
func Load() {
	// Get current workdir.
	WorkDir = os.Getenv("USERS_WORKDIR")

	// Load default config
	data, err := Asset("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := ini.Load(data)
	if err != nil {
		log.Fatal(err)
	}

	// Load server config.
	err = cfg.Section("server").MapTo(&Server)
	if err != nil {
		log.Fatal(err)
	}
	getFromEnvironment("server", &Server)

	// Load database config.
	err = cfg.Section("database").MapTo(&Database)
	if err != nil {
		log.Fatal(err)
	}
	getFromEnvironment("database", &Database)

	// Load logs section.
	err = cfg.Section("logs").MapTo(&Logs)
	if err != nil {
		log.Fatal(err)
	}
	getFromEnvironment("logs", &Logs)

	// Load login section.
	err = cfg.Section("login").MapTo(&Login)
	if err != nil {
		log.Fatal(err)
	}
	getFromEnvironment("login", &Login)

	// Load keys used for signing.
	err = loadSigningKeys()
	if err != nil {
		log.Printf("Could not load signing keys (error: %s). You can generate them with `genkeys` command.\n",
			err.Error())
	}
}

// loadSigningKeys loads private and public RSA key used for signing authentication tokens.
func loadSigningKeys() error {
	// Open file for private key.
	privKeyPath := filepath.Join(WorkDir, "conf", "privkey.pem")
	privFile, err := os.Open(privKeyPath)
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
	Login.SigningPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		return err
	}

	// Open file for public key
	pubKeyPath := filepath.Join(WorkDir, "conf", "pubkey.pem")
	pubFile, err := os.Open(pubKeyPath)
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
	Login.SigningPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		return err
	}

	// Check if private key is valid.
	err = Login.SigningPrivateKey.Validate()
	if err != nil {
		return err
	}

	// Check if private and public keys match.
	if !Login.SigningPrivateKey.PublicKey.Equal(Login.SigningPublicKey) {
		return errors.New("private and public keys mismatch")
	}

	return nil
}
