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

func Parse(str string) (float64, error) {
	r, _ := regexp.Compile(`(?i)^((?:\d+)?\.?\d+) *(ms|seconds?|s|minutes?|m|hours?|h|days?|d|years?|y)?$`)
	match := r.FindStringSubmatch(str)
	if match == nil {
		return 0, fmt.Errorf("Invalid string parsed: %s", str)
	}

	n, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return 0, fmt.Errorf("Number not parsed: ", str)
	}

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
		return float64(n * y), nil
	case "days":
		fallthrough
	case "day":
		fallthrough
	case "d":
		return float64(n * d), nil
	case "hours":
		fallthrough
	case "hour":
		fallthrough
	case "h":
		return float64(n * h), nil
	case "minutes":
		fallthrough
	case "minute":
		fallthrough
	case "m":
		return float64(n * m), nil
	case "seconds":
		fallthrough
	case "second":
		fallthrough
	case "s":
		return float64(n * s), nil
	case "ms":
		return float64(n), nil
	}
	return 0, nil
}

func Short(ms float64) string {
	if ms >= d {
		return fmt.Sprintf("%dd", round(ms/d))
	}
	if ms >= h {
		return fmt.Sprintf("%dh", round(ms/h))
	}
	if ms >= m {
		return fmt.Sprintf("%dm", round(ms/m))
	}
	if ms >= s {
		return fmt.Sprintf("%ds", round(ms/s))
	}
	return fmt.Sprintf("%gms", ms)
}

func Long(ms float64) string {
	if float64(ms) == d {
		return fmt.Sprintf("%d day", round(ms/d))
	}
	if float64(ms) > d {
		return fmt.Sprintf("%d days", round(ms/d))
	}
	if float64(ms) == h {
		return fmt.Sprintf("%d hour", round(ms/h))
	}
	if float64(ms) > h {
		return fmt.Sprintf("%d hours", round(ms/h))
	}
	if float64(ms) == m {
		return fmt.Sprintf("%d minute", round(ms/m))
	}
	if float64(ms) > m {
		return fmt.Sprintf("%d minutes", round(ms/m))
	}
	if float64(ms) == s {
		return fmt.Sprintf("%d second", round(ms/s))
	}
	if float64(ms) > s {
		return fmt.Sprintf("%d seconds", round(ms/s))
	}
	return fmt.Sprintf("%g ms", ms)
}

func round(n float64) int64 {
	dec := n - math.Floor(n)
	if dec >= 0.5 {
		return int64(math.Ceil(n))
	}
	return int64(math.Floor(n))
}
