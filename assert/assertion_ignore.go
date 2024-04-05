package assert

// SetIgnoredFields sets the fields to be ignored by the Equal and Equivalent assertions.
func (a *Assertions) SetIgnoredFields(fields ...string) {
	if a.IgnoredFields == nil {
		a.IgnoredFields = make(map[string]bool)
	}
	for _, field := range fields {
		a.IgnoredFields[field] = true
	}
}

// GetIgnoredFields returns the fields to be ignored by the Equal and Equivalent assertions.
func (a *Assertions) GetIgnoredFields() []string {
	fields := make([]string, 0, len(a.IgnoredFields))
	for field := range a.IgnoredFields {
		fields = append(fields, field)
	}
	return fields
}
