package algorithm

import (
	s "gomoku/structures"
	g "gomoku/game"

	"math"
)

func CopyGoban2(ctx s.SContext) s.Tgoban {
	newGoban := make([][]s.Tnumber, ctx.NSize)
	for Y, line := range ctx.Goban {
		newGoban[Y] = make([]s.Tnumber, ctx.NSize)
		for X, nb := range line {
			newGoban[Y][X] = nb
		}
	}
	return newGoban
}

// func CopyCases(ctx s.SContext) map[s.SVertex][]s.SVertex {
// 	cases := make(map[s.SVertex][]s.SVertex)
// 	for key, elem := range ctx.CasesNonNull {
// 		cases[key] = elem
// 	}
// 	return cases
// }

func swapPlayer(ctx *s.SContext) {
	if ctx.CurrentPlayer == 1 {
		ctx.CurrentPlayer = 2
	} else {
		ctx.CurrentPlayer = 1
	}
}

func AlphaBetaPruning2(ctx s.SContext, depth int8) (s.SVertex, int32) {
	goban := CopyGoban2(ctx)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  nil,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize}

	neighbors := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(neighbors, ctx.CasesNonNull)

	alpha := int32(-2147483648)
	beta := int32(2147483647)

	vertex := s.SVertex{X: -1, Y: -1}
	maxEval := int32(-2147483648)

	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			newNeighbors := getNeighbors(tmp_ctx, neighbor) // fct getNeighbors a faire
			tmp := tmp_ctx.CurrentPlayer
			swapPlayer(&tmp_ctx)
			eval := minimax(tmp_ctx, newNeighbors, depth-1, alpha, beta, false)
			tmp_ctx.CurrentPlayer = tmp
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0
			if eval >= maxEval {
				maxEval = eval
				vertex = neighbor
			}
		}
	}
	
	return vertex, maxEval
}

func Heuristic2(ctx s.SContext, isMaximazingPlayer bool) int32 {
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
			if (ctx.Goban[y][x] != 0) {
				// tmp := ctx.CurrentPlayer

				if g.VictoryConditionAlign(&ctx, x, y, nil) {
					return true
				}
				
				// if (tmp == 1) {
				// 	ctx.CurrentPlayer = 2
				// } else {
				// 	ctx.CurrentPlayer = 1
				// }
	
				// if g.VictoryConditionAlign(&ctx, x, y, nil) {
				// 	return true
				// }
	
				// ctx.CurrentPlayer = tmp
			}
		}
	}
	
	return false
}

func minimax(tmp_ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32, isMaximazingPlayer bool) int32 {
	if depth == 0 || victoryCondition(tmp_ctx) {
		return Heuristic2(tmp_ctx, isMaximazingPlayer)
	}

	if isMaximazingPlayer {
		maxEval := int32(-2147483648)

		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(tmp_ctx.CurrentPlayer)
				newNeighbors := getNeighbors(tmp_ctx, neighbor) // fct getNeighbors a faire
				tmp := tmp_ctx.CurrentPlayer
				swapPlayer(&tmp_ctx)
				eval := minimax(tmp_ctx, newNeighbors, depth-1, alpha, beta, false)
				tmp_ctx.CurrentPlayer = tmp
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0
	
				maxEval = int32(math.Max(float64(maxEval), float64(eval)))
				alpha = int32(math.Max(float64(alpha), float64(maxEval)))

				if alpha >= beta {
					break
				}
			}
		}
		return maxEval
	} else {
		minEval := int32(2147483647)

		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				// fmt.Printf("Min: %v | %d\n", neighbor, depth)
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(tmp_ctx.CurrentPlayer)
				newNeighbors := getNeighbors(tmp_ctx, neighbor) // fct getNeighbors a faire
				tmp := tmp_ctx.CurrentPlayer
				swapPlayer(&tmp_ctx)
				eval := minimax(tmp_ctx, newNeighbors, depth-1, alpha, beta, true)
				tmp_ctx.CurrentPlayer = tmp
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0
				minEval = int32(math.Min(float64(minEval), float64(eval)))
				beta = int32(math.Min(float64(beta), float64(minEval)))

				if beta <= alpha {
					break
				}
			}
		}
		return minEval
	}
}
