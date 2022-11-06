package color

import "testing"

type colorRGBTest struct {
	got, expected string
}

var colorRGBTests = []colorRGBTest{
	{"rgb(255, 0, 0)", "\x1b[38;2;255;0;0m"},
	{"rgb(255,0,0)", "\x1b[38;2;255;0;0m"},
}

func TestColorsRGB(t *testing.T) {
	for _, test := range colorRGBTests {
		if output := ColorRGB(test.got); output != test.expected {
			t.Errorf("Output %q not equal to expected %q for %q", output, test.expected, test.got)
		}
	}
}
