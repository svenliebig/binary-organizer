/*
Copyright © 2024 Sven Liebig
*/
package cmd

import (
	"fmt"
	"os"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("👻 bo(o) is wondering 💭 what it can do(o) for you?")

		// TODO here we want to have some prompts instead of help
		cmd.Help()

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}

func init() {
}
