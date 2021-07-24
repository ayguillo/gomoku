package game

import (
	s "gomoku/structures"
)

func checkDoubleThreeVertical(ctx *s.SContext, case_x int, case_y int) bool {
	leftDoubleThree := checkLeftDoubleThree(ctx, case_x, case_y)
	rightDoubleThree := checkRightDoubleThree(ctx, case_x, case_y)

	leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y)
	rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y)

	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y)

	if (leftDoubleThree >= 2 || rightDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
		return false
	}

	return true
}

func checkVertical(ctx *s.SContext, case_x int, case_y int, leftDoubleThree int, rightDoubleThree int, upDoubleThree int, downDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if (upDoubleThree >= 2) {
		if (leftDoubleThree >= 2 || rightDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}

		if (leftDoubleThree >= 1 && rightDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) {
			return false
		}
	}

	if (downDoubleThree >= 2) {
		if (leftDoubleThree >= 2 || rightDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}

		if (leftDoubleThree >= 1 && rightDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) {
			return false
		}
	}

	if (upDoubleThree >= 1 && downDoubleThree >= 1) {
		if (leftDoubleThree >= 2 || rightDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}

		if (leftDoubleThree >= 1 && rightDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) {
			return false
		}
	}

	return true
}