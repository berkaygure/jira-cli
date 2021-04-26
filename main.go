package main

import (
	"fmt"
	"os"

	"github.com/berkaygure/jira-cli/api"
	root "github.com/berkaygure/jira-cli/cmd"
	c "github.com/berkaygure/jira-cli/pkg"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	configFile string = "jira"
)

func main() {
	// Create configuration
	client, _ := initialization()

	// Login into Jira
	sessionId, isItNew := api.Login(client)

	// Save settings
	if isItNew {
		viper.Set("sessionId", sessionId)
		viper.WriteConfig()
	}

	// Execute root commands
	root.Execute(client)
}

func initialization() (client *c.Client, url string) {
	// Configure viper
	homedir, _ := homedir.Dir()
	viper.SetConfigName(configFile)
	viper.SetConfigType("json")
	viper.AddConfigPath(homedir)

	// Create config file if it's not exists
	createConfigFile(homedir + "/" + configFile + ".json")

	// Read settings
	viper.ReadInConfig()

	url = viper.GetString("url")

	if url == "" {
		fmt.Println("Enter your JIRA url: ")
		fmt.Scanf("%s\n", &url)
		viper.Set("url", url)
	}

	// save settings
	viper.WriteConfig()

	client = c.NewClient(url, viper.GetString("sessionId"), viper.GetString("username"))

	return
}

// Creates configuration file
func createConfigFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if e := viper.SafeWriteConfig(); e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
	}
}
