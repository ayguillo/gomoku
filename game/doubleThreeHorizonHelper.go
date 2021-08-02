package game

import (
	s "gomoku/structures"
)

func checkLeftDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 1

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
	upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y, 3)
	downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y, 3)

	leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y, 3)
	rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y, 3)

	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y, 3)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y, 3)

	if upDoubleThree >= 3 || downDoubleThree >= 3 || upDoubleThree >= 2 && downDoubleThree >= 2 ||
		leftUpDiagDoubleThree >= 3 || rightDownDiagDoubleThree >= 3 || leftUpDiagDoubleThree >= 2 && rightDownDiagDoubleThree >= 2 ||
		leftDownDiagDoubleThree >= 3 || rightUpDiagDoubleThree >= 3 || leftDownDiagDoubleThree >= 2 && rightUpDiagDoubleThree >= 2 {
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

func checkHorizon(ctx *s.SContext, case_x int, case_y int, leftDoubleThree int, rightDoubleThree int, upDoubleThree int, downDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if leftDoubleThree >= 2 {
		if (upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			upDoubleThree >= 1 && downDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) ||
			!loopHorizonLeftPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if rightDoubleThree >= 2 {
		if upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			upDoubleThree >= 1 && downDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1 ||
			!loopHorizonRightPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if leftDoubleThree >= 1 && rightDoubleThree >= 1 {
		if upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			upDoubleThree >= 1 && downDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1 ||
			!loopHorizonLeftPiece(ctx, case_x, case_y) || !loopHorizonRightPiece(ctx, case_x, case_y) {
			return false
		}
	}

	return true
}
