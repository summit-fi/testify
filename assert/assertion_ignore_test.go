package assert

import "testing"

func TestAssertions_GetIgnoredFields(t *testing.T) {
	mockT := new(testing.T)
	test := New(mockT)
	test.SetIgnoredFields("a", "b")
	Equal(t, []string{"a", "b"}, test.GetIgnoredFields())
}

func TestAssertions_SetIgnoredFields(t *testing.T) {
	mockT := new(testing.T)
	test := New(mockT)
	test.SetIgnoredFields("c", "b")
	Equal(t, []string{"c", "b"}, test.GetIgnoredFields())
}
