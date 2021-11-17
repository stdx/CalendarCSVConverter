package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/extrame/xls"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type EventParseArgs struct {
	EventCols map[EventField]string
	RowRange  *Range
}

type Range struct {
	StartRow int
	EndRow   int
}

func (r *Range) String() string {
	return fmt.Sprintf("%d:%d", r.StartRow, r.EndRow)
}

type EventField int

const (
	Title EventField = iota
	Description
)

type Event struct {
	Title string
}

func (r *Event) String() string {
	return fmt.Sprintf("%s", r.Title)
}

const (
	TitleColFlag       = "title"
	DescriptionColFlag = "description"
	RangeFlag          = "range"
)

func main() {

	eventArgs, err := parseEventArgs()
	if err != nil {
		panic(err)
	}

	if flag.NArg() != 1 {
		panic("Need exactly one input")
	}
	xlFile, err := xls.Open(flag.Arg(0), "utf-8")
	if err != nil {
		panic(err)
		return
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
				case Title:
					events[i].Title = fieldVal
				}
			}
		}
	}

	for _, event := range events {
		fmt.Println(event)
	}

}

func parseEventArgs() (*EventParseArgs, error) {
	rowRangeArg := flag.String("range", "", "The row range for the event fields")
	titleColArg := flag.String(TitleColFlag, "", "Column for "+TitleColFlag)
	flag.Parse()

	args := &EventParseArgs{}

	rowRange, err := parseRange(*rowRangeArg)
	if err != nil {
		return nil, err
	}
	args.RowRange = rowRange

	args.EventCols = map[EventField]string{}
	if strings.TrimSpace(*titleColArg) == "" {
		return nil, errors.New("missing required column for " + TitleColFlag)
	}
	args.EventCols[Title] = *titleColArg

	return args, nil
}

func parseRange(rangeArg string) (*Range, error) {
	if strings.TrimSpace(rangeArg) == "" {
		return nil, errors.New("missing required range")
	}
	r := regexp.MustCompile(`(?P<StartRow>\d+):(?P<EndRow>\d+)`)
	matches := r.FindStringSubmatch(rangeArg)
	if len(matches) == 0 {
		return nil, errors.New("invalid format for range of arg")
	}
	startRow, _ := strconv.Atoi(matches[r.SubexpIndex("StartRow")])
	endRow, _ := strconv.Atoi(matches[r.SubexpIndex("EndRow")])

	argRange := &Range{
		StartRow: startRow,
		EndRow:   endRow,
	}
	return argRange, nil
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
			return 0, errors.New("invalid character: " + string(char))
		}
		val += pos * int(math.Pow(float64(base), float64(pow)))
	}

	return val, nil
}
