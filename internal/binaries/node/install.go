package node

import (
	"context"

	"github.com/svenliebig/binary-organizer/internal/binaries"
)

func (n binary) Install(ctx context.Context, v binaries.Version) error {
	return nil
}

func (n binary) IsInstalled(ctx context.Context, v binaries.Version) (string, bool) {
	return "", false
}
