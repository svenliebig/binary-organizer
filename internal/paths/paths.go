package paths

import (
	"strings"

	"github.com/svenliebig/seq"
)

type Paths interface {
	Seq() seq.Seq[string]
}

type paths struct {
	value []string
}

// returns a new paths instance based on the value of a $PATH environment variable.
//
// Example:
//
//	paths := paths.FromValue(os.Getenv("PATH"))
func FromValue(value string) (p paths) {
	r, err := seq.Collect(
		seq.Unique(
			seq.From(strings.Split(value, ":")),
		),
	)

	if err != nil {
		return
	}

	p.value = sort(r)

	return
}

func (p paths) Seq() seq.Seq[string] {
	return seq.From(p.value)
}

func sort(s []string) []string {
	if len(s) < 2 {
		return s
	}

	swapped := true
	for swapped {
		swapped = false

		v := &(s[0])
		for i := 1; i < len(s); i++ {
			c := &(s[i])

			if len(*v) > len(*c) {
				t := *v
				*v = *c
				*c = t
				swapped = true
			}

			v = c
		}
	}

	return s
}
