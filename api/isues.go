package api

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/berkaygure/jira-cli/pkg"
)

type Jql struct {
	Project  string
	assignee string
}

type issueResponse struct {
	Issues []struct {
		Key    string
		Fields struct {
			Summary string
		}
	}
}

func SearchIssues(client *pkg.Client, q Jql) map[string]string {
	response, _ := client.Get("rest/api/2/search?jql=" + build_query(q))

	var issueResponse issueResponse

	jsonErr := json.Unmarshal([]byte(response), &issueResponse)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	var mappedResult = make(map[string]string)

	for _, issues := range issueResponse.Issues {
		mappedResult[issues.Key] = issues.Fields.Summary
	}

	return mappedResult

}

func build_query(q Jql) string {
	return url.QueryEscape("assignee=currentuser() and (status != \"Done\")  and project=" + q.Project)
}
