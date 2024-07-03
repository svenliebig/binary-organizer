package binaries

import (
	"fmt"
	"regexp"
)

type Version interface {
	String() string
	Matches(Version) bool
}

type version struct {
	major string
	minor string
	patch string
}

var fullVersionRegex = regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)$`)
var partialVersionRegex = regexp.MustCompile(`^v?(\d+)\.(\d+)$`)
var xVersionRegex = regexp.MustCompile(`^v?(\d+)$`)

func VersionFrom(s string) (Version, error) {
	if matches := fullVersionRegex.FindStringSubmatch(s); len(matches) != 0 {
		return &version{
			major: matches[1],
			minor: matches[2],
			patch: matches[3],
		}, nil
	}

	if matches := partialVersionRegex.FindStringSubmatch(s); len(matches) != 0 {
		return &version{
			major: matches[1],
			minor: matches[2],
			patch: "x",
		}, nil
	}

	if matches := xVersionRegex.FindStringSubmatch(s); len(matches) != 0 {
		return &version{
			major: matches[1],
			minor: "x",
			patch: "x",
		}, nil
	}

	return nil, fmt.Errorf("invalid version string: %q", s)
}

func (v *version) String() string {
	return v.major + "." + v.minor + "." + v.patch
}

func (v *version) Matches(other Version) bool {
	otherVersion, ok := other.(*version)

	if !ok {
		return false
	}

	if v.major != otherVersion.major {
		return false
	}

	if v.minor != otherVersion.minor && v.minor != "x" {
		return false
	}

	if v.patch != otherVersion.patch && v.patch != "x" {
		return false
	}

	return true
}
