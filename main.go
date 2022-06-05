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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mumoshu/variant/cmd"
)

func main() {

	data := ""

	files, err := ioutil.ReadDir("../tasks")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		if filepath.Ext(f.Name()) == ".yaml" {
			content, _ := ioutil.ReadFile("../tasks/" + f.Name())

			data = data + string(content) + "\n"
		}

	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	text := string(data)
	cmd.YAML(text)
}
