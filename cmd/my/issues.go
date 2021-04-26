package my

import (
	"fmt"

	"github.com/berkaygure/jira-cli/api"
	"github.com/berkaygure/jira-cli/pkg"

	"github.com/spf13/cobra"
)

var filter = &api.Jql{
	Project: "",
}

func NewMyIssuesCmd(restClient *pkg.Client) *cobra.Command {

	myIssuesCmd := &cobra.Command{
		Use:   "issues",
		Short: "Shows all my issues",
		Run: func(cmd *cobra.Command, args []string) {

			response := api.SearchIssues(restClient, *filter)

			for k, v := range response {
				fmt.Printf("[%s]\t%s\n", k, v)
			}
		},
	}

	myIssuesCmd.Flags().StringVarP(&filter.Project, "project", "p", "", "project SIMX")

	return myIssuesCmd
}
