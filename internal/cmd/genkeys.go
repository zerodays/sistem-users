package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/urfave/cli/v2"
	"github.com/zerodays/sistem-users/internal/logger"
	"github.com/zerodays/sistem-users/internal/util"
	"os"
	"path/filepath"
)

const (
	flagBits   = "bits"
	flagOutDir = "outDir"
)

var GenKeys = &cli.Command{
	Name:  "genkeys",
	Usage: "Generate keys for signing authentication tokens.",
	Description: `Generates private and public RSA key pair used for signing authentication token. Outputs
	 to <outDir>/privkey.pem and <outDir>/pubkey.pem and will overwrite existing files.`,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  flagBits,
			Usage: "Size of RSA key to generate.",
			Value: 2048,
		},
		&cli.StringFlag{
			Name:  flagOutDir,
			Usage: "Output directory for generated keys.",
			Value: "conf",
		},
	},
	Action: genKeys,
}

func genKeys(c *cli.Context) error {
	// Generate private key.
	privKey, err := rsa.GenerateKey(rand.Reader, c.Int(flagBits))
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
		return nil
	}

	// Create output directory.
	outDir := c.String(flagOutDir)
	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
		return nil
	}

	// Open file for private key.
	privKeyPath := filepath.Join(outDir, "privkey.pem")
	privKeyFile, err := os.Create(privKeyPath)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
		return nil
	}
	defer privKeyFile.Close()

	// Marshal private key and write it to file.
	privKeyBytes, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
		return nil
	}
	privKeyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privKeyBytes,
	}
	err = pem.Encode(privKeyFile, privKeyBlock)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}

	// Open file for public key.
	pubKeyPath := filepath.Join(outDir, "pubkey.pem")
	pubKeyFile, err := os.Create(pubKeyPath)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
		return nil
	}
	defer pubKeyFile.Close()

	// Marshal public key and write it to file.
	err = util.EncodePubKey(&privKey.PublicKey, pubKeyFile)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
		return nil
	}

	return nil
}
