package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/NoF0rte/gophish-cli/pkg/api"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var client *api.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gophish-cli",
	Short: "A CLI to interact with the Gophish API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client = api.NewClient(viper.GetString("url"), viper.GetString("api-key"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		template, err := client.GetTemplateByID(1)
		if err != nil {
			panic(err)
		}

		fmt.Println(template)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("url", "u", "https://localhost:3333", "The URL to the Gophish server")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))

	rootCmd.PersistentFlags().StringP("api-key", "T", "", "A valid Gophish API key")
	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name "gophish-cli" (without extension).
		viper.AddConfigPath(cwd)
		viper.AddConfigPath(home)
		viper.SetConfigName("gophish-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.ReadInConfig()
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("[!] Error: %v\n", err)
		os.Exit(1)
	}
}
