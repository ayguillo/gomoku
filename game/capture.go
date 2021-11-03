package game

import (
	d "gomoku/display"
	s "gomoku/structures"
)

func diagLeftCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) (bool, []s.SVertex) {
	count_stone := 0
	current_case_x := case_x + 1
	current_case_y := case_y + 1
	ret_value := false
	var ret_vertex []s.SVertex
	color := [4]uint8{226, 196, 115, 255}

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
				d.TraceStone(float64(current_case_x-1), float64(current_case_y-1), ctx, visu, color, true)
				d.TraceStone(float64(current_case_x-2), float64(current_case_y-2), ctx, visu, color, true)
			}
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case_x - 1,
				Y: current_case_y - 1})
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case_x - 2,
				Y: current_case_y - 2})
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y-1][current_case_x-1] = 0
			ctx.Goban[current_case_y-2][current_case_x-2] = 0
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case_x - 1, Y: current_case_y - 1})
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case_x - 2, Y: current_case_y - 2})
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
				d.TraceStone(float64(current_case_x+1), float64(current_case_y+1), ctx, visu, color, true)
				d.TraceStone(float64(current_case_x+2), float64(current_case_y+2), ctx, visu, color, true)
			}
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case_x + 1,
				Y: current_case_y + 1})
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case_x + 2,
				Y: current_case_y + 2})
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y+1][current_case_x+1] = 0
			ctx.Goban[current_case_y+2][current_case_x+2] = 0
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case_x + 1, Y: current_case_y + 1})
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case_x + 2, Y: current_case_y + 2})
			ret_value = true
		}
	}
	return ret_value, ret_vertex
}

func diagRightCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) (bool, []s.SVertex) {
	count_stone := 0
	current_case_x := case_x + 1
	current_case_y := case_y - 1
	ret_value := false
	color := [4]uint8{226, 196, 115, 255}
	var ret_vertex []s.SVertex

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
				d.TraceStone(float64(current_case_x-1), float64(current_case_y+1), ctx, visu, color, true)
				d.TraceStone(float64(current_case_x-2), float64(current_case_y+2), ctx, visu, color, true)
			}
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case_x - 1,
				Y: current_case_y + 1})
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case_x - 2,
				Y: current_case_y + 2})
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y+1][current_case_x-1] = 0
			ctx.Goban[current_case_y+2][current_case_x-2] = 0
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case_x - 1, Y: current_case_y + 1})
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case_x - 2, Y: current_case_y + 2})
			ret_value = true
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
				d.TraceStone(float64(current_case_x+1), float64(current_case_y-1), ctx, visu, color, true)
				d.TraceStone(float64(current_case_x+2), float64(current_case_y-2), ctx, visu, color, true)
			}
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case_x + 1,
				Y: current_case_y - 1})
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case_x + 2,
				Y: current_case_y - 2})
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y-1][current_case_x+1] = 0
			ctx.Goban[current_case_y-2][current_case_x+2] = 0
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case_x + 1, Y: current_case_y - 1})
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case_x + 2, Y: current_case_y - 2})
			ret_value = true
		}
	}
	return ret_value, ret_vertex
}

func diagonalCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) (bool, []s.SVertex) {
	val_ret := false
	var ret_vertex []s.SVertex
	ret_value_left, ret_vertex_left := diagLeftCapture(ctx, visu, case_x, case_y, capture, viz)
	if ret_value_left == true {
		val_ret = true
		for _, vertex_left := range ret_vertex_left {
			ret_vertex = append(ret_vertex, vertex_left)
		}

	}
	ret_value_right, ret_vertex_right := diagRightCapture(ctx, visu, case_x, case_y, capture, viz)
	if ret_value_right == true {
		val_ret = true
		for _, vertex_right := range ret_vertex_right {
			ret_vertex = append(ret_vertex, vertex_right)
		}
	}
	return val_ret, ret_vertex
}

func horizontalCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) (bool, []s.SVertex) {
	count_stone := 0
	current_case := case_x + 1
	ret_value := false
	color := [4]uint8{226, 196, 115, 255}
	var ret_vertex []s.SVertex
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
				d.TraceStone(float64(current_case-1), float64(case_y), ctx, visu, color, true)
				d.TraceStone(float64(current_case-2), float64(case_y), ctx, visu, color, true)
			}
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case - 1,
				Y: case_y})
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case - 2,
				Y: case_y})
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[case_y][current_case-1] = 0
			ctx.Goban[case_y][current_case-2] = 0
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case - 1, Y: case_y})
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case - 2, Y: case_y})
			ret_value = true
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
				d.TraceStone(float64(current_case+1), float64(case_y), ctx, visu, color, true)
				d.TraceStone(float64(current_case+2), float64(case_y), ctx, visu, color, true)
			}
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case + 1,
				Y: case_y})
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: current_case + 2,
				Y: case_y})
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[case_y][current_case+1] = 0
			ctx.Goban[case_y][current_case+2] = 0
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case + 1, Y: case_y})
			ret_vertex = append(ret_vertex, s.SVertex{X: current_case + 2, Y: case_y})
			ret_value = true
		}
	}
	return ret_value, ret_vertex
}

func verticalCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) (bool, []s.SVertex) {
	count_stone := 0
	current_case := case_y + 1
	ret_value := false
	color := [4]uint8{226, 196, 115, 255}
	var ret_vertex []s.SVertex
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
				d.TraceStone(float64(case_x), float64(current_case-1), ctx, visu, color, true)
				d.TraceStone(float64(case_x), float64(current_case-2), ctx, visu, color, true)
			}
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: case_x,
				Y: current_case - 1})
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: case_x,
				Y: current_case - 2})
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case-1][case_x] = 0
			ctx.Goban[current_case-2][case_x] = 0
			ret_vertex = append(ret_vertex, s.SVertex{X: case_x, Y: current_case - 1})
			ret_vertex = append(ret_vertex, s.SVertex{X: case_x, Y: current_case - 2})
			ret_value = true
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
				d.TraceStone(float64(case_x), float64(current_case+1), ctx, visu, color, true)
				d.TraceStone(float64(case_x), float64(current_case+2), ctx, visu, color, true)
			}
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: case_x,
				Y: current_case + 1})
			ctx.CasesNonNull = append(ctx.CasesNonNull, s.SVertex{
				X: case_x,
				Y: current_case + 2})
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case+1][case_x] = 0
			ctx.Goban[current_case+2][case_x] = 0
			ret_vertex = append(ret_vertex, s.SVertex{X: case_x, Y: current_case + 1})
			ret_vertex = append(ret_vertex, s.SVertex{X: case_x, Y: current_case + 2})
			ret_value = true
		}
	}
	return ret_value, ret_vertex
}

func Capture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, viz bool) (bool, []s.SVertex) {
	var capture uint8
	val_ret := false
	var ret_vertex []s.SVertex
	if ctx.CurrentPlayer == 1 {
		capture = 2
	} else {
		capture = 1
	}
	ret_value_horiz, ret_vertex_horiz := horizontalCapture(ctx, visu, case_x, case_y, capture, viz)
	if ret_value_horiz == true {
		val_ret = true
		for _, vertex_horiz := range ret_vertex_horiz {
			ret_vertex = append(ret_vertex, vertex_horiz)
		}
	}
	ret_value_vert, ret_vertex_vert := verticalCapture(ctx, visu, case_x, case_y, capture, viz)
	if ret_value_vert == true {
		val_ret = true
		for _, vertex_vert := range ret_vertex_vert {
			ret_vertex = append(ret_vertex, vertex_vert)
		}
	}
	ret_value_diag, ret_vertex_diag := diagonalCapture(ctx, visu, case_x, case_y, capture, viz)
	if ret_value_diag == true {
		val_ret = true
		for _, vertex_diag := range ret_vertex_diag {
			ret_vertex = append(ret_vertex, vertex_diag)
		}
	}
	return val_ret, ret_vertex
}
