package output

import "csv2csv/internal/pkg/core"

type EventWriter interface {
	Write([]core.Event) error
}
