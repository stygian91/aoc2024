package main

import "core:fmt"
import "core:os"
import "core:slice"
import "core:strconv"
import "core:strings"

part1 :: proc() {
	// contents, err := os.read_entire_file_or_err("./demo.txt")
	contents, err := os.read_entire_file_or_err("./main.txt")
	if err != nil {
		fmt.eprintfln("Could not read file: %s", err)
	}

	lines := strings.split(string(contents), "\n")
	left := [dynamic]int{}
	right := [dynamic]int{}

	for line in lines {
		if len(line) == 0 {
			continue
		}

		nums := strings.split(line, "   ")
		a, oka := strconv.parse_int(nums[0])
		if !oka {
			fmt.eprintfln("Error parsing %s", nums[0])
			return
		}
		b, okb := strconv.parse_int(nums[1])
		if !okb {
			fmt.eprintfln("Error parsing %s", nums[1])
			return
		}

		append(&left, a)
		append(&right, b)
	}

	slice.sort(left[:])
	slice.sort(right[:])

	sum := 0
	for i := 0; i < len(left); i += 1 {
		diff := abs(left[i] - right[i])
		sum += diff
	}

	fmt.printfln("Part 1: %d", sum)
}

main :: proc() {
	part1()
}
