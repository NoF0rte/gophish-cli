package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addTemplateCmd represents the add command
var addTemplateCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new e-mail template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	templatesCmd.AddCommand(addTemplateCmd)
}
