/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const NOTION_API_BASE_URL = "https://api.notion.com"
const NOTION_API_VERSION = "2022-02-22"
const NOTION_SEARCH_URL = "https://api.notion.com/v1/search"
const NOTION_PAGE_URL = "https://api.notion.com/v1/pages"
const NOTION_BLOCK_URL = "https://api.notion.com/v1/blocks"

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
}
