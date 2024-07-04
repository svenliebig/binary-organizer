package shell

import (
	"fmt"
	"os"
	"path/filepath"
)

func WritePath(p Path) error {
	exe, err := os.Executable()

	if err != nil {
		return fmt.Errorf("could not get executable: %w", err)
	}

	dir := filepath.Dir(exe)
	filename := filepath.Join(dir, ".path")

	err = os.WriteFile(filename, []byte(fmt.Sprintf("%s\n", p.Export())), 0644)

	if err != nil {
		return fmt.Errorf("could not write path: %w", err)
	}

	return nil
}
