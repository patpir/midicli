package commands

import (
	"github.com/urfave/cli"

	"github.com/patpir/midicli/messages"
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
	initialized, err := pipeline.IsInitialized(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if !initialized {
		return cli.NewExitError(messages.PipelineNotInitialized, 1)
	}

	transformationBlock, err := blockArguments(c.Args()).toBlock()
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	p, err := pipeline.ReadFromFile(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	_, existingTransformation := p.FindTransformationByName(transformationBlock.Name)
	if existingTransformation != nil {
		return cli.NewExitError(messages.PipelineInsertDuplicateTransformation, 1)
	}

	p.AddTransformation(transformationBlock)
	p.WriteToFile(filepath)

	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}

