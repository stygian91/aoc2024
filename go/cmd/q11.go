package cmd

import (
	"aoc2024/q11"

	"github.com/spf13/cobra"
)

var q11Cmd = &cobra.Command{
	Use:   "q11",
	Short: "",
	Long:  "",
}

func init() {
	aocCmd(q11Cmd, q11.Part1, q11.Part2)
	rootCmd.AddCommand(q11Cmd)
}
