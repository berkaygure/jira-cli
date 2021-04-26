package api

import (
	"encoding/json"
	"fmt"
	"log"

	client "github.com/berkaygure/jira-cli/pkg"

	"golang.org/x/crypto/ssh/terminal"
)

func isAuthorized(client *client.Client) bool {
	if client.SessionId != "" {
		return true
	}

	return false
	// Get ME
	//rest/api/2/myself
}

// Log into Jira and return with a flag to detect that the user's session ID has changed
func Login(client *client.Client) (string, bool) {

	if isAuthorized(client) {
		return client.SessionId, false
	}

	username := client.Username

	if username == "" {
		fmt.Println("Enter your username: ")
		fmt.Scanf("%s\n", &username)
	}

	fmt.Println("Enter your password: ")
	password, _ := terminal.ReadPassword(0)

	response, _ := client.Post("rest/auth/1/session", map[string]string{"username": username, "password": string(password)})

	var jsonResponse struct {
		Session struct {
			Name  string
			Value string
		}
	}

	jsonErr := json.Unmarshal([]byte(response), &jsonResponse)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return jsonResponse.Session.Value, true
}

type Me struct {
	Name         string
	DisplayName  string
	TimeZone     string
	Locale       string
	EmailAddress string
}

func GetMe(client *client.Client) Me {
	response, _ := client.Get("rest/api/2/myself")

	var jsonResponse Me

	jsonErr := json.Unmarshal([]byte(response), &jsonResponse)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return jsonResponse
}
