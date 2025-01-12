package main

import "../common"
import "core:fmt"
import "core:os"
import "core:slice"

run :: proc(path: string) {
	contents, err := os.read_entire_file_or_err(path)
	if err != nil {
		fmt.eprintfln("Could not read file: %s", err)
	}

	lists, parseErr := common.parse(string(contents))
	if parseErr == common.ParseError.NumberParseError {
		fmt.eprintfln("Error parsing")
		return
	}

	slice.sort(lists.left[:])
	slice.sort(lists.right[:])

	sum := 0
	for i := 0; i < len(lists.left); i += 1 {
		diff := abs(lists.left[i] - lists.right[i])
		sum += diff
	}

	fmt.printfln("Part 1: %d", sum)
}

main :: proc() {
	if len(os.args) != 2 {
		fmt.eprintln("Expected exactly 1 argument: the path of the input file.")
		return
	}

	run(os.args[1])
}
