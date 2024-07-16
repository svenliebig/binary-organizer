package node

import (
	"fmt"
	"os"
	"testing"

	"github.com/svenliebig/binary-organizer/internal/shell"
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

func TestMatcherWithPath(t *testing.T) {
	t.Run("should find node 16", func(t *testing.T) {
		os.Setenv("PATH", "~/workspace/.cache/go/bin:~/workspace/software/bin:~/workspace/software/go/bin:/opt/homebrew/bin:~/workspace/software/jdk-17/Contents/Home/bin:/usr/local/bin:/System/Cryptexes/App/usr/bin:/usr/bin:/bin:/usr/sbin:/sbin:/var/run/com.apple.security.cryptexd/codex.system/bootstrap/usr/local/bin:/var/run/com.apple.security.cryptexd/codex.system/bootstrap/usr/bin:/var/run/com.apple.security.cryptexd/codex.system/bootstrap/usr/appleinternal/bin:/Library/Apple/usr/bin:/usr/local/share/dotnet:~/.dotnet/tools:~/workspace/software/node/node-v16.20.2-darwin-arm64/bin")

		p := shell.NewPath()

		b := binary{}
		nodePaths := p.Find(func(s string) bool {
			_, matches := b.Matches(s)
			return matches
		})

		if len(nodePaths) == 0 {
			t.Errorf("expected to find node binary")
		}

		if len(nodePaths) > 1 {
			t.Errorf("expected to find only one node binary")
		}

		if nodePaths[0] != "~/workspace/software/node/node-v16.20.2-darwin-arm64/bin" {
			t.Errorf("expected to find node binary in ~/workspace/software/node/node-v16.20.2-darwin-arm64/bin")
		}
	})
}
