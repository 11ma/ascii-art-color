package color

import "testing"

type colorsHslTest struct {
	got, expected string
}

var colorsHslTests = []colorsHslTest{
	{"hsl(0, 100%, 50%)", "\x1b[38;2;255;0;0m"},
	{"hsl(0,100%,50%)", "\x1b[38;2;255;0;0m"},
}

func TestColorsHsl(t *testing.T) {
	for _, test := range colorsHslTests {
		if output := ColorHSL(test.got); output != test.expected {
			t.Errorf("Output %q not equal to expected %q for %q", output, test.expected, test.got)
		}
	}
}
