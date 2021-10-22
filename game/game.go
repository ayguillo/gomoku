package game

import (
	s "gomoku/structures"
)

func Placement(ctx *s.SContext, case_x int, case_y int) int {
	capture := false
	if len(ctx.Capture) > 0 && ctx.ActiveCapture {
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
	if !CheckDoubleThree(ctx, case_x, case_y) && ctx.ActiveDoubleThrees {
		return 1
	}
	if ctx.Goban[int(case_y)][int(case_x)] == 0 {
		ctx.Goban[int(case_y)][int(case_x)] = s.Tnumber(ctx.CurrentPlayer)
		return 0
	} else {
		return -1
	}
}
