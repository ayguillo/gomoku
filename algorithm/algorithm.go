package algorithm

import (
	h "gomoku/heuristic"
	s "gomoku/structures"
)

var initDepth int8
var depthStock int8
var alphaStock int32
var betaStock int32

func AlphaBetaPruning(ctx s.SContext, depth int8) (s.SVertex, int32) {
	initDepth = depth

	neighbors := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(neighbors, ctx.CasesNonNull)

	var data []stockData2

	impMove := h.CheckImpMoove(ctx, neighbors)

	if impMove != nil {
		neighbors = impMove
	}

	neighbors = sortNeighbors(ctx, neighbors)
	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			data = append(data, initMax(ctx, depth, neighbor))
		}
	}

	vertex := s.SVertex{X: -1, Y: -1}
	maxEval := int32(-2147483648)
	maxDepth := depth

	for _, value := range data {
		eval := value.Heur
		if eval > maxEval {
			maxEval = eval
			vertex = value.Vertex
			maxDepth = value.Depth
		}
	}

	println("END:", maxEval, depth-maxDepth, "\n")

	return vertex, maxEval
}

func initMax(ctx s.SContext, depth int8, neighbor s.SVertex) stockData2 {
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

	check, _ := VictoryCondition(ctx)

	if check {
		eval = h.CalcHeuristic(tmp_ctx)
	} else {
		newNeighbors := getNeighbors(tmp_ctx, neighbor)
		tmp := tmp_ctx.CurrentPlayer
		swapPlayer(&tmp_ctx)

		eval = minimax(tmp_ctx, newNeighbors, depth-1, -2147483648) // minimax ab pruning shorter
		// eval = pvs(tmp_ctx, newNeighbors, depth-1, -2147483648, 2147483647, 1) // negascout
		// eval = minimax2(tmp_ctx, newNeighbors, depth-1, -2147483648, 2147483647, false) // minimax ab prun
		// eval = negaAlphaBeta(tmp_ctx, newNeighbors, depth, -2147483648, 2147483647) // negamax ab purn

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

func minimax(ctx s.SContext, neighbors []s.SVertex, depth int8, i int32) int32 {
	check, _ := VictoryCondition(ctx)

	if depth <= 0 || check {
		swapPlayer(&ctx)
		heur := h.CalcHeuristic(ctx)
		return heur * int32(depth+1)
	}

	j := int32(-2147483648)
	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			newNeighbors := getNeighbors(ctx, neighbor)

			tmp := ctx.CurrentPlayer
			swapPlayer(&ctx)

			j = max(j, minimax(ctx, newNeighbors, depth-1, j))

			ctx.CurrentPlayer = tmp
			ctx.Goban[neighbor.Y][neighbor.X] = 0

			if -j <= i {
				return -j
			}
		}
	}
	return -j
}

func pvs(ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32, swap int32) int32 {
	check, _ := VictoryCondition(ctx)
	if depth <= 0 || check {
		swapPlayer(&ctx)
		heur := h.CalcHeuristic(ctx)
		swapPlayer(&ctx)
		return heur
	}

	i := 0
	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			newNeighbors := getNeighbors(ctx, neighbor)

			tmp := ctx.CurrentPlayer
			swapPlayer(&ctx)
			score := int32(0)

			if i != 0 {
				score = -pvs(ctx, newNeighbors, depth-1, -alpha-1, -alpha, -swap)

				if alpha < score && score < beta {
					score = -pvs(ctx, newNeighbors, depth-1, -beta, -score, -swap)
				}
			} else {
				score = -pvs(ctx, newNeighbors, depth-1, -beta, -alpha, -swap)
			}

			ctx.CurrentPlayer = tmp
			ctx.Goban[neighbor.Y][neighbor.X] = 0

			alpha = max(alpha, score)

			if alpha >= beta {
				break
			}
			i++
		}
	}

	return alpha
}

func minimax2(ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32, isMax bool) int32 {
	check, _ := VictoryCondition(ctx)
	if depth <= 0 || check {
		swapPlayer(&ctx)
		heur := h.CalcHeuristic(ctx)
		if isMax {
			return -heur
		}
		return heur
	}

	if isMax {
		maxEval := int32(-2147483648)
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
				newNeighbors := getNeighbors(ctx, neighbor)

				tmp := ctx.CurrentPlayer
				swapPlayer(&ctx)

				eval := minimax2(ctx, newNeighbors, depth-1, alpha, beta, false)

				ctx.CurrentPlayer = tmp
				ctx.Goban[neighbor.Y][neighbor.X] = 0

				maxEval = max(maxEval, eval)
				alpha = max(alpha, eval)

				if beta <= alpha {
					break
				}
			}
		}
		return maxEval
	} else {
		minEval := int32(2147483647)
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
				newNeighbors := getNeighbors(ctx, neighbor)

				tmp := ctx.CurrentPlayer
				swapPlayer(&ctx)

				eval := minimax2(ctx, newNeighbors, depth-1, alpha, beta, true)

				ctx.CurrentPlayer = tmp
				ctx.Goban[neighbor.Y][neighbor.X] = 0

				minEval = min(minEval, eval)
				beta = min(beta, eval)

				if beta <= alpha {
					break
				}
			}
		}
		return minEval
	}
}

// func negaAlphaBeta(ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32) int32 {
// 	if depth <= 0 || VictoryCondition(ctx) {
// 		swapPlayer(&ctx)
// 		heur := h.CalcHeuristic(ctx)
// 		swapPlayer(&ctx)
// 		return heur
// 	}

// 	best := int32(-2147483648)

// 	for _, neighbor := range neighbors {
// 		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
// 		if placement >= 1 {
// 			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
// 			newNeighbors := getNeighbors(ctx, neighbor)

// 			tmp := ctx.CurrentPlayer
// 			swapPlayer(&ctx)

// 			value := -negaAlphaBeta(ctx, newNeighbors, depth-1, -beta, -alpha)

// 			ctx.CurrentPlayer = tmp
// 			ctx.Goban[neighbor.Y][neighbor.X] = 0

// 			if value > best {
// 				best = value
// 				if best > alpha {
// 					alpha = best
// 					if alpha > beta {
// 						return best
// 					}
// 				}
// 			}

// 		}
// 	}

// 	return best
// }

// func reAlphaBeta(ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32, isMaximazingPlayer bool) int32 {
// 	if depth <= 0 || VictoryCondition(ctx) {
// 		swapPlayer(&ctx)
// 		heur := h.CalcHeuristic(ctx)
// 		swapPlayer(&ctx)
// 		return heur
// 	}

// 	var value int32

// 	if isMaximazingPlayer {
// 		value = int32(-2147483648)

// 		for _, neighbor := range neighbors {
// 			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
// 			if placement >= 1 {
// 				value = max(value, reAlphaBeta(ctx, neighbors, depth-1, alpha, beta, isMaximazingPlayer))
// 				if value >= beta {
// 					return value
// 				}

// 			}
// 		}
// 	} else {
// 		value = int32(2147483647)

// 		for _, neighbor := range neighbors {
// 			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
// 			if placement >= 1 {
// 			}
// 		}
// 	}

// 	return value
// }
