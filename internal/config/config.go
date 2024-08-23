package config

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/svenliebig/binary-organizer/internal/boo"
	"github.com/svenliebig/binary-organizer/internal/logging"
)

var (
	ErrorNoDefaultVersion = errors.New("no default version found")
)

func init() {
	viper.SetConfigName("boo")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config")

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.SetDefault("path", path.Join(home, "workspace", "software"))
	viper.SetDefault("defaults.node", "unset")
}

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

func Create() (*Config, error) {
	logging.Debug("creating configuration file")
	err := viper.SafeWriteConfig()

	if err != nil {
		logging.Error("could not write configuration file", err)
		return nil, err
	}

	return Load()
}

// Reads the boo user configuration file if present, if not
// it creates a default configuration before returning it.
func Load() (*Config, error) {
	// TODO find or create a configuration file in:
	// - ~/.boo.toml
	// - ~/.config/boo.toml
	// - ~/.config/boo/boo.toml

	logging.Debug("reading configuration file")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logging.Debug("config file not found by viper")
			return nil, boo.ErrConfigFileNotExists
		}
		return nil, err
	}

	// TODO read from configuration file
	return &Config{
		BinaryRoot: viper.GetString("path"),
		Defaults:   viper.GetStringMapString("defaults"),
	}, nil
}
