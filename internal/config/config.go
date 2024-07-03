package config

// TODO maybe core interface?
type Config struct {
	BinaryRoot string
}

// returns the default version for the given binary identifier.
func (c Config) DefaultVersion(identifier string) (string, error) {
	// TODO read from configuration file
	return "20", nil
}
