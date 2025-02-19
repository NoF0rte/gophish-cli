package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/NoF0rte/gophish-client/api/models"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

// groupsAddCmd represents the add command
var groupsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new groups",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		groupPaths, _ := cmd.Flags().GetStringSlice("groups")
		name, _ := cmd.Flags().GetString("name")
		targetsFile, _ := cmd.Flags().GetString("targets-file")
		targetsArr, _ := cmd.Flags().GetStringArray("targets")

		if dir != "" {
			files, err := filepath.Glob(filepath.Join(dir, "*.yaml"))
			checkError(err)

			groupPaths = append(groupPaths, files...)

			files, err = filepath.Glob(filepath.Join(dir, "*.yml"))
			checkError(err)

			groupPaths = append(groupPaths, files...)
		}

		if len(groupPaths) != 0 {
			for _, g := range groupPaths {
				group, err := models.GroupFromFile(g, variables)
				checkError(err)

				fmt.Printf("[+] Adding group \"%s\"\n", group.Name)

				_, err = client.CreateGroup(group)
				checkError(err)
			}
		} else {
			var err error
			var targets []*models.Target
			if targetsFile != "" {
				targets, err = models.TargetsFromFile(targetsFile)
				checkError(err)
			} else if len(targetsArr) != 0 {
				reader := strings.NewReader(strings.Join(targetsArr, "\n"))
				err = gocsv.UnmarshalWithoutHeaders(reader, &targets)
				if err != nil {
					checkError(fmt.Errorf("failed to parse targets: %v", err))
				}
			}

			group := &models.Group{
				Name:    name,
				Targets: targets,
			}

			_, err = client.CreateGroup(group)
			checkError(err)
		}
	},
}

func init() {
	groupsCmd.AddCommand(groupsAddCmd)

	groupsAddCmd.Flags().StringP("dir", "d", "", "Directory containing the groups to add. Both .yaml and .yml files will be used.")
	groupsAddCmd.Flags().StringSliceP("groups", "g", []string{}, "The paths of the group files to add. Specify multiple times to add more groups.")
	groupsAddCmd.Flags().StringP("name", "n", "", "Name of the group to add")
	groupsAddCmd.Flags().String("targets-file", "", "The path to the file containing the targets. Can be CSV or YAML")
	groupsAddCmd.Flags().StringArrayP("targets", "t", []string{}, "Comma separated value representing the target. Specify multiple times for more targets. Format: FirstName,LastName,Email,Position")
}
