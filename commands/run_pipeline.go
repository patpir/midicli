package commands

import (
	"github.com/urfave/cli"

	"github.com/patpir/miditf/blocks"

	"github.com/patpir/midicli/pipeline"
)


func RunPipeline() cli.Command {
	return cli.Command{
		Name: "run-pipeline",
		Aliases: []string{"run"},
		Usage: "execute a pipeline consisting of source(s), transformation(s) and visualization(s)",
		Action: runPipeline,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "pipeline-file, f",
				Value: "MidiPipeline.json",
				Usage: "execute pipeline defined in `FILE`",
			},
		},
	}
}


func runPipeline(c *cli.Context) error {
	filepath := c.String("pipeline-file")

	pipelineDefinition, err := pipeline.ReadFromFile(filepath)
	if err != nil {
		return err
	}

	ch := make(chan blocks.PipelineResult, 1)

	pipeline := pipelineDefinition.ToPerformablePipeline()
	go func() {
		pipeline.Perform(ch)
		close(ch)
	}()

	for result := range ch {
		_ = result
	}

	return nil
}


