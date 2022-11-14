package cmd

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to GoPhish and attempt to retrieve the user's API key",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

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
			panic(err)
		}

		fmt.Println(apiKey)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "U", "admin", "The user's username")
	loginCmd.Flags().StringP("password", "p", "", "The user's password. Use '-' to be prompted for the password.")
	loginCmd.MarkFlagRequired("password")
}
