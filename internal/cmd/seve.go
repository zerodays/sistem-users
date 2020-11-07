package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/zerodays/sistem-users/internal/config"
	"github.com/zerodays/sistem-users/internal/logger"
	"github.com/zerodays/sistem-users/internal/router"
	"net/http"
)

var Serve = &cli.Command{
	Name:   "serve",
	Usage:  "Start the sever.",
	Action: serve,
}

func serve(_ *cli.Context) error {
	// Check that signing keys are valid.
	if config.Login.SigningPublicKey == nil || config.Login.SigningPrivateKey == nil {
		logger.Log.Fatal().Msg("Can not run server without valid private and public signing keys.")
	}

	// Load router.
	r := router.NewRouter()

	// Start listening for connections.
	listenAddress := fmt.Sprintf("%s:%d", config.Server.ListenAddress(), config.Server.Port())
	logger.Log.Fatal().Err(http.ListenAndServe(listenAddress, r)).Send()

	return nil
}
