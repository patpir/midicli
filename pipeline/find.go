package pipeline


func (p *pipeline) FindSourceByName(name string) (int, *Block) {
	return findBlockByName(p.Sources, name)
}

func (p *pipeline) FindTransformationByName(name string) (int, *Block) {
	return findBlockByName(p.Transformations, name)
}

func (p *pipeline) FindVisualizationByName(name string) (int, *Block) {
	return findBlockByName(p.Visualizations, name)
}

func findBlockByName(blocks []*Block, name string) (int, *Block) {
	for i, block := range blocks {
		if block.Name == name {
			return i, block
		}
	}

	return -1, nil
}

