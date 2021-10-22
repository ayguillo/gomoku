package game

import (
	s "gomoku/structures"
)

func CounterVertical(ctx *s.SContext, case_x int, case_y int, capturePlayer int) bool {
	current_case := case_y
	case_capture := false
	count_stone := 0
	for ; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == s.Tnumber(capturePlayer) {
			if current_case >= 3 {
				if ctx.Goban[current_case-3][case_x] != s.Tnumber(0) {
					return false
				}
			}
			case_capture = true
			break
		} else if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
	}
	if case_capture == true && count_stone == 2 {
		ctx.Capture = append(ctx.Capture, s.SVertex{X: case_x, Y: current_case - 3})
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
		ctx.Capture = append(ctx.Capture, s.SVertex{X: case_x, Y: current_case})
		return true
	}
	if case_capture == false && count_stone == 2 {
		if case_y > 0 {
			if ctx.Goban[case_y-1][case_x] == s.Tnumber(capturePlayer) {
				ctx.Capture = append(ctx.Capture, s.SVertex{X: case_x, Y: case_y + 2})
				return true
			}
		}
		if current_case >= 0 && current_case < int(ctx.NSize) {
			if ctx.Goban[current_case][case_x] == s.Tnumber(capturePlayer) {
				ctx.Capture = append(ctx.Capture, s.SVertex{X: case_x, Y: current_case + 3})
				return true
			}
		}
	}
	return false
}

func CounterHorizontal(ctx *s.SContext, case_x int, case_y int, capturePlayer int) bool {
	current_case := case_x
	case_capture := false
	count_stone := 0
	for ; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == s.Tnumber(capturePlayer) {
			if current_case >= 3 {
				if ctx.Goban[case_y][current_case-3] != s.Tnumber(0) {
					return false
				}
			}
			case_capture = true
			break
		} else if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			count_stone++
		} else {
			break
		}
	}
	if case_capture == true && count_stone == 2 {
		ctx.Capture = append(ctx.Capture, s.SVertex{X: current_case - 3, Y: case_y})
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
		ctx.Capture = append(ctx.Capture, s.SVertex{X: current_case, Y: case_y})
		return true
	}
	if case_capture == false && count_stone == 2 {
		if case_x > 0 {
			if ctx.Goban[case_y][case_x-1] == s.Tnumber(capturePlayer) {
				ctx.Capture = append(ctx.Capture, s.SVertex{X: case_x + 2, Y: case_y})
				return true
			}
		}
		if current_case >= 0 && current_case < int(ctx.NSize) {
			if ctx.Goban[case_y][current_case] == s.Tnumber(capturePlayer) {
				ctx.Capture = append(ctx.Capture, s.SVertex{X: current_case + 3, Y: case_y})
				return true
			}
		}
	}
	return false
}

func CounterDiagRight(ctx *s.SContext, case_x int, case_y int, capturePlayer int) bool {
	current_case_x, current_case_y := case_x, case_y
	case_capture := false
	count_stone := 0
	for current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize) {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
			if current_case_y >= 3 && current_case_x >= 3 {
				if ctx.Goban[current_case_y-3][current_case_x-3] != s.Tnumber(0) {
					return false
				}
			}
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
		ctx.Capture = append(ctx.Capture, s.SVertex{X: current_case_x - 3, Y: current_case_y - 3})
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
		ctx.Capture = append(ctx.Capture, s.SVertex{X: current_case_x, Y: current_case_y})
		return true
	}
	if case_capture == false && count_stone == 2 {
		if case_x > 0 && case_y > 0 {
			if ctx.Goban[case_y-1][case_x-1] == s.Tnumber(capturePlayer) {
				ctx.Capture = append(ctx.Capture, s.SVertex{X: case_x + 2, Y: case_y + 2})
				return true
			}
		}
		if (current_case_x >= 0 && current_case_x < int(ctx.NSize)) && (current_case_y >= 0 && current_case_y < int(ctx.NSize)) {
			if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
				ctx.Capture = append(ctx.Capture, s.SVertex{X: current_case_x + 3, Y: current_case_y + 3})
				return true
			}
		}
	}
	return false
}

func CounterDiagLeft(ctx *s.SContext, case_x int, case_y int, capturePlayer int) bool {
	current_case_x, current_case_y := case_x, case_y
	case_capture := false
	count_stone := 0

	for current_case_x < int(ctx.NSize) && current_case_y >= 0 {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
			if current_case_x >= 3 && current_case_y < int(ctx.NSize)-3 {
				if ctx.Goban[current_case_y+3][current_case_x-3] != s.Tnumber(0) {
					return false
				}
			}
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
		ctx.Capture = append(ctx.Capture, s.SVertex{X: current_case_x - 3, Y: current_case_y + 3})
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
		ctx.Capture = append(ctx.Capture, s.SVertex{X: current_case_x, Y: current_case_y})
		return true
	}
	if case_capture == false && count_stone == 2 {
		if case_x > 0 && case_y < int(ctx.NSize)-1 {
			if ctx.Goban[case_y+1][case_x-1] == s.Tnumber(capturePlayer) {
				ctx.Capture = append(ctx.Capture, s.SVertex{X: case_x + 2, Y: case_y - 2})
				return true
			}
		}
		if (current_case_x >= 0 && current_case_x < int(ctx.NSize)) && (current_case_y >= 0 && current_case_y < int(ctx.NSize)) {
			if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
				ctx.Capture = append(ctx.Capture, s.SVertex{X: current_case_x + 3, Y: current_case_y - 3})
				return true
			}
		}
	}
	return false
}

func CounterDiag(ctx *s.SContext, case_x int, case_y int, capturePlayer int) bool {
	if CounterDiagRight(ctx, case_x, case_y, capturePlayer) == true {
		return true
	}
	if CounterDiagLeft(ctx, case_x, case_y, capturePlayer) == true {
		return true
	}
	return false
}
