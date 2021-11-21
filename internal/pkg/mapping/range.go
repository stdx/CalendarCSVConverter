package mapping

import "fmt"

type Range struct {
	StartRow int
	EndRow   int
}

func (r *Range) String() string {
	return fmt.Sprintf("%d:%d", r.StartRow, r.EndRow)
}
