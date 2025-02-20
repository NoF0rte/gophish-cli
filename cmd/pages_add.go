/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
)

// pagesAddCmd represents the add command
var pagesAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new page",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		pagePaths, _ := cmd.Flags().GetStringSlice("pages")

		if dir != "" {
			files, err := filepath.Glob(filepath.Join(dir, "*.yaml"))
			checkError(err)

			pagePaths = append(pagePaths, files...)

			files, err = filepath.Glob(filepath.Join(dir, "*.yml"))
			checkError(err)

			pagePaths = append(pagePaths, files...)
		}

		for _, p := range pagePaths {
			page, err := models.PageFromFile(p, variables)
			checkError(err)

			fmt.Printf("[+] Adding landing page \"%s\"\n", page.Name)

			_, err = client.CreateLandingPage(page)
			checkError(err)
		}
	},
}

func init() {
	pagesCmd.AddCommand(pagesAddCmd)

	pagesAddCmd.Flags().StringP("dir", "d", "", "Directory containing the pages to add. Both .yaml and .yml files will be used.")
	pagesAddCmd.Flags().StringSliceP("pages", "p", []string{}, "The paths of the page files to add. Specify multiple times to add more pages.")
}
