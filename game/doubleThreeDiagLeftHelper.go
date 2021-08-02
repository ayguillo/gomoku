package game

import (
	s "gomoku/structures"
)

func checkLefUptDiagDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 0

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
	upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y, 3)
	downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y, 3)

	leftDoubleThree := checkLeftDoubleThree(ctx, case_x, case_y, 2)
	rightDoubleThree := checkRightDoubleThree(ctx, case_x, case_y, 2)

	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y, 3)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y, 3)

	if upDoubleThree >= 3 || downDoubleThree >= 3 || upDoubleThree >= 2 && downDoubleThree >= 2 ||
		leftDoubleThree >= 3 || rightDoubleThree >= 3 || leftDoubleThree >= 2 && rightDoubleThree >= 2 ||
		leftDownDiagDoubleThree >= 3 || rightUpDiagDoubleThree >= 3 || leftDownDiagDoubleThree >= 2 && rightUpDiagDoubleThree >= 2 {
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

func checkDiagLeft(ctx *s.SContext, case_x int, case_y int, leftDoubleThree int, rightDoubleThree int, upDoubleThree int, downDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if leftUpDiagDoubleThree >= 2 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			leftDoubleThree >= 1 && rightDoubleThree >= 1 || upDoubleThree >= 1 && downDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1 ||
			!loopDiagLeftUpPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if rightDownDiagDoubleThree >= 2 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			leftDoubleThree >= 1 && rightDoubleThree >= 1 || upDoubleThree >= 1 && downDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1 ||
			!loopDiagRightDownPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			leftDoubleThree >= 1 && rightDoubleThree >= 1 || upDoubleThree >= 1 && downDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1 ||
			!loopDiagLeftUpPiece(ctx, case_x, case_y) || !loopDiagRightDownPiece(ctx, case_x, case_y) {
			return false
		}
	}

	return true
}
