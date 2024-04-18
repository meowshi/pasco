package types

import (
	"fmt"
	"strconv"
)

type Cell struct {
	Column string
	Row    int
}

func NewCell(column string, row int) Cell {
	return Cell{
		Column: column,
		Row:    row,
	}
}

func NewCellFromString(cell string) Cell {
	row, _ := strconv.Atoi(cell[1:])
	return NewCell(string(cell[0]), row)
}

func (c *Cell) AddColumn(n int) *Cell {
	col := []byte(c.Column)

	for {
		if int(col[len(col)-1])+n <= int('z') {
			col[len(col)-1] += byte(n)
			break
		}

		n = n - int(('z' - col[len(col)-1])) - 1
		if len(col) < 2 {
			col = append([]byte{'a'}, col...)
			continue
		}
		for i := len(col) - 2; i >= 0; i-- {
			if col[i]+1 > 'z' {
				col[i] = 'a'
				if i == 0 {
					col = append([]byte{'a'}, col...)
					break
				}
				continue
			}
			col[i] += 1
			break
		}
	}

	c.Column = string(col)
	return c
}

func (c *Cell) AddRow(n int) *Cell {
	c.Row += n

	return c
}

func (c *Cell) ToString() string {
	return fmt.Sprintf("%s%d", c.Column, c.Row)
}
