package common

import "core:strconv"
import "core:strings"


Lists :: struct {
	left:  [dynamic]int,
	right: [dynamic]int,
}

ParseError :: enum {
	None,
	NumberParseError,
}

parse_int_err :: proc(input: string) -> (int, ParseError) {
	if res, ok := strconv.parse_int(input); ok {
		return res, .None
	}

	return 0, .NumberParseError
}

parse :: proc(contents: string) -> (lists: Lists, err: ParseError) {
	lines := strings.split(string(contents), "\n")
	lists.left = [dynamic]int{}
	lists.right = [dynamic]int{}

	for line in lines {
		if len(line) == 0 {
			continue
		}

		nums := strings.split(line, "   ")
		a := parse_int_err(nums[0]) or_return
		b := parse_int_err(nums[1]) or_return

		append(&lists.left, a)
		append(&lists.right, b)
	}

	return lists, .None
}
