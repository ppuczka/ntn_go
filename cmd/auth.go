/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const NOTION_API_BASE_URL = "https://api.notion.com" 
const NOTION_API_VERSION = "2022-02-22"

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize tool in notion.so ",
	
	Run: func(cmd *cobra.Command, args []string) {
		auth(cmd)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	authCmd.Flags().StringP("username", "u", "", "Provide notion registration email")
	authCmd.Flags().StringP("token", "t", os.Getenv("NOTION_TOKEN"), "Provide notion integrations token")
	// authCmd.MarkFlagRequired("password")
	// authCmd.MarkFlagRequired("token")
}

func auth(cmd *cobra.Command) {
	
	// username, err := cmd.Flags().GetString("username")
	// if err != nil {
	// 	fmt.Printf("error %s", err)
	// }
	
	token, err := cmd.Flags().GetString("token")
	if err != nil {
		fmt.Printf("error %s", err)
	}

	url := "https://api.notion.com/v1/search"

	payload := strings.NewReader("\"page_size\":100\n\"query\":\"External tasks\",\n\"sort\":{\n\"direction\":\"ascending\",\n\"timestamp\":\"last_edited_time\"}\n")
	fmt.Print(payload)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Notion-Version", NOTION_API_VERSION)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error occured: %e", err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	
	var prettyJSON bytes.Buffer
    error := json.Indent(&prettyJSON, body, "", "\t")
    if error != nil {
        log.Println("JSON parse error: ", error)
        return
    }

	fmt.Println(string(prettyJSON.String()))
	
}