package shell

import (
	"fmt"
	"strings"

	"github.com/svenliebig/seq"
)

// represents a list of path values.
type paths struct {
	value []string
}

// returns a new paths instance based on the value of a $PATH environment variable.
//
// Example:
//
//	paths := paths.pathFromValue(os.Getenv("PATH"))
func pathFromValue(value string) *paths {
	p := &paths{}

	r, err := seq.Collect(
		seq.Unique(
			seq.From(strings.Split(value, ":")),
		),
	)

	if err != nil {
		return p
	}

	// TODO maybe we shouldn't sort
	p.value = r

	return p
}

func (p *paths) seq() seq.Seq[string] {
	return seq.From(p.value)
}

func (p *paths) add(path string) {
	p.value = append([]string{path}, p.value...)
}

func (p *paths) remove(path string) {
	r, err := seq.Collect(
		seq.Filter(
			seq.From(p.value),
			func(s string) (bool, error) {
				return s != path, nil
			},
		),
	)

	if err != nil {
		fmt.Println("error while removing path, this should not be possible:", err)
		panic(err)
	}

	p.value = r
}

// func sort(s []string) []string {
// 	if len(s) < 2 {
// 		return s
// 	}
//
// 	swapped := true
// 	for swapped {
// 		swapped = false
//
// 		v := &(s[0])
// 		for i := 1; i < len(s); i++ {
// 			c := &(s[i])
//
// 			if len(*v) > len(*c) {
// 				t := *v
// 				*v = *c
// 				*c = t
// 				swapped = true
// 			}
//
// 			v = c
// 		}
// 	}
//
// 	return s
// }
