package pipeline

import (
	"fmt"
	"os"
	"strings"

	"github.com/patpir/miditf/blocks"
)


type ResultWriter interface {
	WriteResult(result string, filename string) error
}

func (p *pipeline) Run(filename_pattern string, writer ResultWriter) error {
	performable := p.ToPerformablePipeline()

	ch := make(chan blocks.PipelineResult, 1)

	go func() {
		performable.Perform(ch)
		close(ch)
	}()

	log := newLogger(os.Stderr)
	i := 0
	for result := range ch {
		i += 1
		if result.Err != nil {
			log.printErrorTrace(&result)
		} else {
			filename := filename_pattern
			filename = strings.Replace(filename, "{source}", result.Source.Comment(), -1)
			filename = strings.Replace(filename, "{visualization}", result.Visualization.Comment(), -1)
			filename = strings.Replace(filename, "{n}", fmt.Sprintf("%d", i), -1)
			err := writer.WriteResult(result.Output, filename)
			if err != nil {
				log.printWriteError(filename, err)
			} else {
				log.printSuccessTrace(&result, filename)
			}
		}
	}

	return nil
}

