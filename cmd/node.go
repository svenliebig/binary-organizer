/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/svenliebig/binary-organizer/internal/binaries"
	"github.com/svenliebig/binary-organizer/internal/service"
	"github.com/svenliebig/binary-organizer/internal/shell"
)

// nodeCmd represents the node command
var (
	nodeCmd = &cobra.Command{
		Use:         "node",
		Short:       "Manage node versions",
		Args:        cobra.MaximumNArgs(1),
		Annotations: annotationSourceTrue,
		RunE: func(cmd *cobra.Command, args []string) error {
			nodeBinary, err := binaries.Get("node")

			if err != nil {
				// 😱 An error occurred that you shouldn't encounter. We're sorry! Please run the command again with --debug to get more information and report the issue on GitHub.
				// --debug should also print the platform and arch.
				return fmt.Errorf("could not get binary: %w", err)
			}

			if len(args) == 0 {
				fmt.Print("👻 bo(o) is wondering 💭 what it can do(o) for you?\n\n")
				cmd.Help()
				return nil
			}

			s, err := service.New(nodeBinary)

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

			fmt.Printf("\n👻 bo(o) is trying to select %s v%s for you 💪\n\n", nodeBinary.Identifier(), version.String())

			binpath, ok := s.IsInstalled(version)

			if !ok {
				fmt.Printf("😨 %s v%s is not installed yet. Try the command 'boo %s list' to list all available versions of the binary.\n\n", nodeBinary.Identifier(), version.String(), nodeBinary.Identifier())

				// TODO implement installation
				// if err := nodeBinary.Install(context.Background(), *version); err != nil {
				// 	return fmt.Errorf("could not install %s v%s: %w", nodeBinary.Identifier(), version.String(), err)
				// }

				return nil
			}

			p := shell.NewPath()

			// remove other node versions
			for _, pth := range p.Find(func(p string) bool {
				_, ok := nodeBinary.Matches(p)
				return ok
			}) {
				p.Remove(pth)
			}

			p.Add(binpath)

			err = shell.WritePath(p)

			if err != nil {
				return fmt.Errorf("could not write path: %w", err)
			}

			fmt.Println("✅ bo(o) has set up your environment with the selected node version 🎉")

			return nil
		},
	}
	nodeListCmd = &cobra.Command{
		Use:         "list",
		Short:       "Lists all installed node versions",
		Annotations: annotationSourceFalse,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("\n👻 bo(o) is looking for installed node versions 🧐\n\n")

			nodeBinary, err := binaries.Get("node")

			if err != nil {
				return fmt.Errorf("could not get binary: %w", err)
			}

			s, err := service.New(nodeBinary)

			if err != nil {
				return fmt.Errorf("could not create service: %w", err)
			}

			versions, err := s.Versions()

			// TODO handle boo. Errors
			if err != nil {
				return fmt.Errorf("could not get versions: %w", err)
			}

			if len(versions) == 0 {
				fmt.Println("🫣 No node versions found.")
				return nil
			}

			p := shell.NewPath()
			var selected binaries.Version

			nodePaths := p.Find(func(p string) bool {
				_, ok := nodeBinary.Matches(p)
				return ok
			})

			if len(nodePaths) > 0 {
				v, _ := nodeBinary.Matches(nodePaths[len(nodePaths)-1])
				selected = v
			}

			fmt.Printf("🔎 Found the following node versions:\n\n")
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
)

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(nodeListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
