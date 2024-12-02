package cmd

import (
	"aoc2024/q01"

	"github.com/spf13/cobra"
)

var q01partArg *int

var q01Cmd = &cobra.Command{
	Use:   "q01",
	Short: "",
	Long:  ``,
}

func init() {
	aocCmd(q01Cmd, q01.Part1, q01.Part2)
	rootCmd.AddCommand(q01Cmd)
}
