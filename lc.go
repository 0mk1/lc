package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func parseArgs(c *cli.Context) {
}

func countLines() {
}

func countLinesInFile() {
}

func main() {
	app := cli.NewApp()
	app.Name = "code line counter"
	app.Usage = "lc <folder/file>"

	app.Action = func(c *cli.Context) error {
		// TODO: Add https://github.com/derekparker/delve to debug
		fmt.Printf("Hello %q", c.Args().Get(0))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
