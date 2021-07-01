package game

import (
	d "gomoku/display"
	s "gomoku/structures"
)

func Placement(ctx *s.SContext, case_x int, case_y int) bool {
	if ctx.Goban[int(case_y)][int(case_x)] == 0 {
		ctx.Goban[int(case_y)][int(case_x)] = s.Tnumber(ctx.CurrentPlayer)
		return true
	} else {
		return false
	}
}

func diagLeftCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) bool {
	count_stone := 0
	current_case_x := case_x + 1
	current_case_y := case_y + 1
	ret_value := false
	for current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capture) {
			count_stone++
			current_case_x++
			current_case_y++
		} else {
			break
		}
	}
	if count_stone == 2 && current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if viz == true {
				d.TraceStone(float64(current_case_x-1), float64(current_case_y-1), ctx, visu, true)
				d.TraceStone(float64(current_case_x-2), float64(current_case_y-2), ctx, visu, true)
			}
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y-1][current_case_x-1] = 0
			ctx.Goban[current_case_y-2][current_case_x-2] = 0
			ret_value = true
		}
	}
	count_stone = 0
	current_case_x = case_x - 1
	current_case_y = case_y - 1
	for current_case_x >= 0 && current_case_y >= 0 {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capture) {
			count_stone++
			current_case_x--
			current_case_y--
		} else {
			break
		}
	}
	if count_stone == 2 && current_case_x >= 0 && current_case_y >= 0 {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if viz == true {
				d.TraceStone(float64(current_case_x+1), float64(current_case_y+1), ctx, visu, true)
				d.TraceStone(float64(current_case_x+2), float64(current_case_y+2), ctx, visu, true)
			}
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y+1][current_case_x+1] = 0
			ctx.Goban[current_case_y+2][current_case_x+2] = 0
			ret_value = true
		}
	}
	return ret_value
}

func diagRightCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) bool {
	count_stone := 0
	current_case_x := case_x + 1
	current_case_y := case_y - 1
	ret_value := false
	for current_case_x < int(ctx.NSize) && current_case_y >= 0 {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capture) {
			count_stone++
			current_case_x++
			current_case_y--
		} else {
			break
		}
	}
	if count_stone == 2 && current_case_x < int(ctx.NSize) && current_case_y >= 0 {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if viz == true {
				d.TraceStone(float64(current_case_x-1), float64(current_case_y+1), ctx, visu, true)
				d.TraceStone(float64(current_case_x-2), float64(current_case_y+2), ctx, visu, true)
			}
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y+1][current_case_x-1] = 0
			ctx.Goban[current_case_y+2][current_case_x-2] = 0
			return ret_value
		}
	}
	count_stone = 0
	current_case_x = case_x - 1
	current_case_y = case_y + 1
	for current_case_x >= 0 && current_case_y < int(ctx.NSize) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capture) {
			count_stone++
			current_case_x--
			current_case_y++
		} else {
			break
		}
	}
	if count_stone == 2 && current_case_x < int(ctx.NSize) && current_case_y >= 0 {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if viz == true {
				d.TraceStone(float64(current_case_x+1), float64(current_case_y-1), ctx, visu, true)
				d.TraceStone(float64(current_case_x+2), float64(current_case_y-2), ctx, visu, true)
			}
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y-1][current_case_x+1] = 0
			ctx.Goban[current_case_y-2][current_case_x+2] = 0
			return ret_value
		}
	}
	return ret_value
}

func diagonalCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) bool {
	val_ret := false
	if diagLeftCapture(ctx, visu, case_x, case_y, capture, viz) == true {
		val_ret = true
	}
	if diagRightCapture(ctx, visu, case_x, case_y, capture, viz) == true {
		val_ret = true
	}
	return val_ret
}

func horizontalCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) bool {
	count_stone := 0
	current_case := case_x + 1
	ret_value := false
	for ; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == s.Tnumber(capture) {
			count_stone++
		} else {
			break
		}
	}
	if count_stone == 2 && current_case < int(ctx.NSize) {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			if viz == true {
				d.TraceStone(float64(current_case-1), float64(case_y), ctx, visu, true)
				d.TraceStone(float64(current_case-2), float64(case_y), ctx, visu, true)
			}
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[case_y][current_case-1] = 0
			ctx.Goban[case_y][current_case-2] = 0
			return ret_value
		}
	}
	count_stone = 0
	current_case = case_x - 1
	for ; current_case >= 0; current_case-- {
		if ctx.Goban[case_y][current_case] == s.Tnumber(capture) {
			count_stone++
		} else {
			break
		}
	}
	if count_stone == 2 && current_case >= 0 {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			if viz == true {
				d.TraceStone(float64(current_case+1), float64(case_y), ctx, visu, true)
				d.TraceStone(float64(current_case+2), float64(case_y), ctx, visu, true)
			}
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[case_y][current_case+1] = 0
			ctx.Goban[case_y][current_case+2] = 0
			return ret_value
		}
	}
	return ret_value
}

func verticalCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) bool {
	count_stone := 0
	current_case := case_y + 1
	ret_value := false
	for ; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == s.Tnumber(capture) {
			count_stone++
		} else {
			break
		}
	}
	if count_stone == 2 && current_case < int(ctx.NSize) {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if viz == true {
				d.TraceStone(float64(case_x), float64(current_case-1), ctx, visu, true)
				d.TraceStone(float64(case_x), float64(current_case-2), ctx, visu, true)
			}
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case-1][case_x] = 0
			ctx.Goban[current_case-2][case_x] = 0
			return ret_value
		}
	}
	count_stone = 0
	current_case = case_y - 1
	for ; current_case >= 0; current_case-- {
		if ctx.Goban[current_case][case_x] == s.Tnumber(capture) {
			count_stone++
		} else {
			break
		}
	}
	if count_stone == 2 && current_case >= 0 {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if viz == true {
				d.TraceStone(float64(case_x), float64(current_case+1), ctx, visu, true)
				d.TraceStone(float64(case_x), float64(current_case+2), ctx, visu, true)
			}
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case+1][case_x] = 0
			ctx.Goban[current_case+2][case_x] = 0
			return ret_value
		}
	}
	return ret_value
}

func Capture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, viz bool) bool {
	var capture uint8
	val_ret := false
	if ctx.CurrentPlayer == 1 {
		capture = 2
	} else {
		capture = 1
	}
	if horizontalCapture(ctx, visu, case_x, case_y, capture, viz) == true {
		val_ret = true
	}
	if verticalCapture(ctx, visu, case_x, case_y, capture, viz) == true {
		val_ret = true
	}
	if diagonalCapture(ctx, visu, case_x, case_y, capture, viz) == true {
		val_ret = true
	}
	return val_ret
}
