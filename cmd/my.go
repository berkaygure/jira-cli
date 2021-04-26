package cmd

import (
	"github.com/berkaygure/jira-cli/cmd/my"

	"github.com/spf13/cobra"
)

func NewMyCmd() *cobra.Command {
	myCmd := &cobra.Command{
		Use:   "my",
		Short: "Everything related with me",
	}

	myCmd.AddCommand(my.NewMyIssuesCmd(restClient))

	return myCmd
}
