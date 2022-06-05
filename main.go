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
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mumoshu/variant/cmd"
)

// content holds our static web server content.
//go:embed tasks/*.yaml
var tasks embed.FS

var yamlExt string = ".yaml"

var configDir string = ".magento-cli"

func main() {

	data := ""

	// fetches basic configuration, this uses embed.FS and sources from ./tasks/
	fs.WalkDir(tasks, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())

		content, err := fs.ReadFile(tasks, path)
		if err == nil {
			data = data + string(content) + "\n"
		}

		return nil

	})

	// override mechanism, .magento-cli directories can contain additional configs
	if _, err := os.Stat(configDir); !os.IsNotExist(err) {
		files, err := ioutil.ReadDir(configDir)
		if err != nil {
			log.Fatal(err)
		} else {
			for _, f := range files {

				if filepath.Ext(f.Name()) == yamlExt {
					content, _ := ioutil.ReadFile(configDir + f.Name())

					data = data + string(content) + "\n"
				}

			}
		}
	}

	// feeds configuration into paraser and executes
	text := string(data)
	cmd.YAML(text)
}
