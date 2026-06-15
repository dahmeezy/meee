package coloring

import "testing"

func TestSub2color(t *testing.T) {
	result := Sub2color("Hello", "el")
	expected := []bool{false, true, true, false, false}

	if len(result) != len(expected) {
		t.Errorf("Got %v , expected %v", len(result), len(expected))
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("got %v, expected %v", result, expected)

		}

	}
}

func TestGetcolor(t *testing.T) {

	cases := []struct {
		input    string
		expected string
	}{
		// test cases here
		{input: "red", expected: "\033[31m"},
		{input: "green", expected: "\033[32m"},
		{input: "yellow", expected: "\033[33m"},
		{input: "blue", expected: "\033[34m"},
		{input: "purple", expected: "\033[35m"},
		{input: "white", expected: "\033[36m"},
		{input: "", expected: "\033[0m"},
	}

	for _, v := range cases {
		if v.input != v.expected {
			t.Errorf("Got %v , expected %v", v.input, v.expected)
			return
		}
	}

}
