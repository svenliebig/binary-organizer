package shell

import (
	"fmt"
	"os"
	"strings"

	"github.com/svenliebig/seq"
)

type matcher = func(string) bool

type Path interface {
	// collects all the path strings that match the given matcher.
	Find(m matcher) []string

	// removes the given path from the PATH variable, if it exists.
	Remove(p string)

	// add a path to the PATH variable.
	Add(p string)

	// creates a valid bash export statement for the PATH variable.
	//
	// example:
	//
	//	export PATH="/usr/local/bin:/usr/bin:/bin"
	Export() string
}

type path struct {
	content *paths
}

func NewPath() Path {
	return path{
		content: pathFromValue(os.Getenv("PATH")),
	}
}

func (p path) Export() string {
	s := p.content.seq()
	a, err := seq.Collect(s)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("export PATH=\"%s\"", strings.Join(a, ":\\\n"))
}

func (p path) Find(m matcher) []string {
	l, err := seq.Collect(
		seq.Filter(
			p.content.seq(),
			func(s string) (bool, error) {
				return m(s), nil
			},
		),
	)

	if err != nil {
		fmt.Println("error while finding paths, this should not be possible:", err)
		panic(err)
	}

	return l
}

func (p path) Add(path string) {
	p.content.add(path)
}

func (p path) Remove(path string) {
	p.content.remove(path)
}
