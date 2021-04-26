package cmd

import (
	"fmt"

	"github.com/berkaygure/jira-cli/api"
	"github.com/spf13/cobra"
)

func NewWhoAmICmd() *cobra.Command {
	whoami := &cobra.Command{
		Use:   "whoami",
		Short: "Shows authenticated user's  details",
		Run: func(cmd *cobra.Command, args []string) {
			me := api.GetMe(restClient)
			fmt.Printf("%s (%s) <%s> from %s\n", me.DisplayName, me.Name, me.EmailAddress, me.TimeZone)
		},
	}

	return whoami
}
