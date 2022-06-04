/*
Magento-cli a cli application for managing local magneto instances

Magento-cli provides a wrapper for the bin/magento command in a directory,
and is intended to provide convenience commands like `serve` and `sql` leveraging
the local php web server, or docker-compose. If it doesn't match a command,
it will pass through the command to the `bin/magento` command, providing a nice
workflow around this tool.
*/
package main

import (
	"github.com/superterran/magneto-cli/cmd"
)

func main() {
	cmd.Execute()
}
