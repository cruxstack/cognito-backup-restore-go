package cmd

import (
	"fmt"

	"github.com/cruxstack/cognito-backup-restore-go/internal/cognito"
	"github.com/urfave/cli/v2"
)

func newListCmd() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "list user pools",
		Action: func(c *cli.Context) error {
			pools, err := cognito.ListUserpools()
			if err != nil {
				return err
			}

			for _, p := range pools {
				fmt.Printf("id: %s name: %s\n", *p.Id, *p.Name)
			}

			return nil
		},
	}
}
