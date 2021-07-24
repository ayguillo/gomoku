package game

import (
	s "gomoku/structures"
)

func checkDoubleThreeHorizon(ctx *s.SContext, case_x int, case_y int) bool {
	upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y)
	downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y)

	leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y)
	rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y)

	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y)

	if (upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightDownDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
		return false
	}

	return true
}

func checkHorizon(ctx *s.SContext, case_x int, case_y int, leftDoubleThree int, rightDoubleThree int, upDoubleThree int, downDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if (leftDoubleThree >= 2) {
		if (upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}

		if (upDoubleThree >= 1 && downDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) {
			return false
		}
	}

	if (rightDoubleThree >= 2) {
		if (upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}

		if (upDoubleThree >= 1 && downDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) {
			return false
		}
	}

	if (leftDoubleThree >= 1 && rightDoubleThree >= 1) {
		if (upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}

		if (upDoubleThree >= 1 && downDoubleThree >= 1 || leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) {
			return false
		}
	}

	return true
}