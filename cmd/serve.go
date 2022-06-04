// Cmd serve brings up a web server, and its docker friends too
package cmd

import (
	"github.com/spf13/cobra"
)

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

}

// runServe is the main flow for the serve command.
func runServe(cmd *cobra.Command, args []string) error {

	return MainServeFlow(args)
}

// MainServeFlow will run the local php server
func MainServeFlow(args []string) error {

	return nil
}
