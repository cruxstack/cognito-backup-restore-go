package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cbr",
	Short: "cli tool for backing up and restoring aws cognito user",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("command failed", "error", err)
		os.Exit(1)
	}
}
