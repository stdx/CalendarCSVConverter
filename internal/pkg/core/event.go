package core

import (
	"fmt"
)

type Event struct {
	Title       string
	Description string
	StartDate   string
	StartTime   string
}

func (e *Event) String() string {
	return fmt.Sprintf("%s;%s;%s;%s", e.Title, e.Description, e.StartDate, e.StartTime)
}

func (e *Event) SetField(f EventField, val string) {
	switch f {
	case Title:
		e.Title = val
	case Description:
		e.Description = val
	case StartDate:
		e.StartDate = val
	case StartTime:
		e.StartTime = val
	}
}
