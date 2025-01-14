package part1

import cmnalloc "../../common/alloc"
import cmn "../common"
import "core:fmt"
import "core:os"

MAX_INT :: (1 << (size_of(int) * 8 - 1)) - 1

process_file :: proc(path: string) {
	contents, read_err := os.read_entire_file_or_err(path)
	defer delete(contents)
	if read_err != nil {
		fmt.eprintfln("Error while reading file %s", path)
		return
	}

	claws, parse_ok := cmn.parse(string(contents))
	defer delete(claws)
	if !parse_ok {
		fmt.eprintln("Parsing error")
		return
	}

	sum := 0
	for claw in claws {
		min_cost := calc_min_cost(claw)
		if min_cost == MAX_INT {
			continue
		}

		sum += min_cost
	}

	fmt.printfln("Answer: %d", sum)
}

calc_min_cost :: proc(claw: cmn.Claw) -> int {
	min_cost := MAX_INT
	tablea, tableb := precompute_buttons(claw)
	defer delete(tablea)
	defer delete(tableb)

	for a, i in tablea {
		b, j := find_match(a, claw.Prize, &tableb)
		if j == -1 {continue}

		cost := calc_cost(i, j)
		if cost < min_cost {
			min_cost = cost
		}
	}

	return min_cost
}

find_match :: proc(
	a: cmn.Vector2,
	target: cmn.Vector2,
	tableb: ^[dynamic]cmn.Vector2,
) -> (
	cmn.Vector2,
	int,
) {
	diff := cmn.Vector2{target.x - a.x, target.y - a.y}
	if diff.x < 0 || diff.y < 0 {
		return cmn.Vector2{}, -1
	}

	for b, i in tableb {
		if b == diff {
			return b, i
		}
	}

	return cmn.Vector2{}, -1
}

calc_cost :: proc(a_count, b_count: int) -> int {
	return a_count * 3 + b_count
}

precompute_buttons :: proc(claw: cmn.Claw) -> ([dynamic]cmn.Vector2, [dynamic]cmn.Vector2) {
	sumsa := [dynamic]cmn.Vector2{}
	sumsb := [dynamic]cmn.Vector2{}
	sa := cmn.Vector2{}
	sb := cmn.Vector2{}

	for sa.x < claw.Prize.x && sa.y < claw.Prize.y {
		append(&sumsa, sa)
		sa = cmn.Vector2{sa.x + claw.A.x, sa.y + claw.A.y}
	}

	for sb.x < claw.Prize.x && sb.y < claw.Prize.y {
		append(&sumsb, sb)
		sb = cmn.Vector2{sb.x + claw.B.x, sb.y + claw.B.y}
	}

	return sumsa, sumsb
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
