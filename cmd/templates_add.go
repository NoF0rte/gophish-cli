package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

// addTemplateCmd represents the add command
var addTemplateCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new e-mail template(s)",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		templatePaths, _ := cmd.Flags().GetStringSlice("templates")

		if dir != "" {
			files, err := filepath.Glob(filepath.Join(dir, "*.yaml"))
			checkError(err)

			templatePaths = append(templatePaths, files...)

			files, err = filepath.Glob(filepath.Join(dir, "*.yml"))
			checkError(err)

			templatePaths = append(templatePaths, files...)
		}

		for _, t := range templatePaths {
			fmt.Printf("[+] Adding template %s\n", t)

			_, err := client.CreateTemplateFromFile(t, variables)
			checkError(err)
		}
	},
}

func init() {
	templatesCmd.AddCommand(addTemplateCmd)

	addTemplateCmd.Flags().StringP("dir", "d", "", "Directory containing the templates to add. Both .yaml and .yml files will be used.")
	addTemplateCmd.Flags().StringSliceP("templates", "t", []string{}, "The paths of the templates to add. Specify multiple times to add more templates.")
}
