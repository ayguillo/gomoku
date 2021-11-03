package algorithm

import (
	"fmt"
	h "gomoku/heuristic"
	s "gomoku/structures"
)

func AlphaBetaPruning3(ctx s.SContext, depth int8) (s.SVertex, int32) {
	neighbors := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(neighbors, ctx.CasesNonNull)

	alpha := int32(-2147483648)
	beta := int32(2147483647)

	var data []stockData2

	fmt.Printf("%d\n%v\n", len(neighbors), neighbors)
	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			data = append(data, initMax3(ctx, depth, neighbor))
		}
	}

	fmt.Printf("%v\n", data)
	vertex := s.SVertex{X: -1, Y: -1}
	maxEval := int32(-2147483648)
	maxDepth := depth
	maxAlpha := alpha
	maxBeta := beta

	for _, value := range data {
		eval := value.Heur
		if eval >= maxEval {
			maxEval = eval
			vertex = value.Vertex
			maxDepth = value.Depth
			maxAlpha = value.Alpha
			maxBeta = value.Beta
		}
	}

	println("END:", maxEval, depth-maxDepth, maxAlpha, maxBeta, "\n")

	return vertex, maxEval
}

func initMax3(ctx s.SContext, depth int8, neighbor s.SVertex) stockData2 {
	var eval int32

	goban := CopyGoban(ctx)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  nil,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize,
		ActiveCapture: ctx.ActiveCapture,
	}

	tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)

	if VictoryCondition(tmp_ctx) {
		eval = h.CalcHeuristic(tmp_ctx)
	} else {
		newNeighbors := getNeighbors(tmp_ctx, neighbor)
		tmp := tmp_ctx.CurrentPlayer
		swapPlayer(&tmp_ctx)

		eval = minimax3(tmp_ctx, newNeighbors, depth-1, -2147483648)
		// eval = negaAlphaBeta(tmp_ctx, newNeighbors, depth, -2147483648, 2147483647)

		tmp_ctx.CurrentPlayer = tmp
		tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0
	}

	ret := stockData2{
		Heur:   eval,
		Vertex: neighbor,
		Depth:  depthStock,
		Alpha:  alphaStock,
		Beta:   betaStock,
	}

	return ret
}

func negaAlphaBeta(ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32) int32 {
	if depth <= 0 || VictoryCondition(ctx) {
		swapPlayer(&ctx)
		heur := h.CalcHeuristic(ctx)
		swapPlayer(&ctx)
		return heur
	}

	best := int32(-2147483648)

	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			newNeighbors := getNeighbors(ctx, neighbor)

			tmp := ctx.CurrentPlayer
			swapPlayer(&ctx)

			value := -negaAlphaBeta(ctx, newNeighbors, depth-1, -beta, -alpha)

			ctx.CurrentPlayer = tmp
			ctx.Goban[neighbor.Y][neighbor.X] = 0

			if value > best {
				best = value
				if best > alpha {
					alpha = best
					if alpha > beta {
						return best
					}
				}
			}

		}
	}

	return best
}

func reAlphaBeta(ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32, isMaximazingPlayer bool) int32 {
	if depth <= 0 || VictoryCondition(ctx) {
		swapPlayer(&ctx)
		heur := h.CalcHeuristic(ctx)
		swapPlayer(&ctx)
		return heur
	}

	var value int32

	if isMaximazingPlayer {
		value = int32(-2147483648)

		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				value = max(value, reAlphaBeta(ctx, neighbors, depth-1, alpha, beta, isMaximazingPlayer))
				if value >= beta {
					return value
				}

			}
		}
	} else {
		value = int32(2147483647)

		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
			}
		}
	}

	return value
}

func minimax3(ctx s.SContext, neighbors []s.SVertex, depth int8, i int32) int32 {
	if depth <= 0 || VictoryCondition(ctx) {
		swapPlayer(&ctx)
		heur := h.CalcHeuristic(ctx)
		swapPlayer(&ctx)
		return heur
	}

	j := int32(-2147483648)
	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			newNeighbors := getNeighbors(ctx, neighbor)

			tmp := ctx.CurrentPlayer
			swapPlayer(&ctx)

			j = max(j, minimax3(ctx, newNeighbors, depth-1, j))

			ctx.CurrentPlayer = tmp
			ctx.Goban[neighbor.Y][neighbor.X] = 0

			if -j <= i {
				return -j
			}
		}
	}
	return -j
}