package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/spf13/cobra"
)

// groupsAddCmd represents the add command
var groupsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new groups",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		groupPaths, _ := cmd.Flags().GetStringSlice("groups")

		if dir != "" {
			files, err := filepath.Glob(filepath.Join(dir, "*.yaml"))
			checkError(err)

			groupPaths = append(groupPaths, files...)

			files, err = filepath.Glob(filepath.Join(dir, "*.yml"))
			checkError(err)

			groupPaths = append(groupPaths, files...)
		}

		for _, g := range groupPaths {
			group, err := models.GroupFromFile(g, variables)
			checkError(err)

			fmt.Printf("[+] Adding group \"%s\"\n", group.Name)

			_, err = client.CreateGroup(group)
			checkError(err)
		}
	},
}

func init() {
	groupsCmd.AddCommand(groupsAddCmd)

	groupsAddCmd.Flags().StringP("dir", "d", "", "Directory containing the groups to add. Both .yaml and .yml files will be used.")
	groupsAddCmd.Flags().StringSliceP("groups", "g", []string{}, "The paths of the groups to add. Specify multiple times to add more groups.")
}
