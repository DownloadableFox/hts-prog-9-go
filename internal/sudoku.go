package internal

import (
	"errors"
	"fmt"
)

type Sudoku struct {
	Board [9][9]int
}

func SudukoFromString(s string) (Sudoku, error) {
	var current int
	var sudoku Sudoku

	for _, c := range s {
		if current >= 81 {
			return sudoku, errors.New("too many characters in input string")
		}

		if c == ',' {
			current++
		} else {
			sudoku.Board[current/9][current%9] = int(c - '0')
		}
	}

	return sudoku, nil
}

func (s *Sudoku) String() string {
	var str string

	for r, row := range s.Board {
		for c, cell := range row {
			str += fmt.Sprintf("%d", cell)

			if c*9+r < 80 {
				str += ","
			}
		}
	}

	return str
}

func (s *Sudoku) Solve(solved chan<- Sudoku) {
	var stack Stack[Sudoku]
	var dp = make(map[Sudoku]bool)
	stack.Push(*s)

	for !stack.Empty() {
		current := stack.Pop()

		if current.IsSolved() {
			if _, ok := dp[*current]; ok {
				continue
			}

			dp[*current] = true
			solved <- *current
			continue
		}

		var empty_cells []struct {
			values []int
			row    int
			col    int
		}

		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if current.Get(row, col) == 0 {
					var values []int

					for i := 1; i <= 9; i++ {
						if current.IsLegal(row, col, i) {
							values = append(values, i)
						}
					}

					empty_cells = append(empty_cells, struct {
						values []int
						row    int
						col    int
					}{values, row, col})
				}
			}
		}

		// sort empty_cells by the number of possible values
		for i := 0; i < len(empty_cells); i++ {
			for j := i + 1; j < len(empty_cells); j++ {
				if len(empty_cells[i].values) < len(empty_cells[j].values) {
					empty_cells[i], empty_cells[j] = empty_cells[j], empty_cells[i]
				}
			}
		}

		for _, cell := range empty_cells {
			for _, value := range cell.values {
				new_sudoku := *current
				new_sudoku.Set(cell.row, cell.col, value)
				stack.Push(new_sudoku)
			}
		}
	}

	fmt.Println("no more solutions")
	close(solved)
}

func (s *Sudoku) IsSolved() bool {
	for _, row := range s.Board {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}

	return true
}

func (s *Sudoku) IsLegal(row, col, value int) bool {
	for i := 0; i < 9; i++ {
		if s.Board[row][i] == value || s.Board[i][col] == value {
			return false
		}
	}

	startRow := row - row%3
	startCol := col - col%3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s.Board[startRow+i][startCol+j] == value {
				return false
			}
		}
	}

	return true
}

func (s *Sudoku) Get(row, col int) int {
	return s.Board[row][col]
}

func (s *Sudoku) Set(row, col, value int) {
	s.Board[row][col] = value
}
