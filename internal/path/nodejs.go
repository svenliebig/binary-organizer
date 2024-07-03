package path

import (
	"fmt"
	"regexp"
	"strings"
)

var operatingSystems = []string{"darwin", "win", "linux"}
var architectures = map[string][]string{
	"darwin": {"arm64", "x64"},
}

type NodeJS struct {
	OS           string
	Architecture string
	Version      string
}

func IsNodeJS(p string) (*NodeJS, bool) {

	r := fmt.Sprintf(`node-v(\d+)\.(\d+)\.(\d+)-(%s)-([\w\d]*)`, strings.Join(operatingSystems, "|"))
	regexp := regexp.MustCompile(r)

	matches := regexp.FindStringSubmatch(p)

	if len(matches) != 0 {
		version := fmt.Sprintf("%s.%s.%s", matches[1], matches[2], matches[3])
		opers := matches[4]
		arch := matches[5]

		fmt.Println(version, opers, arch)

		return &NodeJS{
			OS:           opers,
			Architecture: "x64",
			Version:      version,
		}, true
	}

	return nil, false
}
