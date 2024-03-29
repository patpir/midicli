package commands

import (
	"github.com/urfave/cli"
)


func Add() cli.Command {
	return cli.Command{
		Name: "add",
		Usage: "add a source, transformation or visualization to the pipeline",
		Subcommands: []cli.Command{
			AddSource(),
			AddTransformation(),
			AddVisualization(),
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

