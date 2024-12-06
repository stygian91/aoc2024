package cmd

import (
	"aoc2024/q06"

	"github.com/spf13/cobra"
)

var q06Cmd = &cobra.Command{
	Use:   "q06",
	Short: "",
	Long:  "",
}

func init() {
	aocCmd(q06Cmd, q06.Part1, q06.Part2)
	rootCmd.AddCommand(q06Cmd)
}
