package heuristic

import (
	s "gomoku/structures"
)

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

func CalcHeuristic(ctx s.SContext) int32 {
	value := 0
	// gotFiveInRow := false
	// // gotFiveInRowOpp := false
	// gotLiveEmptyFour := false
	// gotLiveEmptyFourOpp := false
	// gotLiveEmptyThree := false
	// gotLiveEmptyThreeOpp := false
	// gotLiveEmptyTwo := false
	// gotLiveEmptyTwoOpp := false
	gotFive, gotFiveOpp, gotFour, gotFourOpp, gotThree, gotThreeOpp, gotTwo, gotTwoOpp := 0, 0, 0, 0, 0, 0, 0, 0
	gotFourMid, gotFourMidOpp, gotThreeMid, gotThreeMidOpp, gotTwoMid, gotTwoMidOpp := 0, 0, 0, 0, 0, 0
	// gotFourBlock, gotFourBlockOpp, gotThreeBlock, gotThreeBlockOpp, gotTwoBlock, gotTwoBlockOpp := 0, 0, 0, 0, 0, 0

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
				if nb_align >= 5 {
					if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
						gotFive++
					} else {
						gotFiveOpp++
					}
					continue
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
		}
	}

	value = 100000*(gotFive-gotFiveOpp) + 15000*(gotFour-gotFourOpp) + 10000*(gotFourMid-gotFourMidOpp) + 10000*(gotThree-gotThreeOpp) + 200*(gotThreeMid-gotThreeMidOpp) + 200*(gotTwo-gotTwoOpp) + 50*(gotTwoMid-gotTwoMidOpp)
	// value = 1000000*(gotFive) + 15000*(gotFour) + 10000*(gotFourMid) + 10000*(gotThree) + 500*(gotThreeMid) + 200*(gotTwo) + 100*(gotTwoMid)

	// valueMe := 150000*gotFive + 5000*gotFour + 1000*gotFourMid + 500*gotThree + 200*gotThreeMid + 50*gotTwo + 20*gotTwoMid
	// valueOp := 150000*gotFiveOpp + 10000*gotFourOpp + 5000*gotFourMidOpp + 1000*gotThreeOpp + 200*gotThreeMidOpp + 50*gotTwoOpp + 20*gotTwoMidOpp

	// value = valueMe - valueOp

	// if nb_capture >= 5 {
	// 	value += 100000
	// } else if nb_capture == 4 || nb_capture == 3 {
	// 	value += nb_capture * 2000
	// } else {
	// 	value += nb_capture * 1000
	// }
	// if nb_capture_enemy >= 5 {
	// 	value -= 150000
	// } else if nb_capture_enemy == 4 || nb_capture_enemy == 3 {
	// 	value -= nb_capture_enemy * 3000
	// } else {
	// 	value -= nb_capture_enemy * 2000
	// }

	// println("on est :", ctx.CurrentPlayer)
	// fmt.Println("Me:", gotFive, gotFour, gotFourMid, gotThree, gotThreeMid, gotTwo, gotTwoMid, valueMe)
	// fmt.Println("Op:", gotFiveOpp, gotFourOpp, gotFourMidOpp, gotThreeOpp, gotThreeMidOpp, gotTwoOpp, gotTwoMidOpp, valueOp)
	// println()
	// // fmt.Println("Block me", gotBlockFour, gotBlockThree, gotBlockTwo)
	// // fmt.Println("Block opp", gotBlockFourOpp, gotThreeBlockOpp, gotTwoBlockOpp)

	return int32(value)
}