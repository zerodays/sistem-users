package config

import "encoding/hex"

type hexString string

var (
	Server struct {
		ListenAddress string `ini:"LISTEN_ADDRESS"`
		Port          int    `ini:"PORT"`
		BaseAddress   string `ini:"BASE_ADDRESS"`
	}

	Database struct {
		Host     string `ini:"HOST"`
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
	}

	WorkDir string
	Secret  hexString
)

// Bytes returns bytes array that is represented by hex string s.
func (s hexString) Bytes() []byte {
	bytes, _ := hex.DecodeString(string(s))
	return bytes
}
