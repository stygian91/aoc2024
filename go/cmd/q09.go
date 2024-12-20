package cmd

import (
	"aoc2024/q09"

	"github.com/spf13/cobra"
)

var q09Cmd = &cobra.Command{
	Use:   "q09",
	Short: "",
	Long:  "",
}

func init() {
	aocCmd(q09Cmd, q09.Part1, q09.Part2)
	rootCmd.AddCommand(q09Cmd)
}
