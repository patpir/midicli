package commands

import (
	"os"

	"github.com/urfave/cli"

	"github.com/patpir/midicli/messages"
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
			cli.StringFlag{
				Name: "output-format, o",
				Value: "{source}-{visualization}.txt",
				Usage: "the pattern for the filenames to which the resulting visualizations will be written - available placeholders are `{source}` (name of the source), `{visualization}` (name of the visualization), and `{n}` (consecutive number)",
			},
			cli.BoolFlag{
				Name: "print-visualization, p",
				Usage: "print the visualization results to standard output instead of writing to file",
			},
		},
	}
}



type stdoutWriter struct { }
func (w *stdoutWriter) WriteResult(result string, filename string) error {
	os.Stdout.WriteString(result)
	return nil
}

type fileWriter struct { }
func (w *fileWriter) WriteResult(result string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(result)

	return nil
}

func runPipeline(c *cli.Context) error {
	filepath := c.String("pipeline-file")
	initialized, err := pipeline.IsInitialized(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if !initialized {
		return cli.NewExitError(messages.PipelineNotInitialized, 1)
	}

	pipelineDefinition, err := pipeline.ReadFromFile(filepath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	if len(pipelineDefinition.SourceList()) == 0 {
		return cli.NewExitError(messages.PipelineWithoutSources, 1)
	}
	if len(pipelineDefinition.VisualizationList()) == 0 {
		return cli.NewExitError(messages.PipelineWithoutVisualizations, 1)
	}

	var writer pipeline.ResultWriter
	if c.Bool("print-visualization") {
		writer = &stdoutWriter{}
	} else {
		writer = &fileWriter{}
	}

	return pipelineDefinition.Run(c.String("output-format"), writer)
}


