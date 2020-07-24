package checker

import "testing"

func TestInArray(t *testing.T) {
	var needle string = "vaddy.net"
	var haystack = []string{"vaddy.net", "hoge.vaddy.net"}

	if !in_array(needle, haystack) {
		t.Fail()
	}
}

func TestInArrayNotFound(t *testing.T) {
	var needle string = "foo"
	var haystack = []string{"vaddy.net", "hoge.vaddy.net"}

	if in_array(needle, haystack) {
		t.Fail()
	}
}
