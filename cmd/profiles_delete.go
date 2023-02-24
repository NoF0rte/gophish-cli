package cmd

import (
	"fmt"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
)

// profilesDeleteCmd represents the delete command
var profilesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete sending profiles",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")

		var err error
		var resp *models.GenericResponse
		if id > 0 {
			resp, err = client.DeleteSendingProfileByID(int64(id))
		} else if name != "" {
			resp, err = client.DeleteSendingProfileByName(name)
		}

		checkError(err)

		data, err := resp.ToJson()
		checkError(err)

		fmt.Println(data)
	},
}

func init() {
	profilesCmd.AddCommand(profilesDeleteCmd)

	profilesDeleteCmd.Flags().StringP("name", "n", "", "Delete the template by name.")
	profilesDeleteCmd.Flags().Int("id", 0, "Delete the template by ID.")
}
