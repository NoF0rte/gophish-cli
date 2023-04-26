package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// pagesExportCmd represents the export command
var pagesExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export landing pages",
	Run: func(cmd *cobra.Command, args []string) {
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

		dir, _ := cmd.Flags().GetString("dir")
		contentFiles, _ := cmd.Flags().GetBool("content-files")

		fmt.Printf("[+] Exporting %d pages...\n", len(pages))

		replaceRe := regexp.MustCompile(`[ /]`)
		for _, p := range pages {
			name := strings.ToLower(replaceRe.ReplaceAllString(p.Name, "-"))
			name = filepath.Clean(name)

			if contentFiles {
				if p.HTML != "" {
					p.HTMLFile = fmt.Sprintf("%s.html", name)

					err := os.WriteFile(filepath.Join(dir, p.HTMLFile), []byte(p.HTML), 0644)
					checkError(err)

					p.HTML = ""
				}
			}

			data, err := yaml.Marshal(p)
			checkError(err)

			err = os.WriteFile(filepath.Join(dir, fmt.Sprintf("%s.yaml", name)), data, 0644)
			checkError(err)
		}
	},
}

func init() {
	pagesCmd.AddCommand(pagesExportCmd)

	pagesExportCmd.Flags().StringP("name", "n", "", "Export the page by name.")
	pagesExportCmd.Flags().Int("id", 0, "Export the page by ID")
	pagesExportCmd.Flags().StringP("regex", "r", "", "Export the pages with the name matching the regex.")
	pagesExportCmd.Flags().Bool("content-files", false, "Create separate page content files for each template.")
	pagesExportCmd.Flags().StringP("dir", "d", "", "Directory to export the pages. Defaults to the current directory.")
}
