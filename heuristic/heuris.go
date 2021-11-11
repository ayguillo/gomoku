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
	gotFive, gotFiveOpp, gotFour, gotFourOpp, gotThree, gotThreeOpp, gotTwo, gotTwoOpp := 0, 0, 0, 0, 0, 0, 0, 0
	gotFourMid, gotFourMidOpp, gotThreeMid, gotThreeMidOpp, gotTwoMid, gotTwoMidOpp := 0, 0, 0, 0, 0, 0

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

	value = 100000*(gotFive-gotFiveOpp) + 10000*(gotFour-gotFourOpp) + 1000*(gotFourMid-gotFourMidOpp) + 1000*(gotThree-gotThreeOpp) + 100*(gotThreeMid-gotThreeMidOpp) + 100*(gotTwo-gotTwoOpp) + 10*(gotTwoMid-gotTwoMidOpp)

	// value = 6000*(gotFive-gotFiveOpp) + 4800*(gotFour-gotFourOpp) + 500*(gotFourMid-gotFourMidOpp) + 500*(gotThree-gotThreeOpp) + 100*(gotThreeMid-gotThreeMidOpp) + 25*(gotTwo-gotTwoOpp) + 5*(gotTwoMid-gotTwoMidOpp)
	// value = 1000000*(gotFive) + 100000*(gotFour) + 10000*(gotFourMid) + 10000*(gotThree) + 1000*(gotThreeMid) + 100*(gotTwo) + 10*(gotTwoMid)

	// valueMe := 1000000*gotFive + 100000*gotFour + 10000*gotFourMid + 10000*gotThree + 1000*gotThreeMid
	// valueOp := 150000*gotFiveOpp + 80000*gotFourOpp + 8000*gotFourMidOpp + 8000*gotThreeOpp + 800*gotThreeMidOpp

	// valueMe := 150000*gotFive + 5000*gotFour + 1000*gotFourMid + 1000*gotThree + 200*gotThreeMid + 50*gotTwo
	// valueOp := 150000*gotFiveOpp + 5000*2*gotFourOpp + 2000*gotFourMidOpp + 2000*gotThreeOpp + 400*gotThreeMidOpp + 100*gotTwoOpp

	// value = valueMe - valueOp

	if ctx.ActiveCapture {
		nb_capture := ctx.NbCaptureP1
		nb_capture_enemy := ctx.NbCaptureP2

		if ctx.CurrentPlayer == 2 {
			nb_capture = ctx.NbCaptureP2
			nb_capture_enemy = ctx.NbCaptureP1
		}

		if nb_capture >= 5 {
			value += 100000
		} else if nb_capture == 4 {
			value += 10000
		} else if nb_capture == 3 {
			value += 500
		} else if nb_capture == 2 {
			value += 250
		} else if nb_capture == 1 {
			value += 100
		}
		if nb_capture_enemy >= 5 {
			value -= 100000
		} else if nb_capture_enemy == 4 {
			value -= 10000
		} else if nb_capture_enemy == 3 {
			value -= 500
		} else if nb_capture_enemy == 2 {
			value -= 250
		} else if nb_capture_enemy == 1 {
			value -= 100
		}
	}

	return int32(value)
}
