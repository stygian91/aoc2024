package cmd

import (
	"aoc2024/q01"

	"github.com/spf13/cobra"
)

var partArg *int

var q01Cmd = &cobra.Command{
	Use:   "q01",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if *partArg == 1 {
			q01.Part1()
		} else {
			q01.Part2()
		}
	},
}

func init() {
	partArg = q01Cmd.Flags().Int("part", 1, "")
	rootCmd.AddCommand(q01Cmd)
}
