package main

import (
	"csv2csv/internal/pkg/config"
	"csv2csv/internal/pkg/mapping"
	"csv2csv/internal/pkg/output"
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

	events, err := mapping.ReadEvents(parseConfig)
	if err != nil {
		return err
	}

	g := &output.GoogleCsvEventWriter{}
	g.Write(events)
	return nil
}
