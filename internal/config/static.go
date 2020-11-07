package config

import "crypto/rsa"

var (
	Server struct {
		ListenAddress string `ini:"LISTEN_ADDRESS"`
		Port          int    `ini:"PORT"`
		BaseAddress   string `ini:"BASE_ADDRESS"`
	}

	Database struct {
		Host     string `ini:"HOST"`
		Port     int    `ini:"PORT"`
		User     string `ini:"USER"`
		Password string `ini:"PASSWORD"`
		DbName   string `ini:"DB_NAME"`
		SslMode  string `ini:"SSL_MODE"`
	}

	Logs struct {
		LogLevel       int    `ini:"LOG_LEVEL"`
		FileLogging    bool   `ini:"FILE_LOGGING"`
		ConsoleLogging bool   `ini:"CONSOLE_LOGGING"`
		LogPath        string `ini:"LOG_PATH"`

		MaxSize    int `ini:"MAX_SIZE"`
		MaxAge     int `ini:"MAX_AGE"`
		MaxBackups int `ini:"MAX_BACKUPS"`
	}

	Login struct {
		TokenTtl int `ini:"TOKEN_TTL"`

		SigningPrivateKeyLocation string `ini:"SIGNING_PRIVATE_KEY_LOCATION"`
		SigningPublicKeyLocation  string `ini:"SIGNING_PUBLIC_KEY_LOCATION"`

		// Private and public keys used for signing tokens.
		SigningPrivateKey *rsa.PrivateKey
		SigningPublicKey  *rsa.PublicKey
	}
)
