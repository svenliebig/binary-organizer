package shell

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/svenliebig/binary-organizer/internal/logging"
)

func WritePath(p Path) error {
	defer logging.Fn("shell.WritePath")()

	exe, err := os.Executable()

	if err != nil {
		logging.Error("could not get executable", err)
		return err
	}

	dir := filepath.Dir(exe)
	filename := filepath.Join(dir, ".path")

	err = os.WriteFile(filename, []byte(fmt.Sprintf("%s\n", p.Export())), 0644)

	if err != nil {
		logging.Error("could not write path to file", err)
		return err
	}

	logging.Info("wrote path to file", filename)

	return nil
}
