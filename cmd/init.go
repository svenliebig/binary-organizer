/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes 👻 bo(o) and set up your environment with the defaults.",
	Long: `Initializes 👻 bo(o) and set up your environment with the defaults, if
you don't have a configuration file yet, it will create a configuration file
in your ~/.config/boo directory. Then the configuration file will be used
to set up your $PATH variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")

		// TODO find or create a configuration file in:
		// - ~/.boo.toml
		// - ~/.config/boo.toml
		// - ~/.config/boo/boo.toml

		// TODO read path

		// TODO apply configuration

		// TODO write path

		// p := path.NewPathVariable()
		// err = os.WriteFile(".path", []byte(fmt.Sprintf("%s\n", p.Export())), 0644)
		// if err != nil {
		// 	fmt.Println("Error writing the file")
		// 	return
		// }
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
