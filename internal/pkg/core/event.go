package core

import "fmt"

type Event struct {
	Title       string
	Description string
}

func (e *Event) String() string {
	return fmt.Sprintf("%s;%s", e.Title, e.Description)
}

func (e *Event) SetField(f EventField, val string) {
	switch f {
	case Title:
		e.Title = val
	case Description:
		e.Description = val
	}
}
