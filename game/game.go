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
	// ctx.Capture = s.SVertex{X: -1, Y: -1}
	if ctx.Goban[int(case_y)][int(case_x)] == 0 {
		ctx.Goban[int(case_y)][int(case_x)] = s.Tnumber(ctx.CurrentPlayer)
		return true
	} else {
		return false
	}
}
