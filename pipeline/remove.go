package pipeline


func (p *pipeline) RemoveSourceAt(index int) {
	p.Sources = append(p.Sources[:index], p.Sources[index+1:]...)
}

func (p *pipeline) RemoveTransformationAt(index int) {
	p.Transformations = append(p.Transformations[:index], p.Transformations[index+1:]...)
}

func (p *pipeline) RemoveVisualizationAt(index int) {
	p.Visualizations = append(p.Visualizations[:index], p.Visualizations[index+1:]...)
}

