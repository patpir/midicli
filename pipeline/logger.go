package pipeline

import (
	"fmt"
	"io"

	"github.com/patpir/miditf/blocks"

	"github.com/patpir/midicli/messages"
)

type logger struct {
	out io.Writer
}

type categoryBlock struct {
	category string
	block blocks.Block
}


func newLogger(out io.Writer) *logger {
	return &logger{
		out: out,
	}
}

func (w *logger) printErrorTrace(result *blocks.PipelineResult) {
	messages.PrintPipelineRunErrorHeader(w.out)

	allBlocks := []categoryBlock{}
	if source := result.Source; source != nil {
		allBlocks = append(allBlocks, categoryBlock{ "Source", source })
	}
	for _, transformation := range result.Transformations {
		allBlocks = append(allBlocks, categoryBlock{ "Transformation", transformation })
	}
	if visualization := result.Visualization; visualization != nil {
		allBlocks = append(allBlocks, categoryBlock{ "Visualization", visualization })
	}

	okBlocks := allBlocks[:len(allBlocks)-1]
	lastBlock := allBlocks[len(allBlocks)-1]

	for _, block := range okBlocks {
		messages.PrintPipelineRunBlockSuccess(w.out, block.category, block.block.Comment())
	}
	messages.PrintPipelineRunBlockError(w.out, lastBlock.category, lastBlock.block.Comment(), result.Err)
}

func (w *logger) printWriteError(filepath string, err error) {
	messages.PrintPipelineResultWriteError(w.out, filepath, err)
}

func (w *logger) printSuccessTrace(result *blocks.PipelineResult, filename string) {
	stages := []string{
		result.Source.Comment(),
	}
	if len(result.Transformations) == 1 {
		stages = append(stages, result.Transformations[0].Comment())
	} else if len(result.Transformations) > 1 {
		stages = append(stages, fmt.Sprintf("%d transformations", len(result.Transformations)))
	}
	stages = append(stages, result.Visualization.Comment())

	trace := ""
	for _, stage := range stages {
		trace += fmt.Sprintf(" -> %s", stage)
	}
	trace = trace[4:]
	messages.PrintPipelineResultLocation(w.out, trace, filename)
}

