package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// TestMode var determines if certain flows actually complete or not for unit testing
var TestMode = false

var rootCmd = CreateRootCmd()

func CreateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "magento",
		TraverseChildren: true,
		Short:            "magento-cli tool",
		Long: `Running magento and adobe commerce locally with a cli tool.
		
		usage: magento serve`,
		RunE: func(md *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(InitConfig)

	TestMode = strings.HasSuffix(os.Args[0], ".test")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is loaded from working dir)")

}

func InitConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".mage")
	}

	viper.AddConfigPath(".")
	viper.SetEnvPrefix("MAGE")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
