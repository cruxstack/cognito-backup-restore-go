package cmd

import (
	"github.com/cruxstack/cognito-backup-restore-go/internal/cognito"
	"github.com/urfave/cli/v2"
)

func newRestoreCmd() *cli.Command {
	return &cli.Command{
		Name:  "restore",
		Usage: "restore users from a backup",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "pool-id",
				Aliases:  []string{"p"},
				Usage:    "pool id",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "in",
				Aliases: []string{"i"},
				Usage:   "input path",
				Value:   "backup.json",
			},
		},
		Action: func(c *cli.Context) error {
			return cognito.RestoreUsers(c.String("pool-id"), c.String("in"))
		},
	}
}
