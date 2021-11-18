package doubleThree

import (
	s "gomoku/structures"
)

func checkDiagLeftUpCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) bool {
	if isInRange(case_x-1, case_y-1) && isInRange(case_x-2, case_y-2) && isInRange(case_x-3, case_y-3) && goban[case_y-3][case_x-3] == player && goban[case_y-2][case_x-2] == opp && goban[case_y-1][case_x-1] == opp {
		return true
	}
	return false
}

func checkDiagRightUpCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) bool {
	if isInRange(case_x+1, case_y-1) && isInRange(case_x+2, case_y-2) && isInRange(case_x+3, case_y-3) && goban[case_y-3][case_x+3] == player && goban[case_y-2][case_x+2] == opp && goban[case_y-1][case_x+1] == opp {
		return true
	}
	return false
}

func checkUpCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) bool {
	if isInRange(case_x, case_y-1) && isInRange(case_x, case_y-2) && isInRange(case_x, case_y-3) && goban[case_y-3][case_x] == player && goban[case_y-2][case_x] == opp && goban[case_y-1][case_x] == opp {
		return true
	}
	return false
}

func checkDownCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) bool {
	if isInRange(case_x, case_y+1) && isInRange(case_x, case_y+2) && isInRange(case_x, case_y+3) && goban[case_y+3][case_x] == player && goban[case_y+2][case_x] == opp && goban[case_y+1][case_x] == opp {
		return true
	}
	return false
}

func checkLeftCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) bool {
	if isInRange(case_x-1, case_y) && isInRange(case_x-2, case_y) && isInRange(case_x-3, case_y) && goban[case_y][case_x-3] == player && goban[case_y][case_x-2] == opp && goban[case_y][case_x-1] == opp {
		return true
	}
	return false
}

func checkRightCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) bool {
	if isInRange(case_x+1, case_y) && isInRange(case_x+2, case_y) && isInRange(case_x+3, case_y) && goban[case_y][case_x+3] == player && goban[case_y][case_x+2] == opp && goban[case_y][case_x+1] == opp {
		return true
	}
	return false
}

func checkDiagLeftDownCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) bool {
	if isInRange(case_x-1, case_y+1) && isInRange(case_x-2, case_y+2) && isInRange(case_x-3, case_y+3) && goban[case_y+3][case_x-3] == player && goban[case_y+2][case_x-2] == opp && goban[case_y+1][case_x-1] == opp {
		return true
	}
	return false
}

func checkDiagRightDownCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) bool {
	if isInRange(case_x+1, case_y+1) && isInRange(case_x+2, case_y+2) && isInRange(case_x+3, case_y+3) && goban[case_y+3][case_x+3] == player && goban[case_y+2][case_x+2] == opp && goban[case_y+1][case_x+1] == opp {
		return true
	}
	return false
}

func isCapture(goban s.Tgoban, case_x int, case_y int, player uint8) bool {
	var opp uint8 = 1

	if player == 1 {
		opp = 2
	}

	if checkDiagLeftUpCapture(goban, case_x, case_y, s.Tnumber(player), s.Tnumber(opp)) || checkDiagRightUpCapture(goban, case_x, case_y, s.Tnumber(player), s.Tnumber(opp)) ||
		checkUpCapture(goban, case_x, case_y, s.Tnumber(player), s.Tnumber(opp)) || checkDownCapture(goban, case_x, case_y, s.Tnumber(player), s.Tnumber(opp)) || 
		checkLeftCapture(goban, case_x, case_y, s.Tnumber(player), s.Tnumber(opp)) || checkRightCapture(goban, case_x, case_y, s.Tnumber(player), s.Tnumber(opp)) ||
		checkDiagLeftDownCapture(goban, case_x, case_y, s.Tnumber(player), s.Tnumber(opp)) || checkDiagRightDownCapture(goban, case_x, case_y, s.Tnumber(player), s.Tnumber(opp)) {
			return true
	}
	return false
}