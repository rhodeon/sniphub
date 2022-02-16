package forms

type errors map[string][]string

func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}

// Returns the first error message from the specified field
func (e errors) Get(field string) string {
	nestedMap := e[field]
	if len(nestedMap) == 0 {
		return ""
	}

	return nestedMap[0]
}
