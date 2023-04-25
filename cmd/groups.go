package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List, add, or delete groups",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		re, _ := cmd.Flags().GetString("regex")
		asSummary, _ := cmd.Flags().GetBool("summary")

		var err error
		var groups []*models.Group
		if id > 0 {
			var group *models.Group
			if asSummary {
				group, err = client.GetGroupSummaryByID(id)
			} else {
				group, err = client.GetGroupByID(id)
			}
			checkError(err)

			if group != nil {
				groups = append(groups, group)
			}
		} else if name != "" {
			var group *models.Group
			if asSummary {
				group, err = client.GetGroupSummaryByName(name)
			} else {
				group, err = client.GetGroupByName(name)

			}
			checkError(err)

			if group != nil {
				groups = append(groups, group)
			}
		} else if re != "" {
			if asSummary {
				groups, err = client.GetGroupsSummaryByRegex(re)
			} else {
				groups, err = client.GetGroupsByRegex(re)
			}
			checkError(err)
		} else {
			if asSummary {
				groups, err = client.GetGroupsSummary()
			} else {
				groups, err = client.GetGroups()
			}
			checkError(err)
		}

		if len(groups) == 0 {
			fmt.Println("[!] No groups found")
			return
		}

		if len(groups) == 1 {
			data, err := groups[0].ToJSON()
			checkError(err)

			fmt.Println(data)
			return
		}

		data, err := json.MarshalIndent(groups, "", "  ")
		checkError(err)

		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(groupsCmd)

	groupsCmd.Flags().StringP("name", "n", "", "Get the group by name.")
	groupsCmd.Flags().Int("id", 0, "Get the group by ID")
	groupsCmd.Flags().StringP("regex", "r", "", "List the groups with the name matching the regex.")
	groupsCmd.Flags().Bool("summary", false, "Display group(s) as summaries")
}
