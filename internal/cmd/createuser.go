package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/zerodays/sistem-users/internal/logger"
	"github.com/zerodays/sistem-users/internal/models/user"
)

const (
	flagEmail    = "email"
	flagName     = "name"
	flagPassword = "password"
)

var CreateUser = &cli.Command{
	Name:  "createuser",
	Usage: "Create new user.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     flagEmail,
			Usage:    "Email for new user.",
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagName,
			Usage:    "Name for new user.",
			Required: true,
		},

		&cli.StringFlag{
			Name:     flagPassword,
			Usage:    "Password fro new user.",
			Required: true,
		},
	},
	Action: createUser,
}

func createUser(c *cli.Context) error {
	u, err := user.New(c.String(flagEmail), c.String(flagName), c.String(flagPassword))
	if err != nil {
		logger.Log.Fatal().Err(err).Send()
		return nil
	}

	fmt.Printf("Successfully create new user with UID \"%s\".\n", u.UID)

	return nil
}
