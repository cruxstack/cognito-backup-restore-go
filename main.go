package main

import (
	"log/slog"
	"os"

	"github.com/cruxstack/cognito-backup-restore-go/cmd"
)

func main() {
	app := cmd.NewApp()
	if err := app.Run(os.Args); err != nil {
		slog.Error("command failed", "error", err)
		os.Exit(1)
	}
}
