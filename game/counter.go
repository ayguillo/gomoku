package game

import (
	s "gomoku/structures"
)

func counterVertical(ctx s.SContext, case_x int, case_y int, capturePlayer int) bool {
	current_case := case_y
	case_capture := false
	count_stone := 0
	for ; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == s.Tnumber(capturePlayer) {
			case_capture = true
			break
		} else if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
	}
	if case_capture == true && count_stone == 2 {
		return true
	}
	if count_stone < 2 {
		count_stone = 0
		current_case = case_y
		for ; current_case >= 0; current_case-- {
			if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
				count_stone++
			} else {
				break
			}
		}
	}
	if count_stone == 2 && case_capture == true {
		return true
	}
	if case_capture == false && count_stone == 2 {
		if case_y > 0 {
			if ctx.Goban[case_y-1][case_x] == s.Tnumber(capturePlayer) {
				return true
			}
		} else if current_case >= 0 {
			if ctx.Goban[current_case][case_x] == s.Tnumber(capturePlayer) {
				return true
			}
		}
	}
	return false
}

func counterHorizontal(ctx s.SContext, case_x int, case_y int, capturePlayer int) bool {
	current_case := case_x
	case_capture := false
	count_stone := 0
	for ; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == s.Tnumber(capturePlayer) {
			case_capture = true
			break
		} else if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
	}
	if case_capture == true && count_stone == 2 {
		return true
	}
	if count_stone < 2 {
		count_stone = 0
		current_case = case_x
		for ; current_case >= 0; current_case-- {
			if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
				count_stone++
			} else {
				break
			}
		}
	}
	if count_stone == 2 && case_capture == true {
		return true
	}
	if case_capture == false && count_stone == 2 {
		if case_x > 0 {
			if ctx.Goban[case_y][case_x-1] == s.Tnumber(capturePlayer) {
				return true
			}
		} else if current_case >= 0 {
			if ctx.Goban[case_y][current_case] == s.Tnumber(capturePlayer) {
				return true
			}
		}
	}
	return false
}

func counterDiagRight(ctx s.SContext, case_x int, case_y int, capturePlayer int) bool {
	current_case_x, current_case_y := case_x, case_y
	case_capture := false
	count_stone := 0
	for current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
			case_capture = true
			break
		} else if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
		current_case_x++
		current_case_y++
	}
	if case_capture == true && count_stone == 2 {
		return true
	}
	if count_stone < 2 {
		count_stone = 0
		current_case_x, current_case_y = case_x, case_y
		for current_case_x >= 0 && current_case_y >= 0 {
			if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
				count_stone++
			} else {
				break
			}
			current_case_x--
			current_case_y--
		}
	}
	if count_stone == 2 && case_capture == true {
		return true
	}
	if case_capture == false && count_stone == 2 {
		if case_x > 0 && case_y > 0 {
			if ctx.Goban[case_y-1][case_x-1] == s.Tnumber(capturePlayer) {
				return true
			}
		} else if current_case_x >= 0 && current_case_y >= 0 {
			if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
				return true
			}
		}
	}
	return false
}

func counterDiagLeft(ctx s.SContext, case_x int, case_y int, capturePlayer int) bool {
	current_case_x, current_case_y := case_x, case_y
	case_capture := false
	count_stone := 0
	for current_case_x < int(ctx.NSize) && current_case_y >= 0 {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
			case_capture = true
			break
		} else if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
		current_case_x++
		current_case_y--
	}
	if case_capture == true && count_stone == 2 {
		return true
	}
	if count_stone < 2 {
		count_stone = 0
		current_case_x, current_case_y = case_x, case_y
		for current_case_x >= 0 && current_case_y < int(ctx.NSize) {
			if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
				count_stone++
			} else {
				break
			}
			current_case_x--
			current_case_y++
		}
	}
	if count_stone == 2 && case_capture == true {
		return true
	}
	if case_capture == false && count_stone == 2 {
		if case_x > 0 && case_y < int(ctx.NSize)-1 {
			if ctx.Goban[case_y+1][case_x-1] == s.Tnumber(capturePlayer) {
				return true
			}
		} else if current_case_x >= 0 && current_case_y >= 0 {
			if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
				return true
			}
		}
	}
	return false
}

func counterDiag(ctx s.SContext, case_x int, case_y int, capturePlayer int) bool {
	if counterDiagRight(ctx, case_x, case_y, capturePlayer) == true {
		return true
	}
	if counterDiagLeft(ctx, case_x, case_y, capturePlayer) == true {
		return true
	}
	return false
}
