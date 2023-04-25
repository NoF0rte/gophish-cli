package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
)

// pagesCmd represents the pages command
var pagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "List, add, or delete landing pages",
	Run: func(cmd *cobra.Command, args []string) {
		showContent, _ := cmd.Flags().GetBool("show-content")
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		re, _ := cmd.Flags().GetString("regex")

		var err error
		var pages []*models.Page
		if id > 0 {
			page, err := client.GetLandingPageByID(id)
			checkError(err)

			if page != nil {
				pages = append(pages, page)
			}
		} else if name != "" {
			page, err := client.GetLandingPageByName(name)
			checkError(err)

			if page != nil {
				pages = append(pages, page)
			}
		} else if re != "" {
			pages, err = client.GetLandingPagesByRegex(re)
			checkError(err)
		} else {
			pages, err = client.GetLandingPages()
			checkError(err)
		}

		if len(pages) == 0 {
			fmt.Println("[!] No pages found")
			return
		}

		if !showContent {
			for _, p := range pages {
				p.HTML = ""
			}
		}

		if len(pages) == 1 {
			data, err := pages[0].ToJSON()
			checkError(err)

			fmt.Println(data)
			return
		}

		data, err := json.MarshalIndent(pages, "", "  ")
		checkError(err)

		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(pagesCmd)

	pagesCmd.Flags().StringP("name", "n", "", "Get the page by name.")
	pagesCmd.Flags().Int("id", 0, "Get the page by ID")
	pagesCmd.Flags().StringP("regex", "r", "", "List the pages with the name matching the regex.")
	pagesCmd.Flags().Bool("show-content", false, "Show the pages's HTML content in output")
}
