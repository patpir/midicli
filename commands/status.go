package commands

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"

	"github.com/patpir/midicli/messages"
	"github.com/patpir/midicli/pipeline"
)


func Status() cli.Command {
	return cli.Command{
		Name: "status",
		Usage: "display the sources, transformations and visualizations of a pipeline",
		Action: performStatus,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "pipeline-file, f",
				Value: "MidiPipeline.json",
				Usage: "status of pipeline in `FILE`",
			},
		},
	}
}


func performStatus(c *cli.Context) error {
	filepath := c.String("pipeline-file")
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

	printBlockList("sources", "midicli add source", p.SourceList())
	printBlockList("transformations", "midicli add transformation", p.TransformationList())
	printBlockList("visualizations", "midicli add visualization", p.VisualizationList())

	return nil
}

func printBlockList(category string, command string, blockList []*pipeline.Block) {
	fmt.Printf("%s:\n", strings.ToUpper(category))
	if len(blockList) == 0 {
		fmt.Printf("  No %s defined - try '%s'\n", category, command)
	} else {
		for _, block := range blockList {
			fmt.Printf("  %s: type \"%s\"\n", block.Comment(), block.TypeId())
			args := block.Arguments()
			if len(args) == 0 {
				fmt.Println("    (no arguments)")
			} else {
				for name, value := range args {
					fmt.Printf("     -  %-10s = %s\n", name, value)
				}
			}
		}
	}
	fmt.Println()
}

