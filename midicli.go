package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	// register sources, transformations and visualizations from miditf
	_ "github.com/patpir/miditf/blocks"
	_ "github.com/patpir/miditf/sources"
	_ "github.com/patpir/miditf/transform"
	_ "github.com/patpir/miditf/visualize"

	"github.com/patpir/midicli/commands"
)


func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		commands.InitCmd(),
		commands.ResetCmd(),
		commands.ListCmd(),
		commands.AddCmd(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

