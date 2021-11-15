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

func playerCanBeCaptured(goban s.Tgoban, x int, y int) bool {
	player := goban[y][x]
	opp := 2
	if player == 2 {
		opp = 1
	}

	leftUp := checkDiagLeftUpCapture(goban, x, y, s.Tnumber(player), s.Tnumber(opp))
	rightUp := checkDiagRightUpCapture(goban, x, y, s.Tnumber(player), s.Tnumber(opp))
	up := checkUpCapture(goban, x, y, s.Tnumber(player), s.Tnumber(opp))
	down := checkDownCapture(goban, x, y, s.Tnumber(player), s.Tnumber(opp))
	left := checkLeftCapture(goban, x, y, s.Tnumber(player), s.Tnumber(opp))
	right := checkRightCapture(goban, x, y, s.Tnumber(player), s.Tnumber(opp))
	leftDown := checkDiagLeftDownCapture(goban, x, y, s.Tnumber(player), s.Tnumber(opp))
	rightDown := checkDiagRightDownCapture(goban, x, y, s.Tnumber(player), s.Tnumber(opp))

	if leftUp != nil || rightUp != nil || up != nil || down != nil || left != nil || right != nil || leftDown != nil || rightDown != nil {
		return true
	}

	return false
}

func victoryCondition(goban s.Tgoban, captureP0 int, captureP1 int) (bool, s.Tnumber) {
	if isCapture {
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
					if isCapture && playerCanBeCaptured(goban, x, y) {
						return false, 0
					}
					return true, goban[y][x]
				}
			}
		}
	}

	return false, 0
}
