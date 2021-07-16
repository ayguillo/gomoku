package game

import (
	s "gomoku/structures"
)

func checkLefUptDiagDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	piece := 0
	count := 0

	for case_x >= 0 && case_y >= 0 && count < 4 && piece < 2 {
		if (ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer)) {
			piece++
		} else if (ctx.Goban[case_y][case_x] != 0) {
			return 0
		}

		count++
		case_y--
		case_x--
	}

	return piece
}

func checkRightUpDiagDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	piece := 0
	count := 0

	for case_x < int(ctx.NSize) && case_y >= 0 && count < 4 && piece < 2 {
		if (ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer)) {
			piece++
		} else if (ctx.Goban[case_y][case_x] != 0) {
			return 0
		}

		count++
		case_y--
		case_x++
	}

	return piece
}

func checkLeftDownDiagDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	piece := 0
	count := 0

	for case_x >= 0 && case_y < int(ctx.NSize) && count < 4 && piece < 2 {
		if (ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer)) {
			piece++
		} else if (ctx.Goban[case_y][case_x] != 0) {
			return 0
		}

		count++
		case_y++
		case_x--
	}

	return piece
}

func checkRightDownDiagDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	piece := 0
	count := 0

	for case_x < int(ctx.NSize) && case_y < int(ctx.NSize) && count < 4 && piece < 2 {
		if (ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer)) {
			piece++
		} else if (ctx.Goban[case_y][case_x] != 0) {
			return 0
		}

		count++
		case_y++
		case_x++
	}

	return piece
}

func checkUpDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	piece := 0
	count := 0

	for case_y >= 0 && count < 4 && piece < 2 {
		if (ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer)) {
			piece++
		} else if (ctx.Goban[case_y][case_x] != 0) {
			return 0
		}

		count++
		case_y--
	}

	return piece
}

func checkDownDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	piece := 0
	count := 0

	for case_y < int(ctx.NSize) && count < 4 && piece < 2 {
		if (ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer)) {
			piece++
		} else if (ctx.Goban[case_y][case_x] != 0) {
			return 0
		}

		count++
		case_y++
	}

	return piece
}

func checkLeftDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	piece := 0
	count := 0

	for case_x >= 0 && count < 4 && piece < 2 {
		if (ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer)) {
			piece++
		} else if (ctx.Goban[case_y][case_x] != 0) {
			return 0
		}

		count++
		case_x--
	}

	return piece
}

func checkRightDoubleThree(ctx *s.SContext, case_x int, case_y int) int {
	piece := 0
	count := 0

	for case_x < int(ctx.NSize) && count < 4 && piece < 2 {
		if (ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer)) {
			piece++
		} else if (ctx.Goban[case_y][case_x] != 0) {
			return 0
		}

		count++
		case_x++
	}

	return piece
}

func checkHorizon(ctx *s.SContext, case_x int, case_y int, leftDoubleThree int, rightDoubleThree int, upDoubleThree int, downDoubleThree int, leftUpDiagDoubleThree int, rightDownDiagDoubleThree int, leftDownDiagDoubleThree int, rightUpDiagDoubleThree int) bool {
	if (leftDoubleThree >= 2) {
		if (upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}
	}

	if (leftDoubleThree >= 1 && rightDoubleThree >= 1) {
		if (upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}
	}

	if (rightDoubleThree >= 2) {
		if (upDoubleThree >= 2 || downDoubleThree >= 2 || leftUpDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2 || leftDownDiagDoubleThree >= 2 || rightUpDiagDoubleThree >= 2) {
			return false
		}
	}

	return true
}



func checkDoubleThree(ctx *s.SContext, case_x int, case_y int) bool {
	leftDoubleThree := checkLeftDoubleThree(ctx, case_x, case_y)
	rightDoubleThree := checkRightDoubleThree(ctx, case_x, case_y)
	
	upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y)
	downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y)

	leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y)
	rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y)

	leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y)
	rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y)


	if (checkHorizon(ctx, case_x, case_y, leftDoubleThree, rightDoubleThree, upDoubleThree, downDoubleThree, leftUpDiagDoubleThree, rightDownDiagDoubleThree, leftDownDiagDoubleThree, rightUpDiagDoubleThree) == false) {
		return false
	}
	// if (/*captureIsNotACapture() && */ horizontal + vertical + left + right >= 4) {
	// 	return false
	// }
	return true
}