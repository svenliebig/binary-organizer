package node

import (
	"path"

	"github.com/svenliebig/binary-organizer/internal/binaries"
)

type binary struct {
}

func (n binary) Identifier() string {
	return "node"
}

func init() {
	binaries.Register(&binary{})
}

func (n binary) BinPath(root string) string {
	return path.Join(root, "bin")
}
