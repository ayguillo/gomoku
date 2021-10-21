package game

import (
	s "gomoku/structures"
)

func Placement(ctx *s.SContext, case_x int, case_y int) bool {
	if ctx.Capture.X != -1 {
		if case_x != ctx.Capture.X || case_y != ctx.Capture.Y {
			return false
		}
	}
	if (ctx.Goban[int(case_y)][int(case_x)] == 0) && checkDoubleThree(ctx, case_x, case_y) {
		ctx.Goban[int(case_y)][int(case_x)] = s.Tnumber(ctx.CurrentPlayer)
		return true
	} else {
		return false
	}
}
