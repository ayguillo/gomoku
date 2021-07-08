package game

import (
	s "gomoku/structures"
)

func checkLefUptDiagDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	return 0
}

func checkRightUpDiagDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	return 0
}

func checkLeftDownDiagDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	return 0
}

func checkRightDownDiagDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	return 0
}

func checkUpDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	return 0
}

func checkDownDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	return 0
}

func checkLeftDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	return 0
}

func checkRightDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	piece := 0
	count := 0
	for case_x < int(ctx.NSize) && count < 4 && piece < 2 {
		if (ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer)) {
			piece++
		} else if (ctx.Goban[case_y][case_x] != 0) {
			println("ENNEMI", piece)
			return 1
		}

		count++
		case_x++
	}

	if (piece >= 2) {
		println("RIGHT_TREE")
	}

	return 0
}

func checkDoubleThree(ctx *s.SContext, case_x int, case_y int) bool {
	if (/*captureIsNotACapture() && */(checkLefUptDiagDoubleThree(ctx, case_x, case_y) + checkRightUpDiagDoubleThree(ctx, case_x, case_y) + checkLeftDownDiagDoubleThree(ctx, case_x, case_y) + checkRightDownDiagDoubleThree(ctx, case_x, case_y) + checkUpDoubleThree(ctx, case_x, case_y) + checkDownDoubleThree(ctx, case_x, case_y) + checkLeftDoubleThree(ctx, case_x, case_y) + checkRightDoubleThree(ctx, case_x, case_y) >= 2)) {
		return false
	}
	return true
}