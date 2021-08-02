package game

import (
	s "gomoku/structures"
)

func checkDoubleThree(ctx *s.SContext, case_x int, case_y int) bool {
	leftDoubleThree := checkLeftDoubleThree(ctx, case_x, case_y, 2)
	rightDoubleThree := checkRightDoubleThree(ctx, case_x, case_y, 2)

	upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y, 2)
	downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y, 2)

	leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y, 2)
	rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y, 2)

	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y, 2)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y, 2)

	if !checkHorizon(ctx, case_x, case_y, leftDoubleThree, rightDoubleThree, upDoubleThree, downDoubleThree, leftUpDiagDoubleThree, rightDownDiagDoubleThree, leftDownDiagDoubleThree, rightUpDiagDoubleThree) ||
		!checkVertical(ctx, case_x, case_y, leftDoubleThree, rightDoubleThree, upDoubleThree, downDoubleThree, leftUpDiagDoubleThree, rightDownDiagDoubleThree, leftDownDiagDoubleThree, rightUpDiagDoubleThree) ||
		!checkDiagLeft(ctx, case_x, case_y, leftDoubleThree, rightDoubleThree, upDoubleThree, downDoubleThree, leftUpDiagDoubleThree, rightDownDiagDoubleThree, leftDownDiagDoubleThree, rightUpDiagDoubleThree) ||
		!checkDiagRight(ctx, case_x, case_y, leftDoubleThree, rightDoubleThree, upDoubleThree, downDoubleThree, leftUpDiagDoubleThree, rightDownDiagDoubleThree, leftDownDiagDoubleThree, rightUpDiagDoubleThree) {
		return false
	}

	return true
}
