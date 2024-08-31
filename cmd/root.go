/*
Copyright © 2024 Sven Liebig
*/
package cmd

import (
	"os"

	_ "github.com/svenliebig/binary-organizer/internal/binaries/node"
	"github.com/svenliebig/binary-organizer/internal/logging"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "boo",
	Short: "👻 bo(o) is a tool to manage your binary file pathes",
	Long: `👻 bo(o) is a tool to manage different versions of your binary files
in your $PATH variable. It helps you to switch between different versions,
for example if you need a different version of node or python for a project.

To achieve that, bo(o) reads the content of the $PATH variable and overrides
it with the desired pathes.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("debug", cmd.Flags().Lookup("debug"))

		if viper.GetBool("debug") {
			logging.SetLevel(logging.DebugLevel)
			logging.Info("🐞 debug mode enabled")
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("👻 bo(o) is wondering 💭 what it can do(o) for you?")

		// TODO here we want to have some prompts instead of help
		cmd.Help()

		return nil
	},
}

func Execute() {
	c, err := rootCmd.ExecuteC()

	if err != nil {
		logging.Error("could not execute command", err)
		os.Exit(1)
	}

	if c != nil {
		if shouldSource(c) {
			os.Exit(20)
		} else {
			os.Exit(21)
		}
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debug output")
}
