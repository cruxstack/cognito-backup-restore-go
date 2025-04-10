package cmd

import (
	"github.com/cruxstack/cognito-backup-restore-go/internal/cognito"
	"github.com/spf13/cobra"
)

var (
	restorePoolId string
	restoreInPath string
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore users from a backup",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cognito.RestoreUsers(restorePoolId, restoreInPath)
	},
}

func init() {
	restoreCmd.Flags().StringVarP(&restorePoolId, "pool-id", "p", "", "pool id")
	restoreCmd.Flags().StringVarP(&restoreInPath, "in", "i", "backup.json", "input path")
	rootCmd.AddCommand(restoreCmd)
}
