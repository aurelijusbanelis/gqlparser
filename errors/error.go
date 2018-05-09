package errors

import "fmt"

type Syntax struct {
	Message   string     `json:"message"`
	Locations []Location `json:"locations,omitempty"`
}

type Location struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

func (err *Syntax) Error() string {
	if err == nil {
		return "<nil>"
	}
	str := fmt.Sprintf("Syntax Error: %s", err.Message)
	for _, loc := range err.Locations {
		str += fmt.Sprintf(" (line %d, column %d)", loc.Line, loc.Column)
	}
	return str
}