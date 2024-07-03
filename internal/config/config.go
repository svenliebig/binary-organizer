package config

import "errors"

var (
	ErrorNoDefaultVersion = errors.New("no default version found")
)

// TODO maybe core interface?
type Config struct {
	BinaryRoot string
	Defaults   map[string]string
}

// returns the default version for the given binary identifier.
func (c Config) DefaultVersion(identifier string) (string, error) {
	if v, ok := c.Defaults[identifier]; ok {
		return v, nil
	}

	return "", ErrorNoDefaultVersion
}
