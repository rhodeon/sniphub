package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Form struct {
	Values url.Values
	Errors errors
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// New creates a new instance of a Form pointer from a submitted form values.
func New(data url.Values) *Form {
	return &Form{
		Values: data,
		Errors: map[string][]string{},
	}
}

// Valid checks if any error exists in the form.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required checks if the specified fields are populated with non-whitespace characters.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Values.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MaxLength adds an error if the field value(s) have a higher character count than the maximum limit.
func (f *Form) MaxLength(max int, fields ...string) {
	for _, field := range fields {
		value := f.Values.Get(field)
		if utf8.RuneCountInString(value) > max {
			f.Errors.Add(field, fmt.Sprintf("This field must not have over %d characters", max))
		}
	}
}

// MinLength adds an error if the field value(s) have a lower character count than the minimum limit.
func (f *Form) MinLength(min int, fields ...string) {
	for _, field := range fields {
		value := f.Values.Get(field)
		if utf8.RuneCountInString(value) < min {
			f.Errors.Add(field, fmt.Sprintf("This field must have at least %d characters", min))
		}
	}

}

// MatchesPattern checks if the field values matches the given pattern.
func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Values.Get(field)
	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}
