/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/svenliebig/binary-organizer/internal/binaries"
	"github.com/svenliebig/binary-organizer/internal/config"
	"github.com/svenliebig/binary-organizer/internal/shell"
)

// nodeCmd represents the node command
var (
	nodeCmd = &cobra.Command{
		Use:   "node",
		Short: "Manage node versions",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("node called")
		},
	}
	nodeListCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all installed node versions",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("\n👻 bo(o) is looking for installed node versions 🧐\n\n")

			config, err := config.Load()

			if err != nil {
				return fmt.Errorf("could not load configuration: %w", err)
			}

			// TODO test what happens if the directory does not exist

			entries, err := os.ReadDir(path.Join(config.BinaryRoot, "node"))

			if err != nil {
				return fmt.Errorf("could not read directory: %w", err)
			}

			nodeBinary, err := binaries.Get("node")

			if err != nil {
				return fmt.Errorf("could not get binary: %w", err)
			}

			versions := make([]binaries.Version, 0, len(entries))
			for _, entry := range entries {
				if !entry.IsDir() {
					continue
				}

				v, ok := nodeBinary.Matches(entry.Name())

				if ok {
					versions = append(versions, v)
				}
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
