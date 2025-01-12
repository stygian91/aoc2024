package part2

import "../common"
import "core:fmt"
import "core:os"
import "core:slice"

run :: proc(path: string) {
	contents, read_err := os.read_entire_file_or_err(path)
	if read_err != nil {
		fmt.eprintfln("Error while reading file %s", path)
		return
	}

	level_lists, parse_ok := common.parse(string(contents))
	defer delete(level_lists)
	if !parse_ok {
		fmt.eprintln("Error while parsing")
		return
	}

	cnt := 0
	for levels in level_lists {
		unsafe := common.is_unsafe(levels[:])
		if unsafe {unsafe = try_removing_one(levels[:])}
		if !unsafe {cnt += 1}
	}

	fmt.printfln("Answer: %d", cnt)
}

try_removing_one :: proc(level: []int) -> bool {
	for i := 0; i < len(level); i += 1 {
		spliced := slice.concatenate([][]int{level[0:i], level[i + 1:]})
		defer delete(spliced)
		if !common.is_unsafe(spliced) {
			return false
		}
	}

	return true
}

main :: proc() {
	if len(os.args) != 2 {
		fmt.eprintln("Expected exactly 1 argument: the path of the input file.")
		return
	}

	run(os.args[1])
}