// Cmd serve brings up a web server, and its docker friends too
package cmd

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Command = "php -S " + BaseUrl

// BuildImageDirname tells the tool which directory to itereate through to find Dockerfiles. defaults the present working
// directory, but a good practice is to mint a .mach.yaml and set this to `images` or the like when building an IaC repo.
var BaseUrl string = "127.0.0.1:8200"

var serveCmd = CreateServeCmd()

func CreateServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serves magento",
		Long:  `serves magento either using built-in php sever and/or docker`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServe(cmd, args)
		},
	}
	return cmd
}

func init() {

	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&BaseUrl, "base-url", BaseUrl, "base url for php server")
	viper.SetDefault("BaseUrl", BaseUrl)
	viper.BindPFlag("base-url", serveCmd.Flags().Lookup("base-url"))

}

// runServe is the main flow for the serve command.
func runServe(cmd *cobra.Command, args []string) error {

	BaseUrl = viper.GetString("base-url")

	return MainServeFlow(args)
}

// MainServeFlow will run the local php server
func MainServeFlow(args []string) error {

	cmd := exec.Command("sh", "-c", Command)
	pipe, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		// handle error
		return nil
	}
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	fmt.Println(line)
	for err == nil {
		fmt.Println(line)
		line, err = reader.ReadString('\n')
	}
	return nil
}
