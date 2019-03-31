package commands

import (
	"github.com/urfave/cli"

	"github.com/patpir/midicli/messages"
	"github.com/patpir/midicli/pipeline"
)


func RemoveTransformation() cli.Command {
	return cli.Command{
		Name: "transformation",
		Aliases: []string{"t"},
		Usage: "remove the transformation with the given name from the pipeline",
		ArgsUsage: "<name>",
		Action: removeTransformation,
	}
}


func removeTransformation(c *cli.Context) error {
	filepath := c.Parent().String("pipeline-file")
	initialized, err := pipeline.IsInitialized(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if !initialized {
		return cli.NewExitError(messages.PipelineNotInitialized, 1)
	}

	p, err := pipeline.ReadFromFile(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	i, existingTransformation := p.FindTransformationByName(c.Args()[0])
	if existingTransformation == nil {
		return cli.NewExitError(messages.PipelineTransformationNotFound, 1)
	}

	p.RemoveTransformationAt(i)
	err = p.WriteToFile(filepath)

	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}

