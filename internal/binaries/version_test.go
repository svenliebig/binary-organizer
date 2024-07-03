package binaries

import (
	"fmt"
	"testing"
)

func TestVersionFrom(t *testing.T) {
	cases := map[string]string{
		"1.0.0":  "1.0.0",
		"1.0.1":  "1.0.1",
		"1.1.0":  "1.1.0",
		"1.1.1":  "1.1.1",
		"1.1":    "1.1.x",
		"2.0":    "2.0.x",
		"2.1":    "2.1.x",
		"2":      "2.x.x",
		"v1.0.0": "1.0.0",
		"v1.0":   "1.0.x",
		"v2":     "2.x.x",
	}

	for v, expected := range cases {
		t.Run(fmt.Sprintf("should get a valid version for %s", v), func(t *testing.T) {
			version, err := VersionFrom(v)

			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}

			if version.String() != expected {
				t.Errorf("expected %s, got %s", expected, version.String())
			}
		})
	}
}

func TestVersionMatches(t *testing.T) {
	positiveCases := map[string]string{
		"1.0.0": "1.0.0",
		"1.0.1": "1.0.1",
		"1.1.0": "1.1.0",
		"1.1.1": "1.1.1",
		"1.1":   "1.1.0",
		"1.2":   "1.2.1",
		"1":     "1.0.0",
		"2":     "2.1.0",
		"23":    "23.21.3",
	}

	for v, matches := range positiveCases {
		t.Run(fmt.Sprintf("should match %s to %v", v, matches), func(t *testing.T) {
			version, err := VersionFrom(v)

			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}

			other, err := VersionFrom(v)

			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}

			actual := version.Matches(other)

			if actual != true {
				t.Errorf("expected %v, got %v", true, actual)
			}
		})
	}

	negativeCases := map[string]string{
		"1.0.0": "1.0.1",
		"1.0.1": "1.0.0",
		"1.1.0": "1.1.1",
		"1.1.1": "1.1.0",
		"1.1":   "1.2.0",
		"1.2":   "1.1.0",
		"1":     "2.0.0",
		"2":     "1.0.0",
		"23":    "24.0.0",
	}

	for v, matches := range negativeCases {
		t.Run(fmt.Sprintf("should not match %s to %v", v, matches), func(t *testing.T) {
			version, err := VersionFrom(v)

			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}

			other, err := VersionFrom(matches)

			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}

			actual := version.Matches(other)

			if actual != false {
				t.Errorf("expected %v, got %v", false, actual)
			}
		})
	}
}
