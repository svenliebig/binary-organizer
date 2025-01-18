/*
Copyright © 2024 Sven Liebig
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/svenliebig/binary-organizer/internal/binaries"
	"github.com/svenliebig/binary-organizer/internal/boo"
	"github.com/svenliebig/binary-organizer/internal/logging"
	"github.com/svenliebig/binary-organizer/internal/service"
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
	Annotations:   annotationSourceTrue,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		defer logging.Fn("initCmd.RunE")()

		boo.Intro("is initializing the default versions of your binaries...📦")

		binarySeq := binaries.All()
		p := shell.NewPath()

		for b := range binarySeq.Iterator() {
			s, err := service.New(b)

			if err != nil {
				logging.Error("err while trying to create a service", err)
				return err
			}

			v, err := s.GetDefaultVersion()

			if err != nil {
				if errors.Is(err, boo.ErrNoDefaultVersion) {
					boo.Bodyf("  ⚠️ %s has no default version set", b.Identifier())
					continue
				} else {
					return err
				}
			}

			err = s.SetVersion(v, p)

			if err != nil {
				if errors.Is(err, boo.ErrVersionNotInstalled) {
					boo.Bodyf("  ⚠️ %s version %s is not installed and can't be set", b.Identifier(), v.String())
				} else {
					return err
				}
			} else {
				boo.Bodyf("  ✏️ using %s in version %s", b.Identifier(), v.String())
			}
		}

		err := shell.WritePath(p)

		if err != nil {
			return err
		}

		boo.Outro("has setup your environment 🎉\n\nTip: to supress this message, you can use the --silent (or -s) flag like this: \n\n  boo init -s")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
