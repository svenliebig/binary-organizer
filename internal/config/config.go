package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/svenliebig/binary-organizer/internal/boo"
	"github.com/svenliebig/binary-organizer/internal/logging"
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
	viper.SetDefault("defaults.maven", "unset")
	viper.SetDefault("defaults.java", "unset")
}

type Config struct {
	// the root path where the binaries are stored.
	BinaryRoot string
	Defaults   map[string]string
}

// returns the default version for the given binary identifier.
// if no default version is found, it returns an error.
//
// possible errors:
//   - boo.ErrNoDefaultVersion
func (c Config) DefaultVersion(identifier string) (string, error) {
	defer logging.Fn("config.DefaultVersion")()

	if v, ok := c.Defaults[identifier]; ok {
		return v, nil
	}

	return "", boo.ErrNoDefaultVersion
}

func Create() (*Config, error) {
	defer logging.Fn("config.Create")()

	logging.Debug("creating configuration file")
	err := viper.SafeWriteConfig()

	// TODO exclude some values from the configuration file
	// because for example if I have root cmd flags and put
	// them into the env of viper, they are saved into the
	// configuration file as well. Things like --quite or --debug.

	if err != nil {
		logging.Error("could not write configuration file", err)
		return nil, err
	}

	return Load()
}

// Reads the boo user configuration file if present, if not
// it creates a default configuration before returning it.
func Load() (*Config, error) {
	defer logging.Fn("config.Load")()

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

		logging.Error("could not read configuration file", err)
		return nil, err
	}

	// TODO read from configuration file
	return &Config{
		BinaryRoot: viper.GetString("path"),
		Defaults:   viper.GetStringMapString("defaults"),
	}, nil
}
