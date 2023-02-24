package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
)

// campaignsCmd represents the campaigns command
var campaignsCmd = &cobra.Command{
	Use:   "campaigns",
	Short: "List, add, or delete campaigns",
	Run: func(cmd *cobra.Command, args []string) {
		showContent, _ := cmd.Flags().GetBool("show-content")
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		re, _ := cmd.Flags().GetString("regex")

		var err error
		var campaigns []*models.Campaign
		if id > 0 {
			fmt.Println("Not implemented")
			return
			// template, err := client.GetTemplateByID(id)
			// checkError(err)

			// if template != nil {
			// 	campaigns = append(campaigns, template)
			// }
		} else if name != "" {
			// template, err := client.GetTemplateByName(name)
			// checkError(err)

			// if template != nil {
			// 	campaigns = append(campaigns, template)
			// }
			fmt.Println("Not implemented")
			return
		} else if re != "" {
			// campaigns, err = client.GetTemplatesByRegex(re)
			// checkError(err)
			fmt.Println("Not implemented")
			return
		} else {
			campaigns, err = client.GetCampaigns()
			checkError(err)
		}

		if len(campaigns) == 0 {
			fmt.Println("[!] No campaigns found")
			return
		}

		if !showContent {
			for _, c := range campaigns {
				c.Template.HTML = ""
				c.Template.Text = ""
			}
		}

		if len(campaigns) == 1 {
			data, err := campaigns[0].ToJSON()
			checkError(err)

			fmt.Println(data)
			return
		}

		data, err := json.MarshalIndent(campaigns, "", "  ")
		checkError(err)

		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(campaignsCmd)

	campaignsCmd.Flags().StringP("name", "n", "", "Get the campaign by name.")
	campaignsCmd.Flags().Int("id", 0, "Get the campaign by ID")
	campaignsCmd.Flags().StringP("regex", "r", "", "List the campaigns with the name matching the regex.")
	campaignsCmd.Flags().Bool("show-content", false, "Show the campaign's template content in output")
}
