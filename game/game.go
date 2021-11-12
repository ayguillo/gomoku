package game

import (
	s "gomoku/structures"
)

func Placement(ctx *s.SContext, case_x int, case_y int) int {
	capture := false
	if ctx.ActiveCapture && len(ctx.Capture) > 0 {
		for _, cap := range ctx.Capture {
			if case_x == cap.X && case_y == cap.Y {
				capture = true
				break
			}
		}
		if capture == false {
			return 2
		}
	}
	if ctx.ActiveDoubleThrees && DoubleThree(s.SVertex{X: case_x, Y: case_y}, ctx.Goban, ctx.CurrentPlayer, ctx.ActiveCapture) {
		return 1
	}
	if ctx.Goban[int(case_y)][int(case_x)] == 0 {
		ctx.Goban[int(case_y)][int(case_x)] = s.Tnumber(ctx.CurrentPlayer)
		return 0
	} else {
		return -1
	}
}
