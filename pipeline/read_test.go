package pipeline

import (
	"testing"
	"strings"

	"github.com/stretchr/testify/assert"
)


func TestReadEmpty(t *testing.T) {
	const emptyStream = ""
	reader := strings.NewReader(emptyStream)

	pipeline, err := Read(reader)
	assert.Nil(t, err)
	assert.NotNil(t, pipeline)
	assert.Equal(t, 0, len(pipeline.SourceList()))
	assert.Equal(t, 0, len(pipeline.TransformationList()))
	assert.Equal(t, 0, len(pipeline.VisualizationList()))
}

func TestReadBlocks(t *testing.T) {
	pipeline, err := ReadFromFile("../testdata/pipeline.json")
	assert.Nil(t, err)
	assert.NotNil(t, pipeline)
	assert.Equal(t, 2, len(pipeline.SourceList()))
	assert.Equal(t, 2, len(pipeline.TransformationList()))
	assert.Equal(t, 2, len(pipeline.VisualizationList()))
}

