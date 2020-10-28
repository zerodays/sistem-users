package main

import (
	"github.com/urfave/cli/v2"
	"github.com/zerodays/sistem-users/internal/cmd"
	"github.com/zerodays/sistem-users/internal/config"
	"github.com/zerodays/sistem-users/internal/database"
	"github.com/zerodays/sistem-users/internal/logger"
	"os"
)

func main() {
	// Load configuration.
	config.Load()

	// Initialize logger instance.
	logger.Init()

	// Initialize database.
	database.Init()
	defer database.Close()

	// Create new CLI app.
	app := cli.NewApp()

	// Basic info.
	app.Name = "Sistem users microservice"
	app.Authors = []*cli.Author{
		{
			Name:  "Vid Drobnič",
			Email: "vid.drobnic@gmail.com",
		},
	}
	app.Version = "0.0.1"

	// Commands for CLI app.
	app.Commands = []*cli.Command{
		cmd.Serve,
		cmd.GenKeys,
		cmd.CreateUser,
	}

	// Run the app.
	err := app.Run(os.Args)
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
	}
}
