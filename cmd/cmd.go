/*
Copyright © 2024 Sven Liebig
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/svenliebig/binary-organizer/internal/binaries"
	"github.com/svenliebig/binary-organizer/internal/boo"
	"github.com/svenliebig/binary-organizer/internal/service"
	"github.com/svenliebig/binary-organizer/internal/shell"
	"github.com/svenliebig/seq"
)

func createCommand(identifier string) *cobra.Command {
	root := &cobra.Command{
		Use:         identifier,
		Short:       "Manage " + identifier + " versions",
		Args:        cobra.MaximumNArgs(1),
		Annotations: annotationSourceTrue,
		RunE: func(cmd *cobra.Command, args []string) error {
			binary, err := binaries.Get(identifier)

			if err != nil {
				// 😱 An error occurred that you shouldn't encounter. We're sorry! Please run the command again with --debug to get more information and report the issue on GitHub.
				// --debug should also print the platform and arch.
				return fmt.Errorf("could not get binary: %w", err)
			}

			if len(args) == 0 {
				boo.Intro("is wondering 💭 what it can do(o) for you?")
				cmd.Help()
				return nil
			}

			s, err := service.New(binary)

			if err != nil {
				return fmt.Errorf("could not create service: %w", err)
			}

			version, err := binaries.VersionFrom(args[0])

			if err != nil {
				fmt.Printf("\n👻 bo(o) has issues to understand the version you provided 🤯\n\n")
				fmt.Printf("🚫 You provided %q, but normally bo(o) expects versions like:\n\n", args[0])
				fmt.Println("  - 14.17.0")
				fmt.Println("  - v16.0")
				fmt.Println("  - 17")
				fmt.Printf("\n👻 bo(o) is sorry for the inconvenience 😔\n\n")

				return nil
			}

			fmt.Printf("\n👻 bo(o) is trying to select %s v%s for you 💪\n\n", binary.Identifier(), version.String())

			err = s.SetVersion(version)

			if errors.Is(err, boo.ErrVersionNotInstalled) {
				fmt.Printf("😨 %s v%s is not installed yet. Try the command 'boo %s list' to list all available versions of the binary.\n\n", binary.Identifier(), version.String(), binary.Identifier())

				return err
			}

			if err != nil {
				return fmt.Errorf("could not write path: %w", err)
			}

			fmt.Printf("\n✅ bo(o) has set up your environment with the selected %s version 🎉", identifier)

			return nil
		},
	}

	list := &cobra.Command{
		Use:         "list",
		Short:       fmt.Sprintf("Lists all installed %s versions", identifier),
		Annotations: annotationSourceFalse,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("\n👻 bo(o) is looking for installed %s versions 🧐\n\n", identifier)

			binary, err := binaries.Get(identifier)

			if err != nil {
				return fmt.Errorf("could not get binary: %w", err)
			}

			s, err := service.New(binary)

			if err != nil {
				return fmt.Errorf("could not create service: %w", err)
			}

			versions, err := s.Versions()

			// TODO handle boo. Errors
			if err != nil {
				return fmt.Errorf("could not get versions: %w", err)
			}

			if len(versions) == 0 {
				fmt.Printf("\n🫣 No %s versions found.", identifier)
				return nil
			}

			p := shell.NewPath()
			var selected binaries.Version

			paths := p.Find(func(p string) bool {
				_, ok := binary.Matches(p)
				return ok
			})

			if len(paths) > 0 {
				v, _ := binary.Matches(paths[len(paths)-1])
				selected = v
			}

			fmt.Printf("🔎 Found the following %s versions:\n\n", identifier)
			for _, v := range versions {
				if v.Matches(selected) {
					fmt.Println("👉", v, "(selected)")
					continue
				}

				fmt.Println("  ", v)
			}

			return nil
		},
	}

	root.AddCommand(list)
	return root
}

func init() {
	cmds := seq.Map(
		seq.Map(
			binaries.All(),
			func(b binaries.Binary) (string, error) {
				return b.Identifier(), nil
			},
		),
		func(s string) (*cobra.Command, error) {
			return createCommand(s), nil
		},
	)

	for cmd, _ := range cmds.Iterator() {
		rootCmd.AddCommand(cmd)
	}
}
