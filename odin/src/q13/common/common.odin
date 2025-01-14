package common

import "core:fmt"
import "core:strconv"
import "core:strings"

Vector2 :: struct {
	x: int,
	y: int,
}

Claw :: struct {
	A:     Vector2,
	B:     Vector2,
	Prize: Vector2,
}

@(require_results)
parse :: proc(contents: string) -> (res: [dynamic]Claw, ok: bool) {
	lines := strings.split(contents, "\n")
	defer delete(lines)
	claw_count := (len(lines) + 1) / 4

	for i := 0; i < claw_count; i += 1 {
		offset := i * 4
		claw: Claw
		claw.A = parse_button(lines[offset]) or_return
		claw.B = parse_button(lines[offset + 1]) or_return
		claw.Prize = parse_prize(lines[offset + 2]) or_return
		append(&res, claw)
	}

	return res, true
}

@(require_results)
parse_button :: proc(line: string) -> (res: Vector2, ok: bool) {
	x_idx := strings.index(line, "X+")
	x_space := strings.last_index(line, ", ")
	if x_idx == -1 || x_space == -1 {
		return
	}

	x_sub := strings.substring(line, x_idx + 2, x_space) or_return
	y_sub := strings.substring_from(line, x_space + 4) or_return

	res.x = strconv.parse_int(x_sub) or_return
	res.y = strconv.parse_int(y_sub) or_return

	return res, true
}

@(require_results)
parse_prize :: proc(line: string) -> (res: Vector2, ok: bool) {
	x_idx := strings.index(line, "X=")
	x_space := strings.last_index(line, ", ")
	if x_idx == -1 || x_space == -1 {return}

	x_sub := strings.substring(line, x_idx + 2, x_space) or_return
	y_sub := strings.substring_from(line, x_space + 4) or_return

	res.x = strconv.parse_int(x_sub) or_return
	res.y = strconv.parse_int(y_sub) or_return

	return res, true
}
