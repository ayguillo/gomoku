package heuristic

import (
	g "gomoku/game"
	s "gomoku/structures"
)

type heurCount struct {
	GotFive      int
	GotFour      int
	GotFourMid   int
	GotThree     int
	GotThreeMid  int
	GotThreeOpen int
	GotTwo       int
	GotTwoMid    int
}

type heurStock struct {
	Player   heurCount
	Opponent heurCount
}

func countPlays(ctx s.SContext) heurStock {
	gotThreeOpen, gotThreeOpenOpp := 0, 0
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
							gotThreeOpen += 1
							gotThree += 1
						} else {
							gotThreeOpenOpp += 1
							gotThreeOpp += 1
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

	player := heurCount{
		GotFive:      gotFive,
		GotFour:      gotFour,
		GotFourMid:   gotFourMid,
		GotThree:     gotThree,
		GotThreeMid:  gotThreeMid,
		GotThreeOpen: gotThreeOpen,
		GotTwo:       gotTwo,
		GotTwoMid:    gotTwoMid,
	}

	opponent := heurCount{
		GotFive:      gotFiveOpp,
		GotFour:      gotFourOpp,
		GotFourMid:   gotFourMidOpp,
		GotThree:     gotThreeOpp,
		GotThreeMid:  gotThreeMidOpp,
		GotThreeOpen: gotThreeOpenOpp,
		GotTwo:       gotTwoOpp,
		GotTwoMid:    gotTwoMidOpp,
	}

	return heurStock{Player: player, Opponent: opponent}
}

func CheckImpMoove(ctx s.SContext, neighbors []s.SVertex) []s.SVertex {
	heurPlays := countPlays(ctx)
	player := heurPlays.Player
	opponent := heurPlays.Opponent

	if player.GotFour > 0 || player.GotFourMid > 0 {
		println("got four")
		return getWinPlaysFive(ctx, neighbors)
	} else if opponent.GotFour > 0 {
		println("got four opp")
		return getCounterPlaysFour(ctx, neighbors, opponent.GotFour)
	} else if opponent.GotFourMid > 0 {
		println("got four mid opp")
		return getCounterPlaysFourMid(ctx, neighbors, opponent.GotFourMid)
	} else if player.GotThree == 0 && opponent.GotThree > 0 {
		println("got three op")
		return getCounterPlaysThree(ctx, neighbors, opponent.GotThree)
	}

	return nil
}

func placementHeuristicCopy(ctx s.SContext, case_x int, case_y int) uint8 {
	// Retour 2 pour une obligation de capture, 1 pour un coup ok, 0 sinon
	capture := false
	if ctx.ActiveCapture && len(ctx.Capture) > 0 {
		for _, cap := range ctx.Capture {
			if case_x == cap.X && case_y == cap.Y {
				capture = true
				break
			}
		}
		if capture == false {
			return 2
		}
	}
	if ctx.ActiveDoubleThrees > 0 && !g.CheckDoubleThree(&ctx, case_x, case_y) {
		return 0
	}
	if case_y < 0 || case_y > int(ctx.NSize) {
		return 0
	}
	if case_x < 0 || case_x > int(ctx.NSize) {
		return 0
	}
	if ctx.Goban[int(case_y)][int(case_x)] == 0 {
		return 1
	} else {
		return 0
	}
}

func getWinPlaysFive(ctx s.SContext, neighbors []s.SVertex) []s.SVertex {
	var ret []s.SVertex

	for _, neighbor := range neighbors {
		placement := placementHeuristicCopy(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			heurPlays := countPlays(ctx)

			if heurPlays.Player.GotFive > 0 {
				ret = append(ret, neighbor)
			}

			ctx.Goban[neighbor.Y][neighbor.X] = 0
		}
	}

	if ret == nil {
		return neighbors
	}

	return ret
}

func getCounterPlaysFour(ctx s.SContext, neighbors []s.SVertex, count int) []s.SVertex {
	var ret []s.SVertex

	for _, neighbor := range neighbors {
		placement := placementHeuristicCopy(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			heurPlays := countPlays(ctx)

			if heurPlays.Opponent.GotFour < count {
				ret = append(ret, neighbor)
			}

			ctx.Goban[neighbor.Y][neighbor.X] = 0
		}
	}

	if ret == nil {
		return neighbors
	}

	return ret
}

func getCounterPlaysFourMid(ctx s.SContext, neighbors []s.SVertex, count int) []s.SVertex {
	var ret []s.SVertex

	for _, neighbor := range neighbors {
		placement := placementHeuristicCopy(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			heurPlays := countPlays(ctx)

			if heurPlays.Opponent.GotFourMid < count {
				ret = append(ret, neighbor)
			}

			ctx.Goban[neighbor.Y][neighbor.X] = 0
		}
	}

	if ret == nil {
		return neighbors
	}

	return ret
}

func getCounterPlaysThree(ctx s.SContext, neighbors []s.SVertex, count int) []s.SVertex {
	var ret []s.SVertex

	for _, neighbor := range neighbors {
		placement := placementHeuristicCopy(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			heurPlays := countPlays(ctx)

			if heurPlays.Opponent.GotThree < count {
				ret = append(ret, neighbor)
			}

			ctx.Goban[neighbor.Y][neighbor.X] = 0
		}
	}

	if ret == nil {
		return neighbors
	}

	return ret
}

func getCounterPlaysThreeOpen(ctx s.SContext, neighbors []s.SVertex, count int) []s.SVertex {
	var ret []s.SVertex

	for _, neighbor := range neighbors {
		placement := placementHeuristicCopy(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			heurPlays := countPlays(ctx)

			if heurPlays.Opponent.GotThreeOpen < count {
				ret = append(ret, neighbor)
			}

			ctx.Goban[neighbor.Y][neighbor.X] = 0
		}
	}

	if ret == nil {
		return neighbors
	}

	return ret
}
