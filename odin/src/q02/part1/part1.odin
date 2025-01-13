package part1

import cmnalloc "../../common/alloc"
import "../common"
import "core:fmt"
import "core:os"

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
		if !common.is_unsafe(levels[:]) {cnt += 1}
	}

	fmt.printfln("Answer: %d", cnt)
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
