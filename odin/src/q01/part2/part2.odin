package main

import "../common"
import "core:fmt"
import "core:os"
import "core:slice"

count_unique :: proc(list: []int) -> map[int]int {
	res := make(map[int]int)
	for el in list {
		res[el] += 1
	}

	return res
}

run :: proc(path: string) {
	contents, err := os.read_entire_file_or_err(path)
	if err != nil {
		fmt.eprintfln("Could not read file: %s", err)
	}

	lists, parseErr := common.parse(string(contents))
	defer delete(lists.left)
	defer delete(lists.right)

	switch e in parseErr {
	case common.NumberParseError:
		fmt.eprintfln("Error while parsing input: %s", e.input)
		return
	}

	sum := 0
	right_counts := count_unique(lists.right[:])
	defer delete(right_counts)

	for num in lists.left {
		sum += num * right_counts[num]
	}

	fmt.printfln("Part 2: %d", sum)
}

main :: proc() {
	if len(os.args) != 2 {
		fmt.eprintln("Expected exactly 1 argument: the path of the input file.")
		return
	}

	run(os.args[1])
}