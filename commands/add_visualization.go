package commands

import (
	"github.com/urfave/cli"

	"github.com/patpir/midicli/messages"
	"github.com/patpir/midicli/pipeline"
)


func AddVisualization() cli.Command {
	return cli.Command{
		Name: "visualization",
		Aliases: []string{"v"},
		Usage: "add a visualization to the pipeline",
		Action: addVisualization,
	}
}


func addVisualization(c *cli.Context) error {
	filepath := c.Parent().String("pipeline-file")
	initialized, err := pipeline.IsInitialized(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if !initialized {
		return cli.NewExitError(messages.PipelineNotInitialized, 1)
	}

	visualizationBlock, err := blockArguments(c.Args()).toBlock()
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	p, err := pipeline.ReadFromFile(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	p.AddVisualization(visualizationBlock)
	p.WriteToFile(filepath)

	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}

