package require

// Assertions provides assertion methods around the
// TestingT interface.
type Assertions struct {
	t TestingT
	// IgnoredFields is a map of fields to ignore when comparing structs.
	IgnoredFields map[string]bool
}

// New makes a new Assertions object for the specified TestingT.
func New(t TestingT) *Assertions {
	return &Assertions{
		t: t,
	}
}

//go:generate sh -c "cd ../_codegen && go build && cd - && ../_codegen/_codegen -output-package=require -template=require_forward.go.tmpl -include-format-funcs"
