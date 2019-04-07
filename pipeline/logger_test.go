package pipeline

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/patpir/miditf/blocks"
)


func checkErrorTrace(t *testing.T, result *blocks.PipelineResult) {
	var buf bytes.Buffer
	log := newLogger(&buf)

	log.printErrorTrace(result)
	logContent := buf.String()
	logLines := strings.Split(logContent, "\n")
	logLines = logLines[1:len(logLines)-1]
	categories := []string{
		"Source",
	}
	for i := 0; i < len(result.Transformations); i++ {
		categories = append(categories, "Transformation")
	}
	categories = append(categories, "Visualization")

	for i, logLine := range logLines[:len(logLines)-1] {
		assert.Contains(t, logLine, "OK")
		assert.NotContains(t, logLine, "ERROR")
		assert.Contains(t, logLine, fmt.Sprintf("test name %d", i))
		assert.NotContains(t, logLine, fmt.Sprintf("test error %d", i))
		assert.Contains(t, logLine, categories[i])
	}

	i := len(logLines)-1
	logLine := logLines[i]
	assert.Contains(t, logLine, "ERROR")
	assert.NotContains(t, logLine, "OK")
	assert.Contains(t, logLine, fmt.Sprintf("test error %d", i))
	assert.Contains(t, logLine, fmt.Sprintf("test name %d", i))
	assert.Contains(t, logLine, categories[i])
}


func TestPrintErrorTraceSource(t *testing.T) {
	result := &blocks.PipelineResult{
		Source: blocks.NewBlock("test type 0", "test name 0", nil),
		Transformations: []blocks.Block{},
		Visualization: nil,
		Err: errors.New("test error 0"),
		Output: "",
	}

	checkErrorTrace(t, result)
}

func TestPrintErrorTraceTransformation(t *testing.T) {
	result := &blocks.PipelineResult{
		Source: blocks.NewBlock("test type 0", "test name 0", nil),
		Transformations: []blocks.Block{
			blocks.NewBlock("test type 1", "test name 1", nil),
		},
		Visualization: nil,
		Err: errors.New("test error 1"),
		Output: "",
	}

	checkErrorTrace(t, result)
}

func TestPrintErrorTraceVisualization(t *testing.T) {
	result := &blocks.PipelineResult{
		Source: blocks.NewBlock("test type 0", "test name 0", nil),
		Transformations: []blocks.Block{
			blocks.NewBlock("test type 1", "test name 1", nil),
		},
		Visualization: blocks.NewBlock("test type 2", "test name 2", nil),
		Err: errors.New("test error 2"),
		Output: "",
	}

	checkErrorTrace(t, result)
}

func TestPrintWriteError(t *testing.T) {
	var buf bytes.Buffer
	log := newLogger(&buf)
	log.printWriteError("/path/to/file", errors.New("error message"))

	logContent := buf.String()
	assert.Contains(t, logContent, "/path/to/file")
	assert.Contains(t, logContent, "error message")
}

func TestPrintSuccessTraceWithoutTransformation(t *testing.T) {
	var buf bytes.Buffer
	log := newLogger(&buf)

	result := &blocks.PipelineResult{
		Source: blocks.NewBlock("test type 0", "test source", nil),
		Transformations: []blocks.Block{},
		Visualization: blocks.NewBlock("test type 2", "test visualization", nil),
		Err: nil,
		Output: "",
	}

	log.printSuccessTrace(result, "/path/to/file")
	logContent := buf.String()
	assert.Contains(t, logContent, "/path/to/file")
	assert.Contains(t, logContent, "test source")
	assert.NotContains(t, logContent, "transformation")
	assert.Contains(t, logContent, "test visualization")
}

func TestPrintSuccessTraceWithOneTransformation(t *testing.T) {
	var buf bytes.Buffer
	log := newLogger(&buf)

	result := &blocks.PipelineResult{
		Source: blocks.NewBlock("test type 0", "test source", nil),
		Transformations: []blocks.Block{
			blocks.NewBlock("test type 1", "test transformation", nil),
		},
		Visualization: blocks.NewBlock("test type 2", "test visualization", nil),
		Err: nil,
		Output: "",
	}

	log.printSuccessTrace(result, "/path/to/file")
	logContent := buf.String()
	assert.Contains(t, logContent, "/path/to/file")
	assert.Contains(t, logContent, "test source")
	assert.Contains(t, logContent, "test transformation")
	assert.Contains(t, logContent, "test visualization")
}

func TestPrintSuccessTraceWithMultipleTransformations(t *testing.T) {
	var buf bytes.Buffer
	log := newLogger(&buf)

	result := &blocks.PipelineResult{
		Source: blocks.NewBlock("test type 0", "test source", nil),
		Transformations: []blocks.Block{
			blocks.NewBlock("test type 1", "test transformation 1", nil),
			blocks.NewBlock("test type 1", "test transformation 2", nil),
		},
		Visualization: blocks.NewBlock("test type 2", "test visualization", nil),
		Err: nil,
		Output: "",
	}

	log.printSuccessTrace(result, "/path/to/file")
	logContent := buf.String()
	assert.Contains(t, logContent, "/path/to/file")
	assert.Contains(t, logContent, "test source")
	assert.NotContains(t, logContent, "test transformation")
	assert.Contains(t, logContent, "transformations")
	assert.Contains(t, logContent, "test visualization")
}

