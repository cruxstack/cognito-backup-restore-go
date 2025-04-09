package main

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/cruxstack/cognito-backup-restore-go/internal/cognito"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list user pools",
	RunE: func(cmd *cobra.Command, args []string) error {
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

func init() {
	rootCmd.AddCommand(listCmd)
}
