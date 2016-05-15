package exploration

import (
	"testing"
)

func TestSymbol(t *testing.T) {
	cases := []struct {
		sense    Sense
		expected rune
	}{
		{Sense(0), 'N'},
		{Sense(1), 'E'},
		{Sense(2), 'S'},
		{Sense(3), 'W'},
		{Sense(4), 'N'},
		{Sense(5), 'E'},
	}

	for _, c := range cases {
		if c.expected != c.sense.Symbol() {
			t.Errorf("The sense %v should have symbol %v, instead of %v", c.sense, c.expected, c.sense.Symbol())
		}
	}
}
