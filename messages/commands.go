package messages

import (
	"fmt"
)

func PrintCommandArgumentUnexpected(arg string) {
	fmt.Printf("Unexpected argument \"%s\"\n", arg)
}

