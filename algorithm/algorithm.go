package algorithm

import (
	h "gomoku/heuristic"
	s "gomoku/structures"
)

var initDepth int8
var initPlayer uint8
var depthStock int8
var alphaStock int32
var betaStock int32

func rePrun(ctx s.SContext, depth int8) (s.SVertex, int32, int8) {
	initDepth = depth
	initPlayer = ctx.CurrentPlayer

	neighbors := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(neighbors, ctx.CasesNonNull)

	var data []stockData2

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

	return vertex, maxEval, maxDepth
}

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

	if vertex.X == -1 || vertex.Y == -1 {
		vertex, maxEval, maxDepth = rePrun(ctx, depth)
	}

	println("END:", maxEval, depth-maxDepth, "\n")

	return vertex, maxEval
}

func initMax(ctx s.SContext, depth int8, neighbor s.SVertex) stockData2 {
	var eval int32

	goban := CopyGoban(ctx)
	tmp_ctx := s.SContext{
		Goban:              goban,
		CurrentPlayer:      ctx.CurrentPlayer,
		CasesNonNull:       nil,
		Capture:            ctx.Capture,
		NbCaptureP1:        ctx.NbCaptureP1,
		NbCaptureP2:        ctx.NbCaptureP2,
		NSize:              ctx.NSize,
		ActiveCapture:      ctx.ActiveCapture,
		ActiveDoubleThrees: ctx.ActiveDoubleThrees,
	}

	tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)

	var captureP1 int = 0
	var captureP2 int = 0
	var captureVertex []s.SVertex = nil

	if ctx.ActiveCapture {
		captureP1, captureP2 = ctx.NbCaptureP1, ctx.NbCaptureP2
		captureVertex = CaptureAlgo(&ctx, neighbor.X, neighbor.Y)
	}

	newNeighbors := getNeighbors(tmp_ctx, neighbor)
	tmp := tmp_ctx.CurrentPlayer
	swapPlayer(&tmp_ctx)

	// eval = minimax(tmp_ctx, newNeighbors, depth-1, -2147483648)
	eval = minimaxRecursive(tmp_ctx, newNeighbors, depth-1, -2147483648, 2147483647, false)

	if ctx.ActiveCapture && captureVertex != nil {
		revertCapture(&ctx, captureVertex, captureP1, captureP2, tmp)
	}

	tmp_ctx.CurrentPlayer = tmp
	tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0

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
		if ctx.CurrentPlayer == initPlayer {
			heur *= -1
		}
		return heur * int32(depth+1)
	}

	j := int32(-2147483648)
	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)

			var captureP1 int = 0
			var captureP2 int = 0
			var captureVertex []s.SVertex = nil

			if ctx.ActiveCapture {
				captureP1, captureP2 = ctx.NbCaptureP1, ctx.NbCaptureP2
				captureVertex = CaptureAlgo(&ctx, neighbor.X, neighbor.Y)
			}

			newNeighbors := getNeighbors(ctx, neighbor)

			tmp := ctx.CurrentPlayer
			swapPlayer(&ctx)

			j = max(j, minimax(ctx, newNeighbors, depth-1, j))

			if ctx.ActiveCapture && captureVertex != nil {
				revertCapture(&ctx, captureVertex, captureP1, captureP2, tmp)
			}

			ctx.Goban[neighbor.Y][neighbor.X] = 0
			ctx.CurrentPlayer = tmp

			if -j <= i {
				return -j
			}
		}
	}
	return -j
}

const maxInt = int32(2147483647)
const minInt = int32(-2147483648)

func minimaxRecursive(ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32, maximizingPlayer bool) int32 {
	check, _ := VictoryCondition(ctx)

	if depth <= 0 || check {
		swapPlayer(&ctx)
		heur := h.CalcHeuristic(ctx)
		if ctx.CurrentPlayer == initPlayer {
			heur *= -1
		}
		return heur * int32(depth+1)
	}

	if maximizingPlayer {
		maxValue := minInt
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)

				newNeighbors := getNeighbors(ctx, neighbor)

				tmp := ctx.CurrentPlayer
				swapPlayer(&ctx)

				value := minimaxRecursive(ctx, newNeighbors, depth-1, alpha, beta, false)

				ctx.Goban[neighbor.Y][neighbor.X] = 0
				ctx.CurrentPlayer = tmp

				if value > maxValue {
					maxValue = value
				}

				alpha = max(alpha, maxValue)
				if alpha >= beta {
					break
				}
			}
		}
		return maxValue
	} else {
		minValue := maxInt
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
				newNeighbors := getNeighbors(ctx, neighbor)

				tmp := ctx.CurrentPlayer
				swapPlayer(&ctx)

				value := minimaxRecursive(ctx, newNeighbors, depth-1, alpha, beta, true)

				ctx.Goban[neighbor.Y][neighbor.X] = 0
				ctx.CurrentPlayer = tmp

				if value < minValue {
					minValue = value
				}

				beta = min(beta, minValue)
				if beta <= alpha {
					break
				}
			}
		}
		return minValue
	}
}
