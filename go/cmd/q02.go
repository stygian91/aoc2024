package cmd

import (
	"aoc2024/q02"

	"github.com/spf13/cobra"
)

var q02Cmd = &cobra.Command{
	Use:   "q02",
	Short: "",
	Long:  ``,
}

func init() {
	aocCmd(q02Cmd, q02.Part1, q02.Part2)
	rootCmd.AddCommand(q02Cmd)
}
