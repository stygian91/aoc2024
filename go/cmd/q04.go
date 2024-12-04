package cmd

import (
	"aoc2024/q04"

	"github.com/spf13/cobra"
)

var q04Cmd = &cobra.Command{
	Use:   "q04",
	Short: "",
	Long:  "",
}

func init() {
	aocCmd(q04Cmd, q04.Part1, q04.Part2)
	rootCmd.AddCommand(q04Cmd)
}
