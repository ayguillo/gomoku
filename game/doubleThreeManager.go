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

	if (piece >= 2) {
		println("LEFT_UP_THREE")
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

	if (piece >= 2) {
		println("RIGHT_UP_THREE")
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

	if (piece >= 2) {
		println("LEFT_DOWN_THREE")
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

	if (piece >= 2) {
		println("RIGHT_DOWN_THREE")
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

	if (piece >= 2) {
		println("UP_THREE")
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

	if (piece >= 2) {
		println("DOWN_THREE")
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

	if (piece >= 2) {
		println("LEFT_THREE")
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

	if (piece >= 2) {
		println("RIGHT_TREE")
	}

	return piece
}

func checkDoubleThree(ctx *s.SContext, case_x int, case_y int) bool {
	// horizontal := 0
	// vertical := 0
	// left := 0
	// right := 0

	// leftDoubleThree := checkLeftDoubleThree(ctx, case_x, case_y)
	// rightDoubleThree := checkRightDoubleThree(ctx, case_x, case_y)
	
	// upDoubleThree := checkUpDoubleThree(ctx, case_x, case_y)
	// downDoubleThree := checkDownDoubleThree(ctx, case_x, case_y)

	// leftUpDiagDoubleThree := checkLefUptDiagDoubleThree(ctx, case_x, case_y)
	// rightDownDiagDoubleThree := checkRightDownDiagDoubleThree(ctx, case_x, case_y)

	// leftDownDiagDoubleThree := checkLeftDownDiagDoubleThree(ctx, case_x, case_y)
	// rightUpDiagDoubleThree := checkRightUpDiagDoubleThree(ctx, case_x, case_y)

	

	// if (/*captureIsNotACapture() && */ horizontal + vertical + left + right >= 4) {
	// 	return false
	// }
	return true
}