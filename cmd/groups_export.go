package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// groupsExportCmd represents the export command
var groupsExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export groups",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		re, _ := cmd.Flags().GetString("regex")
		asYAML, _ := cmd.Flags().GetBool("yaml")
		asCSV, _ := cmd.Flags().GetBool("csv")

		var err error
		var groups []*models.Group
		if id > 0 {
			group, err := client.GetGroupByID(id)
			checkError(err)

			if group != nil {
				groups = append(groups, group)
			}
		} else if name != "" {
			group, err := client.GetGroupByName(name)
			checkError(err)

			if group != nil {
				groups = append(groups, group)
			}
		} else if re != "" {
			groups, err = client.GetGroupsByRegex(re)
			checkError(err)
		} else {
			groups, err = client.GetGroups()
			checkError(err)
		}

		if len(groups) == 0 {
			fmt.Println("[!] No groups found")
			return
		}

		dir, _ := cmd.Flags().GetString("dir")

		fmt.Printf("[+] Exporting %d groups...\n", len(groups))

		for _, g := range groups {
			name := sanitize(g.Name)

			if asYAML || asCSV {
				var data []byte

				if asYAML {
					g.TargetsFile = fmt.Sprintf("%s-targets.yaml", name)
					data, err = yaml.Marshal(&g.Targets)
				} else {
					g.TargetsFile = fmt.Sprintf("%s-targets.csv", name)
					data, err = gocsv.MarshalBytes(&g.Targets)
				}

				checkError(err)

				err = os.WriteFile(filepath.Join(dir, g.TargetsFile), data, 0644)
				checkError(err)

				g.Targets = nil
			}

			data, err := yaml.Marshal(g)
			checkError(err)

			err = os.WriteFile(filepath.Join(dir, fmt.Sprintf("%s.yaml", name)), data, 0644)
			checkError(err)
		}
	},
}

func init() {
	groupsCmd.AddCommand(groupsExportCmd)

	groupsExportCmd.Flags().StringP("name", "n", "", "Export the group by name.")
	groupsExportCmd.Flags().Int("id", 0, "Export the group by ID")
	groupsExportCmd.Flags().StringP("regex", "r", "", "Export the groups with the name matching the regex.")
	groupsExportCmd.Flags().StringP("dir", "d", "", "Directory to export the groups. Defaults to the current directory.")
	groupsExportCmd.Flags().Bool("yaml", false, "Export targets file as a separate yaml file")
	groupsExportCmd.Flags().Bool("csv", false, "Export targets file as a separate csv file")
}
