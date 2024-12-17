package q07

import (
	"aoc2024/inputs"
	"fmt"
	"math"
)

func Part1() {
	// content, err := inputs.GetInputFile("q07/demo.txt")
	content, err := inputs.GetInputFile("q07/main.txt")
	if err != nil {
		panic(err)
	}

	eqs, err := Parse(content)
	if err != nil {
		panic(err)
	}

	answer := 0

	for _, eq := range eqs {
		if !hasPossibleCorrect(eq) {
			continue
		}

		answer += eq.Result
	}

	fmt.Printf("Part 1 answer: %d\n", answer)
}

func hasPossibleCorrect(eq Equation) bool {
	permCount := math.Pow(2, float64(len(eq.Operands)-1))
	permutations := GenerateOperationPermutations(uint(permCount))

	for _, permutation := range permutations {
		ans := ProcessOperands(eq.Operands, permutation)
		if ans == eq.Result {
			return true
		}
	}

	return false
}
