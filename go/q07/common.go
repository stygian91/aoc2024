package q07

import (
	"strconv"
	"strings"
)

type Equation struct {
	Result   int
	Operands []int
}

type Operation interface {
	Calc(a, b int) int
}

type Add struct{}

func (this Add) Calc(a, b int) int { return a + b }
func (this Add) String() string    { return "Add" }

type Mul struct{}

func (this Mul) Calc(a, b int) int { return a * b }
func (this Mul) String() string    { return "Mul" }

func Parse(str string) ([]Equation, error) {
	eqs := []Equation{}

	for _, line := range strings.Split(str, "\n") {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ": ")
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			return []Equation{}, err
		}

		operands := []int{}
		for _, operandPart := range strings.Split(parts[1], " ") {
			op, err := strconv.Atoi(operandPart)
			if err != nil {
				return []Equation{}, err
			}
			operands = append(operands, op)
		}

		eqs = append(eqs, Equation{
			Result:   result,
			Operands: operands,
		})
	}

	return eqs, nil
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
