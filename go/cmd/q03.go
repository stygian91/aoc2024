package cmd

import (
	"aoc2024/q03"

	"github.com/spf13/cobra"
)

var q03Cmd = &cobra.Command{
	Use:   "q03",
	Short: "",
	Long:  ``,
}

func init() {
	aocCmd(q03Cmd, q03.Part1, q03.Part2)
	rootCmd.AddCommand(q03Cmd)
}
