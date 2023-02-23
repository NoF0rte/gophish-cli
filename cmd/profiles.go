package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/NoF0rte/gophish-cli/pkg/api/models"
	"github.com/spf13/cobra"
)

// profilesCmd represents the profiles command
var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List, add, or delete sending profiles",
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
			profiles, err = client.GetSendingProfileByRegex(re)
			checkError(err)
		} else {
			profiles, err = client.GetSendingProfiles()
			checkError(err)
		}

		if len(profiles) == 0 {
			fmt.Println("[!] No profiles found")
			return
		}

		if len(profiles) == 1 {
			data, err := profiles[0].ToJSON()
			checkError(err)

			fmt.Println(data)
			return
		}

		data, err := json.MarshalIndent(profiles, "", "  ")
		checkError(err)

		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)

	profilesCmd.Flags().StringP("name", "n", "", "Get the profile by name.")
	profilesCmd.Flags().Int("id", 0, "Get the profile by ID")
	profilesCmd.Flags().StringP("regex", "r", "", "List the profiles with the name matching the regex.")
}
