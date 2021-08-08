package game

import (
	s "gomoku/structures"
)

func checkHorizonDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	if case_x-1 >= 0 && ctx.Goban[case_y][case_x-1] == s.Tnumber(ctx.CurrentPlayer) {
		if case_x-2 < 0 || (case_x-2 >= 0 && isEnnemyCase(ctx, case_x-2, case_y)) {
			return -1
		}

		if case_x+1 < int(ctx.NSize) && ctx.Goban[case_y][case_x+1] == s.Tnumber(ctx.CurrentPlayer) {
			if case_x+2 >= int(ctx.NSize) || (case_x+2 < int(ctx.NSize) && isEnnemyCase(ctx, case_x+2, case_y)) {
				return -1
			}
			return 1
		}

		if case_x+2 < int(ctx.NSize) && ctx.Goban[case_y][case_x+2] == s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x+1] == 0 {
			if case_x+3 >= int(ctx.NSize) || (case_x+3 < int(ctx.NSize) && isEnnemyCase(ctx, case_x+3, case_y)) {
				return -1
			}
			return 1
		}
	}

	if case_x-2 >= 0 && ctx.Goban[case_y][case_x-2] == s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x-1] == 0 {
		if case_x-3 < 0 || (case_x-3 >= 0 && isEnnemyCase(ctx, case_x-3, case_y)) {
			return -1
		}

		if case_x+1 < int(ctx.NSize) && ctx.Goban[case_y][case_x+1] == s.Tnumber(ctx.CurrentPlayer) {
			if case_x+2 >= int(ctx.NSize) || (case_x+2 < int(ctx.NSize) && isEnnemyCase(ctx, case_x+2, case_y)) {
				return -1
			}
			return 1
		}
	}

	return 0
}

func checkLeftDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 1

	if case_x+1 >= int(ctx.NSize) || case_x+1 < int(ctx.NSize) && isEnnemyCase(ctx, case_x+1, case_y) {
		return -1
	}

	for case_x >= 0 && count < 4 && piece < totalPiece {
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
		case_x--
	}

	if lp == 1 && (case_x <= 0 || ctx.Goban[case_y][case_x] != s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x] != 0) {
		println("INFO: R: obstrued")
		return -1
	}

	return piece
}

func checkRightDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 1

	if case_x-1 < 0 || case_x-1 >= 0 && isEnnemyCase(ctx, case_x-1, case_y) {
		return -1
	}

	for case_x < int(ctx.NSize) && count < 4 && piece < totalPiece {
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
		case_x++
	}

	if lp == 1 && (case_x >= int(ctx.NSize) || ctx.Goban[case_y][case_x] != s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x] != 0) {
		println("INFO: R: obstrued")
		return -1
	}

	return piece
}

func checkHorizonPiece(ctx *s.SContext, case_x int, case_y int) bool {
	verticalDoubleThree := checkVerticalDoubleThree(ctx, case_x, case_y)
	upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y, 3)
	downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y, 3)

	diagLeftDoubleThree := checkDiagLeftDoubleThree(ctx, case_x, case_y)
	leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y, 3)
	rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y, 3)

	diagRightDoubleThree := checkDiagRightDoubleThree(ctx, case_x, case_y)
	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y, 3)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y, 3)

	if upDoubleThree >= 3 || downDoubleThree >= 3 || verticalDoubleThree == 1 ||
		leftUpDiagDoubleThree >= 3 || rightDownDiagDoubleThree >= 3 || diagLeftDoubleThree == 1 ||
		leftDownDiagDoubleThree >= 3 || rightUpDiagDoubleThree >= 3 || diagRightDoubleThree == 1 {
		return false
	}

	return true
}

func loopHorizonLeftPiece(ctx *s.SContext, case_x int, case_y int) bool {
	piece := 0
	count := 0

	for case_x >= 0 && count < 4 && piece < 2 {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			if !checkHorizonPiece(ctx, case_x, case_y) {
				return false
			}
		} else if ctx.Goban[case_y][case_x] != 0 {
			return true
		}

		count++
		case_x--
	}

	return true
}

func loopHorizonRightPiece(ctx *s.SContext, case_x int, case_y int) bool {
	piece := 0
	count := 0

	for case_x < int(ctx.NSize) && count < 4 && piece < 2 {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			if !checkHorizonPiece(ctx, case_x, case_y) {
				return false
			}
		} else if ctx.Goban[case_y][case_x] != 0 {
			return true
		}

		count++
		case_x++
	}

	return true
}

func checkHorizon(ctx *s.SContext, case_x int, case_y int, horizonDoubleThree int, leftDoubleThree int, rightDoubleThree int, verticalDoubleThree int, upDoubleThree int, downDoubleThree int, diagLeftDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, diagRightDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if leftDoubleThree >= 2 {
		println("je suis la")
		if upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			verticalDoubleThree == 1 || diagLeftDoubleThree == 1 || diagRightDoubleThree == 1 ||
			!loopHorizonLeftPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if rightDoubleThree >= 2 {
		if upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			verticalDoubleThree == 1 || diagLeftDoubleThree == 1 || diagRightDoubleThree == 1 ||
			!loopHorizonRightPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if horizonDoubleThree == 1 {
		if upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			verticalDoubleThree == 1 || diagLeftDoubleThree == 1 || diagRightDoubleThree == 1 ||
			!loopHorizonLeftPiece(ctx, case_x, case_y) || !loopHorizonRightPiece(ctx, case_x, case_y) {
			return false
		}
	}

	return true
}
