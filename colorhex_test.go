package color

import "testing"

type colorsHexTest struct {
	got, expected string
}

var colorsHexTests = []colorsHexTest{
	{"#ff0000", "\x1b[38;2;255;0;0m"},
	{"#f00", "\x1b[38;2;255;0;0m"},
	{"#00f", "\x1b[38;2;0;0;255m"},
	{"#FFA500", "\x1b[38;2;255;165;0m"},
}

func TestColorsHex(t *testing.T) {
	for _, test := range colorsHexTests {
		if output := ColorHex(test.got); output != test.expected {
			t.Errorf("Output %q not equal to expected %q for %q", output, test.expected, test.got)
		}
	}
}
