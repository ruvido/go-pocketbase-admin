package admin

import (
	"log"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:  "pbadmin",
	Version: version,
	Short: "pbadmin - a simple CLI to straightforward pocketbase aministration",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.[toml,yaml,json])")
	viper.SetDefault("author", "ruvido <ruvido@gmail.com>")
	viper.SetDefault("license", "MIT")
}

func initConfig() {

	extList    := []string{"toml", "yaml", "json"}
	configName := "config"
	defaultDir := os.Getenv("HOME")+"/.config/pbadmin/"

	switch {
	case cfgFile != "":
		// Use config file from the flag.
		log.Println(cfgFile)
		viper.SetConfigFile(cfgFile)

	case fileExists(".", configName, extList):
		log.Println(configName)
		viper.AddConfigPath(".")
		viper.SetConfigName(configName)

	case fileExists(defaultDir, configName, extList):
		log.Println(defaultDir+configName)
		viper.AddConfigPath(defaultDir)
		viper.SetConfigName(configName)

	default:
		log.Fatal("Config file not found")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}


	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else if {
	// 	// Search config in local folder with name ".cobra" (without extension).
	// 	viper.AddConfigPath(".")
	// 	viper.SetConfigName("config")
	// } else {
	// 	homeDir, err := os.UserHomeDir()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func fileExists(dir, name string, extlist []string ) bool {
	for _, ext := range extlist {
		fileInfo, err := os.Stat(filepath.Join(dir, name+"."+ext))
		if err == nil && !fileInfo.IsDir() {
			return true
		}
	}
	return false
}


