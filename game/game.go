package game

import (
	"fmt"
	d "gomoku/display"
	s "gomoku/structures"
)

func addToNeigbors(ctx *s.SContext, case_x_neigbors int, case_y_neigbors int, current_x int, current_y int, visu *s.SVisu) {
	vertex := s.SVertex{X: current_x, Y: current_y}
	new_vertex := s.SVertex{X: case_x_neigbors, Y: case_y_neigbors}
	array_neigh := ctx.CasesNonNull[vertex]
	array_neigh = append(array_neigh, new_vertex)
	ctx.CasesNonNull[vertex] = array_neigh
	// Loop display
	fmt.Println("Add", case_x_neigbors, case_y_neigbors)
	color := [4]uint8{83, 51, 237, 1}
	for _, neighbor := range ctx.CasesNonNull[vertex] {
		if neighbor == new_vertex {
			d.TraceStone(float64(neighbor.X), float64(neighbor.Y), ctx, visu, color, false)
		}
	}
}

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

func diagLeftCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) bool {
	count_stone := 0
	current_case_x := case_x + 1
	current_case_y := case_y + 1
	ret_value := false
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
			addToNeigbors(ctx, current_case_x-1, current_case_y-1, current_case_x, current_case_y, visu)
			addToNeigbors(ctx, current_case_x-2, current_case_y-2, case_x, case_y, visu)
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
				d.TraceStone(float64(current_case_x+1), float64(current_case_y+1), ctx, visu, color, true)
				d.TraceStone(float64(current_case_x+2), float64(current_case_y+2), ctx, visu, color, true)
			}
			addToNeigbors(ctx, current_case_x+1, current_case_y+1, current_case_x, current_case_y, visu)
			addToNeigbors(ctx, current_case_x+2, current_case_y+2, case_x, case_y, visu)
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
	color := [4]uint8{226, 196, 115, 255}
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
			addToNeigbors(ctx, current_case_x-1, current_case_y+1, current_case_x, current_case_y, visu)
			addToNeigbors(ctx, current_case_x-2, current_case_y+2, case_x, case_y, visu)
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y+1][current_case_x-1] = 0
			ctx.Goban[current_case_y+2][current_case_x-2] = 0
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
			addToNeigbors(ctx, current_case_x+1, current_case_y-1, current_case_x, current_case_y, visu)
			addToNeigbors(ctx, current_case_x+2, current_case_y-2, case_x, case_y, visu)
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case_y-1][current_case_x+1] = 0
			ctx.Goban[current_case_y-2][current_case_x+2] = 0
			ret_value = true
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
	color := [4]uint8{226, 196, 115, 255}
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
			addToNeigbors(ctx, current_case-1, case_y, case_x, case_y, visu)
			addToNeigbors(ctx, current_case-2, case_y, current_case+1, case_y, visu)
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[case_y][current_case-1] = 0
			ctx.Goban[case_y][current_case-2] = 0
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
			addToNeigbors(ctx, current_case+1, case_y, current_case, case_y, visu)
			addToNeigbors(ctx, current_case+2, case_y, case_x, case_y, visu)
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[case_y][current_case+1] = 0
			ctx.Goban[case_y][current_case+2] = 0
			ret_value = true
		}
	}
	return ret_value
}

func verticalCapture(ctx *s.SContext, visu *s.SVisu, case_x int, case_y int, capture uint8, viz bool) bool {
	count_stone := 0
	current_case := case_y + 1
	ret_value := false
	color := [4]uint8{226, 196, 115, 255}
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
			addToNeigbors(ctx, case_x, current_case-1, case_x, case_y, visu)
			addToNeigbors(ctx, case_x, current_case-2, case_x, current_case+1, visu)
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case-1][case_x] = 0
			ctx.Goban[current_case-2][case_x] = 0
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
			addToNeigbors(ctx, case_x, current_case+1, case_x, current_case, visu)
			addToNeigbors(ctx, case_x, current_case+2, case_x, case_y, visu)
			if ctx.CurrentPlayer == 1 {
				ctx.NbCaptureP1++
			} else {
				ctx.NbCaptureP2++
			}
			ctx.Goban[current_case+1][case_x] = 0
			ctx.Goban[current_case+2][case_x] = 0
			ret_value = true
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
