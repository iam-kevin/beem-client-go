package textutils

import (
	"fmt"
)

// Instance of the phone number
type PhoneNumber struct {
	val  string
	code string
}

// Returns the true formatted phone number
func (p *PhoneNumber) Value() string {
	return p.val
}

// When stringifying the number, it returns the masked version.
// To use the value explicity, user `.Value()`
//
// This comes in handy so that we don't accidently log it
func (p *PhoneNumber) String() string {
	return fmt.Sprintf(`%s***%s`, p.code, p.val[len(p.val)-3:])
}
