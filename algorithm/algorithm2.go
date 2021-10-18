package algorithm

import (
	g "gomoku/game"
	s "gomoku/structures"
	"math"
)

func CopyGoban(ctx s.SContext) s.Tgoban {
	newGoban := make([][]s.Tnumber, ctx.NSize)
	for Y, line := range ctx.Goban {
		newGoban[Y] = make([]s.Tnumber, ctx.NSize)
		for X, nb := range line {
			newGoban[Y][X] = nb
		}
	}
	return newGoban
}

func CopyCases(ctx s.SContext) map[s.SVertex][]s.SVertex {
	cases := make(map[s.SVertex][]s.SVertex)
	for key, elem := range ctx.CasesNonNull {
		cases[key] = elem
	}
	return cases
}

func swapPlayer(ctx *s.SContext) {
	if ctx.CurrentPlayer == 1 {
		ctx.CurrentPlayer = 2
	} else {
		ctx.CurrentPlayer = 1
	}
}

func Heuristic2(ctx s.SContext) int32 {
	value := 0
	// gotFiveInRow := false
	// gotLiveFour := false
	// gotDeadFour := false
	// gotLiveThree := false
	// gotDeadThree := false
	// gotLiveTwo := false
	// gotDeadTwo := false

	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
				value += Heuristic(ctx, x, y)
			}
		}
	}

	return int32(value)
}

func victoryCondition(ctx s.SContext) bool {
	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] != 0 {
				if g.VictoryConditionAlign(&ctx, x, y, nil) {
					return true
				}
			}
		}
	}

	return false
}

func AlphaBetaPruning(ctx s.SContext, depth int) (s.SVertex, int32) {
	goban := CopyGoban(ctx)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  nil,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize}

	alpha := int32(-2147483648)
	beta := int32(2147483647)

	vertex := s.SVertex{X: -1, Y: -1}
	maxEval := int32(-2147483648)

	for stone := range ctx.CasesNonNull {
		for _, neighbor := range ctx.CasesNonNull[stone] {
			println("oui")
			placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)              // place ma pierre
				newNeighbors := FindNeighborsExplo(ctx, neighbor.X, neighbor.Y, ctx.CasesNonNull) // fct getNeighbors a faire
				tmpCurrentPlayer := tmp_ctx.CurrentPlayer
				swapPlayer(&tmp_ctx)                                            // Je change de player
				eval := min_player(tmp_ctx, newNeighbors, depth-1, alpha, beta) // mon algo
				tmp_ctx.CurrentPlayer = tmpCurrentPlayer                        // Je rechange de player
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0                       // j'enleve ma piece
				if eval >= maxEval {
					maxEval = eval
					vertex = neighbor
				}
			}
		}
	}

	return vertex, maxEval
}

// 1st player
func max_player(ctx s.SContext, neighbors map[s.SVertex][]s.SVertex, depth int, alpha int32, beta int32) int32 {
	if depth == 0 {
		return Heuristic2(ctx)
	}

	maxEval := int32(-1 << 31)

	for stone := range neighbors {
		for _, neighbor := range neighbors[stone] {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
				newNeighbors := FindNeighborsExplo(ctx, neighbor.X, neighbor.Y, neighbors) // fct getNeighbors a faire
				tmp := ctx.CurrentPlayer
				swapPlayer(&ctx)
				// tmp_vertex, tmp_u := max_player(tmp_ctx, alpha, beta, explor, max_explor)
				eval := max_player(ctx, newNeighbors, depth-1, alpha, beta)
				ctx.CurrentPlayer = tmp
				ctx.Goban[neighbor.Y][neighbor.X] = 0
				// Alpha beta prunning a ajouter

				maxEval = int32(math.Max(float64(maxEval), float64(eval)))
				alpha = int32(math.Max(float64(alpha), float64(maxEval)))

				if alpha >= beta {
					break
				}
			}
		}
	}
	return maxEval
}

// 2nd player
func min_player(ctx s.SContext, neighbors map[s.SVertex][]s.SVertex, depth int, alpha int32, beta int32) int32 {
	if depth == 0 {
		return Heuristic2(ctx)
	}

	minEval := int32(1<<31 - 1)

	for stone := range neighbors {
		for _, neighbor := range neighbors[stone] {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
				newNeighbors := FindNeighborsExplo(ctx, neighbor.X, neighbor.Y, neighbors) // fct getNeighbors a faire
				tmp := ctx.CurrentPlayer
				swapPlayer(&ctx)
				// tmp_vertex, tmp_u := max_player(tmp_ctx, alpha, beta, explor, max_explor)
				eval := max_player(ctx, newNeighbors, depth-1, alpha, beta)
				ctx.CurrentPlayer = tmp
				ctx.Goban[neighbor.Y][neighbor.X] = 0
				// Alpha beta prunning a ajouter

				minEval = int32(math.Min(float64(minEval), float64(eval)))
				beta = int32(math.Min(float64(beta), float64(minEval)))

				if beta <= alpha {
					return minEval
				}
			}
		}
	}
	return minEval
}
