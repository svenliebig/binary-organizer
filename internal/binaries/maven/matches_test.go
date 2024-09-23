package maven

import (
	"fmt"
	"os"
	"testing"

	"github.com/svenliebig/binary-organizer/internal/shell"
)

func TestMatcher(t *testing.T) {
	cases := map[string]bool{
		"/apache-maven-2.2.0/bin": true,
		"/apache-maven-3.6.3/bin": true,
		"/apache-maven-3.8.6/bin": true,
		"/apache-maven-3.9.6/bin": true,
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
	t.Run("should find maven 3.8.6", func(t *testing.T) {
		os.Setenv("PATH", "~/workspace/.cache/go/bin:~/workspace/software/bin:~/workspace/software/go/bin:/opt/homebrew/bin:~/workspace/software/jdk-17/Contents/Home/bin:/usr/local/bin:/System/Cryptexes/App/usr/bin:/usr/bin:/bin:/usr/sbin:/sbin:/var/run/com.apple.security.cryptexd/codex.system/bootstrap/usr/local/bin:/var/run/com.apple.security.cryptexd/codex.system/bootstrap/usr/bin:/var/run/com.apple.security.cryptexd/codex.system/bootstrap/usr/appleinternal/bin:/Library/Apple/usr/bin:/usr/local/share/dotnet:~/.dotnet/tools:~/workspace/software/maven/maven-3.8.6/bin")

		p := shell.NewPath()

		b := binary{}
		paths := p.Find(func(s string) bool {
			_, matches := b.Matches(s)
			return matches
		})

		if len(paths) == 0 {
			t.Errorf("expected to find binary")
		}

		if len(paths) > 1 {
			t.Errorf("expected to find only one binary")
		}

		if paths[0] != "~/workspace/software/maven/maven-3.8.6/bin" {
			t.Errorf("expected to find maven binary in ~/workspace/software/maven/maven-3.8.6/bin")
		}
	})
}
