package messages

import (
	"fmt"
	"io"
)

const PipelineNotInitialized = "Could not find pipeline file. Use \"midicli init\" to initialize a new pipeline."
const PipelineAlreadyInitialized = "Pipeline is already initialized. Use 'midicli reset' to return to a fresh state."
const PipelineInsertDuplicateSource = "The pipeline already contains a source with this name. Please choose another name!"
const PipelineInsertDuplicateTransformation = "The pipeline already contains a transformation with this name. Please choose another name!"
const PipelineInsertDuplicateVisualization = "The pipeline already contains a visualization with this name. Please choose another name!"
const PipelineSourceNotFound = "Could not find a source with this name in the pipeline!"
const PipelineTransformationNotFound = "Could not find a transformation with this name in the pipeline!"
const PipelineVisualizationNotFound = "Could not find a visualization with this name in the pipeline!"
const PipelineWithoutSources = "Cannot run a pipeline without sources!"
const PipelineWithoutVisualizations = "Cannot run a pipeline without visualizations!"
const PipelineRunError = "Error while running pipeline:"

func PrintPipelineResetSuccessful(filepath string) {
	fmt.Printf("Reset pipeline at \"%s\".\n", filepath)
}

func PrintPipelineRunErrorHeader(out io.Writer) {
	fmt.Fprintln(out, PipelineRunError)
}

func PrintPipelineRunBlockSuccess(out io.Writer, category string, name string) {
	fmt.Fprintf(out, "[ OK  ] %s \"%s\"\n", category, name)
}

func PrintPipelineRunBlockError(out io.Writer, category string, name string, err error) {
	fmt.Fprintf(out, "[ERROR] %s \"%s\": %s\n", category, name, err.Error())
}

func PrintPipelineResultWriteError(out io.Writer, filepath string, err error) {
	fmt.Fprintf(out, "ERROR: Could not open \"%s\" to write result of pipeline: %s\n", filepath, err.Error())
}

func PrintPipelineResultLocation(out io.Writer, trace string, filepath string) {
	fmt.Fprintf(out, "Wrote result of %s to %s\n", trace, filepath)
}

