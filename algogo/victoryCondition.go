package algogo

import (
	s "gomoku/structures"
)

func checkVictory(goban s.Tgoban, player s.Tnumber, x int, y int, i int, j int) int {
	count := 0
	for k := 0; k < 5 && y >= 0 && y < 19 && x >= 0 && x < 19 && goban[y][x] == player; k++ {
		x += i
		y += j
		count += 1
	}

	return count
}

func checkHorizontalVictory(goban s.Tgoban, x int, y int) bool {
	if checkVictory(goban, goban[y][x], x, y, 1, 0)+checkVictory(goban, goban[y][x], x, y, -1, 0) > 5 {
		return true
	}

	return false
}

func checkVerticalVictory(goban s.Tgoban, x int, y int) bool {
	if checkVictory(goban, goban[y][x], x, y, 0, 1)+checkVictory(goban, goban[y][x], x, y, 0, -1) > 5 {
		return true
	}

	return false
}

func checkDiagVictory(goban s.Tgoban, x int, y int) bool {
	if checkVictory(goban, goban[y][x], x, y, 1, 1)+checkVictory(goban, goban[y][x], x, y, -1, -1) > 5 || checkVictory(goban, goban[y][x], x, y, -1, 1)+checkVictory(goban, goban[y][x], x, y, 1, -1) > 5 {
		return true
	}

	return false
}

func CheckAlignVictory(goban s.Tgoban, x int, y int) bool {
	if checkHorizontalVictory(goban, x, y) || checkVerticalVictory(goban, x, y) || checkDiagVictory(goban, x, y) {
		return true
	}
	return false
}

func victoryCondition(goban s.Tgoban, capture bool, captureP0 int, captureP1 int) (bool, s.Tnumber) {
	if capture {
		if captureP0 >= 5 {
			return true, 1
		} else if captureP1 >= 5 {
			return true, 2
		}
	}

	for y := range goban {
		for x := range goban[y] {
			if goban[y][x] != 0 {
				if CheckAlignVictory(goban, x, y) {
					return true, goban[y][x]
				}
			}
		}
	}

	return false, 0
}
