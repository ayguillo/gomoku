package heuristic

import (
	s "gomoku/structures"
)

func horizontalHeuristic(ctx s.SContext, case_x int, case_y int, player s.Tnumber) (uint8, bool, bool, bool) {
	empty_cases_left := 0
	empty_cases_right := 0
	align_cases := uint8(1)
	enemy_player := s.Tnumber(2)
	block := false
	end := false
	middle := false
	place_ok := false

	if player == 2 {
		enemy_player = 1
	}
	for current_case := case_x + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == player {
			if empty_cases_right != 0 {
				middle = true
			}
			align_cases++
		} else if ctx.Goban[case_y][current_case] == enemy_player {
			block = true
			break
		} else {
			empty_cases_right++
		}
		if empty_cases_right >= 5-int(align_cases) {
			place_ok = true
			break
		}
		if align_cases >= 5 {
			end = true
			break
		}
	}
	if end == true {
		return align_cases, place_ok, block, middle
	}
	for current_case := case_x - 1; current_case >= 0; current_case-- {
		if ctx.Goban[case_y][current_case] == player {
			align_cases++
			if empty_cases_left != 0 {
				middle = true
			}
		} else if ctx.Goban[case_y][current_case] == enemy_player {
			block = true
			break
		} else {
			empty_cases_left++
		}
		if empty_cases_left >= 5-int(align_cases) {
			place_ok = true
			break
		}
		if align_cases >= 5 {
			break
		}
	}
	if empty_cases_left+empty_cases_right >= 5-int(align_cases) {
		place_ok = true
	}
	return align_cases, place_ok, block, middle
}

func verticalHeuristic(ctx s.SContext, case_x int, case_y int, player s.Tnumber) (uint8, bool, bool, bool) {
	empty_cases_top := 0
	empty_cases_down := 0
	align_cases := uint8(1)
	enemy_player := s.Tnumber(1)
	block := false
	end := false
	middle := false
	place_ok := false

	if player == 1 {
		enemy_player = 2
	}
	for current_case := case_y + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == player {
			if empty_cases_down != 0 {
				middle = true
			}
			align_cases++
		} else if ctx.Goban[case_y][current_case] == enemy_player {
			block = true
			break
		} else {
			empty_cases_down++
		}
		if empty_cases_down >= 5-int(align_cases) {
			place_ok = true
			break
		}
		if align_cases >= 5 {
			end = true
			break
		}
	}
	if end == true {
		return align_cases, place_ok, block, middle
	}
	for current_case := case_y - 1; current_case >= 0; current_case-- {
		if ctx.Goban[current_case][case_x] == player {
			align_cases++
			if empty_cases_top != 0 {
				middle = true
			}
		} else if ctx.Goban[case_y][current_case] == enemy_player {
			block = true
			break
		} else {
			empty_cases_top++
		}
		if empty_cases_top >= 5-int(align_cases) {
			place_ok = true
			break
		}
		if align_cases >= 5 {
			break
		}
	}
	if empty_cases_down+empty_cases_top >= 5-int(align_cases) {
		place_ok = true
	}
	return align_cases, place_ok, block, middle
}

func diagRightHeuristic(ctx s.SContext, case_x int, case_y int, player s.Tnumber) (uint8, bool, bool, bool) {
	empty_cases_top := 0
	empty_cases_down := 0
	align_cases := uint8(1)
	enemy_player := s.Tnumber(1)
	block := false
	end := false
	middle := false
	place_ok := false

	if player == 1 {
		enemy_player = 2
	}
	for current_case_x, current_case_y := case_x+1, case_y-1; current_case_x < int(ctx.NSize) && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == player {
			if empty_cases_top != 0 {
				middle = true
			}
			align_cases++
		} else if ctx.Goban[current_case_y][current_case_x] == enemy_player {
			block = true
			break
		} else {
			empty_cases_top++
		}
		if empty_cases_top >= 5-int(align_cases) {
			place_ok = true
			break
		}
		if align_cases >= 5 {
			end = true
			break
		}
		current_case_x++
		current_case_y--
	}
	if end == true {
		return align_cases, place_ok, block, middle
	}
	for current_case_x, current_case_y := case_x-1, case_y+1; current_case_x >= 0 && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == player {
			if empty_cases_down != 0 {
				middle = true
			}
			align_cases++
		} else if ctx.Goban[current_case_y][current_case_x] == enemy_player {
			block = true
			break
		} else {
			empty_cases_down++
		}
		if empty_cases_down >= 5-int(align_cases) {
			place_ok = true
			break
		}
		if align_cases >= 5 {
			end = true
			break
		}
		current_case_x--
		current_case_y++
	}
	if empty_cases_top+empty_cases_down >= 5-int(align_cases) {
		place_ok = true
	}
	return align_cases, place_ok, block, middle
}

func diagLefttHeuristic(ctx s.SContext, case_x int, case_y int, player s.Tnumber) (uint8, bool, bool, bool) {
	empty_cases_top := 0
	empty_cases_down := 0
	align_cases := uint8(1)
	enemy_player := s.Tnumber(1)
	block := false
	end := false
	middle := false
	place_ok := false

	if player == 1 {
		enemy_player = 2
	}
	for current_case_x, current_case_y := case_x+1, case_y+1; current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == player {
			if empty_cases_down != 0 {
				middle = true
			}
			align_cases++
		} else if ctx.Goban[current_case_y][current_case_x] == enemy_player {
			block = true
			break
		} else {
			empty_cases_down++
		}
		if empty_cases_down >= 5-int(align_cases) {
			place_ok = true
			break
		}
		if align_cases >= 5 {
			end = true
			break
		}
		current_case_x++
		current_case_y++
	}
	if end == true {
		return align_cases, place_ok, block, middle
	}
	for current_case_x, current_case_y := case_x-1, case_y-1; current_case_x >= 0 && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == player {
			if empty_cases_top != 0 {
				middle = true
			}
			align_cases++
		} else if ctx.Goban[current_case_y][current_case_x] == enemy_player {
			block = true
			break
		} else {
			empty_cases_top++
		}
		if empty_cases_top >= 5-int(align_cases) {
			place_ok = true
			break
		}
		if align_cases >= 5 {
			end = true
			break
		}
		current_case_x--
		current_case_y--
	}
	if empty_cases_top+empty_cases_down >= 5-int(align_cases) {
		place_ok = true
	}
	return align_cases, place_ok, block, middle
}
