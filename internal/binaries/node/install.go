package node

import (
	"context"
	"os"
	"path"

	"github.com/svenliebig/binary-organizer/internal/binaries"
)

func (n binary) Install(ctx context.Context, v binaries.Version) error {
	panic("install not implemented")
}

func (n binary) IsInstalled(ctx context.Context, root string, v binaries.Version) (string, bool) {
	entries, err := os.ReadDir(path.Join(root, n.Identifier()))

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if vers, ok := n.Matches(entry.Name()); ok {
			if v.Matches(vers) {
				return path.Join(root, n.Identifier(), entry.Name(), "bin"), true
			}
		}
	}

	return "", false
}
