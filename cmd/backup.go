package cmd

import (
	"github.com/cruxstack/cognito-backup-restore-go/internal/cognito"
	"github.com/urfave/cli/v2"
)

func newBackupCmd() *cli.Command {
	return &cli.Command{
		Name:  "backup",
		Usage: "backup users from a cognito pool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "pool-id",
				Aliases:  []string{"p"},
				Usage:    "pool id",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "out",
				Aliases: []string{"o"},
				Usage:   "output path",
				Value:   "backup.json",
			},
		},
		Action: func(c *cli.Context) error {
			return cognito.BackupUsers(c.String("pool-id"), c.String("out"))
		},
	}
}
