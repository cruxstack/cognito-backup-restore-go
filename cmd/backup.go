package main

import (
	"github.com/cruxstack/cognito-backup-restore-go/internal/cognito"
	"github.com/spf13/cobra"
)

var (
	backupPoolId  string
	backupOutPath string
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup user from a cognito pool",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cognito.BackupUsers(backupPoolId, backupOutPath)
	},
}

func init() {
	backupCmd.Flags().StringVarP(&backupPoolId, "pool-id", "p", "", "pool id")
	backupCmd.Flags().StringVarP(&backupOutPath, "out", "o", "backup.json", "output path")
	rootCmd.AddCommand(backupCmd)
}
