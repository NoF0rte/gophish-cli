package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// templatesCmd represents the templates command
var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Get or list templates",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		re, _ := cmd.Flags().GetString("regex")

		if id > 0 {
			template, err := client.GetTemplateByID(id)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(template)
			return
		}

		if name != "" {
			template, err := client.GetTemplateByName(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(template)
			return
		}

		if re != "" {
			templates, err := client.GetTemplatesByRegex(re)
			if err != nil {
				fmt.Println(err)
				return
			}

			output, _ := json.Marshal(templates)

			fmt.Println(string(output))
			return
		}

		templates, err := client.GetTemplates()
		if err != nil {
			fmt.Println(err)
			return
		}

		output, _ := json.Marshal(templates)

		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(templatesCmd)

	templatesCmd.Flags().StringP("name", "n", "", "Get the template by name.")
	templatesCmd.Flags().Int("id", 0, "Get the template by ID")
	templatesCmd.Flags().StringP("regex", "r", "", "List the templates with the name matching the regex.")
}
