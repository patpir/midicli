package commands

import (
	"github.com/urfave/cli"
)


func Remove() cli.Command {
	return cli.Command{
		Name: "remove",
		Usage: "remove a source, transformation or visualization from the pipeline",
		Subcommands: []cli.Command{
			RemoveSource(),
			RemoveTransformation(),
			RemoveVisualization(),
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "pipeline-file, f",
				Value: "MidiPipeline.json",
				Usage: "modify pipeline file `FILE`",
			},
		},
	}
}

