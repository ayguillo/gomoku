package algogo

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

func EvaluateMove(ctx s.SContext, case_x int, case_y int) int32 {
	value := 0
	gotFive, gotFiveOpp, gotFour, gotFourOpp, gotThree, gotThreeOpp, gotTwo, gotTwoOpp := 0, 0, 0, 0, 0, 0, 0, 0
	gotFourMid, gotFourMidOpp, gotThreeMid, gotThreeMidOpp, gotTwoMid, gotTwoMidOpp := 0, 0, 0, 0, 0, 0
	gotFourMidPlus, gotFourMidPlusOpp, gotThreeMidPlus, gotThreeMidPlusOpp, gotTwoMidPlus, gotTwoMidPlusOpp := 0, 0, 0, 0, 0, 0
	if ctx.Goban[case_y][case_x] != 0 {
		nb_align, place_ok, block, middle := heuristicAlign(ctx, case_x, case_y, ctx.Goban[case_y][case_x])
		if nb_align >= 5 {
			if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
				gotFive++
			} else {
				gotFiveOpp++
			}
		} else if middle == false && block == false && place_ok == true {
			if nb_align == 4 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotFour++
				} else {
					gotFourOpp++
				}
			} else if nb_align == 3 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotThree++
				} else {
					gotThreeOpp++
				}
			} else if nb_align == 2 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotTwo++
				} else {
					gotTwoOpp++
				}
			}
		} else if middle == true && block == false && place_ok == true {
			if nb_align >= 4 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotFourMidPlus += 1
				} else {
					gotFourMidPlusOpp += 1
				}
			} else if nb_align == 3 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotThreeMidPlus += 1
				} else {
					gotThreeMidPlusOpp += 1
				}
			} else if nb_align == 2 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotTwoMidPlus += 1
				} else {
					gotTwoMidPlusOpp += 1
				}
			}
		} else if middle == true && block == true && place_ok == true {
			if nb_align >= 4 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotFourMid += 1
				} else {
					gotFourMidOpp += 1
				}
			} else if nb_align == 3 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotThreeMid += 1
				} else {
					gotThreeMidOpp += 1
				}
			} else if nb_align == 2 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotTwoMid += 1
				} else {
					gotTwoMidOpp += 1
				}
			}
		} else if middle == false && block == true && place_ok == true { // bloquer + 1 cote libre si place_ok
			if nb_align == 4 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotFourMid += 1
				} else {
					gotFourMidOpp += 1
				}
			} else if nb_align == 3 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotThreeMid += 1
				} else {
					gotThreeMidOpp += 1
				}
			} else if nb_align == 2 {
				if ctx.Goban[case_y][case_x] == s.Tnumber(ctx.CurrentPlayer) {
					gotTwoMid += 1
				} else {
					gotTwoMidOpp += 1
				}
			}
		}

	}

	// value = 1000000*(gotFive-gotFiveOpp) + 100000*(gotFour-gotFourOpp) + 1000*(gotFourMid-gotFourMidOpp) + 1500*(gotThree-gotThreeOpp) + 200*(gotThreeMid-gotThreeMidOpp) + 50*(gotTwo-gotTwoOpp) + 10*(gotTwoMid-gotTwoMidOpp)
	value = 60000*(gotFive-gotFiveOpp) + 4800*(gotFour-gotFourOpp) + 1000*(gotFourMid-gotFourMidOpp) + 1000*(gotThree-gotThreeOpp) + 300*(gotThreeMid-gotThreeMidOpp) + 50*(gotTwo-gotTwoOpp) + 10*(gotTwoMid-gotTwoMidOpp)
	value += 1500*(gotFourMidPlus-gotFourMidPlusOpp) + 750*(gotThreeMidPlus-gotThreeMidPlusOpp) + 30*(gotTwoMidPlus-gotTwoMidPlusOpp)

	if ctx.ActiveCapture {
		nbCapture := ctx.NbCaptureP1
		nbCaptureOpp := ctx.NbCaptureP2

		if ctx.CurrentPlayer == 2 {
			nbCapture = ctx.NbCaptureP2
			nbCaptureOpp = ctx.NbCaptureP1
		}

		if nbCapture >= 5 {
			value += 1 * 5
		} else if nbCaptureOpp >= 5 {
			value -= 1 * 5
		} else if nbCapture == 4 {
			value += 1 * 4
		} else if nbCaptureOpp == 4 {
			value -= 1 * 4
		} else {
			value += 1 * (nbCapture - nbCaptureOpp)
		}
	}

	return int32(value)
}

func EvaluateGoban(ctx s.SContext) int32 {
	value := 0
	gotFive, gotFiveOpp, gotFour, gotFourOpp, gotThree, gotThreeOpp, gotTwo, gotTwoOpp := 0, 0, 0, 0, 0, 0, 0, 0
	gotFourMid, gotFourMidOpp, gotThreeMid, gotThreeMidOpp, gotTwoMid, gotTwoMidOpp := 0, 0, 0, 0, 0, 0
	gotFourMidPlus, gotFourMidPlusOpp, gotThreeMidPlus, gotThreeMidPlusOpp, gotTwoMidPlus, gotTwoMidPlusOpp := 0, 0, 0, 0, 0, 0

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
							gotFourMidPlus += 1
						} else {
							gotFourMidPlusOpp += 1
						}
					} else if nb_align == 3 {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							gotThreeMidPlus += 1
						} else {
							gotThreeMidPlusOpp += 1
						}
					} else if nb_align == 2 {
						if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
							gotTwoMidPlus += 1
						} else {
							gotTwoMidPlusOpp += 1
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

	// value = 1000000*(gotFive-gotFiveOpp) + 100000*(gotFour-gotFourOpp) + 1000*(gotFourMid-gotFourMidOpp) + 1500*(gotThree-gotThreeOpp) + 200*(gotThreeMid-gotThreeMidOpp) + 50*(gotTwo-gotTwoOpp) + 10*(gotTwoMid-gotTwoMidOpp)
	value = 60000*(gotFive-gotFiveOpp) + 4800*(gotFour-gotFourOpp) + 1000*(gotFourMid-gotFourMidOpp) + 1000*(gotThree-gotThreeOpp) + 300*(gotThreeMid-gotThreeMidOpp) + 50*(gotTwo-gotTwoOpp) + 10*(gotTwoMid-gotTwoMidOpp)
	value += 1500*(gotFourMidPlus-gotFourMidPlusOpp) + 750*(gotThreeMidPlus-gotThreeMidPlusOpp) + 30*(gotTwoMidPlus-gotTwoMidPlusOpp)

	if ctx.ActiveCapture {
		nbCapture := ctx.NbCaptureP1
		nbCaptureOpp := ctx.NbCaptureP2

		if ctx.CurrentPlayer == 2 {
			nbCapture = ctx.NbCaptureP2
			nbCaptureOpp = ctx.NbCaptureP1
		}

		println(ctx.CurrentPlayer, nbCapture, nbCaptureOpp)

		if nbCapture >= 5 {
			value += 1 * 5
		} else if nbCaptureOpp >= 5 {
			value -= 1 * 5
		} else if nbCapture == 4 {
			value += 1 * 4
		} else if nbCaptureOpp == 4 {
			value -= 1 * 4
		} else {
			value += 1 * (nbCapture - nbCaptureOpp)
		}
	}

	return int32(value)
}
