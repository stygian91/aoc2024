package q09

import (
	"aoc2024/common/data"
	"slices"
	"strconv"
	"strings"
)

type Block struct {
	Id, Start, End int

	IsFree bool
}

func (this Block) Size() int {
	return data.Absint(this.End - this.Start)
}

func (this Block) String() string {
	if this.IsFree {
		return strings.Repeat(".", this.Size())
	}

	return strings.Repeat(strconv.Itoa(this.Id), this.Size())
}

func ExpandSpace(str string) ([]Block, error) {
	id, cursor := 0, 0
	blocks := []Block{}

	for i, r := range strings.TrimSpace(str) {
		cnt, err := strconv.Atoi(string(r))

		if i%2 == 0 {
			if err != nil {
				return []Block{}, err
			}

			blocks = append(blocks, Block{
				Id:     id,
				Start:  cursor,
				End:    cursor + cnt,
				IsFree: false,
			})
			id++
		} else {
			if err != nil {
				return []Block{}, err
			}
			blocks = append(blocks, Block{
				Start:  cursor,
				End:    cursor + cnt,
				IsFree: true,
			})
		}

		cursor += cnt
	}

	return blocks, nil
}

func CleanupBlocks(blocks []Block) []Block {
	res := blocks

	if len(res) == 1 && res[0].Start == res[0].End {
		return []Block{}
	}

	if len(res) < 2 {
		return res
	}

	for i := len(res) - 2; i >= 0; i-- {
		if res[i+1].Start == res[i+1].End {
			res = slices.Delete(res, i+1, i+2)
			continue
		}

		if res[i].IsFree == res[i+1].IsFree && (res[i].IsFree || res[i].Id == res[i+1].Id) {
			res[i].End = res[i+1].End
			res = slices.Delete(res, i+1, i+2)
		}
	}

	return res
}
