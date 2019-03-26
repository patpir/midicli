package pipeline

import (
	"encoding/json"
	"io"
	"os"
)


func Read(reader io.Reader) (Pipeline, error) {
	decoder := json.NewDecoder(reader)
	content := new(pipeline)
	err := decoder.Decode(content)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return content, nil
}

func ReadFromFile(filepath string) (Pipeline, error) {
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return Read(reader)
}


