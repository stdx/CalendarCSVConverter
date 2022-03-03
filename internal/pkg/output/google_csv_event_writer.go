package output

import (
	"csv2csv/internal/pkg/core"
	"fmt"
)

type GoogleCsvEventWriter struct {
}

func (g *GoogleCsvEventWriter) Write(events []core.Event) error {
	for _, event := range events {
		fmt.Println(event.String())
	}
	return nil
}
