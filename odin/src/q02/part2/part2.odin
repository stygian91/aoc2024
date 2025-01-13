package part2

import cmnalloc "../../common/alloc"
import "../common"
import "core:fmt"
import "core:os"
import "core:slice"

process_file :: proc(path: string) {
	contents, read_err := os.read_entire_file_or_err(path)
	defer delete(contents)
	if read_err != nil {
		fmt.eprintfln("Error while reading file %s", path)
		return
	}

	level_lists, parse_ok := common.parse(string(contents))
	defer {
		for levels in level_lists {
			delete(levels)
		}

		delete(level_lists)
	}
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

run :: proc() {
	if len(os.args) != 2 {
		fmt.eprintln("Expected exactly 1 argument: the path of the input file.")
		return
	}

	process_file(os.args[1])
}

main :: proc () {
	cmnalloc.track_leaks_for_proc(run)
}
