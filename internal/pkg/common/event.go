package common

import "fmt"

type Event struct {
	Title string
}

func (r *Event) String() string {
	return fmt.Sprintf("%s", r.Title)
}
