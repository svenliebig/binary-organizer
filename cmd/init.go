/*
Copyright © 2024 Sven Liebig
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/svenliebig/binary-organizer/internal/config"
	"github.com/svenliebig/binary-organizer/internal/shell"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"initialize"},
	Short:   "Initializes 👻 bo(o) and set up your environment with the defaults",
	Long: `Initializes 👻 bo(o) and set up your environment with the defaults, if
you don't have a configuration file yet, it will create a configuration file
in your ~/.config/boo directory. Then the configuration file will be used
to set up your $PATH variable.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("init called")

		_, err := config.Load()

		if err != nil {
			return fmt.Errorf("could not load configuration: %w", err)
		}

		// read path
		pth := shell.NewPath()

		// TODO apply configuration

		// write path
		// TODO change .path to constant?
		err = os.WriteFile(".path", []byte(fmt.Sprintf("%s\n", pth.Export())), 0644)

		if err != nil {
			return fmt.Errorf("could not write path: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
