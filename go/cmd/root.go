package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc2024",
	Short: "",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func aocCmd(cmd *cobra.Command, part1Fn func(), part2Fn func()) {
	var partArg *int
	partArg = cmd.Flags().Int("part", 1, "")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if *partArg == 1 {
			part1Fn()
		} else {
			part2Fn()
		}
	}
}
