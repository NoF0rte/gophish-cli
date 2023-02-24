package cmd

import (
	"fmt"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
)

// templatesDeleteCmd represents the delete command
var templatesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete e-mail templates",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")

		var err error
		var resp *models.GenericResponse
		if id > 0 {
			resp, err = client.DeleteTemplateByID(int64(id))
		} else if name != "" {
			resp, err = client.DeleteTemplateByName(name)
		}

		checkError(err)

		data, err := resp.ToJson()
		checkError(err)

		fmt.Println(data)
	},
}

func init() {
	templatesCmd.AddCommand(templatesDeleteCmd)

	templatesDeleteCmd.Flags().StringP("name", "n", "", "Delete the template by name.")
	templatesDeleteCmd.Flags().Int("id", 0, "Delete the template by ID.")
}
