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
