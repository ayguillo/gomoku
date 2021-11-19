package algogo

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
				} else if middle == false && block == true && place_ok == true {
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

func checkImpCapture(ctx s.SContext, neighbors []s.SVertex) ([]s.SVertex, []s.SVertex) {
	var playerCap []s.SVertex
	var oppCap []s.SVertex

	player1Cap := ctx.NbCaptureP1
	player2Cap := ctx.NbCaptureP2

	goban := copyGoban(ctx.Goban)

	for _, neighbor := range neighbors {
		placement := placementHeuristicCopy(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)

			capturesVertex := CaptureAlgoCtx(&ctx, neighbor.X, neighbor.Y)

			if capturesVertex != nil {
				playerCap = append(playerCap, capturesVertex...)
			}

			goban[neighbor.Y][neighbor.X] = 0
		}
	}

	ctx.NbCaptureP1 = player1Cap
	ctx.NbCaptureP2 = player2Cap

	return playerCap, oppCap
}

func CheckImpMove(ctx s.SContext, neighbors []s.SVertex) []s.SVertex {
	heurPlays := countPlays(ctx)
	player := heurPlays.Player
	opponent := heurPlays.Opponent

	playerNbCap := ctx.NbCaptureP1
	oppNbCap := ctx.NbCaptureP2

	if ctx.CurrentPlayer == 2 {
		playerNbCap = ctx.NbCaptureP2
		oppNbCap = ctx.NbCaptureP1
	}

	var playerCap []s.SVertex = nil
	var oppCap []s.SVertex = nil

	if isCapture {
		playerCap, oppCap = checkImpCapture(ctx, neighbors)
	}

	if playerNbCap >= 4 && playerCap != nil {
		return playerCap
	} else if player.GotFour > 0 || player.GotFourMid > 0 {
		return getWinPlaysFive(ctx, neighbors)
	} else if oppNbCap >= 4 && oppCap != nil {
		return oppCap
	} else if opponent.GotFourMid > 0 {
		return getCounterPlaysFourMid(ctx, neighbors, opponent.GotFourMid)
	}

	return nil
}

func placementHeuristicCopy(ctx s.SContext, case_x int, case_y int) uint8 {
	if ctx.ActiveDoubleThrees && ctx.Goban[case_y][case_x] != 0 && g.DoubleThree(s.SVertex{X: case_x, Y: case_y}, ctx.Goban, uint8(ctx.Goban[case_y][case_x]), isCapture) {
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
