package ms

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	s float64 = 1000
	m float64 = s * 60
	h float64 = m * 60
	d float64 = h * 24
	y float64 = d * 365.25
)

var parseRgx, _ = regexp.Compile(`(?i)^((?:\d+)?\.?\d+) *(ms|seconds?|s|minutes?|m|hours?|h|days?|d|years?|y)?$`)

func Parse(str string) (float64, error) {
	match := parseRgx.FindStringSubmatch(str)
	if match == nil {
		return 0, fmt.Errorf("Invalid string parsed: %s", str)
	}

	n, _ := strconv.ParseFloat(match[1], 64)

	var t string
	if match[2] != "" {
		t = strings.ToLower(match[2])
	} else {
		t = "ms"
	}

	switch t {
	case "years":
		fallthrough
	case "year":
		fallthrough
	case "y":
		return n * y, nil
	case "days":
		fallthrough
	case "day":
		fallthrough
	case "d":
		return n * d, nil
	case "hours":
		fallthrough
	case "hour":
		fallthrough
	case "h":
		return n * h, nil
	case "minutes":
		fallthrough
	case "minute":
		fallthrough
	case "m":
		return n * m, nil
	case "seconds":
		fallthrough
	case "second":
		fallthrough
	case "s":
		return n * s, nil
	case "ms":
		return n, nil
	}
	return 0, nil
}

func Short(ms float64) string {
	if ms >= d {
		return fmt.Sprintf("%gd", round(ms/d))
	}
	if ms >= h {
		return fmt.Sprintf("%gh", round(ms/h))
	}
	if ms >= m {
		return fmt.Sprintf("%gm", round(ms/m))
	}
	if ms >= s {
		return fmt.Sprintf("%gs", round(ms/s))
	}
	return fmt.Sprintf("%gms", ms)
}

func Long(ms float64) string {
	if float64(ms) == d {
		return fmt.Sprintf("%g day", round(ms/d))
	}
	if float64(ms) > d {
		return fmt.Sprintf("%g days", round(ms/d))
	}
	if float64(ms) == h {
		return fmt.Sprintf("%g hour", round(ms/h))
	}
	if float64(ms) > h {
		return fmt.Sprintf("%g hours", round(ms/h))
	}
	if float64(ms) == m {
		return fmt.Sprintf("%g minute", round(ms/m))
	}
	if float64(ms) > m {
		return fmt.Sprintf("%g minutes", round(ms/m))
	}
	if float64(ms) == s {
		return fmt.Sprintf("%g second", round(ms/s))
	}
	if float64(ms) > s {
		return fmt.Sprintf("%g seconds", round(ms/s))
	}
	return fmt.Sprintf("%g ms", ms)
}

func round(n float64) float64 {
	dec := n - math.Floor(n)
	if dec >= 0.5 {
		return math.Ceil(n)
	}
	return math.Floor(n)
}
