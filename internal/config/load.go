package config

import (
	"os"
	"path"
)

// Reads the boo user configuration file if present, if not
// it creates a default configuration before returning it.
func Load() (Config, error) {

	// TODO find or create a configuration file in:
	// - ~/.boo.toml
	// - ~/.config/boo.toml
	// - ~/.config/boo/boo.toml

	home, err := os.UserHomeDir()

	if err != nil {
		return Config{}, err
	}

	// TODO read from configuration file
	return Config{
		BinaryRoot: path.Join(home, "workspace", "software"),
		Defaults: map[string]string{
			"node": "20",
		},
	}, nil
}
