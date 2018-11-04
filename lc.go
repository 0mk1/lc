package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

var supportedFileTypes = [2]string{"py", "go"}

func getFiles(folderPath string) []string {
	fileList := make([]string, 0)
	err := filepath.Walk(folderPath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
	return fileList
}

func main() {
	app := cli.NewApp()
	app.Name = "code line counter"

	app.Action = func(c *cli.Context) error {
		path := c.Args().First()
		fi, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}

		var files []string
		switch mode := fi.Mode(); {
		case mode.IsDir():
			files = append(files, getFiles(path)...)
		case mode.IsRegular():
			files = append(files, path)
		}

		for _, f := range files {
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
