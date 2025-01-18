package common

import "core:fmt"
import "core:os"
import "core:strings"
import "core:strconv"

Vec2 :: struct {
	x, y: int,
}

Bot :: struct {
	pos, vel: Vec2,
}

parse :: proc(contents: string) -> [dynamic]Bot {
	bots := [dynamic]Bot{}
	lines := strings.split_lines(contents)
	defer delete(lines)

	for line in lines {
		if len(line) == 0 {continue}

		fcom_idx := strings.index(line, ",")
		if fcom_idx == -1 {
			fmt.eprintf("Error parsing line '%s' - did not find first comma.", line)
			os.exit(1)
		}

		sp_idx := strings.index(line, " ")
		if sp_idx == -1 {
			fmt.eprintf("Error parsing line '%s' - no space found.", line)
			os.exit(1)
		}

		bot : Bot
		bot.pos.x = strconv.atoi(line[2:fcom_idx])
		bot.pos.y = strconv.atoi(line[fcom_idx+1:sp_idx])

		lcom_idx := strings.last_index(line, ",")
		if lcom_idx == -1 {
			fmt.eprintf("Error parsing line '%s' - did not find last comma.", line)
			os.exit(1)
		}

		bot.vel.x = strconv.atoi(line[sp_idx+3:lcom_idx])
		bot.vel.y = strconv.atoi(line[lcom_idx+1:])
		append(&bots, bot)
	}

	return bots
}
