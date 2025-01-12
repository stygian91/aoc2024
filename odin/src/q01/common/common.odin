package common

import "core:strconv"
import "core:strings"


Lists :: struct {
	left:  [dynamic]int,
	right: [dynamic]int,
}

NumberParseError :: struct {
	input: string,
}

ParseError :: union {
	NumberParseError,
}

@(require_results)
parse_int_err :: proc(input: string) -> (int, ParseError) {
	if res, ok := strconv.parse_int(input); ok {
		return res, nil
	}

	return 0, NumberParseError{input}
}

@(require_results)
parse :: proc(contents: string) -> (lists: Lists, err: ParseError) {
	lines := strings.split(string(contents), "\n")
	defer delete(lines)
	lists.left = [dynamic]int{}
	lists.right = [dynamic]int{}

	for line in lines {
		if len(line) == 0 {
			continue
		}

		nums := strings.split(line, "   ")
		defer delete(nums)
		a := parse_int_err(nums[0]) or_return
		b := parse_int_err(nums[1]) or_return

		append(&lists.left, a)
		append(&lists.right, b)
	}

	return lists, nil
}
