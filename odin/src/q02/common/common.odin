package common

import "core:strconv"
import "core:strings"

@(require_results)
parse :: proc(contents: string) -> (res: [dynamic][dynamic]int, ok: bool) {
	lines := strings.split(contents, "\n")
	defer delete(lines)

	for line in lines {
		if len(line) == 0 {continue}
		nums := parse_line(line) or_return
		append(&res, nums)
	}

	return res, true
}

@(require_results)
parse_line :: proc(line: string) -> (res: [dynamic]int, ok: bool) {
	num_strs := strings.split(line, " ")
	defer delete(num_strs)

	for num_str in num_strs {
		num := strconv.parse_int(num_str) or_return
		append(&res, num)
	}

	return res, true
}

is_unsafe :: proc(level: []int) -> bool {
	if len(level) < 2 {
		return true
	}

	first_diff := level[0] - level[1]
	if first_diff == 0 || abs(first_diff) > 3 {
		return true
	}

	is_asc := first_diff < 0 ? true : false

	for i := 1; i < len(level) - 1; i += 1 {
		curr := level[i]
		next := level[i + 1]
		diff := curr - next

		if (diff == 0) || (diff > 0 && is_asc) || (diff < 0 && !is_asc) || (abs(diff) > 3) {
			return true
		}
	}

	return false
}
