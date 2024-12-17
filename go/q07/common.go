package q07

import (
	"fmt"
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

type Concat struct{}

func (this Concat) String() string { return "Concat" }
func (this Concat) Calc(a, b int) int {
	v, e := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if e != nil {
		panic(e)
	}

	return v
}

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
