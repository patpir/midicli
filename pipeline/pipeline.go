package pipeline

import (
	"io"
)

// naming is unusual, but clashes with members of
// `pipeline` struct need to be avoided
type Pipeline interface {
	AddSource(source *Block)
	AddTransformation(transformation *Block)
	AddVisualization(visualization *Block)

	Write(writer io.Writer) error
	WriteToFile(filepath string) error

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

