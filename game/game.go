package game

import s "gomoku/structures"

func Placement(ctx *s.SContext, case_x int, case_y int) bool {
	if ctx.Goban[int(case_y-1)][int(case_x-1)] == 0 {
		ctx.Goban[int(case_y-1)][int(case_x-1)] = s.Tnumber(ctx.CurrentPlayer)
		return true
	} else {
		return false
	}
}
