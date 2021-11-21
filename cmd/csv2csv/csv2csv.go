package main

import (
	"csv2csv/internal/pkg/config"
	"csv2csv/internal/pkg/mapping"
	"fmt"
	"github.com/extrame/xls"
	"math"
	"os"
	"strings"
)

type Event struct {
	Title string
}

func (r *Event) String() string {
	return fmt.Sprintf("%s", r.Title)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {

	eventArgs, err := config.FromCmdLine()
	if err != nil {
		return err
	}

	xlFile, err := xls.Open(eventArgs.InputFile, "utf-8")
	if err != nil {
		return err
	}

	sheet := xlFile.GetSheet(0)

	numEvents := eventArgs.RowRange.EndRow - eventArgs.RowRange.StartRow
	events := make([]*Event, numEvents+1) // inclusive last event
	for i := 0; i <= numEvents; i++ {
		events[i] = &Event{}
		rowIndex := eventArgs.RowRange.StartRow + i
		row := sheet.Row(rowIndex)
		for e, colName := range eventArgs.EventCols {
			var fieldVal = ""
			colIndex, err := toColIndex(colName)
			if err == nil {
				if row.Col(colIndex) != "" {
					fieldVal = row.Col(colIndex)
				}
				switch e {
				case mapping.Title:
					events[i].Title = fieldVal
				}
			}
		}
	}

	for _, event := range events {
		fmt.Println(event)
	}

	return nil

}

const (
	base         int = 26
	characterSet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func toColIndex(encoded string) (int, error) {
	var val int
	for index, char := range encoded {
		pow := len(encoded) - (index + 1)
		pos := strings.IndexRune(characterSet, char)
		if pos == -1 {
			return 0, fmt.Errorf("invalid character: %s", string(char))
		}
		val += pos * int(math.Pow(float64(base), float64(pow)))
	}

	return val, nil
}
