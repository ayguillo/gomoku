package game

import (
	s "gomoku/structures"
)

func horizontalAlign(ctx s.SContext, case_x int, case_y int) bool {
	count_stone := 0
	for current_case := case_x; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y-1][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
	}
	for current_case := case_x - 2; current_case >= 0; current_case-- {
		if ctx.Goban[case_y-1][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		return true
	}
	return false
}

func verticalAlign(ctx s.SContext, case_x int, case_y int) bool {
	count_stone := 0
	for current_case := case_y; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x-1] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
	}
	for current_case := case_y - 2; current_case >= 0; current_case-- {
		if ctx.Goban[current_case][case_x-1] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		return true
	}
	return false
}

func diagLeft(ctx s.SContext, case_x int, case_y int) bool {
	count_stone := 0
	for current_case_x, current_case_y := case_x, case_y; current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
			current_case_x++
			current_case_y++
		} else {
			break
		}
	}
	for current_case_x, current_case_y := case_x-2, case_y-2; current_case_x >= 0 && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
			current_case_x--
			current_case_y--
		} else {
			break
		}
	}
	if count_stone >= 4 {
		return true
	}
	return false
}

func diagRight(ctx s.SContext, case_x int, case_y int) bool {
	count_stone := 0
	for current_case_x, current_case_y := case_x, case_y-2; current_case_x < int(ctx.NSize) && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
			current_case_x++
			current_case_y--
		} else {
			break
		}
	}
	for current_case_x, current_case_y := case_x-2, case_y; current_case_x >= 0 && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
			current_case_x--
			current_case_y++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		return true
	}
	return false
}

func diagonalAlign(ctx s.SContext, case_x int, case_y int) bool {
	if diagLeft(ctx, case_x, case_y) == true {
		return true
	}
	if diagRight(ctx, case_x, case_y) == true {
		return true
	}
	return false
}

func VictoryConditionAlign(ctx *s.SContext, case_x int, case_y int) bool {
	ret_value := false
	if horizontalAlign(*ctx, case_x, case_y) == true {
		ret_value = true
	}
	if verticalAlign(*ctx, case_x, case_y) == true {
		ret_value = true
	}
	if diagonalAlign(*ctx, case_x, case_y) == true {
		ret_value = true
	}
	if ret_value == true {
		if ctx.CurrentPlayer == 1 {
			ctx.NbVictoryP1++
		} else {
			ctx.NbVictoryP2++
		}
	}
	return ret_value
}
