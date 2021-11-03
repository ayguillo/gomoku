package algorithm

import (
	h "gomoku/heuristic"
	s "gomoku/structures"
)

type stockData2 struct {
	Heur   int32
	Vertex s.SVertex
	Depth  int8
	Alpha  int32
	Beta   int32
}

var depthStock int8
var alphaStock int32
var betaStock int32

func AlphaBetaPruning2(ctx s.SContext, depth int8) (s.SVertex, int32) {
	neighbors := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(neighbors, ctx.CasesNonNull)

	alpha := int32(-2147483648)
	beta := int32(2147483647)

	var data []stockData2

	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			data = append(data, initMax2(ctx, depth, neighbor, alpha, beta))
		}
	}

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

func initMax2(ctx s.SContext, depth int8, neighbor s.SVertex, alpha int32, beta int32) stockData2 {
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

		eval = minimax2(tmp_ctx, newNeighbors, depth-1, alpha, beta, false)

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

func minimax2(ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32, isMaximazingPlayer bool) int32 {
	if depth <= 0 || VictoryCondition(ctx) {
		depthStock = depth
		alphaStock = alpha
		betaStock = beta
		swapPlayer(&ctx)
		heur := h.CalcHeuristic(ctx)
		println(heur)
		swapPlayer(&ctx)
		return heur
	}

	if isMaximazingPlayer {
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
				newNeighbors := getNeighbors(ctx, neighbor)

				tmp := ctx.CurrentPlayer
				swapPlayer(&ctx)

				alpha = max(alpha, minimax2(ctx, newNeighbors, depth-1, alpha, beta, false))

				ctx.CurrentPlayer = tmp
				ctx.Goban[neighbor.Y][neighbor.X] = 0

				if alpha >= beta {
					break
				}
			}
		}
		return alpha
	} else {
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
				newNeighbors := getNeighbors(ctx, neighbor)

				tmp := ctx.CurrentPlayer
				swapPlayer(&ctx)

				beta = min(beta, minimax2(ctx, newNeighbors, depth-1, alpha, beta, true))

				ctx.CurrentPlayer = tmp
				ctx.Goban[neighbor.Y][neighbor.X] = 0

				if alpha >= beta {
					break
				}
			}
		}
		return beta
	}
}
