package textutils

import (
	"regexp"
	"testing"
)

func TestNormalizePhoneNumber(t *testing.T) {
	valid_inputs := map[string]*regexp.Regexp{
		"255 766 123 456":    regexp.MustCompile(`\b+255766123456\b`),
		"0762 145     123  ": regexp.MustCompile(`\b+255762145123\b`),
	}

	for input, want := range valid_inputs {
		val, err := NewTanzania(input)
		if !want.Match([]byte(val.Value())) || err != nil {
			t.Errorf("NormalizePhoneNumber(%s); got %s but want %s", input, val, want.String())
		}
	}
}
