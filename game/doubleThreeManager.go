package game

import (
	s "gomoku/structures"
)

func isEnnemyCase(ctx *s.SContext, case_x int, case_y int) bool {
	if ctx.Goban[case_y][case_x] != s.Tnumber(ctx.CurrentPlayer) && ctx.Goban[case_y][case_x] != 0 {
		return true
	}

	return false
}

func CheckDoubleThree(ctx *s.SContext, case_x int, case_y int) bool {
	horizonDoubleThree := checkHorizonDoubleThree(ctx, case_x, case_y)
	leftDoubleThree := checkLeftDoubleThree(ctx, case_x, case_y, 2)
	rightDoubleThree := checkRightDoubleThree(ctx, case_x, case_y, 2)

	verticalDoubleThree := checkVerticalDoubleThree(ctx, case_x, case_y)
	upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y, 2)
	downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y, 2)

	diagLeftDoubleThree := checkDiagLeftDoubleThree(ctx, case_x, case_y)
	leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y, 2)
	rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y, 2)

	diagRightDoubleThree := checkDiagRightDoubleThree(ctx, case_x, case_y)
	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y, 2)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y, 2)

	if !checkHorizon(ctx, case_x, case_y, horizonDoubleThree, leftDoubleThree, rightDoubleThree, verticalDoubleThree, upDoubleThree, downDoubleThree, diagLeftDoubleThree, leftUpDiagDoubleThree, rightDownDiagDoubleThree, diagRightDoubleThree, leftDownDiagDoubleThree, rightUpDiagDoubleThree) ||
		!checkVertical(ctx, case_x, case_y, horizonDoubleThree, leftDoubleThree, rightDoubleThree, verticalDoubleThree, upDoubleThree, downDoubleThree, diagLeftDoubleThree, leftUpDiagDoubleThree, rightDownDiagDoubleThree, diagRightDoubleThree, leftDownDiagDoubleThree, rightUpDiagDoubleThree) ||
		!checkDiagLeft(ctx, case_x, case_y, horizonDoubleThree, leftDoubleThree, rightDoubleThree, verticalDoubleThree, upDoubleThree, downDoubleThree, diagLeftDoubleThree, leftUpDiagDoubleThree, rightDownDiagDoubleThree, diagRightDoubleThree, leftDownDiagDoubleThree, rightUpDiagDoubleThree) ||
		!checkDiagRight(ctx, case_x, case_y, horizonDoubleThree, leftDoubleThree, rightDoubleThree, verticalDoubleThree, upDoubleThree, downDoubleThree, diagLeftDoubleThree, leftUpDiagDoubleThree, rightDownDiagDoubleThree, diagRightDoubleThree, leftDownDiagDoubleThree, rightUpDiagDoubleThree) {
		return false
	}

	return true
}
