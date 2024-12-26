package cmd

import (
	"aoc2024/q12"

	"github.com/spf13/cobra"
)

var q12Cmd = &cobra.Command{
	Use:   "q12",
	Short: "",
	Long:  "",
}

func init() {
	aocCmd(q12Cmd, q12.Part1, q12.Part2)
	rootCmd.AddCommand(q12Cmd)
}
