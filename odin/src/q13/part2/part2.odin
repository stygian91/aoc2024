package part2

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

		// sum += min_cost
	}

	fmt.printfln("Answer: %d", sum)
}

calc_min_cost :: proc(claw: cmn.Claw) -> int {
	b :=
		(claw.Prize.y * claw.A.x - claw.A.y * claw.Prize.x) /
		(claw.B.y * claw.A.x - claw.A.y * claw.B.x)
	a := (claw.Prize.x - b * claw.B.x) / claw.A.x

	fmt.printfln("a=%d; b=%d", a, b)

	return MAX_INT
}

calc_cost :: proc(a_count, b_count: int) -> int {
	return a_count * 3 + b_count
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
