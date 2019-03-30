package messages

import (
	"fmt"
)

const PipelineNotInitialized = "Could not find pipeline file. Use \"midicli init\" to initialize a new pipeline."
const PipelineAlreadyInitialized = "Pipeline is already initialized. Use 'midicli reset' to return to a fresh state."
const PipelineInsertDuplicateSource = "The pipeline already contains a source with this name. Please choose another name!"
const PipelineInsertDuplicateTransformation = "The pipeline already contains a transformation with this name. Please choose another name!"
const PipelineInsertDuplicateVisualization = "The pipeline already contains a visualization with this name. Please choose another name!"

func PrintPipelineResetSuccessful(filepath string) {
	fmt.Printf("Reset pipeline at \"%s\".\n", filepath)
}
