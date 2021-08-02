package game

import (
	s "gomoku/structures"
)

func checkRightUpDiagDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 1

	for case_x < int(ctx.NSize) && case_y >= 0 && count < 4 && piece < totalPiece {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
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
		case_x++
	}

	if lp == 1 && (case_x >= int(ctx.NSize) || case_y <= 0 || ctx.Goban[case_y][case_x] != s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x] != 0) {
		println("INFO: RU: obstrued")
		return -1
	}

	return piece
}

func checkLeftDownDiagDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 1

	for case_x >= 0 && case_y < int(ctx.NSize) && count < 4 && piece < totalPiece {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
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
		case_x--
	}

	if lp == 1 && (case_x <= 0 || case_y >= int(ctx.NSize) || ctx.Goban[case_y][case_x] != s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x] != 0) {
		println("INFO: LD: obstrued")
		return -1
	}

	return piece
}

func checkDiagRightPiece(ctx *s.SContext, case_x int, case_y int) bool {
	upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y, 3)
	downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y, 3)

	leftDoubleThree := checkLeftDoubleThree(ctx, case_x, case_y, 2)
	rightDoubleThree := checkRightDoubleThree(ctx, case_x, case_y, 2)

	leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y, 3)
	rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y, 3)

	if upDoubleThree >= 3 || downDoubleThree >= 3 || upDoubleThree >= 2 && downDoubleThree >= 2 ||
		leftDoubleThree >= 3 || rightDoubleThree >= 3 || leftDoubleThree >= 2 && rightDoubleThree >= 2 ||
		leftUpDiagDoubleThree >= 3 || rightDownDiagDoubleThree >= 3 || leftUpDiagDoubleThree >= 2 && rightDownDiagDoubleThree >= 2 {
		return false
	}

	return true
}

func loopDiagRightUpPiece(ctx *s.SContext, case_x int, case_y int) bool {
	piece := 0
	count := 0

	for case_x < int(ctx.NSize) && case_y >= 0 && count < 4 && piece < 2 {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			if !checkDiagRightPiece(ctx, case_x, case_y) {
				return false
			}
		} else if ctx.Goban[case_y][case_x] != 0 {
			return true
		}

		count++
		case_y--
		case_x++
	}

	return true
}

func loopDiagLeftDownpPiece(ctx *s.SContext, case_x int, case_y int) bool {
	piece := 0
	count := 0

	for case_x >= 0 && case_y < int(ctx.NSize) && count < 4 && piece < 2 {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			if !checkDiagRightPiece(ctx, case_x, case_y) {
				return false
			}
		} else if ctx.Goban[case_y][case_x] != 0 {
			return true
		}

		count++
		case_y++
		case_x--
	}

	return true
}

func checkDiagRight(ctx *s.SContext, case_x int, case_y int, leftDoubleThree int, rightDoubleThree int, upDoubleThree int, downDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if rightUpDiagDoubleThree >= 2 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 ||
			leftDoubleThree >= 1 && rightDoubleThree >= 1 || upDoubleThree >= 1 && downDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 ||
			!loopDiagRightUpPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if leftDownDiagDoubleThree >= 2 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 ||
			leftDoubleThree >= 1 && rightDoubleThree >= 1 || upDoubleThree >= 1 && downDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 ||
			!loopDiagLeftDownpPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 ||
			leftDoubleThree >= 1 && rightDoubleThree >= 1 || upDoubleThree >= 1 && downDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 ||
			!loopDiagRightUpPiece(ctx, case_x, case_y) || !loopDiagLeftDownpPiece(ctx, case_x, case_y) {
			return false
		}
	}

	return true
}
