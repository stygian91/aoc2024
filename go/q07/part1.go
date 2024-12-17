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
		if hasPossibleCorrectPart1(eq) {
			answer += eq.Result
		}
	}

	fmt.Printf("Part 1 answer: %d\n", answer)
}

func hasPossibleCorrectPart1(eq Equation) bool {
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

func ProcessOperands(operands []int, operations []Operation) int {
	res := operands[0]

	for i := 1; i < len(operands); i++ {
		res = operations[i-1].Calc(res, operands[i])
	}

	return res
}

var perms [][]Operation

func GenerateOperationPermutations(count uint) [][]Operation {
	if perms == nil {
		perms = [][]Operation{}
	}

	if uint(len(perms)) >= count {
		return perms[0:count]
	}

	for i := uint(len(perms)); i < count; i++ {
		perms = append(perms, PermutationFromInt(i, count))
	}

	return perms[0:count]
}

func PermutationFromInt(seed, count uint) []Operation {
	res := []Operation{}
	w := seed

	for i := uint(0); i < count; i++ {
		if w&1 == 0 {
			res = append(res, Add{})
		} else {
			res = append(res, Mul{})
		}

		w = w >> 1
	}

	return res
}
