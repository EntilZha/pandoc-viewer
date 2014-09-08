package main

import (
	"fmt"
	"log"
	"os"

	"github.com/EntilZha/pandoc-viewer/pandoc"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Pandoc Viewer"
	app.Usage = "Auto compile pandoc files, focus Preview to update, refocus MacVim for continued editing"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get the current working directory")
	}
	app.Commands = []cli.Command{
		{
			Name:      "background",
			ShortName: "b",
			Usage:     "Start service which listens on input folder or current working directory, detecting changes to *.md and recompiling them to *.pdf",
			Action: func(c *cli.Context) {
				fmt.Println("Running pandoc listener...")
				pandoc.RunPandocListener(wd + "/" + c.Args().First())
			},
		},
		{
			Name:      "compile",
			ShortName: "c",
			Usage:     "Compile given file (without .md extension) to pdf",
			Action: func(c *cli.Context) {
				pandoc.CompileAndRefresh(wd + "/" + c.Args().First())
			},
		},
	}
	app.Run(os.Args)
}
