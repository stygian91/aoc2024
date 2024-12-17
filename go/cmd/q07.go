package cmd

import (
	"aoc2024/q07"

	"github.com/spf13/cobra"
)

var q07Cmd = &cobra.Command{
	Use:   "q07",
	Short: "",
	Long:  "",
}

func init() {
	aocCmd(q07Cmd, q07.Part1, q07.Part2)
	rootCmd.AddCommand(q07Cmd)
}
