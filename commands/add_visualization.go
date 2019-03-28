package commands

import (
	"github.com/urfave/cli"

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

	visualizationBlock, err := blockArguments(c.Args()).toBlock()
	if err != nil {
		return err
	}

	pipeline, err := pipeline.ReadFromFile(filepath)
	if err != nil {
		return err
	}

	pipeline.AddVisualization(visualizationBlock)
	pipeline.WriteToFile(filepath)

	return err
}

