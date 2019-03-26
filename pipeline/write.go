package pipeline

import (
	"encoding/json"
	"io"
	"os"
)


func (p *pipeline) Write(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(p)
	return err
}

func (p *pipeline) WriteToFile(filepath string) error {
	writer, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer writer.Close()

	return p.Write(writer)
}

