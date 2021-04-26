package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/spf13/viper"
)

type Client struct {
	host       string
	SessionId  string
	Username   string
	httpClient *http.Client
}

// Makes GET Request
func (c *Client) Get(path string) (response string, status int) {
	status, response = c.request(http.MethodGet, c.host+path, nil)
	return
}

// Makes POST Request
func (c *Client) Post(path string, data map[string]string) (response string, status int) {
	status, response = c.request(http.MethodPost, c.host+path, data)
	return
}

// Makes HTTP request to host with given data
func (c *Client) request(requestType string, url string, data map[string]string) (int, string) {

	json_data, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	if c.SessionId != "" {
		req.AddCookie(&http.Cookie{Name: "JSESSIONID", Value: c.SessionId})
	}

	res, getErr := c.httpClient.Do(req)
	status, response := parseResponse(res, getErr)

	if status == http.StatusUnauthorized {
		viper.Set("sessionId", "")
		viper.WriteConfig()
		fmt.Printf("%d Unauthenticated, please try to login again\n", status)
		os.Exit(1)
	}

	return status, response
}

// Parses HTTP responses to string
func parseResponse(response *http.Response, err error) (int, string) {
	if err != nil {
		log.Fatal(err)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return response.StatusCode, string(body)
}

// Creates New Jira Rest CLient
func NewClient(host string, sessionId string, username string) *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	tr := http.DefaultTransport
	httpClient := &http.Client{Transport: tr, Jar: jar}

	return &Client{
		host:       host,
		SessionId:  sessionId,
		Username:   username,
		httpClient: httpClient,
	}
}
