package cmd

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/NoF0rte/gophish-client/api"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var client *api.Client
var variables map[string]string = make(map[string]string)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gophish-cli",
	Short: "A CLI to interact with the Gophish API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		type variable struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		}
		var variableList []variable

		err := viper.UnmarshalKey("vars", &variableList)
		checkError(err)

		for _, v := range variableList {
			variables[v.Name] = v.Value
		}

		vars, _ := cmd.Flags().GetStringToString("vars")
		for key, value := range vars {
			variables[key] = value
		}

		client = api.NewClient(viper.GetString("url"), viper.GetString("api-key"))
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

	rootCmd.PersistentFlags().StringToStringP("vars", "V", nil, "Variables to use when creating/editing items from files that have replacement variables. Use name=value syntax.")
}

func initConfig() {
	setConfigDefault("vars", []map[string]string{
		{
			"name":  "Example",
			"value": "Value",
		},
	})

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

// If no config file exists, all possible keys in the defaults
// need to be registered with viper otherwise viper will only think
// the keys explicitly set via viper.SetDefault() exist.
func setConfigDefault(key string, value interface{}) {
	valueType := reflect.TypeOf(value)
	valueValue := reflect.ValueOf(value)

	if valueType.Kind() == reflect.Map {
		iter := valueValue.MapRange()
		for iter.Next() {
			k := iter.Key().Interface()
			v := iter.Value().Interface()
			setConfigDefault(fmt.Sprintf("%s.%s", key, k), v)
		}
	} else if valueType.Kind() == reflect.Struct {
		numFields := valueType.NumField()
		for i := 0; i < numFields; i++ {
			structField := valueType.Field(i)
			fieldValue := valueValue.Field(i)

			setConfigDefault(fmt.Sprintf("%s.%s", key, structField.Name), fieldValue.Interface())
		}
	} else {
		viper.SetDefault(key, value)
	}
}
