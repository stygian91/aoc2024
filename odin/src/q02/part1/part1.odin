package part1

import "../common"
import "core:fmt"
import "core:os"

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
		if !common.is_unsafe(levels[:]) {cnt += 1}
	}

  fmt.printfln("Answer: %d", cnt)
}

main :: proc() {
	if len(os.args) != 2 {
		fmt.eprintln("Expected exactly 1 argument: the path of the input file.")
		return
	}

	run(os.args[1])
}
