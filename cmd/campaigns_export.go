package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// campaignsExportCmd represents the export command
var campaignsExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export phishing campaigns",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		re, _ := cmd.Flags().GetString("regex")

		var err error
		var campaigns []*models.Campaign
		if id > 0 {
			campaign, err := client.GetCampaignByID(id)
			checkError(err)

			if campaign != nil {
				campaigns = append(campaigns, campaign)
			}
		} else if name != "" {
			campaign, err := client.GetCampaignByName(name)
			checkError(err)

			if campaign != nil {
				campaigns = append(campaigns, campaign)
			}
		} else if re != "" {
			campaigns, err = client.GetCampaignsByRegex(re)
			checkError(err)
		} else {
			campaigns, err = client.GetCampaigns()
			checkError(err)
		}

		if len(campaigns) == 0 {
			fmt.Println("[!] No campaigns found")
			return
		}

		dir, _ := cmd.Flags().GetString("dir")
		contentFiles, _ := cmd.Flags().GetBool("content-files")
		full, _ := cmd.Flags().GetBool("full")

		fmt.Printf("[+] Exporting %d campaigns...\n", len(campaigns))

		for _, c := range campaigns {
			campaignName := sanitize(c.Name)

			if !full {
				c.Page = &models.Page{
					Name: c.Page.Name,
				}

				c.Results = nil
				c.Timeline = nil

				c.SMTP = &models.SendingProfile{
					Name: c.SMTP.Name,
				}

				c.Template = &models.Template{
					Name: c.Template.Name,
				}
			} else if contentFiles {
				templateName := sanitize(c.Template.Name)
				writeTemplateContentFiles(c.Template, fmt.Sprintf("%s - %s", campaignName, templateName), dir)

				pageName := sanitize(c.Page.Name)
				writePageContentFile(c.Page, fmt.Sprintf("%s - %s", campaignName, pageName), dir)
			}

			data, err := yaml.Marshal(c)
			checkError(err)

			err = os.WriteFile(filepath.Join(dir, fmt.Sprintf("%s.yaml", campaignName)), data, 0644)
			checkError(err)
		}
	},
}

func init() {
	campaignsCmd.AddCommand(campaignsExportCmd)

	campaignsExportCmd.Flags().StringP("name", "n", "", "Export the campaign by name.")
	campaignsExportCmd.Flags().Int("id", 0, "Export the campaign by ID")
	campaignsExportCmd.Flags().StringP("regex", "r", "", "Export the campaigns with the name matching the regex.")
	campaignsExportCmd.Flags().Bool("full", false, "Run a full export. This includes all data for each campaign's template, landing page, sending profile, and results.")
	campaignsExportCmd.Flags().Bool("content-files", false, "Create separate template and landing page content files for each campaign. Only applies when doing a full export.")
	campaignsExportCmd.Flags().StringP("dir", "d", "", "Directory to export the campaigns. Defaults to the current directory.")
}
