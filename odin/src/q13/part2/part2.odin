package part2

import cmnalloc "../../common/alloc"
import cmn "../common"
import "core:fmt"
import "core:math/big"
import "core:os"
import "core:strconv"

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

	sum: big.Int
	big.atoi(&sum, "0")
	defer big.destroy(&sum)

	for &claw in claws {
		claw.Prize.x += 10000000000000
		claw.Prize.y += 10000000000000
		min_cost, ok := calc_min_cost(claw)
		defer big.destroy(&min_cost)

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

	fmt.printfln("Answer: %s", sumstr)
}

to_big :: proc(a: int) -> big.Int {
	buf := [30]u8{}

	str := strconv.itoa(buf[:], a)
	res: big.Int

	big.int_atoi(&res, str)
	return res
}

calc_min_cost :: proc(claw: cmn.Claw) -> (big.Int, bool) {
	tx := to_big(claw.Prize.x)
	ty := to_big(claw.Prize.y)
	xa := to_big(claw.A.x)
	ya := to_big(claw.A.y)
	xb := to_big(claw.B.x)
	yb := to_big(claw.B.y)

	tyxa, yatx, ybxa, yaxb: big.Int
	big.mul(&tyxa, &ty, &xa)
	big.mul(&yatx, &ya, &tx)
	big.mul(&ybxa, &yb, &xa)
	big.mul(&yaxb, &ya, &xb)

	btop, bbottom: big.Int
	big.sub(&btop, &tyxa, &yatx)
	big.sub(&bbottom, &ybxa, &yaxb)

	b: big.Int
	big.div(&b, &btop, &bbottom)

	a, atop, bxb: big.Int
	big.mul(&bxb, &b, &xb)
	big.sub(&atop, &tx, &bxb)
	big.div(&a, &atop, &xa)

	defer big.destroy(
		&tx,
		&ty,
		&xa,
		&ya,
		&xb,
		&yb,
		&tyxa,
		&yatx,
		&ybxa,
		&yaxb,
		&btop,
		&bbottom,
		&b,
		&a,
		&atop,
		&bxb,
	)

	actualX1, actualX2, actualY1, actualY2: big.Int
	defer big.destroy(&actualX1, &actualX2, &actualY1, &actualY2)

	// actualX = a*xa + b*xb
	big.mul(&actualX1, &a, &xa)
	big.mul(&actualX2, &b, &xb)
	big.add(&actualX1, &actualX1, &actualX2)

	// actualY = a*ya + b*yb
	big.mul(&actualY1, &a, &ya)
	big.mul(&actualY2, &b, &yb)
	big.add(&actualY1, &actualY1, &actualY2)

	prizex := to_big(claw.Prize.x)
	prizey := to_big(claw.Prize.y)
	defer big.destroy(&prizex, &prizey)

	xeq, xeqerr := big.eq(&prizex, &actualX1)
	yeq, yeqerr := big.eq(&prizey, &actualY1)
	if xeqerr != nil || yeqerr != nil || !xeq || !yeq {
		return big.Int{}, false
	}

	return calc_cost(&a, &b), true
}

calc_cost :: proc(a_count, b_count: ^big.Int) -> big.Int {
	a3, res: big.Int
	defer big.destroy(&a3)
	big.mul(&a3, a_count, 3)
	big.add(&res, &a3, b_count)
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
