package shell

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/svenliebig/binary-organizer/internal/logging"
)

var env_file = ".env"
var path_file = ".path"

func init() {
	logging.Info("init shell package")
	envfile, err := getEnv()

	if err != nil {
		logging.Error("could not get env path", err)
		return
	}

	if _, err = os.Stat(envfile); !os.IsNotExist(err) {
		err = os.Remove(envfile)
		if err != nil {
			logging.Error("could not delete file", err)
		}
	}

	pathfile, err := getPath()

	if err != nil {
		logging.Error("could not get path path", err)
		return
	}

	if _, err = os.Stat(pathfile); !os.IsNotExist(err) {
		err = os.Remove(pathfile)
		if err != nil {
			logging.Error("could not delete file", err)
		}
	}
}

func WritePath(p Path) error {
	defer logging.Fn("shell.WritePath")()

	filename, err := getPath()

	if err != nil {
		return err
	}

	content := fmt.Sprintf("%s\n", p.Export())
	err = writeFile(filename, content)

	if err != nil {
		logging.Error("could not write path to file", err)
		return err
	}

	logging.Info("wrote path to file", filename)

	return nil
}

func WriteEnv(name, value string) error {
	defer logging.Fn("shell.AppendEnv")()

	filename, err := getEnv()

	if err != nil {
		return err
	}

	content := fmt.Sprintf("export %s=%s\n", name, value)
	err = appendToFile(filename, content)

	if err != nil {
		return err
	}

	logging.Info("wrote to file", filename)

	return nil
}

func writeFile(filename, content string) error {
	defer logging.Fn("shell.WriteFile")()
	return os.WriteFile(filename, []byte(content), 0644)
}

func appendToFile(filename, content string) error {
	defer logging.Fn("shell.AppendToFile")()

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		logging.Errorf("could not open file %s. %v", filename, err)
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)

	if err != nil {
		logging.Error("could not write to file", err)
		return err
	}

	logging.Info("wrote to file", filename)

	return nil
}

func getEnv() (string, error) {
	defer logging.Fn("shell.getEnv")()
	p, err := getProjectPath()

	if err != nil {
		return "", err
	}

	return filepath.Join(p, env_file), nil
}

func getPath() (string, error) {
	defer logging.Fn("shell.getPath")()
	p, err := getProjectPath()

	if err != nil {
		return "", err
	}

	return filepath.Join(p, path_file), nil
}

func getProjectPath() (string, error) {
	defer logging.Fn("shell.getProjectPath")()
	exe, err := os.Executable()

	if err != nil {
		logging.Error("could not get executable", err)
		return "", err
	}

	return filepath.Dir(exe), nil
}
