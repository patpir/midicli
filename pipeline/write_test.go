package pipeline

import (
	"testing"
	"bytes"
	"strings"

	"github.com/stretchr/testify/assert"
)


func TestWriteEmpty(t *testing.T) {
	var buf bytes.Buffer
	p := new(pipeline)
	err := p.Write(&buf)

	output := strings.TrimSpace(buf.String())

	assert.Nil(t, err)
	assert.Equal(t, `{"Sources":null,"Transformations":null,"Visualizations":null}`, output)
}

