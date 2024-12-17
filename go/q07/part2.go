package q07

import (
	"aoc2024/inputs"
	"fmt"
)

func Part2() {
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
		if hasPossibleCorrectPart2(eq) {
			answer += eq.Result
		}
	}

	fmt.Printf("Part 2 answer: %d\n", answer)
}

func hasPossibleCorrectPart2(eq Equation) bool {
	return checkRec(eq.Operands[0], eq, Add{}, 1) || checkRec(eq.Operands[0], eq, Mul{}, 1) || checkRec(eq.Operands[0], eq, Concat{}, 1)
}

func checkRec(result int, eq Equation, operation Operation, idx int) bool {
	newRes := operation.Calc(result, eq.Operands[idx])
	if newRes == eq.Result && idx == len(eq.Operands)-1 {
		return true
	}

	if newRes > eq.Result || idx >= len(eq.Operands)-1 {
		return false
	}

	newIdx := idx + 1

	return checkRec(newRes, eq, Add{}, newIdx) || checkRec(newRes, eq, Mul{}, newIdx) || checkRec(newRes, eq, Concat{}, newIdx)
}
