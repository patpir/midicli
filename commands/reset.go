package commands

import (
	"github.com/urfave/cli"

	"github.com/patpir/midicli/messages"
	"github.com/patpir/midicli/pipeline"
)


func Reset() cli.Command {
	return cli.Command {
		Name: "reset",
		Usage: "reset the pipeline to a clean state",
		Action: performReset,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "pipeline-file, f",
				Value: "MidiPipeline.json",
				Usage: "reset pipeline in `FILE`",
			},
		},
	}
}


func performReset(c *cli.Context) error {
	filepath := c.String("pipeline-file")
	initialized, err := pipeline.IsInitialized(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if !initialized {
		return cli.NewExitError(messages.PipelineNotInitialized, 1)
	}

	p := pipeline.New()
	err = p.WriteToFile(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	messages.PrintPipelineResetSuccessful(filepath)
	return nil
}

