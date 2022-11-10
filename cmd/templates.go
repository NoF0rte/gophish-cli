package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/NoF0rte/gophish-cli/pkg/api/models"
	"github.com/spf13/cobra"
)

// templatesCmd represents the templates command
var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Get or list templates",
	Run: func(cmd *cobra.Command, args []string) {
		showContent, _ := cmd.Flags().GetBool("show-content")
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		re, _ := cmd.Flags().GetString("regex")

		var err error
		var templates []*models.Template
		if id > 0 {
			template, err := client.GetTemplateByID(id)
			checkError(err)

			if template != nil {
				templates = append(templates, template)
			}
		} else if name != "" {
			template, err := client.GetTemplateByName(name)
			checkError(err)

			if template != nil {
				templates = append(templates, template)
			}
		} else if re != "" {
			templates, err = client.GetTemplatesByRegex(re)
			checkError(err)
		} else {
			templates, err = client.GetTemplates()
			checkError(err)
		}

		if len(templates) == 0 {
			fmt.Println("[!] No templates found")
			return
		}

		if !showContent {
			for _, t := range templates {
				t.HTML = ""
				t.Text = ""
			}
		}

		if len(templates) == 1 {
			data, err := templates[0].ToJson()
			checkError(err)

			fmt.Println(data)
			return
		}

		data, err := json.MarshalIndent(templates, "", "  ")
		checkError(err)

		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(templatesCmd)

	templatesCmd.Flags().StringP("name", "n", "", "Get the template by name.")
	templatesCmd.Flags().Int("id", 0, "Get the template by ID")
	templatesCmd.Flags().StringP("regex", "r", "", "List the templates with the name matching the regex.")
	templatesCmd.Flags().Bool("show-content", false, "Show the template content in output")
}
