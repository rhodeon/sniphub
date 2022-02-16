package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	Values url.Values
	Errors errors
}

// Creates a new instance of a Form pointer from a submitted form values.
func New(data url.Values) *Form {
	return &Form{
		Values: data,
		Errors: map[string][]string{},
	}
}

// Checks if the specified fields are populated with non-whitespace characters.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Values.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Adds an error if the field value(s) have a higher character count than the maximum limit.
func (f *Form) MaxLength(max int, fields ...string) {
	for _, field := range fields {
		value := f.Values.Get(field)
		if utf8.RuneCountInString(value) > max {
			f.Errors.Add(field, fmt.Sprintf("This field must not have over %d characters", max))
		}
	}
}

// Checks if any error exists in the form.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
