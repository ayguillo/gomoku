package algogo

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
		} else if ctx.Goban[current_case][case_x] == enemy_player {
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
		} else if ctx.Goban[current_case][case_x] == enemy_player {
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

func heuristicAlign(ctx s.SContext, case_x int, case_y int, player s.Tnumber) (uint8, bool, bool, bool) {
	nb_align_h, place_ok_h, block_h, middle_h := horizontalHeuristic(ctx, case_x, case_y, player)
	if nb_align_h >= 5 && middle_h == false {
		return nb_align_h, place_ok_h, block_h, middle_h
	}
	nb_align_v, place_ok_v, block_v, middle_v := verticalHeuristic(ctx, case_x, case_y, player)
	if nb_align_v >= 5 && middle_v == false {
		return nb_align_v, place_ok_v, block_v, middle_v
	}
	nb_align_l, place_ok_l, block_l, middle_l := diagLefttHeuristic(ctx, case_x, case_y, player)
	if nb_align_l >= 5 && middle_l == false {
		return nb_align_l, place_ok_l, block_l, middle_l
	}
	nb_align_r, place_ok_r, block_r, middle_r := diagRightHeuristic(ctx, case_x, case_y, player)
	if nb_align_r >= 5 && middle_r == false {
		return nb_align_r, place_ok_r, block_r, middle_r
	}
	if 3*nb_align_h > (nb_align_v+nb_align_l+nb_align_r) && place_ok_h == true {
		return nb_align_h, place_ok_h, block_h, middle_h
	} else if 3*nb_align_v > (nb_align_h+nb_align_r+nb_align_l) && place_ok_v == true {
		return nb_align_v, place_ok_v, block_v, middle_v
	} else if 3*nb_align_l > (nb_align_h+nb_align_l+nb_align_v) && place_ok_l == true {
		return nb_align_l, place_ok_l, block_l, middle_l
	} else if 3*nb_align_r > (nb_align_h+nb_align_l+nb_align_v) && place_ok_r == true {
		return nb_align_r, place_ok_r, block_r, middle_r
	} else if 3*nb_align_h > (nb_align_v+nb_align_l+nb_align_r) && place_ok_h == false && block_h == false {
		return nb_align_h, place_ok_h, block_h, middle_h
	} else if 3*nb_align_v > (nb_align_h+nb_align_r+nb_align_l) && place_ok_v == false && block_v == false {
		return nb_align_v, place_ok_v, block_v, middle_v
	} else if 3*nb_align_l > (nb_align_h+nb_align_l+nb_align_v) && place_ok_l == false && block_l == false {
		return nb_align_l, place_ok_l, block_l, middle_l
	} else if 3*nb_align_r > (nb_align_h+nb_align_l+nb_align_v) && place_ok_r == false && block_l == false {
		return nb_align_r, place_ok_r, block_r, middle_r
	}
	nb_align, place_ok, block, middle := nb_align_h, place_ok_h, block_h, middle_h
	return nb_align, place_ok, block, middle
}

func CalcHeuristic(ctx s.SContext, y int, x int) int32 {
	value := 0
	gotFour, gotFourOpp, gotThree, gotThreeOpp, gotTwo, gotTwoOpp := 0, 0, 0, 0, 0, 0
	gotFourMid, gotFourMidOpp, gotThreeMid, gotThreeMidOpp, gotTwoMid, gotTwoMidOpp := 0, 0, 0, 0, 0, 0

	if ctx.Goban[y][x] != 0 {
		nb_align, place_ok, block, middle := heuristicAlign(ctx, x, y, ctx.Goban[y][x])
		if nb_align >= 5 {
			if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
				return 100000
			} else {
				return -100000
			}
		}
		if middle == false && block == false && place_ok == true { // libre
			if nb_align == 4 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotFour++
				} else {
					gotFourOpp++
				}
			} else if nb_align == 3 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotThree++
				} else {
					gotThreeOpp++
				}
			} else if nb_align == 2 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotTwo++
				} else {
					gotTwoOpp++
				}
			}
		} else if middle == true && block == false && place_ok == true {
			if nb_align >= 4 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotFourMid += 1
				} else {
					gotFourMidOpp += 1
				}
			} else if nb_align == 3 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotThreeMid += 1
				} else {
					gotThreeMidOpp += 1
				}
			} else if nb_align == 2 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotTwoMid += 1
				} else {
					gotTwoMidOpp += 1
				}
			}
		} else if middle == true && block == true && place_ok == true {
			if nb_align >= 4 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotFourMid += 1
				} else {
					gotFourMidOpp += 1
				}
			} else if nb_align == 3 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotThreeMid += 1
				} else {
					gotThreeMidOpp += 1
				}
			} else if nb_align == 2 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotTwoMid += 1
				} else {
					gotTwoMidOpp += 1
				}
			}
		} else if middle == false && block == true && place_ok == true { // bloquer + 1 cote libre si place_ok
			if nb_align == 4 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotFourMid += 1
				} else {
					gotFourMidOpp += 1
				}
			} else if nb_align == 3 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotThreeMid += 1
				} else {
					gotThreeMidOpp += 1
				}
			} else if nb_align == 2 {
				if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
					gotTwoMid += 1
				} else {
					gotTwoMidOpp += 1
				}
			}
		}
	}

	value = 10000*(gotFour-gotFourOpp) + 1000*(gotFourMid-gotFourMidOpp) + 1000*(gotThree-gotThreeOpp) + 100*(gotThreeMid-gotThreeMidOpp) + 100*(gotTwo-gotTwoOpp) + 10*(gotTwoMid-gotTwoMidOpp)

	return int32(value)
}
