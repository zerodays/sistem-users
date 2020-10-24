package config

import (
	"crypto/rand"
	"encoding/hex"
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

	// Path to config photos.
	confPath := filepath.Join(WorkDir, "conf", "config.ini")

	// Create default config if needed.
	err := createDefaultConfig(confPath)
	if err != nil {
		log.Fatalf("Could not create default config: %v\n", err)
	}

	// Load config file.
	cfg, err := ini.Load(confPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load server config.
	err = cfg.Section("server").MapTo(&Server)
	if err != nil {
		log.Fatal(err)
	}

	// Load database config.
	err = cfg.Section("database").MapTo(&Database)
	if err != nil {
		log.Fatal(err)
	}

	// Load logs section.
	err = cfg.Section("logs").MapTo(&Logs)
	if err != nil {
		log.Fatal(err)
	}

	// Load login section.
	err = cfg.Section("login").MapTo(&Login)
	if err != nil {
		log.Fatal(err)
	}

	// Load secret.
	err = loadSecret()
	if err != nil {
		log.Fatalf("Could not read secret: %v", err)
	}
}

// createDefaultConfig creates default settings file at `path`
// if settings file does not yet exist.
func createDefaultConfig(path string) error {
	// Check if config photos already exists.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("No config found. Creating default config at \"%s\"\n", path)

		// Create new config file.
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()

		// Get default config file to be written.
		data, err := Asset("config.ini")
		if err != nil {
			return err
		}

		// Write config file.
		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}

// loadSecret loads secret from file if it exists. If secret does not yet exists, new
// secret is created.
func loadSecret() error {
	secretPath := filepath.Join(WorkDir, "conf", "secret.txt")

	if _, err := os.Stat(secretPath); os.IsNotExist(err) {
		// Create new secret since it doesn't exist.
		log.Printf("No secret found. Creating default secret at \"%s\"\n", secretPath)

		// Create new file for secret.
		err := os.MkdirAll(filepath.Dir(secretPath), os.ModePerm)
		if err != nil {
			return err
		}

		f, err := os.Create(secretPath)
		if err != nil {
			return err
		}
		defer f.Close()

		// Get random bytes to be used as a secret.
		secretBytes := make([]byte, 128)
		_, err = rand.Read(secretBytes)
		if err != nil {
			return err
		}

		// Encode secret to string.
		secret := hex.EncodeToString(secretBytes)

		// Write secret to file.
		_, err = f.WriteString(secret)
		if err != nil {
			return err
		}

		// Set secret.
		Secret = hexString(secret)
	} else {
		// Read secret from existing file.
		// Open the file.
		f, err := os.Open(secretPath)
		if err != nil {
			return err
		}
		defer f.Close()

		// Read secret from file.
		secret, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		// Set secret.
		Secret = hexString(secret)
	}

	return nil
}
