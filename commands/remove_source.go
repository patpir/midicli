package commands

import (
	"github.com/urfave/cli"

	"github.com/patpir/midicli/messages"
	"github.com/patpir/midicli/pipeline"
)


func RemoveSource() cli.Command {
	return cli.Command{
		Name: "source",
		Aliases: []string{"s"},
		Usage: "remove the source with the given name from the pipeline",
		ArgsUsage: "<name>",
		Action: removeSource,
	}
}


func removeSource(c *cli.Context) error {
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

	i, existingSource := p.FindSourceByName(c.Args()[0])
	if existingSource == nil {
		return cli.NewExitError(messages.PipelineSourceNotFound, 1)
	}

	p.RemoveSourceAt(i)
	err = p.WriteToFile(filepath)

	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}

