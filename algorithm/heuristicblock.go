package algorithm

import (
	s "gomoku/structures"
)

func blockHorizontal(ctx s.SContext, case_x int, case_y int, opponent s.Tnumber) int {
	blank_first, blank_second := 0, 0
	nb_block := 0
	heuristic := 0
	for current_case := case_x + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			break
		} else if ctx.Goban[case_y][current_case] == 0 {
			blank_first += 1
			if blank_first > 1 {
				break
			}
		} else if ctx.Goban[case_y][current_case] == opponent {
			nb_block++
		}
	}
	for current_case := case_x - 1; current_case >= 0; current_case-- {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			break
		} else if ctx.Goban[case_y][current_case] == 0 {
			blank_second += 1
			if blank_second > 1 {
				break
			}
		} else if ctx.Goban[case_y][current_case] == opponent {
			nb_block++
		}
	}
	if nb_block > 2 {
		if nb_block == 3 {
			heuristic += 100
		} else if nb_block == 4 {
			heuristic += 200
		}
	}
	return heuristic
}

func blockVertical(ctx s.SContext, case_x int, case_y int, opponent s.Tnumber) int {
	blank_first, blank_second := 0, 0
	nb_block := 0
	heuristic := 0
	for current_case := case_y + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			break
		} else if ctx.Goban[current_case][case_x] == 0 {
			blank_first += 1
			if blank_first > 1 {
				break
			}
		} else if ctx.Goban[current_case][case_x] == opponent {
			nb_block++
		}
	}
	for current_case := case_y - 1; current_case >= 0; current_case-- {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			break
		} else if ctx.Goban[current_case][case_x] == 0 {
			blank_second += 1
			if blank_second > 1 {
				break
			}
		} else if ctx.Goban[current_case][case_x] == opponent {
			nb_block++
		}
	}
	if nb_block > 2 {
		if nb_block == 3 {
			heuristic += 100
		} else if nb_block == 4 {
			heuristic += 200
		}
	}
	return heuristic
}

func blockDiagLeft(ctx s.SContext, case_x int, case_y int, opponent s.Tnumber) int {
	blank_first, blank_second := 0, 0
	nb_block := 0
	heuristic := 0
	for current_case_x, current_case_y := case_x+1, case_y+1; current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			break
		} else if ctx.Goban[current_case_y][current_case_x] == 0 {
			blank_first += 1
			if blank_first > 1 {
				break
			}
		} else if ctx.Goban[current_case_y][current_case_x] == opponent {
			nb_block++
		}
		current_case_x++
		current_case_y++
	}
	for current_case_x, current_case_y := case_x-1, case_y-1; current_case_x >= 0 && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			break
		} else if ctx.Goban[current_case_y][current_case_x] == 0 {
			blank_second += 1
			if blank_second > 1 {
				break
			}
		} else if ctx.Goban[current_case_y][current_case_x] == opponent {
			nb_block++
		}
		current_case_x--
		current_case_y--
	}
	if nb_block > 2 {
		if nb_block == 3 {
			heuristic += 100
		} else if nb_block == 4 {
			heuristic += 200
		}
	}
	return heuristic
}

func blockDiagRight(ctx s.SContext, case_x int, case_y int, opponent s.Tnumber) int {
	blank_first, blank_second := 0, 0
	nb_block := 0
	heuristic := 0
	for current_case_x, current_case_y := case_x+1, case_y-1; current_case_x < int(ctx.NSize) && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			break
		} else if ctx.Goban[current_case_y][current_case_x] == 0 {
			blank_first += 1
			if blank_first > 1 {
				break
			}
		} else if ctx.Goban[current_case_y][current_case_x] == opponent {
			nb_block++
		}
		current_case_x++
		current_case_y--
	}
	for current_case_x, current_case_y := case_x-1, case_y+1; current_case_x >= 0 && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			break
		} else if ctx.Goban[current_case_y][current_case_x] == 0 {
			blank_second += 1
			if blank_second > 1 {
				break
			}
		} else if ctx.Goban[current_case_y][current_case_x] == opponent {
			nb_block++
		}
		current_case_x--
		current_case_y++
	}
	if nb_block > 2 {
		if nb_block == 3 {
			heuristic += 100
		} else if nb_block == 4 {
			heuristic += 200
		}
	}
	return heuristic
}
