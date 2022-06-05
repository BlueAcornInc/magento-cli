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
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mumoshu/variant/cmd"
)

// content holds our static web server content.
//go:embed tasks/*.yaml
var tasks embed.FS

func main() {

	data := ""

	fs.WalkDir(tasks, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())

		content, err := fs.ReadFile(tasks, path)
		if err != nil {
			// return err // or panic or ignore
		}

		data = data + string(content) + "\n"

		return nil

	})

	// files, err := ioutil.ReadDir("../tasks")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, f := range files {

	// 	if filepath.Ext(f.Name()) == ".yaml" {
	// 		content, _ := ioutil.ReadFile("../tasks/" + f.Name())

	// 		data = data + string(content) + "\n"
	// 	}

	// }

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	text := string(data)
	cmd.YAML(text)
}
