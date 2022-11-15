package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new sending profiles",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		profilePaths, _ := cmd.Flags().GetStringSlice("profiles")

		if dir != "" {
			files, err := filepath.Glob(filepath.Join(dir, "*.yaml"))
			checkError(err)

			profilePaths = append(profilePaths, files...)

			files, err = filepath.Glob(filepath.Join(dir, "*.yml"))
			checkError(err)

			profilePaths = append(profilePaths, files...)
		}

		for _, p := range profilePaths {
			fmt.Printf("[+] Adding profile %s\n", p)

			_, err := client.CreateSendingProfileFromFile(p, variables)
			checkError(err)
		}
	},
}

func init() {
	profilesCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("dir", "d", "", "Directory containing the sending profiles to add. Both .yaml and .yml files will be used.")
	addCmd.Flags().StringSliceP("profiles", "p", []string{}, "The paths of the profiles to add. Specify multiple times to add more profiles.")
}
