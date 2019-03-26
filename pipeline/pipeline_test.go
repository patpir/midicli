package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestAddSource(t *testing.T) {
	p := new(pipeline)
	block := &Block{
		TypeIdentifier: "some-type",
		Name: "some-name",
	}

	p.AddSource(block)
	assert.Equal(t, 1, len(p.Sources))
	assert.Equal(t, "some-type", p.Sources[0].TypeIdentifier)
	assert.Equal(t, "some-name", p.Sources[0].Name)
}

func TestAddTransformation(t *testing.T) {
	p := new(pipeline)
	block := &Block{
		TypeIdentifier: "some-type",
		Name: "some-name",
	}

	p.AddTransformation(block)
	assert.Equal(t, 1, len(p.Transformations))
	assert.Equal(t, "some-type", p.Transformations[0].TypeIdentifier)
	assert.Equal(t, "some-name", p.Transformations[0].Name)
}

func TestAddVisualization(t *testing.T) {
	p := new(pipeline)
	block := &Block{
		TypeIdentifier: "some-type",
		Name: "some-name",
	}

	p.AddVisualization(block)
	assert.Equal(t, 1, len(p.Visualizations))
	assert.Equal(t, "some-type", p.Visualizations[0].TypeIdentifier)
	assert.Equal(t, "some-name", p.Visualizations[0].Name)
}

