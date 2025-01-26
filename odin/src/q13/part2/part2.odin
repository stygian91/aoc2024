package part2

import cmnalloc "../../common/alloc"
import cmn "../common"
import "core:fmt"
import "core:math/big"
import "core:mem"
import "core:os"
import "core:strconv"
import "core:time"

MAX_INT :: (1 << (size_of(int) * 8 - 1)) - 1

process_file :: proc(path: string) {
	defer free_all(context.temp_allocator)

	sw: time.Stopwatch
	time.stopwatch_start(&sw)

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

	sum: big.Int
	big.atoi(&sum, "0")
	defer big.destroy(&sum)

	for &claw in claws {
		claw.Prize.x += 10000000000000
		claw.Prize.y += 10000000000000
		min_cost, ok := calc_min_cost(claw)

		if ok == false {
			continue
		}

		big.add(&sum, &sum, &min_cost)
	}

	sumstr, err := big.itoa(&sum)
	defer delete(sumstr)

	if err != nil {
		fmt.printfln("Err: %s", err)
		return
	}

	time.stopwatch_stop(&sw)
	duration := time.stopwatch_duration(sw)

	fmt.printfln("Answer: %s", sumstr)
	fmt.printfln("Duration: %f us", time.duration_microseconds(duration))
}

to_big :: proc(a: int, allocator: mem.Allocator) -> big.Int {
	res: big.Int
	// big.int_add_digit(&res, &res, cast(big.DIGIT)a, allocator)
	big.int_set_from_integer(&res, a, allocator = allocator)
	return res
}

calc_min_cost :: proc(claw: cmn.Claw, allocator := context.temp_allocator) -> (big.Int, bool) {
	tx := to_big(claw.Prize.x, allocator)
	ty := to_big(claw.Prize.y, allocator)
	xa := to_big(claw.A.x, allocator)
	ya := to_big(claw.A.y, allocator)
	xb := to_big(claw.B.x, allocator)
	yb := to_big(claw.B.y, allocator)

	tyxa, yatx, ybxa, yaxb: big.Int
	big.mul(&tyxa, &ty, &xa, allocator)
	big.mul(&yatx, &ya, &tx, allocator)
	big.mul(&ybxa, &yb, &xa, allocator)
	big.mul(&yaxb, &ya, &xb, allocator)

	btop, bbottom: big.Int
	big.sub(&btop, &tyxa, &yatx, allocator)
	big.sub(&bbottom, &ybxa, &yaxb, allocator)

	b: big.Int
	big.div(&b, &btop, &bbottom, allocator)

	a, atop, bxb: big.Int
	big.mul(&bxb, &b, &xb, allocator)
	big.sub(&atop, &tx, &bxb, allocator)
	big.div(&a, &atop, &xa, allocator)

	actualX1, actualX2, actualY1, actualY2: big.Int

	big.mul(&actualX1, &a, &xa, allocator)
	big.mul(&actualX2, &b, &xb, allocator)
	big.add(&actualX1, &actualX1, &actualX2, allocator)

	big.mul(&actualY1, &a, &ya, allocator)
	big.mul(&actualY2, &b, &yb, allocator)
	big.add(&actualY1, &actualY1, &actualY2, allocator)

	prizex := to_big(claw.Prize.x, allocator)
	prizey := to_big(claw.Prize.y, allocator)

	xeq, xeqerr := big.eq(&prizex, &actualX1, allocator)
	yeq, yeqerr := big.eq(&prizey, &actualY1, allocator)
	if xeqerr != nil || yeqerr != nil || !xeq || !yeq {
		return big.Int{}, false
	}

	return calc_cost(&a, &b, allocator), true
}

calc_cost :: proc(a_count, b_count: ^big.Int, allocator: mem.Allocator) -> big.Int {
	a3, res: big.Int
	big.mul(&a3, a_count, 3, allocator)
	big.add(&res, &a3, b_count, allocator)
	return res
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
