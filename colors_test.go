package colors

import "testing"

func TestColorsSuccess(t *testing.T) {
	color := Color("bold", "bold")

	if color != "\x1B[1mbold\x1B[22m" {
		t.Error("Test fails, Output Strings did not match")
	} else {
		t.Log("Tests:TestColorsSuccess Passed!")
	}
}

func TestColorsMissing(t *testing.T) {
	color := Color("missing", "missing")

	if color == "missing" {
		t.Log("Tests:TestColorsMissing Passed!")
	} else {
		t.Error("This test must be fail!")
	}
}
