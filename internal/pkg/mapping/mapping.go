package mapping

import (
	"csv2csv/internal/pkg/config"
	"errors"
	"fmt"
	"github.com/extrame/xls"
	"math"
	"strings"
)

const (
	base         int = 26
	characterSet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func MapToEvents(c *config.EventParseConfig) ([]Event, error) {

	xlFile, err := xls.Open(c.InputFile, "utf-8")
	if err != nil {
		return nil, err
	}
	sheet := xlFile.GetSheet(0)
	if sheet == nil {
		return nil, errors.New("could not read sheet 0 in input file")
	}

	rowRange := c.RowRange
	numEvents := rowRange.EndRow - rowRange.StartRow
	events := make([]Event, numEvents+1) // inclusive last event
	for i := 0; i <= numEvents; i++ {
		events[i] = Event{}
		rowIndex := rowRange.StartRow + i
		row := sheet.Row(rowIndex)
		for e, colName := range c.EventCols {
			colIndex, err := toColIndex(colName)
			if err == nil {
				if fieldVal := row.Col(colIndex); fieldVal != "" {
					events[i].SetField(e, fieldVal)
				}
			}
		}
	}

	return events, nil
}

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
