/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ppuczka/ntn_go/model"
	"github.com/spf13/cobra"
)

const NOTION_API_BASE_URL = "https://api.notion.com" 
const NOTION_API_VERSION = "2022-02-22"
const SEARCH_URL = "https://api.notion.com/v1/search"
const CREATE_PAGE_URL = "https://api.notion.com/v1/pages"

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

	// pageurl := "https://api.notion.com/v1/pages/2f94ff6b-8e94-42b6-8032-435a847b8a38"
	// blockUrl := "https://api.notion.com/v1/blocks/352f3ffe-057a-4ecf-bda3-9a65e1cd99b0/children"

	payload := strings.NewReader("{\"query\":\"CLI Snippets\",\"filter\":{\"value\":\"page\",\"property\":\"object\"}}")	

	req := create_notion_request("POST", SEARCH_URL, token, payload)
	res, err := http.DefaultClient.Do(&req)
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
	var pages model.Pages
	var cliPage model.Page
	
	json.Unmarshal(body, &pages)
	// fmt.Println(string(prettyJSON.String()))
	for _, p := range pages.Pages {
		if (strings.Contains(p.Url, "CLI-Snippets")) {
		cliPage = p
		}
	}

	jsonData, _ := json.Marshal(cliPage)
	notionPage := model.CreateSnippetPage(cliPage, "new title", "new text")

	fmt.Println(string(jsonData))

	response := create_notion_page(notionPage, token)
	fmt.Println(response)
}
	
func create_notion_page(newPage model.Page, token string) (response string)  {
	jsonBody, _ := json.Marshal(newPage)
	payload := strings.NewReader(string(jsonBody))

	req := create_notion_request("POST", CREATE_PAGE_URL, token, payload)
	res, _ := http.DefaultClient.Do(&req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func create_notion_request(requestMethod, url, token string, body io.Reader) (request http.Request) {
	req, _ := http.NewRequest(requestMethod, url, body) 

	// req.Header.Add("Accept", "application/json")
	req.Header.Add("Notion-Version", NOTION_API_VERSION)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	
	return *req
}   