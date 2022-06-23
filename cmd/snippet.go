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
	"time"

	"github.com/ppuczka/ntn_go/model"
	"github.com/spf13/cobra"
)

var snippetCmd = &cobra.Command{
	Use:   "snippet",
	Short: "Creates new snippet in notion.so",
	
	Run: func(cmd *cobra.Command, args []string) {
		snippet(cmd)
	},
}

func init() {
	currentTime := time.Now().Format("2017-09-07 17:06:06 Wednesday")
	rootCmd.AddCommand(snippetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	snippetCmd.Flags().StringP("title", "t",  currentTime, "Provide snippet title otherwise current timestamp will be added")
	snippetCmd.Flags().StringP("text", "x", "", "Provide snippet content text")
	snippetCmd.Flags().StringP("token", "o", os.Getenv("NOTION_TOKEN"), "Provide notion integrations token")
	authCmd.MarkFlagRequired("title")
	authCmd.MarkFlagRequired("text")
}

func snippet(cmd *cobra.Command) {
	token, err := cmd.Flags().GetString("token")
	if err != nil {
		fmt.Printf("error %s", err)
	}

	pageTitle, err := cmd.Flags().GetString("title")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	
	pageText, err := cmd.Flags().GetString("text")
	if err != nil {
		fmt.Printf("error %s", err)
	}

	parentPage := search_notion_page(token)
	notionPage := model.CreateSnippetPage(parentPage, pageTitle, pageText)
	
	response := create_notion_page(notionPage, token)
	fmt.Println(response)
}

func search_notion_page(token string)(model.Page) {
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
        return model.Page{}
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
	return cliPage

}
	
func create_notion_page(newPage model.Page, token string) (response string)  {
	jsonBody, _ := json.Marshal(newPage)
	payload := strings.NewReader(string(jsonBody))

	req := create_notion_request("POST", CREATE_PAGE_URL, token, payload)
	res, err := http.DefaultClient.Do(&req)
	if (err != nil) {
		log.Fatal("Error while sending request to notion")
	}
	defer res.Body.Close()
	
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func create_notion_request(requestMethod, url, token string, body io.Reader) (request http.Request) {
	req, _ := http.NewRequest(requestMethod, url, body) 

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Notion-Version", NOTION_API_VERSION)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	
	return *req
}   
