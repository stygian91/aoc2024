package q07_test

import (
	"aoc2024/q07"
	"testing"
)

func TestPermutationFromInt(t *testing.T) {
	checkPermutationSeed(5, 3, []q07.Operation{
		q07.Mul{},
		q07.Add{},
		q07.Mul{},
	}, t)

	checkPermutationSeed(4, 3, []q07.Operation{
		q07.Add{},
		q07.Add{},
		q07.Mul{},
	}, t)

	checkPermutationSeed(15, 4, []q07.Operation{
		q07.Mul{},
		q07.Mul{},
		q07.Mul{},
		q07.Mul{},
	}, t)

	checkPermutationSeed(0, 4, []q07.Operation{
		q07.Add{},
		q07.Add{},
		q07.Add{},
		q07.Add{},
	}, t)
}

func checkPermutationSeed(seed, count uint, expected []q07.Operation, t *testing.T) {
	res := q07.PermutationFromInt(seed, count)

	if len(res) != len(expected) {
		t.Errorf("Permutation len error: expected %d, got %d\n", len(expected), len(res))
		return
	}

	for i := uint(0); i < count; i++ {
		preal := isAdd(res[i])
		pexp := isAdd(expected[i])

		if preal != pexp {
			t.Errorf("Permutation mismatch: seed: %d, count: %d, expected: %s, real: %d\n", seed, count, expected[i], res[i])
		}
	}
}

func isAdd(op q07.Operation) bool {
	_, ok := op.(q07.Add)
	return ok
}
