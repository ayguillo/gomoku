package algorithm

import (
	g "gomoku/game"
	s "gomoku/structures"
)

func horizontalHeuristicAlign(ctx s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) int {
	count_stone := 1
	counter := 0
	count_blank_left, count_blank_right := 0, 0
	if g.CounterVertical(&ctx, case_x, case_y, capturePlayer) == true || g.CounterDiag(&ctx, case_x, case_y, capturePlayer) == true ||
		g.CounterHorizontal(&ctx, case_x, case_y, capturePlayer) == true {
		ctx.Capture = s.SVertex{X: -1, Y: -1}
		counter += 15
	}
	for current_case := case_x + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			if g.CounterVertical(&ctx, case_x, case_y, capturePlayer) == true || g.CounterDiag(&ctx, case_x, case_y, capturePlayer) == true ||
				g.CounterHorizontal(&ctx, case_x, case_y, capturePlayer) == true {
				ctx.Capture = s.SVertex{X: -1, Y: -1}
				counter += 15
				if nbCapture >= 4 {
					counter += 200
				}
			}
			count_stone++
		} else {
			if counter != 0 || ctx.Goban[case_y][current_case] == s.Tnumber(capturePlayer) {
				break
			}
			if ctx.Goban[case_y][current_case] == 0 {
				count_blank_right++
				if count_blank_right >= 4 {
					break
				}
			}
		}
	}
	for current_case := case_x - 1; current_case >= 0; current_case-- {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			if g.CounterVertical(&ctx, case_x, case_y, capturePlayer) == true || g.CounterDiag(&ctx, case_x, case_y, capturePlayer) == true ||
				g.CounterHorizontal(&ctx, case_x, case_y, capturePlayer) == true {
				ctx.Capture = s.SVertex{X: -1, Y: -1}
				counter += 15
				if nbCapture >= 4 {
					counter += 200
				}
			}
			count_stone++
		} else {
			if counter != 0 || ctx.Goban[case_y][current_case] == s.Tnumber(capturePlayer) {
				break
			}
			if ctx.Goban[case_y][current_case] == 0 {
				count_blank_left++
				if count_blank_left >= 4 {
					break
				}
			}
		}
	}
	if count_stone > 1 {
		need_place := 5 - count_stone
		count_blank := 0
		if count_blank_left >= need_place && count_blank_right >= need_place {
			count_blank = 3 * need_place
		} else if count_blank_left >= need_place && count_blank_right < need_place || count_blank_left < need_place && count_blank_right >= need_place {
			count_blank = need_place
		}
		if count_stone > 4 {
			count_stone *= 10
		}
		return 10*count_stone + count_blank - counter
	}
	return 0
}

func verticalHeuristicAlign(ctx s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) int {
	count_stone := 1
	counter := 0
	count_blank_top, count_blank_down := 0, 0
	if g.CounterVertical(&ctx, case_x, case_y, capturePlayer) == true || g.CounterDiag(&ctx, case_x, case_y, capturePlayer) == true ||
		g.CounterHorizontal(&ctx, case_x, case_y, capturePlayer) == true {
		ctx.Capture = s.SVertex{X: -1, Y: -1}
		counter += 15
	}
	for current_case := case_y + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if g.CounterVertical(&ctx, case_x, case_y, capturePlayer) == true || g.CounterDiag(&ctx, case_x, case_y, capturePlayer) == true ||
				g.CounterHorizontal(&ctx, case_x, case_y, capturePlayer) == true {
				ctx.Capture = s.SVertex{X: -1, Y: -1}
				counter += 15
				if nbCapture >= 4 {
					counter += 200
				}
			}
			count_stone++
		} else {
			if counter != 0 || ctx.Goban[case_y][current_case] == s.Tnumber(capturePlayer) {
				break
			}
			if ctx.Goban[current_case][case_x] == 0 {
				count_blank_down++
				if count_blank_down >= 4 {
					break
				}
			}
		}
	}
	for current_case := case_y - 1; current_case >= 0; current_case-- {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if g.CounterVertical(&ctx, case_x, case_y, capturePlayer) == true || g.CounterDiag(&ctx, case_x, case_y, capturePlayer) == true ||
				g.CounterHorizontal(&ctx, case_x, case_y, capturePlayer) == true {
				ctx.Capture = s.SVertex{X: -1, Y: -1}
				counter += 15
				if nbCapture >= 4 {
					counter += 200
				}
			}
			count_stone++
		} else {
			if counter != 0 || ctx.Goban[case_y][current_case] == s.Tnumber(capturePlayer) {
				break
			}
			if ctx.Goban[current_case][case_x] == 0 {
				count_blank_top++
				if count_blank_top >= 4 {
					break
				}
			}
		}
	}
	if count_stone > 1 {
		need_place := 5 - count_stone
		count_blank := 0
		if count_blank_top >= need_place && count_blank_down >= need_place {
			count_blank = 3 * need_place
		} else if count_blank_top >= need_place && count_blank_down < need_place || count_blank_top < need_place && count_blank_down >= need_place {
			count_blank = need_place
		}
		if count_stone > 4 {
			count_stone *= 10
		}
		return 10*count_stone + count_blank - counter
	}
	return 0
}

func diagLeftHeuristic(ctx s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) int {
	count_stone := 1
	counter := 0
	count_blank_top, count_blank_down := 0, 0
	if g.CounterVertical(&ctx, case_x, case_y, capturePlayer) == true || g.CounterHorizontal(&ctx, case_x, case_y, capturePlayer) ||
		g.CounterDiag(&ctx, case_x, case_y, capturePlayer) {
		ctx.Capture = s.SVertex{X: -1, Y: -1}
		counter += 15
		if nbCapture >= 4 {
			counter += 200
		}
	}
	for current_case_x, current_case_y := case_x+1, case_y+1; current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if g.CounterVertical(&ctx, current_case_x, current_case_y, capturePlayer) == true || g.CounterHorizontal(&ctx, current_case_x, current_case_y, capturePlayer) ||
				g.CounterDiag(&ctx, current_case_x, current_case_y, capturePlayer) {
				ctx.Capture = s.SVertex{X: -1, Y: -1}
				counter += 15
				if nbCapture >= 4 {
					counter += 200
				}
			}
			count_stone++
			current_case_x++
			current_case_y++
		} else {
			if counter != 0 || ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
				break
			}
			if ctx.Goban[current_case_y][current_case_x] == 0 {
				count_blank_down++
				current_case_x++
				current_case_y++
				if count_blank_down >= 4 {
					break
				}
			}
		}
	}
	for current_case_x, current_case_y := case_x-1, case_y-1; current_case_x >= 0 && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if g.CounterVertical(&ctx, current_case_x, current_case_y, capturePlayer) == true || g.CounterHorizontal(&ctx, current_case_x, current_case_y, capturePlayer) ||
				g.CounterDiag(&ctx, current_case_x, current_case_y, capturePlayer) {
				ctx.Capture = s.SVertex{X: -1, Y: -1}
				counter += 15
				if nbCapture >= 4 {
					counter += 200
				}
			}
			count_stone++
			current_case_x--
			current_case_y--
		} else {
			if counter != 0 || ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
				break
			}
			if ctx.Goban[current_case_y][current_case_x] == 0 {
				count_blank_top++
				current_case_x--
				current_case_y--
				if count_blank_top >= 4 {
					break
				}
			}
		}
	}
	if count_stone > 1 {
		need_place := 5 - count_stone
		count_blank := 0
		if count_blank_top >= need_place && count_blank_down >= need_place {
			count_blank = 3 * need_place
		} else if count_blank_top >= need_place && count_blank_down < need_place || count_blank_top < need_place && count_blank_down >= need_place {
			count_blank = need_place
		}
		if count_stone > 4 {
			count_stone *= 10
		}
		return 10*count_stone + count_blank - counter
	}
	return 0
}

func diagRightHeuristic(ctx s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) int {
	count_stone := 1
	counter := 0
	count_blank_down, count_blank_top := 0, 0
	if g.CounterVertical(&ctx, case_x, case_y, capturePlayer) == true || g.CounterHorizontal(&ctx, case_x, case_y, capturePlayer) ||
		g.CounterDiag(&ctx, case_x, case_y, capturePlayer) {
		ctx.Capture = s.SVertex{X: -1, Y: -1}
		counter += 15
	}
	for current_case_x, current_case_y := case_x+1, case_y-1; current_case_x < int(ctx.NSize) && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if g.CounterVertical(&ctx, current_case_x, current_case_y, capturePlayer) == true || g.CounterHorizontal(&ctx, current_case_x, current_case_y, capturePlayer) ||
				g.CounterDiag(&ctx, current_case_x, current_case_y, capturePlayer) {
				ctx.Capture = s.SVertex{X: -1, Y: -1}
				counter += 15
				if nbCapture >= 4 {
					counter += 200
				}
			}
			count_stone++
			current_case_x++
			current_case_y--
		} else {
			if counter != 0 || ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
				break
			}
			if ctx.Goban[current_case_y][current_case_x] == 0 {
				count_blank_top++
				current_case_x++
				current_case_y--
				if count_blank_top >= 4 {
					break
				}
			}
		}
	}
	for current_case_x, current_case_y := case_x-1, case_y+1; current_case_x >= 0 && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if g.CounterVertical(&ctx, current_case_x, current_case_y, capturePlayer) == true || g.CounterHorizontal(&ctx, current_case_x, current_case_y, capturePlayer) ||
				g.CounterDiag(&ctx, current_case_x, current_case_y, capturePlayer) {
				ctx.Capture = s.SVertex{X: -1, Y: -1}
				counter += 15
				if nbCapture >= 4 {
					counter += 200
				}
			}
			count_stone++
			current_case_x--
			current_case_y++
		} else {
			if counter != 0 || ctx.Goban[current_case_y][current_case_x] == s.Tnumber(capturePlayer) {
				break
			}
			if ctx.Goban[current_case_y][current_case_x] == 0 {
				count_blank_down++
				current_case_x--
				current_case_y++
				if count_blank_down >= 4 {
					break
				}
			}
		}
	}
	if count_stone > 1 {
		need_place := 5 - count_stone
		count_blank := 0
		if count_blank_top >= need_place && count_blank_down >= need_place {
			count_blank = 3 * need_place
		} else if count_blank_top >= need_place && count_blank_down < need_place || count_blank_top < need_place && count_blank_down >= need_place {
			count_blank = need_place
		}
		if count_stone > 4 {
			count_stone *= 10
		}
		return 10*count_stone + count_blank - counter
	}
	return 0
}

func diagonalHeuristicAlign(ctx s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) int {
	left := diagLeftHeuristic(ctx, case_x, case_y, capturePlayer, nbCapture)
	right := diagRightHeuristic(ctx, case_x, case_y, capturePlayer, nbCapture)
	return left + right
}
