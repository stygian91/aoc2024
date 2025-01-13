package main

import cmnalloc "../../common/alloc"
import "../common"
import "core:fmt"
import "core:mem"
import "core:os"
import "core:slice"

process_file :: proc(path: string) {
	contents, err := os.read_entire_file_or_err(path)
	defer delete(contents)
	if err != nil {
		fmt.eprintfln("Could not read file: %s", err)
		return
	}

	lists, parseErr := common.parse(string(contents))
	defer delete(lists.left)
	defer delete(lists.right)
	switch e in parseErr {
	case common.NumberParseError:
		fmt.eprintfln("Error while parsing input: %s", e.input)
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

run :: proc() {
	if len(os.args) != 2 {
		fmt.eprintln("Expected exactly 1 argument: the path of the input file.")
		return
	}

	process_file(os.args[1])
}

main :: proc() {
	cmnalloc.track_leaks_for_proc(run)
}
