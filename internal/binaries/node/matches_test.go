package node

import (
	"fmt"
	"testing"
)

func TestMatcher(t *testing.T) {
	cases := map[string]bool{
		"/node-v22.3.0-darwin-arm64/bin": true,
		"/node-v21.7.3-darwin-arm64/bin": true,
		"/node-v22.3.0-win-x64/bin":      true,
		"/node-v21.7.3-win-x64/bin":      true,
		"/node-v22.3.0-win-x86/bin":      true,
		"/node-v21.7.3-win-x86/bin":      true,
		"/node-v22.3.0-linux-x64/bin":    true,
		"/node-v21.7.3-linux-x64/bin":    true,
		"/node-v22.3.0-linux-x86/bin":    true,
		"/node-v21.7.3-linux-x86/bin":    true,
		"/node-v22.3.0-linux-s390x/bin":  true,
		"/node-v21.7.3-linux-s390x/bin":  true,
	}

	for p, expected := range cases {
		t.Run(fmt.Sprintf("should return %v for %q", expected, p), func(t *testing.T) {
			b := binary{}
			_, actual := b.Matches(p)

			if actual != expected {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})
	}
}
