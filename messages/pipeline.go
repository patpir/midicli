package messages

import (
	"fmt"
)

const PipelineNotInitialized = "Could not find pipeline file. Use \"midicli init\" to initialize a new pipeline."
const PipelineAlreadyInitialized = "Pipeline is already initialized. Use 'midicli reset' to return to a fresh state."

func PrintPipelineResetSuccessful(filepath string) {
	fmt.Printf("Reset pipeline at \"%s\".\n", filepath)
}

