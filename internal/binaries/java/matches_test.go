package java

import (
	"fmt"
	"os"
	"testing"

	"github.com/svenliebig/binary-organizer/internal/shell"
)

func TestMatcher(t *testing.T) {
	cases := map[string]bool{
		"/jdk-23/bin":     true,
		"/jdk-21.0.4/bin": true,
		"/jdk-17.0.10+7-eclipse-temurin/Contents/Home/bin": true,
		"/jdk-17/Contents/Home/bin":                        true,
		"/jdk-11/Contents/Home/bin":                        true,
		"/jdk-11/bin":                                      true,
		"/jdk1.8.0_411/bin":                                true,
		"/jdk1.6.0_06/bin":                                 true,
	}

	for p, expected := range cases {
		t.Run(fmt.Sprintf("should return matching %v for %q", expected, p), func(t *testing.T) {
			b := binary{}
			_, actual := b.Matches(p)

			if actual != expected {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})
	}

	versioncases := map[string]string{
		"/jdk-23/bin":     "23.0.0",
		"/jdk-21.0.4/bin": "21.0.4",
		"/jdk-17.0.10+7-eclipse-temurin/Contents/Home/bin": "17.0.10",
		"/jdk-17/Contents/Home/bin":                        "17.0.0",
		"/jdk-11/Contents/Home/bin":                        "11.0.0",
		"/jdk-11/bin":                                      "11.0.0",
		"/jdk1.8.0_411/bin":                                "8.0.0",
		"/jdk1.6.0_06/bin":                                 "6.0.0",
	}

	for p, expected := range versioncases {
		t.Run(fmt.Sprintf("should return version %v for %q", expected, p), func(t *testing.T) {
			b := binary{}
			version, _ := b.Matches(p)

			if version.String() != expected {
				t.Errorf("expected %v, got %v", expected, version.String())
			}
		})
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

func TestWithPath(t *testing.T) {
	t.Run("should find java 17", func(t *testing.T) {
		os.Setenv("PATH", "~/workspace/.cache/go/bin:~/workspace/software/bin:~/workspace/software/go/bin:/opt/homebrew/bin:~/workspace/software/java/jdk-17/Contents/Home/bin:/usr/local/bin:/System/Cryptexes/App/usr/bin:/usr/bin:/bin:/usr/sbin:/sbin:/var/run/com.apple.security.cryptexd/codex.system/bootstrap/usr/local/bin:/var/run/com.apple.security.cryptexd/codex.system/bootstrap/usr/bin:/var/run/com.apple.security.cryptexd/codex.system/bootstrap/usr/appleinternal/bin:/Library/Apple/usr/bin:/usr/local/share/dotnet:~/.dotnet/tools:~/workspace/software/maven/maven-3.8.6/bin")

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

		if paths[0] != "~/workspace/software/java/jdk-17/Contents/Home/bin" {
			t.Errorf("expected to find java binary in ~/workspace/software/java/jdk-17/Contents/Home/bin")
		}
	})
}
