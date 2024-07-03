package binaries

import (
	"sync"

	"context"
)

var binaries []Binary
var lock = sync.RWMutex{}

type Binary interface {
	// returns the identifier of the binary.
	Identifier() string

	// checks if the given path matches the binary.
	Matches(path string) (Version, bool)

	// problems here: needs the configuration (like software path), maybe as context.
	Install(context.Context, Version) error

	// checks if the given version of that binary is installed.
	// if it is installed, it returns the path to the binary.
	IsInstalled(context.Context, Version) (string, bool)
}

func Register(b Binary) {
	lock.Lock()
	defer lock.Unlock()

	binaries = append(binaries, b)
}

func All() []Binary {
	lock.RLock()
	defer lock.RUnlock()

	return binaries
}
