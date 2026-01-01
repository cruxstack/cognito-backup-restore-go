package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:  "cbr",
		Usage: "cli tool for backing up and restoring aws cognito users",
		Commands: []*cli.Command{
			newBackupCmd(),
			newRestoreCmd(),
			newListCmd(),
		},
	}
}
