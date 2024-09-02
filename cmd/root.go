/*
Copyright © 2024 Sven Liebig
*/
package cmd

import (
	"errors"
	"os"

	_ "github.com/svenliebig/binary-organizer/internal/binaries/node"
	"github.com/svenliebig/binary-organizer/internal/boo"
	"github.com/svenliebig/binary-organizer/internal/config"
	"github.com/svenliebig/binary-organizer/internal/logging"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boo",
	Short: "👻 bo(o) is a tool to manage your binary file pathes",
	Long: `👻 bo(o) is a tool to manage different versions of your binary files
in your $PATH variable. It helps you to switch between different versions,
for example if you need a different version of node or python for a project.

To achieve that, bo(o) reads the content of the $PATH variable and overrides
it with the desired pathes.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")

		if err != nil {
			logging.Error("could not get debug flag", err)
			return err
		}

		if debug {
			logging.SetLevel(logging.DebugLevel)
			logging.Info("🐞 debug mode enabled")
		}

		trace, err := cmd.Flags().GetBool("trace")

		if err != nil {
			logging.Error("could not get trace flag", err)
			return err
		}

		if trace {
			logging.SetLevel(logging.TraceLevel)
			logging.Info("🔎 trace mode enabled")
		}

		silent, err := cmd.Flags().GetBool("silent")
		boo.Silent(silent)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		boo.Intro("is wondering 💭 what it can do(o) for you?")
		cmd.Help()
		return nil
	},
}

func Execute() {
	c, err := rootCmd.ExecuteC()

	if err != nil {
		handleExecutionError(err)
		os.Exit(1)
	}

	if c != nil {
		if shouldSource(c) {
			logging.Info("sending exit code to source the .path")
			os.Exit(20)
		} else {
			logging.Info("sending exit code to not source the .path")
			os.Exit(21)
		}
	}
}

func handleExecutionError(err error) {
	if errors.Is(err, boo.ErrBinaryDirIsFile) {
		c, _ := config.Load()
		boo.Bodyf("🚫 the configured binary directory (%q) is a file", c.BinaryRoot)
		return
	}

	if errors.Is(err, boo.ErrBinaryDirNotExists) {
		c, _ := config.Load()
		boo.Bodyf("🚫 the configured binary directory (%q) does not exist", c.BinaryRoot)
		return
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("trace", "t", false, "enable trace output")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debug output")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "suppress cli output")
}
