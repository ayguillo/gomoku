package game

import (
	s "gomoku/structures"
)

func checkUpDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 1

	for case_y >= 0 && count < 4 && piece < totalPiece {
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
	}

	if lp == 1 && (case_y <= 0 || ctx.Goban[case_y][case_x] != s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x] != 0) {
		println("INFO: U: obstrued")
		return -1
	}

	return piece
}

func checkDownDoubleThree(ctx *s.SContext, case_x int, case_y int, totalPiece int) int {
	piece := 0
	count := 0
	lp := 1

	for case_y < int(ctx.NSize) && count < 4 && piece < totalPiece {
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
	}

	if lp == 1 && (case_y >= int(ctx.NSize) || ctx.Goban[case_y][case_x] != s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x] != 0) {
		println("INFO: D: obstrued")
		return -1
	}

	return piece
}

func checkVerticalPiece(ctx *s.SContext, case_x int, case_y int) bool {
	leftDoubleThree := checkLeftDoubleThree(ctx, case_x, case_y, 3)
	rightDoubleThree := checkRightDoubleThree(ctx, case_x, case_y, 3)

	leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y, 3)
	rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y, 3)

	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y, 3)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y, 3)

	if leftDoubleThree >= 3 || rightDoubleThree >= 3 || leftDoubleThree >= 2 && rightDoubleThree >= 2 ||
		leftUpDiagDoubleThree >= 3 || rightDownDiagDoubleThree >= 3 || leftUpDiagDoubleThree >= 2 && rightDownDiagDoubleThree >= 2 ||
		leftDownDiagDoubleThree >= 3 || rightUpDiagDoubleThree >= 3 || leftDownDiagDoubleThree >= 2 && rightUpDiagDoubleThree >= 2 {
		return false
	}

	return true
}

func loopVerticalUpPiece(ctx *s.SContext, case_x int, case_y int) bool {
	piece := 0
	count := 0

	for case_y >= 0 && count < 4 && piece < 2 {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			if !checkVerticalPiece(ctx, case_x, case_y) {
				return false
			}
		} else if ctx.Goban[case_y][case_x] != 0 {
			break
		}
		count++
		case_y--
	}

	return true
}

func loopVerticalDownPiece(ctx *s.SContext, case_x int, case_y int) bool {
	piece := 0
	count := 0

	for case_y < int(ctx.NSize) && count < 4 && piece < 2 {
		if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			piece++
			if !checkVerticalPiece(ctx, case_x, case_y) {
				return false
			}
		} else if ctx.Goban[case_y][case_x] != 0 {
			break
		}

		count++
		case_y++
	}

	return true
}

func checkVertical(ctx *s.SContext, case_x int, case_y int, leftDoubleThree int, rightDoubleThree int, upDoubleThree int, downDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if upDoubleThree >= 2 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			leftDoubleThree >= 1 && rightDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1 ||
			!loopVerticalUpPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if downDoubleThree >= 2 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			leftDoubleThree >= 1 && rightDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1 ||
			!loopVerticalDownPiece(ctx, case_x, case_y) {
			return false
		}
	}

	if upDoubleThree >= 1 && downDoubleThree >= 1 {
		if leftDoubleThree >= 2 || rightDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 ||
			leftDoubleThree >= 1 && rightDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1 ||
			!loopVerticalUpPiece(ctx, case_x, case_y) || !loopVerticalDownPiece(ctx, case_x, case_y) {
			return false
		}
	}

	return true
}
