package commands

import (
	"github.com/urfave/cli"

	"github.com/patpir/midicli/pipeline"
)


func AddSource() cli.Command {
	return cli.Command{
		Name: "source",
		Aliases: []string{"s"},
		Usage: "add a source to the pipeline",
		Action: addSource,
	}
}


func addSource(c *cli.Context) error {
	filepath := c.Parent().String("pipeline-file")

	sourceBlock, err := blockArguments(c.Args()).toBlock()
	if err != nil {
		return err
	}

	pipeline, err := pipeline.ReadFromFile(filepath)
	if err != nil {
		return err
	}

	pipeline.AddSource(sourceBlock)
	pipeline.WriteToFile(filepath)

	return err
}

