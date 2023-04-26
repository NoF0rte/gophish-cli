package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// profilesExportCmd represents the export command
var profilesExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export sending profiles",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		re, _ := cmd.Flags().GetString("regex")

		var err error
		var profiles []*models.SendingProfile
		if id > 0 {
			profile, err := client.GetSendingProfileByID(id)
			checkError(err)

			if profile != nil {
				profiles = append(profiles, profile)
			}
		} else if name != "" {
			profile, err := client.GetSendingProfileByName(name)
			checkError(err)

			if profile != nil {
				profiles = append(profiles, profile)
			}
		} else if re != "" {
			profiles, err = client.GetSendingProfilesByRegex(re)
			checkError(err)
		} else {
			profiles, err = client.GetSendingProfiles()
			checkError(err)
		}

		if len(profiles) == 0 {
			fmt.Println("[!] No templates found")
			return
		}

		dir, _ := cmd.Flags().GetString("dir")

		fmt.Printf("[+] Exporting %d sending profiles...\n", len(profiles))

		for _, t := range profiles {
			name := sanitize(t.Name)

			data, err := yaml.Marshal(t)
			checkError(err)

			err = os.WriteFile(filepath.Join(dir, fmt.Sprintf("%s.yaml", name)), data, 0644)
			checkError(err)
		}
	},
}

func init() {
	profilesCmd.AddCommand(profilesExportCmd)

	profilesExportCmd.Flags().StringP("name", "n", "", "Export the sending profile by name.")
	profilesExportCmd.Flags().Int("id", 0, "Export the sending profile by ID")
	profilesExportCmd.Flags().StringP("regex", "r", "", "Export the sending profiles with the name matching the regex.")
	profilesExportCmd.Flags().StringP("dir", "d", "", "Directory to export the profiles. Defaults to the current directory.")
}
