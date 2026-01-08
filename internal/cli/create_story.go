package cli

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

func newCreateStoryCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "create-story <story-key>",
		Short: "Run create-story workflow",
		Long:  `Run the create-story workflow for the specified story key.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			storyKey := args[0]
			ctx := context.Background()
			exitCode := app.Runner.RunSingle(ctx, "create-story", storyKey)
			os.Exit(exitCode)
		},
	}
}
