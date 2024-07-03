package shell

import (
	"os"
	"strings"
	"testing"

	"github.com/svenliebig/seq"
)

func TestPathFind(t *testing.T) {
	t.Run("should find three paths", func(t *testing.T) {
		os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin")

		p := NewPath()

		paths := p.Find(func(s string) bool {
			return strings.Contains(s, "bin")
		})

		if len(paths) != 3 {
			t.Errorf("expected 3 paths, got %d", len(paths))
		}
	})

	t.Run("should find no paths", func(t *testing.T) {
		os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin")

		p := NewPath()

		paths := p.Find(func(s string) bool {
			return strings.Contains(s, "foo")
		})

		if len(paths) != 0 {
			t.Errorf("expected 0 paths, got %d", len(paths))
		}
	})

	t.Run("should find one path", func(t *testing.T) {
		os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin")

		p := NewPath()

		paths := p.Find(func(s string) bool {
			return strings.Contains(s, "local")
		})

		if len(paths) != 1 {
			t.Errorf("expected 1 path, got %d", len(paths))
		}
	})
}

func TestPathRemove(t *testing.T) {
	t.Run("should remove one path", func(t *testing.T) {
		os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin")

		p := NewPath()

		p.Remove("/usr/local/bin")

		ex := p.Export()

		if strings.Contains(ex, "/usr/local/bin") {
			t.Errorf("expected path to be removed, got %s", ex)
		}

		l, err := seq.Collect(p.(path).content.Seq())

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(l) != 2 {
			t.Errorf("expected 2 paths, got %d", len(l))
		}
	})
}

func TestPathAdd(t *testing.T) {
	t.Run("should add one path", func(t *testing.T) {
		os.Setenv("PATH", "/usr/local/bin:/usr/bin:/bin")

		p := NewPath()

		p.Add("/usr/local/foo")

		ex := p.Export()

		if !strings.Contains(ex, "/usr/local/foo") {
			t.Errorf("expected path to be added, got %s", ex)
		}

		l, err := seq.Collect(p.(path).content.Seq())

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(l) != 4 {
			t.Errorf("expected 4 paths, got %d", len(l))
		}
	})
}
