package cmd

import (
	"aoc2024/q10"

	"github.com/spf13/cobra"
)

var q10Cmd = &cobra.Command{
	Use:   "q10",
	Short: "",
	Long:  "",
}

func init() {
	aocCmd(q10Cmd, q10.Part1, q10.Part2)
	rootCmd.AddCommand(q10Cmd)
}
