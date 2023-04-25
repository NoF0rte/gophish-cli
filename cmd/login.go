package cmd

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to GoPhish and attempt to retrieve the user's API key",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		save, _ := cmd.Flags().GetBool("save")

		if password == "-" {
			fmt.Printf("[-] Enter password: ")
			bytePassword, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				checkError(err)
			}
			fmt.Println()

			password = string(bytePassword)
		}

		apiKey, err := client.GetAPIKey(username, password)
		if err != nil {
			checkError(err)
		}

		fmt.Println(apiKey)
		if save {
			viper.Set("api-key", apiKey)
			saveConfig(false)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "U", "admin", "The user's username")
	loginCmd.Flags().StringP("password", "p", "", "The user's password. Use '-' to be prompted for the password.")
	loginCmd.MarkFlagRequired("password")
	loginCmd.Flags().BoolP("save", "s", false, "Save the API key to the config")
}
