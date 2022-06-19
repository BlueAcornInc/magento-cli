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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mumoshu/variant/cmd"
)

var tempDir string

//go:embed tasks/*.yaml
var tasks embed.FS

//go:embed services/*.yaml
var services embed.FS

var yamlExt string = ".yaml"

var configDir string = ".magento-cli"

func main() {
	// put services files in a temp directory
	extractYamlToTemp(services)

	// feeds configuration into paraser and executes
	cmd.YAML(string(loadYaml(tasks)))

	// clean teh temp diretory
	rmTemp()
}

func extractYamlToTemp(objects embed.FS) {

	tempDir, err := ioutil.TempDir("/tmp", configDir+"-")
	if err != nil {
		log.Fatal(err)
	}

	// fetches basic configuration, this uses embed.FS and sources from ./<objects>/
	fs.WalkDir(objects, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())
		if !d.IsDir() == true {

			content, err := fs.ReadFile(objects, path)
			if err == nil {

				filename := tempDir + "/" + path

				_ = os.MkdirAll(filepath.Dir(filename), 0700)
				f, err := os.Create(tempDir + "/" + path)
				if err != nil {
					fmt.Println(err)
					return nil
				}
				_, err = f.WriteString(string(content))
				if err != nil {
					fmt.Println(err)
					f.Close()
					return nil
				}
				err = f.Close()
				if err != nil {
					fmt.Println(err)
					return nil
				}
			}

		}

		return nil

	})
}

func rmTemp() {

	defer os.RemoveAll(tempDir)
}

func loadYaml(objects embed.FS) string {

	data := ""

	// fetches basic configuration, this uses embed.FS and sources from ./<objects>/
	fs.WalkDir(objects, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())

		content, err := fs.ReadFile(objects, path)
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

	return data

}
