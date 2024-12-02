package cmd

import (
	"aoc2024/q02"

	"github.com/spf13/cobra"
)

var q02partArg *int

var q02Cmd = &cobra.Command{
	Use:   "q02",
	Short: "",
	Long: ``,

	Run: func(cmd *cobra.Command, args []string) {
		if *q02partArg == 1 {
			q02.Part1()
		} else {
			q02.Part2()
		}
	},
}

func init() {
	q02partArg = q02Cmd.Flags().Int("part", 1, "")
	rootCmd.AddCommand(q02Cmd)
}
