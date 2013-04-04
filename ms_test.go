package ms

import "testing"

func TestParse(t *testing.T) {
	var tests = []struct {
		str  string
		ms   float64
		name string
	}{
		{"100", 100, "preserve ms"},
		{"1m", 60000, "m to ms"},
		{"1h", 3600000, "h to ms"},
		{"2d", 172800000, "d to ms"},
		{"1s", 1000, "s to ms"},
		{"100ms", 100, "ms to ms"},
		{"1.5h", 5400000, "work with decimals"},
		{"1.5H", 5400000, "case-insensitive"},
		{".5ms", 0.5, "work with decimals starting with ."},
	}

	for _, test := range tests {
		if n, _ := Parse(test.str); n != test.ms {
			t.Errorf("%s failed: Parse(%s) != %d; Actual: %f", test.name, test.str, test.ms, n)
		}
	}

	// Invalid string format
	if _, err := Parse("blerg"); err == nil {
		t.Error("Invalid string allowed to parse")
	} else if err.Error() != "Invalid string parsed: blerg" {
		t.Error("Expected parse error did not occur")
	}
}

func TestShort(t *testing.T) {
	var tests = []struct {
		ms   float64
		str  string
		name string
	}{
		{500, "500ms", "Milliseconds"},
		{1000, "1s", "Seconds"},
		{10000, "10s", "Seconds"},
		{60 * 1000, "1m", "Minutes"},
		{60 * 10000, "10m", "Minutes"},
		{60 * 60 * 1000, "1h", "Hours"},
		{60 * 60 * 10000, "10h", "Hours"},
		{24 * 60 * 60 * 1000, "1d", "Days"},
		{24 * 60 * 60 * 10000, "10d", "Days"},
		{234234234, "3d", "Rounding"},
	}

	for _, test := range tests {
		if s := Short(test.ms); s != test.str {
			t.Errorf("%s failed: Short(%f) != %s; Actual: %s", test.name, test.ms, test.str, s)
		}
	}
}

func TestLong(t *testing.T) {
	var tests = []struct {
		ms   float64
		str  string
		name string
	}{
		{500, "500 ms", "Milliseconds"},
		{1000, "1 second", "Seconds"},
		{10000, "10 seconds", "Seconds"},
		{60 * 1000, "1 minute", "Minutes"},
		{60 * 10000, "10 minutes", "Minutes"},
		{60 * 60 * 1000, "1 hour", "Hours"},
		{60 * 60 * 10000, "10 hours", "Hours"},
		{24 * 60 * 60 * 1000, "1 day", "Days"},
		{24 * 60 * 60 * 10000, "10 days", "Days"},
		{234234234, "3 days", "Rounding"},
	}

	for _, test := range tests {
		if s := Long(test.ms); s != test.str {
			t.Errorf("%s failed: Long(%f) != %s; Actual: %s", test.name, test.ms, test.str, s)
		}
	}
}
