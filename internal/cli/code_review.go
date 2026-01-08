package cli

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

func newCodeReviewCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "code-review <story-key>",
		Short: "Run code-review workflow",
		Long:  `Run the code-review workflow for the specified story key.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			storyKey := args[0]
			ctx := context.Background()
			exitCode := app.Runner.RunSingle(ctx, "code-review", storyKey)
			os.Exit(exitCode)
		},
	}
}
