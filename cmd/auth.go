/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize tool in notion.so ",
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auth called")
		auth(args)
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
	authCmd.Flags().BoolP("username", "u", true, "Provide notion registration email")
	authCmd.Flags().BoolP("password", "p", true, "Provide notion password")
}

func auth(params []string) {
	for _, p := range params {
		fmt.Println(p)
	}
}