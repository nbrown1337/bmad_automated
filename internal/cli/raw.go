package cli

import (
	"context"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func newRawCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "raw <prompt>",
		Short: "Run an arbitrary prompt",
		Long: `Run an arbitrary prompt directly with Claude.
Useful for testing or one-off commands.

Example:
  bmad-automate raw "List all Go files in the project"`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			prompt := strings.Join(args, " ")
			ctx := context.Background()
			exitCode := app.Runner.RunRaw(ctx, prompt)
			os.Exit(exitCode)
		},
	}
}
