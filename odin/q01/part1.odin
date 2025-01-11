package main

import "core:fmt"
import "core:os"
import "core:slice"

part1 :: proc() {
	// contents, err := os.read_entire_file_or_err("./demo.txt")
	contents, err := os.read_entire_file_or_err("./main.txt")
	if err != nil {
		fmt.eprintfln("Could not read file: %s", err)
	}

	lists, parseErr := parse(string(contents))
	if parseErr == .NumberParseError {
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
