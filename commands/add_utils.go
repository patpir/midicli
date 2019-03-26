package commands

import (
	"errors"
	"strings"

	"github.com/urfave/cli"

	"github.com/patpir/midicli/pipeline"
)


type blockArguments cli.Args


var missingArgumentsError error = errors.New("Type identifier and name are required")
var invalidArgumentError error = errors.New("Argument does not match name=value pattern")


func (args blockArguments) toBlock() (*pipeline.Block, error) {
	if len(args) < 2 {
		return nil, missingArgumentsError
	}

	namedArgs := make(map[string]interface{})
	for _, arg := range args[2:] {
		if !isNamedArgument(arg) {
			return nil, invalidArgumentError
		}
		name, value := splitNameValue(arg)
		namedArgs[name] = value
	}
	return &pipeline.Block{
		TypeIdentifier: args[0],
		Name: args[1],
		Args: namedArgs,
	}, nil
}

func isNamedArgument(arg string) bool {
	return strings.Contains(arg, "=")
}

func splitNameValue(arg string) (string, string) {
	parts := strings.SplitN(arg, "=", 2)
	return parts[0], parts[1]
}

