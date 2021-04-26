package cmd

import (
	"github.com/berkaygure/jira-cli/pkg"

	"github.com/spf13/cobra"
)

var restClient *pkg.Client // All sub commands may want to access it

func Execute(client *pkg.Client) error {
	restClient = client

	// Define root commands
	rootCmd := &cobra.Command{
		Use:   "jira <command> <subcommand> [flags]",
		Short: "jira is a CLI too for Jira",
	}

	// Register all sub commands
	rootCmd.AddCommand(NewWhoAmICmd())
	rootCmd.AddCommand(NewMyCmd())

	return rootCmd.Execute()
}
