package part1

import cmnalloc "../../common/alloc"
import "../common"
import "core:fmt"
import "core:os"

// demo values:
// WIDTH :: 11
// HEIGHT :: 7
// main values
WIDTH :: 101
HEIGHT :: 103

ITER :: 100
MIDDLE_X :: WIDTH / 2
MIDDLE_Y :: HEIGHT / 2

process_file :: proc(path: string) {
	contents, read_err := os.read_entire_file_or_err(path)
	defer delete(contents)
	if read_err != nil {
		fmt.eprintfln("Error while reading file %s", path)
		return
	}

	bots := common.parse(string(contents))
	defer delete(bots)

	counts := [4]int{}
	for &bot in bots {
		bot.pos = bot_pos_after(bot, ITER)
		quadrant := bot_get_quadrant(bot) or_continue
		counts[quadrant] += 1
	}

	answer := counts[0] * counts[1] * counts[2] * counts[3]

	fmt.printfln("Answer: %d", answer)
}

bot_pos_after :: proc(bot: common.Bot, n: int) -> common.Vec2 {
	x := (bot.pos.x + (bot.vel.x * n)) % WIDTH
	y := (bot.pos.y + (bot.vel.y * n)) % HEIGHT
	if x < 0 {x = WIDTH + x}
	if y < 0 {y = HEIGHT + y}
	return common.Vec2{x, y}
}

bot_get_quadrant :: proc(bot: common.Bot) -> (res: u8, in_quadrant: bool) {
	if bot.pos.x == MIDDLE_X || bot.pos.y == MIDDLE_Y {return}

	if bot.pos.x < MIDDLE_X {
		if bot.pos.y < MIDDLE_Y {return 0, true} else {return 2, true}
	} else {
		if bot.pos.y < MIDDLE_Y {return 1, true} else {return 3, true}
	}
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
