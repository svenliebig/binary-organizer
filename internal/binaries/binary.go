package binaries

import (
	"errors"
	"sync"

	"github.com/svenliebig/binary-organizer/internal/logging"
	"github.com/svenliebig/seq"
)

var ErrBinaryNotFound = errors.New("binary not found")

var binaries []Binary
var lock = sync.RWMutex{}

type Binary interface {
	// returns the identifier of the binary.
	Identifier() string

	// checks if the given path matches the binary.
	Matches(path string) (Version, bool)

	// returns the path to the bin directory of the binary.
	// given is the root path of the version directory.
	BinPath(string) string
}

func Register(b Binary) {
	defer logging.Fn("binaries.Register")()

	lock.Lock()
	defer lock.Unlock()

	binaries = append(binaries, b)
}

func All() seq.Seq[Binary] {
	defer logging.Fn("binaries.All")()

	lock.RLock()
	defer lock.RUnlock()

	return seq.From(binaries)
}

func Get(identifier string) (Binary, error) {
	defer logging.Fn("binaries.Get")()

	lock.RLock()
	defer lock.RUnlock()

	for _, b := range binaries {
		if b.Identifier() == identifier {
			return b, nil
		}
	}

	return nil, ErrBinaryNotFound
}
