package part2

import cmnalloc "../../common/alloc"
import cmn "../common"
import "core:fmt"
import "core:os"
import "core:strconv"
import "core:math/big"

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
	for &claw in claws {
		claw.Prize.x += 10000000000000
		claw.Prize.y += 10000000000000
		min_cost := calc_min_cost(claw)
		if min_cost == MAX_INT {
			continue
		}

		sum += min_cost
	}

	fmt.printfln("Answer: %d", sum)
}

to_big :: proc(a: int) -> big.Int {
	buf := [20]u8{}

	str := strconv.itoa(buf[:], a)
	res : big.Int

	big.int_atoi(&res, str)
	return res
}

calc_min_cost :: proc(claw: cmn.Claw) -> int {
	// b :=
	// 	(claw.Prize.y * claw.A.x - claw.A.y * claw.Prize.x) /
	// 	(claw.B.y * claw.A.x - claw.A.y * claw.B.x)
	// a := (claw.Prize.x - b * claw.B.x) / claw.A.x
	//
	// fmt.printfln("a=%d; b=%d", a, b)

	tx := to_big(claw.Prize.x)
	ty := to_big(claw.Prize.y)
	xa := to_big(claw.A.x)
	ya := to_big(claw.A.y)
	xb := to_big(claw.B.x)
	yb := to_big(claw.B.y)

	tyxa, yatx, ybxa, yaxb : big.Int
	big.int_mul(&tyxa, &ty, &xa)
	big.int_mul(&yatx, &ya, &tx)
	big.int_mul(&ybxa, &yb, &xa)
	big.int_mul(&yaxb, &ya, &xb)

	btop, bbottom : big.Int
	big.int_sub(&btop, &tyxa, &yatx)
	big.int_sub(&bbottom, &ybxa, &yaxb)

	b : big.Int
	big.int_div(&b, &btop, &bbottom)

	a, atop, bxb : big.Int
	big.int_mul(&bxb, &b, &xb)
	big.int_sub(&atop, &tx, &bxb)
	big.int_div(&a, &atop, &xa)

	astr, aerr := big.int_to_string(&a)
	bstr, berr := big.int_to_string(&b)

	if aerr != nil || berr != nil {
		return MAX_INT
	}

	aint := strconv.atoi(astr)
	bint := strconv.atoi(bstr)

	return calc_cost(aint, bint)
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
