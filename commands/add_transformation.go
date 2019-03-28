package commands

import (
	"github.com/urfave/cli"

	"github.com/patpir/midicli/pipeline"
)


func AddTransformation() cli.Command {
	return cli.Command{
		Name: "transformation",
		Aliases: []string{"t"},
		Usage: "add a transformation to the pipeline",
		Action: addTransformation,
	}
}


func addTransformation(c *cli.Context) error {
	filepath := c.Parent().String("pipeline-file")

	transformationBlock, err := blockArguments(c.Args()).toBlock()
	if err != nil {
		return err
	}

	pipeline, err := pipeline.ReadFromFile(filepath)
	if err != nil {
		return err
	}

	pipeline.AddTransformation(transformationBlock)
	pipeline.WriteToFile(filepath)

	return err
}

