package commands

import (
	"github.com/urfave/cli"

	"github.com/patpir/midicli/messages"
	"github.com/patpir/midicli/pipeline"
)


func AddSource() cli.Command {
	return cli.Command{
		Name: "source",
		Aliases: []string{"s"},
		Usage: "add a source to the pipeline",
		ArgsUsage: "<TYPE> <NAME> <ARG=VALUE ...>",
		Action: addSource,
	}
}


func addSource(c *cli.Context) error {
	filepath := c.Parent().String("pipeline-file")
	initialized, err := pipeline.IsInitialized(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if !initialized {
		return cli.NewExitError(messages.PipelineNotInitialized, 1)
	}

	sourceBlock, err := blockArguments(c.Args()).toBlock()
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	p, err := pipeline.ReadFromFile(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	_, existingSource := p.FindSourceByName(sourceBlock.Name)
	if existingSource != nil {
		return cli.NewExitError(messages.PipelineInsertDuplicateSource, 1)
	}

	p.AddSource(sourceBlock)
	err = p.WriteToFile(filepath)

	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}

