package heuristic

import (
	s "gomoku/structures"
)

func heuristicAlign(ctx s.SContext, case_x int, case_y int, player s.Tnumber) (uint8, bool, bool, bool) {
	nb_align_h, place_ok_h, block_h, middle_h := horizontalHeuristic(ctx, case_x, case_y, player)
	if nb_align_h == 5 && middle_h == false {
		return nb_align_h, place_ok_h, block_h, middle_h
	}
	nb_align_v, place_ok_v, block_v, middle_v := verticalHeuristic(ctx, case_x, case_y, player)
	if nb_align_v == 5 && middle_v == false {
		return nb_align_v, place_ok_v, block_v, middle_v
	}
	nb_align_l, place_ok_l, block_l, middle_l := diagLefttHeuristic(ctx, case_x, case_y, player)
	if nb_align_l == 5 && middle_l == false {
		return nb_align_l, place_ok_l, block_l, middle_l
	}
	nb_align_r, place_ok_r, block_r, middle_r := diagRightHeuristic(ctx, case_x, case_y, player)
	if nb_align_r == 5 && middle_r == false {
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

func Heuristic3(ctx s.SContext) int32 {
	value := 0
	// gotFiveInRow := false
	// gotFiveInRowOpp := false
	gotLiveFour := false
	gotLiveFourOpp := false
	gotLiveEmptyFour := false
	gotLiveEmptyFourOpp := false
	gotDeadFour := false
	gotDeadFourOpp := false
	gotLiveThree := false
	gotLiveThreeOpp := false
	gotLiveEmptyThree := false
	gotLiveEmptyThreeOpp := false
	gotDeadThree := false
	gotDeadThreeOpp := false
	gotLiveTwo := false
	gotLiveTwoOpp := false
	gotLiveEmptyTwo := false
	gotLiveEmptyTwoOpp := false
	gotFive, gotFiveOpp, gotFour, gotFourOpp, gotThree, gotThreeOpp, gotTwo, gotTwoOpp := 0, 0, 0, 0, 0, 0, 0, 0
	gotFourMid, gotFourOppMid, gotThreeMid, gotThreeOppMid, gotTwoMid, gotTwoOppMid := 0, 0, 0, 0, 0, 0
	gotBlockFour, gotBlockFourOpp, gotBlockThree, gotThreeBlockOpp, gotBlockTwo, gotTwoBlockOpp := 0, 0, 0, 0, 0, 0

	nb_capture := ctx.NbCaptureP1
	nb_capture_enemy := ctx.NbCaptureP2
	if ctx.CurrentPlayer == 2 {
		nb_capture = ctx.NbCaptureP2
		nb_capture_enemy = ctx.NbCaptureP1
	}
	if nb_capture >= 5 {
		value += 100000
	} else if nb_capture == 4 || nb_capture == 3 {
		value += nb_capture * 2000
	} else {
		value += nb_capture * 1000
	}
	if nb_capture_enemy >= 5 {
		value -= 150000
	} else if nb_capture_enemy == 4 || nb_capture_enemy == 3 {
		value -= nb_capture_enemy * 3000
	} else {
		value -= nb_capture_enemy * 2000
	}
	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] != 0 {
				nb_align, place_ok, block, middle := heuristicAlign(ctx, x, y, ctx.Goban[y][x])
				if middle == false && block == false {
					if nb_align >= 5 {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							gotFive++
						} else {
							gotFiveOpp++
						}
						// if gotFiveInRow == false && ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
						// 	value += 100000
						// 	gotFiveInRow = true
						// } else if gotFiveInRowOpp == false && ctx.Goban[y][x] != s.Tnumber(ctx.CurrentPlayer) {
						// 	value -= 150000
						// 	gotFiveInRowOpp = true
						// }
					} else if nb_align == 4 && place_ok == true {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							if gotLiveFour == false {
								value += 15000
								gotLiveFour = true
								gotFour++
							} else {
								gotFour++
								value += 10000
							}
						} else {
							if gotLiveFourOpp == false {
								gotLiveFourOpp = true
								value -= 18000
								gotFourOpp += 1
							} else {
								value -= 12000
								gotFour++
							}
						}
					} else if nb_align == 3 && place_ok == true {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							if gotLiveThree == false {
								value += 5000
								gotLiveThree = true
								gotThree += 1
							} else {
								value += 4000
								gotThree++
							}
						} else {
							if gotLiveThreeOpp == false {
								gotLiveThreeOpp = true
								value -= 7000
								gotThreeOpp += 1
							} else {
								value -= 6000
								gotThreeOpp++
							}
						}
					} else if nb_align == 2 && place_ok == true {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							if gotLiveTwo == false {
								value += 2000
								gotLiveTwo = true
								gotTwo += 1
							} else {
								value += 1000
								gotTwo++
							}
						} else {
							if gotLiveTwoOpp == false {
								value -= 3000
								gotLiveTwoOpp = true
								gotTwoOpp++
							} else {
								value -= 2000
								gotTwoOpp++
							}
						}
					}
				} else if middle == true && block == false {
					if nb_align == 5 {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							if gotLiveEmptyFour == false {
								value += 20000
								gotLiveEmptyFour = true
							} else {
								value += 17000
							}
						} else {
							if gotLiveEmptyFourOpp == false {
								value -= 25000
								gotLiveEmptyFourOpp = true
							} else {
								value -= 22000
							}
						}
					} else if nb_align == 4 && place_ok == true {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							if gotLiveEmptyFour == false {
								value += 18000
								gotLiveEmptyFour = true
							} else {
								value += 15000
							}
						} else {
							if gotLiveEmptyFourOpp == false {
								value -= 23000
								gotLiveEmptyFourOpp = true
							} else {
								value -= 20000
							}
						}
					} else if nb_align == 3 && place_ok == true {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							if gotLiveEmptyThree == false {
								value += 8000
								gotLiveEmptyThree = true
							} else {
								value += 6000
							}
						} else {
							if gotLiveEmptyThreeOpp == false {
								value -= 10000
								gotLiveEmptyThreeOpp = true
							} else {
								value -= 8000
							}
						}
					} else if nb_align == 2 && place_ok == true {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							if gotLiveEmptyTwo == false {
								value += 3000
								gotLiveEmptyTwo = true
							} else {
								value += 2000
							}
						} else {
							if gotLiveEmptyTwoOpp == false {
								value -= 5000
							} else {
								value -= 4000
							}
						}
					}
				} else if block == true {
					if nb_align == 4 && place_ok == false {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							if gotDeadFour == false {
								value -= 300
								gotDeadFour = true
								gotBlockFour += 1
							} else {
								gotBlockFour += 1
								value -= 200
							}
						} else {
							if gotDeadFourOpp == false {
								value += 50000
								gotBlockFourOpp += 1
								gotDeadFourOpp = true
							} else {
								gotBlockFourOpp += 1
								value += 45000
							}
						}
					} else if nb_align == 3 && place_ok == false {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							if gotDeadThree == false {
								value -= 300
								gotDeadThree = true
								gotBlockThree += 1
							} else {
								gotBlockThree += 1
								value -= 200
							}
						} else {
							if gotDeadThreeOpp == false {
								gotThreeBlockOpp += 1
								value += 10000
								gotDeadThreeOpp = true
							} else {
								gotThreeBlockOpp += 1
								value += 8000
							}
						}
					} else if nb_align == 2 && place_ok == false {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							gotBlockTwo += 1
							value += 1000
						} else {
							gotTwoBlockOpp += 1
							value -= 1500
						}
					} else if nb_align == 4 && place_ok == true {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							gotFourMid += 1
						} else {
							gotFourOppMid += 1
						}
					} else if nb_align == 3 && place_ok == true {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							gotThreeMid += 1
						} else {
							gotThreeOppMid += 1
						}
					} else if nb_align == 2 && place_ok == true {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							gotTwoMid += 1
						} else {
							gotTwoOppMid += 1
						}
					}
				}
			}
		}
	}
	// fmt.Println("Me", gotFive, gotFour, gotFourMid, gotThree, gotThreeMid, gotTwo, gotTwoMid)
	// fmt.Println("Opp", gotFiveOpp, gotFourOpp, gotFourOppMid, gotThreeOpp, gotThreeOppMid, gotTwoOpp, gotTwoOppMid)
	// // fmt.Println("Block me", gotBlockFour, gotBlockThree, gotBlockTwo)
	// // fmt.Println("Block opp", gotBlockFourOpp, gotThreeBlockOpp, gotTwoBlockOpp)
	value = 6000*(gotFive-gotFiveOpp) + 4800*(gotFour-gotFourOpp) + 500*(gotFourMid-gotFourOppMid) + 500*(gotThree-gotThreeOpp) + 200*(gotThreeMid-gotThreeOppMid) + 50*(gotTwo-gotTwoOpp) + 10*(gotTwoMid-gotTwoOppMid)
	return int32(value)
}
