package game

import (
	s "gomoku/structures"
)

func checkDiagLeftDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	if case_x-1 >= 0 && case_y-1 >= 0 && ctx.Goban[case_y-1][case_x-1] == s.Tnumber(ctx.CurrentPlayer) {
		if case_x-2 < 0 && case_y-2 < 0 || (case_x-2 >= 0 && case_y-2 >= 0 && isEnnemyCase(ctx, case_x-2, case_y-2)) {
			return -1
		}

		if case_x+1 < int(ctx.NSize) && case_y+1 < int(ctx.NSize) && ctx.Goban[case_y+1][case_x+1] == s.Tnumber(ctx.CurrentPlayer) {
			if case_x+2 >= int(ctx.NSize) && case_y+2 >= int(ctx.NSize) || (case_x+2 < int(ctx.NSize) && case_y+2 < int(ctx.NSize) && isEnnemyCase(ctx, case_x+2, case_y+2)) {
				return -1
			}
			return 1
		}

		if case_x+2 < int(ctx.NSize) && case_y+2 < int(ctx.NSize) && ctx.Goban[case_y+2][case_x+2] == s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y+1][case_x+1] == 0 {
			if case_x+3 >= int(ctx.NSize) && case_y+3 >= int(ctx.NSize) || (case_x+3 < int(ctx.NSize) && case_y+3 < int(ctx.NSize) && isEnnemyCase(ctx, case_x+3, case_y+3)) {
				return -1
			}
			return 1
		}
	}

	if case_x-2 >= 0 && case_y-2 >= 0 && ctx.Goban[case_y-2][case_x-2] == s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y-1][case_x-1] == 0 {
		if case_x-3 < 0 && case_y-3 < 0 || (case_x-3 >= 0 && case_y-3 >= 0 && isEnnemyCase(ctx, case_x-3, case_y-3)) {
			return -1
		}

		if case_x+1 < int(ctx.NSize) && case_y+1 < int(ctx.NSize) && ctx.Goban[case_y+1][case_x+1] == s.Tnumber(ctx.CurrentPlayer) {
			if case_x+2 >= int(ctx.NSize) && case_y+2 >= int(ctx.NSize) || (case_x+2 < int(ctx.NSize) && case_y+2 < int(ctx.NSize) && isEnnemyCase(ctx, case_x+2, case_y+2)) {
				return -1
			}
			return 1
		}
	}

	return 0
}

func checkLefUptDiagDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 0

	if case_x+1 >= int(ctx.NSize) && case_y+1 >= int(ctx.NSize) || case_x+1 < int(ctx.NSize) && case_y+1 < int(ctx.NSize) && isEnnemyCase(ctx, case_x+1, case_y+1) {
		return -1
	}

	for case_x >= 0 && case_y >= 0 && count < 4 && piece < totalPiece {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			lp = 1
		} else if ctx.Goban[case_y][case_x] != 0 {
			if lp == 1 {
				return -1
			} else {
				return piece
			}
		} else if count != 0 {
			lp = 0
		}

		count++
		case_y--
		case_x--
	}

	if lp == 1 && (case_x <= 0 || case_y <= 0 || ctx.Goban[case_y][case_x] != s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x] != 0) {
		println("INFO: LU: obstrued")
		return -1
	}

	return piece
}

func checkRightDownDiagDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 0

	if case_x-1 < 0 && case_y-1 < 0 || case_x-1 >= 0 && case_y-1 >= 0 && isEnnemyCase(ctx, case_x-1, case_y-1) {
		return -1
	}

	for case_x < int(ctx.NSize) && case_y < int(ctx.NSize) && count < 4 && piece < totalPiece {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			lp = 1
		} else if ctx.Goban[case_y][case_x] != 0 {
			if lp == 1 {
				return -1
			} else {
				return piece
			}
		} else if count != 0 {
			lp = 0
		}

		count++
		case_y++
		case_x++
	}

	if lp == 1 && (case_x >= int(ctx.NSize) || case_y >= int(ctx.NSize) || ctx.Goban[case_y][case_x] != s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x] != 0) {
		println("INFO: RD: obstrued")
		return -1
	}

	return piece
}

func checkDiagLeftPiece(ctx *s.SContext, case_x int, case_y int) bool {
	verticalDoubleThree := checkVerticalDoubleThree(ctx, case_x, case_y)
	upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y, 3)
	downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y, 3)

	horizonDoubleThree := checkHorizonDoubleThree(ctx, case_x, case_y)
	leftDoubleThree := checkLeftDoubleThree(ctx, case_x, case_y, 3)
	rightDoubleThree := checkRightDoubleThree(ctx, case_x, case_y, 3)

	diagRightDoubleThree := checkDiagRightDoubleThree(ctx, case_x, case_y)
	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y, 3)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y, 3)

	if upDoubleThree >= 3 || downDoubleThree >= 3 || verticalDoubleThree == 1 ||
		leftDoubleThree >= 3 || rightDoubleThree >= 3 || horizonDoubleThree == 1 ||
		leftDownDiagDoubleThree >= 3 || rightUpDiagDoubleThree >= 3 || diagRightDoubleThree == 1 {
		return false
	}

	return true
}

func loopDiagLeftUpPiece(ctx *s.SContext, case_x int, case_y int) bool {
	piece := 0
	count := 0

	for case_x >= 0 && case_y >= 0 && count < 4 && piece < 2 {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			if !checkDiagLeftPiece(ctx, case_x, case_y) {
				return false
			}
		} else if ctx.Goban[case_y][case_x] != 0 {
			return true
		}

		count++
		case_y--
		case_x--
	}

	return true
}

func loopDiagRightDownPiece(ctx *s.SContext, case_x int, case_y int) bool {
	piece := 0
	count := 0

	for case_x < int(ctx.NSize) && case_y < int(ctx.NSize) && count < 4 && piece < 2 {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			if !checkDiagLeftPiece(ctx, case_x, case_y) {
				return false
			}
		} else if ctx.Goban[case_y][case_x] != 0 {
			break
		}

		count++
		case_y++
		case_x++
	}

	return true
}

func checkDiagLeft(ctx *s.SContext, case_x int, case_y int, horizonDoubleThree int, leftDoubleThree int, rightDoubleThree int, verticalDoubleThree int, upDoubleThree int, downDoubleThree int, diagLeftDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, diagRightDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if leftUpDiagDoubleThree >= 2 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			horizonDoubleThree == 1 || verticalDoubleThree == 1 || diagRightDoubleThree == 1 ||
			!loopDiagLeftUpPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if rightDownDiagDoubleThree >= 2 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			horizonDoubleThree == 1 || verticalDoubleThree == 1 || diagRightDoubleThree == 1 ||
			!loopDiagRightDownPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if diagLeftDoubleThree == 1 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			horizonDoubleThree == 1 || verticalDoubleThree == 1 || diagRightDoubleThree == 1 ||
			!loopDiagLeftUpPiece(ctx, case_x, case_y) || !loopDiagRightDownPiece(ctx, case_x, case_y) {
			return false
		}
	}

	return true
}
