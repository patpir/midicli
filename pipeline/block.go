package pipeline

import (
)

type Block struct {
	TypeIdentifier string
	Name string
	Args map[string]interface{}
}


// Implement `github.com/patpir/miditf/blocks.Block` interface for `block` struct

func (b *Block) TypeId() string {
	return b.TypeIdentifier
}

func (b *Block) Comment() string {
	return b.Name
}

func (b *Block) Arguments() map[string]interface{} {
	return b.Args
}

