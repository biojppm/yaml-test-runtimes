package main

import (
	"fmt"
	"io"
	"os"

	"go.yaml.in/yaml/v4"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	events, err := yaml.ParserGetEvents(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(events)
}
