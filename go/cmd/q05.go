package cmd

import (
	"aoc2024/q05"

	"github.com/spf13/cobra"
)

var q05Cmd = &cobra.Command{
	Use:   "q05",
	Short: "",
	Long:  "",
}

func init() {
	aocCmd(q05Cmd, q05.Part1, q05.Part2)
	rootCmd.AddCommand(q05Cmd)
}
