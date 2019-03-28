package commands

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"

	"github.com/patpir/midicli/pipeline"
)


func Init() cli.Command {
	return cli.Command {
		Name: "init",
		Usage: "initializes a new pipeline",
		Action: performInit,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "pipeline-file, f",
				Value: "MidiPipeline.json",
				Usage: "initialize pipeline in `FILE`",
			},
		},
	}
}


func performInit(c *cli.Context) error {
	filepath := c.String("pipeline-file")
	initialized, err := pipeline.IsInitialized(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if initialized {
		return cli.NewExitError(errors.New("Pipeline is already initialized. Use 'midicli reset' to return to a fresh state."), 1)
	}

	p := pipeline.New()
	err = p.WriteToFile(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Printf("Initialized pipeline at \"%s\".\n", filepath)
	return nil
}

