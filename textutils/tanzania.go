package textutils

import (
	"fmt"
	"regexp"
	"strings"
)

// Creates a new number
func NewTanzania(phone string) (*PhoneNumber, error) {
	phone = strings.Replace(phone, " ", "", -1)

	r, _ := regexp.Compile(`^((\+)?(255)|0)(\d{9})$`)
	// check if phone has +255 / 255

	if !r.Match([]byte(phone)) {
		return nil, fmt.Errorf("phone doesnt match a valid tanzanian pattern")
	}

	ph := r.ReplaceAllString(phone, "+255${4}")
	return &PhoneNumber{
		val:  ph,
		code: "+255",
	}, nil
}
