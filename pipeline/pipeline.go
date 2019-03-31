package pipeline

import (
	"io"

	"github.com/patpir/miditf/blocks"
)

// naming is unusual, but clashes with members of
// `pipeline` struct need to be avoided
type Pipeline interface {
	AddSource(source *Block)
	AddTransformation(transformation *Block)
	AddVisualization(visualization *Block)

	FindSourceByName(name string) (int, *Block)
	FindTransformationByName(name string) (int, *Block)
	FindVisualizationByName(name string) (int, *Block)

	RemoveSourceAt(index int)
	RemoveTransformationAt(index int)
	RemoveVisualizationAt(index int)

	Write(writer io.Writer) error
	WriteToFile(filepath string) error

	ToPerformablePipeline() *blocks.Pipeline

	SourceList() []*Block
	TransformationList() []*Block
	VisualizationList() []*Block
}


// `pipeline` struct is not exported itself
// however, its members are exported for JSON encoding/decoding
type pipeline struct {
	Sources          []*Block
	Transformations  []*Block
	Visualizations   []*Block
}


func New() Pipeline {
	return &pipeline{
		Sources: []*Block{},
		Transformations: []*Block{},
		Visualizations: []*Block{},
	}
}

// Implement `Pipeline` interface for `pipeline` struct

func (p *pipeline) AddSource(source *Block) {
	p.Sources = append(p.Sources, source)
}

func (p *pipeline) AddTransformation(transformation *Block) {
	p.Transformations = append(p.Transformations, transformation)
}

func (p *pipeline) AddVisualization(visualization *Block) {
	p.Visualizations = append(p.Visualizations, visualization)
}

func (p *pipeline) SourceList() []*Block {
	return p.Sources
}

func (p *pipeline) TransformationList() []*Block {
	return p.Transformations
}

func (p *pipeline) VisualizationList() []*Block {
	return p.Visualizations
}


func (p *pipeline) ToPerformablePipeline() *blocks.Pipeline {
	pipeline := blocks.NewPipeline()

	for _, source := range p.Sources {
		pipeline.AddSource(source)
	}

	for _, transformation := range p.Transformations {
		pipeline.AddTransformation(transformation)
	}

	for _, visualization := range p.Visualizations {
		pipeline.AddVisualization(visualization)
	}

	return pipeline
}

