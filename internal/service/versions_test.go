package service

import (
	"testing"

	"github.com/svenliebig/binary-organizer/internal/binaries"
	"github.com/svenliebig/binary-organizer/internal/config"
)

var node binaries.Binary

func init() {
	n, err := binaries.Get("node")

	if err != nil {
		panic(err)
	}

	node = n
}

func TestVersions(t *testing.T) {
	t.Run("should return installed node versions", func(t *testing.T) {
		s := service{
			binary: node,
			config: &config.Config{
				BinaryRoot: getValidBinaryTestdataDir(),
			},
		}

		versions, err := s.Versions()

		if err != nil {
			t.Fatalf("could not get versions: %v", err)
		}

		if len(versions) != 2 {
			t.Fatalf("expected 2 versions, but got %v", len(versions))
		}
	})
}
