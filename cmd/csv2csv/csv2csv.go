package main

import (
	"csv2csv/internal/pkg/config"
	"csv2csv/internal/pkg/mapping"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {

	parseConfig, err := config.FromCmdLine()
	if err != nil {
		return err
	}

	events, err := mapping.MapToEvents(parseConfig)
	if err != nil {
		return err
	}

	for _, event := range events {
		fmt.Println(event)
	}

	return nil
}
