package game

import (
	s "gomoku/structures"
)

func checkDiagRight(ctx *s.SContext, case_x int, case_y int, leftDoubleThree int, rightDoubleThree int, upDoubleThree int, downDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if (leftUpDiagDoubleThree >= 2) {
		if (leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}

		if (leftDoubleThree >= 1 && rightDoubleThree >= 1 || upDoubleThree >= 1 && downDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) {
			return false
		}
	}

	if (rightDownDiagDoubleThree >= 2) {
		if (leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}

		if (leftDoubleThree >= 1 && rightDoubleThree >= 1 || upDoubleThree >= 1 && downDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) {
			return false
		}
	}

	if (leftUpDiagDoubleThree >= 1 && rightDownDiagDoubleThree >= 1) {
		if (leftDoubleThree >= 2 || rightDoubleThree >= 2 || upDoubleThree >= 2 || downDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}

		if (leftDoubleThree >= 1 && rightDoubleThree >= 1 || upDoubleThree >= 1 && downDoubleThree >= 1 || rightUpDiagDoubleThree >= 1 && leftDownDiagDoubleThree >= 1) {
			return false
		}
	}

	return true
}