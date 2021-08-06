package algorithm

import (
	s "gomoku/structures"
)

func diagLeftCaptureHeuristic(ctx s.SContext, case_x int, case_y int, capture uint8, score_capture int) int {
	count_stone := 0
	current_case_x := case_x + 1
	current_case_y := case_y + 1
	ret_value := 0
	for (current_case_x < int(ctx.NSize) && current_case_x > 0) && (current_case_y < int(ctx.NSize) && current_case_y > 0) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capture) {
			count_stone++
			current_case_x++
			current_case_y++
		} else {
			break
		}
	}
	if count_stone == 2 && (current_case_x < int(ctx.NSize) && current_case_x > 0) && (current_case_y < int(ctx.NSize) && current_case_y > 0) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			// addToNeigbors(ctx, current_case_x-1, current_case_y-1, current_case_x, current_case_y, visu) A ajouter sans visu
			// addToNeigbors(ctx, current_case_x-2, current_case_y-2, case_x, case_y, visu) A ajouter sans visu
			ret_value += 20
		} else if ctx.Goban[current_case_y][current_case_x] == 0 {
			ret_value += 10
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
			// addToNeigbors(ctx, current_case_x+1, current_case_y+1, current_case_x, current_case_y, visu)  A ajouter sans visu
			// addToNeigbors(ctx, current_case_x+2, current_case_y+2, case_x, case_y, visu) A ajouter sans
			ret_value += 20
		} else if ctx.Goban[current_case_y][current_case_x] == 0 {
			ret_value += 10
		}
	}
	if ret_value != 0 {
		ret_value += score_capture
	}
	return ret_value
}

func diagRightCaptureHeuristic(ctx s.SContext, case_x int, case_y int, capture uint8, score_capture int) int {
	count_stone := 0
	current_case_x := case_x + 1
	current_case_y := case_y - 1
	ret_value := 0
	for (current_case_x < int(ctx.NSize) && current_case_x > 0) && (current_case_y < int(ctx.NSize) && current_case_y > 0) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capture) {
			count_stone++
			current_case_x++
			current_case_y--
		} else {
			break
		}
	}
	if count_stone == 2 && (current_case_x < int(ctx.NSize) && current_case_x > 0) && (current_case_y < int(ctx.NSize) && current_case_y > 0) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			// addToNeigbors(ctx, current_case_x-1, current_case_y+1, current_case_x, current_case_y, visu) A ajouter sans visu
			// addToNeigbors(ctx, current_case_x-2, current_case_y+2, case_x, case_y, visu) A ajouter sans visu
			ret_value += 20
		} else if ctx.Goban[current_case_y][current_case_x] == 0 {
			ret_value += 10
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
	if count_stone == 2 && (current_case_x < int(ctx.NSize) && current_case_x > 0) && (current_case_y < int(ctx.NSize) && current_case_y > 0) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			// addToNeigbors(ctx, current_case_x+1, current_case_y-1, current_case_x, current_case_y, visu) A ajouter sans visu
			// addToNeigbors(ctx, current_case_x+2, current_case_y-2, case_x, case_y, visu) A ajouter sans visu
			ret_value += 20
		} else if ctx.Goban[current_case_y][current_case_x] == 0 {
			ret_value += 10
		}
	}
	if ret_value != 0 {
		ret_value += score_capture
	}
	return ret_value
}

func diagonalCaptureHeuristic(ctx s.SContext, case_x int, case_y int, capture uint8, score_capture int) int {
	left := diagLeftCaptureHeuristic(ctx, case_x, case_y, capture, score_capture)
	right := diagRightCaptureHeuristic(ctx, case_x, case_y, capture, score_capture)
	return left + right
}

func horizontalCaptureHeuristic(ctx s.SContext, case_x int, case_y int, capture uint8, score_capture int) int {
	count_stone := 0
	current_case := case_x + 1
	ret_value := 0
	for ; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == s.Tnumber(capture) {
			count_stone++
		} else {
			break
		}
	}
	if count_stone == 2 && current_case < int(ctx.NSize) {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			// addToNeigbors(ctx, current_case-1, case_y, case_x, case_y, visu) A ajouter sans visu
			// addToNeigbors(ctx, current_case-2, case_y, current_case+1, case_y, visu) A ajouter sans visu
			ret_value += 20
		} else if ctx.Goban[case_y][current_case] == 0 {
			ret_value += 10
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
			// addToNeigbors(ctx, current_case+1, case_y, current_case, case_y, visu) A ajouter sans visu
			// addToNeigbors(ctx, current_case+2, case_y, case_x, case_y, visu) A ajouter sans visu
			ret_value += 20
		} else if ctx.Goban[case_y][current_case] == 0 {
			ret_value += 10
		}
	}
	if ret_value != 0 {
		ret_value += score_capture
	}
	return ret_value
}

func verticalCaptureHeuristic(ctx s.SContext, case_x int, case_y int, capture uint8, score_capture int) int {
	count_stone := 0
	current_case := case_y + 1
	ret_value := 0
	for ; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == s.Tnumber(capture) {
			count_stone++
		} else {
			break
		}
	}
	if count_stone == 2 && current_case < int(ctx.NSize) {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			// addToNeigbors(ctx, case_x, current_case-1, case_x, case_y, visu) A ajouter sans visu
			// addToNeigbors(ctx, case_x, current_case-2, case_x, current_case+1, visu) A ajouter sans
			ret_value += 20
		} else if ctx.Goban[current_case][case_x] == 0 {
			ret_value += 10
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
			// addToNeigbors(ctx, case_x, current_case+1, case_x, current_case, visu) A ajouter sans visu
			// addToNeigbors(ctx, case_x, current_case+2, case_x, case_y, visu) A ajouter sans visu
			ret_value += 20
		} else if ctx.Goban[current_case][case_x] == 0 {
			ret_value += 10
		}
	}
	if ret_value != 0 {
		ret_value += score_capture
	}
	return ret_value
}
