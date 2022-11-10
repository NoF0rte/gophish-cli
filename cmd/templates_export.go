package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/NoF0rte/gophish-cli/pkg/api/models"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// templatesExportCmd represents the export command
var templatesExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export email templates",
	Run: func(cmd *cobra.Command, args []string) {
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

		dir, _ := cmd.Flags().GetString("dir")
		contentFiles, _ := cmd.Flags().GetBool("content-files")

		fmt.Printf("[+] Exporting %d templates to %s\n", len(templates), dir)

		replaceRe := regexp.MustCompile(`[ /]`)
		for _, t := range templates {
			name := strings.ToLower(replaceRe.ReplaceAllString(t.Name, "-"))
			name = filepath.Clean(name)

			if contentFiles {
				if t.Text != "" {
					t.TextFile = fmt.Sprintf("%s.txt", name)

					err := os.WriteFile(filepath.Join(dir, t.TextFile), []byte(t.Text), 0644)
					checkError(err)

					t.Text = ""
				}

				if t.HTML != "" {
					t.HTMLFile = fmt.Sprintf("%s.html", name)

					err := os.WriteFile(filepath.Join(dir, t.HTMLFile), []byte(t.HTML), 0644)
					checkError(err)

					t.HTML = ""
				}
			}

			data, err := yaml.Marshal(t)
			checkError(err)

			err = os.WriteFile(filepath.Join(dir, fmt.Sprintf("%s.yaml", name)), data, 0644)
			checkError(err)
		}
	},
}

func init() {
	templatesCmd.AddCommand(templatesExportCmd)

	templatesExportCmd.Flags().StringP("name", "n", "", "Export the template by name.")
	templatesExportCmd.Flags().Int("id", 0, "Export the template by ID")
	templatesExportCmd.Flags().StringP("regex", "r", "", "Export the templates with the name matching the regex.")
	templatesExportCmd.Flags().Bool("content-files", false, "Create separate template content files for each template.")
	templatesExportCmd.Flags().StringP("dir", "d", "", "Directory to export the templates.")
	templatesExportCmd.MarkFlagRequired("dir")
}