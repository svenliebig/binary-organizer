package path

import (
	"fmt"
	"os"
	"strings"

	"github.com/svenliebig/binary-organizer/internal/paths"
	"github.com/svenliebig/seq"
)

type PathVariable struct {
	p paths.Paths
}

func NewPathVariable() PathVariable {
	return PathVariable{
		p: paths.FromValue(os.Getenv("PATH")),
	}
}

func (p PathVariable) Export() string {
	s := p.p.Seq()
	a, err := seq.Collect(s)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("export PATH=\"%s\"", strings.Join(a, ":\\\n"))
}
