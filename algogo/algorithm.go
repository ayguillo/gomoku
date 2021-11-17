package algogo

import (
	"fmt"
	s "gomoku/structures"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

var initDepth = uint8(0)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buildContext(node node, player uint8) s.SContext {
	var ctx s.SContext

	ctx.Goban = node.goban
	ctx.NSize = 19
	ctx.CurrentPlayer = player
	ctx.NbCaptureP1 = int(node.captures.Capture0)
	ctx.NbCaptureP2 = int(node.captures.Capture1)
	ctx.ActiveCapture = isCapture

	return ctx
}

func minimaxRecursive(node *node, depth uint8, alpha int, beta int, maximizingPlayer bool) int {
	check, _ := victoryCondition(node.goban, int(node.captures.Capture0), int(node.captures.Capture1))
	if depth <= 0 || (check && depth != initDepth) {
		opp := uint8(2)
		if node.player == 2 {
			opp = 1
		}
		if node.maximizingPlayer {
			return -int(EvaluateGoban(buildContext(*node, opp))) / int(node.depth)
		} else {
			return int(EvaluateGoban(buildContext(*node, opp))) / int(node.depth)
		}
	}

	generateTree(node, node.neighbors)

	if maximizingPlayer {
		maxValue := minInt
		for _, child := range node.children {
			child.goban[child.coord.Y][child.coord.X] = s.Tnumber(child.player)

			if isCapture {
				for _, capture := range child.capturesVertex {
					child.goban[capture.Y][capture.X] = 0
				}
			}

			value := minimaxRecursive(child, depth-1, alpha, beta, false)
			if value > maxValue {
				node.bestMove = child
				maxValue = value
			}

			child.goban[child.coord.Y][child.coord.X] = 0
			if isCapture {
				opp := s.Tnumber(2)
				if child.player == 2 {
					opp = 1
				}

				for _, capture := range child.capturesVertex {
					child.goban[capture.Y][capture.X] = opp
				}
			}

			alpha = max(alpha, maxValue)
			if alpha >= beta {
				break
			}

		}
		return maxValue
	} else {
		minValue := maxInt
		for _, child := range node.children {
			child.goban[child.coord.Y][child.coord.X] = s.Tnumber(child.player)

			value := minimaxRecursive(child, depth-1, alpha, beta, true)
			if value < minValue {
				node.bestMove = child
				minValue = value
			}

			child.goban[child.coord.Y][child.coord.X] = 0

			beta = min(beta, minValue)
			if beta <= alpha {
				break
			}

		}
		return minValue
	}
}

func getMinimaxValue(node *node, depth uint8, ch chan playData) {
	alpha := minInt
	beta := maxInt

	opp := s.Tnumber(2)
	if node.player == 2 {
		opp = 1
	}

	node.goban[node.coord.Y][node.coord.X] = opp
	value := minimaxRecursive(node, depth-1, alpha, beta, true)
	node.goban[node.coord.Y][node.coord.X] = 0
	ch <- playData{Heur: int32(value), Vertex: node.coord}
}

func startRoutine(node *node, depth uint8) (s.SVertex, int) {
	var ch = make(chan playData)

	generateTree(node, node.neighbors)

	stockPlays := make([]playData, len(node.children))
	println(len(node.children))
	for _, child := range node.children {
		child.goban = copyGoban(child.goban)
		go getMinimaxValue(child, depth, ch)
	}
	a := 0
	for i := range stockPlays {
		stockPlays[i] = <-ch
		a++
	}
	maxValue := maxInt
	stockMove := s.SVertex{X: -1, Y: -1}
	fmt.Printf("data: %v\n", stockPlays)
	for _, play := range stockPlays {
		if int(play.Heur) < maxValue {
			stockMove = play.Vertex
			maxValue = int(play.Heur)
		}
	}

	return stockMove, maxValue
}

func MinimaxTree(ctx s.SContext, depth uint8) (s.SVertex, int) {
	alpha := minInt
	beta := maxInt
	initDepth = depth

	isCapture = ctx.ActiveCapture
	isDoubleThree = ctx.ActiveDoubleThrees
	initPlayer = ctx.CurrentPlayer

	var emptyVertex s.SVertex = s.SVertex{X: -1, Y: -1}

	var neighbors []s.SVertex

	if isCapture && len(ctx.Capture) > 0 {
		neighbors = make([]s.SVertex, len(ctx.Capture))
		copy(neighbors, ctx.Capture)
	} else {
		neighbors = make([]s.SVertex, len(ctx.CasesNonNull))
		copy(neighbors, ctx.CasesNonNull)
	}

	opp := uint8(2)
	if ctx.CurrentPlayer == 2 {
		opp = 1
	}

	root := createNode(0, 0, copyGoban(ctx.Goban), emptyVertex, sortNeighbors(ctx, neighbors), opp, false, uint8(ctx.NbCaptureP1), uint8(ctx.NbCaptureP2), nil, 1)

	multiThreading := true

	return s.SVertex{X: -1, Y: -1}, 0
	if multiThreading {
		return startRoutine(root, depth)
	} else {
		minimaxRecursive(root, depth, alpha, beta, true)

		if root.bestMove != nil {
			return root.bestMove.coord, root.bestMove.value
		} else {
			println("INFO: Reprunning minimax")
			return reMinimaxTree(ctx)
		}
	}
}

////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////
////////////////////////

func reMinimaxTree(ctx s.SContext) (s.SVertex, int) {
	alpha := minInt
	beta := maxInt
	initDepth = ctx.Depth

	isCapture = ctx.ActiveCapture
	isDoubleThree = ctx.ActiveDoubleThrees
	initPlayer = ctx.CurrentPlayer

	var emptyVertex s.SVertex = s.SVertex{X: -1, Y: -1}

	var neighbors []s.SVertex

	if isCapture && len(ctx.Capture) > 0 {
		neighbors = make([]s.SVertex, len(ctx.Capture))
		copy(neighbors, ctx.Capture)
	} else {
		neighbors = make([]s.SVertex, len(ctx.CasesNonNull))
		copy(neighbors, ctx.CasesNonNull)
	}

	opp := uint8(2)
	if ctx.CurrentPlayer == 2 {
		opp = 1
	}

	root := createNode(0, 0, copyGoban(ctx.Goban), emptyVertex, neighbors, opp, false, uint8(ctx.NbCaptureP1), uint8(ctx.NbCaptureP2), nil, 1)
	minimaxRecursive(root, ctx.Depth, alpha, beta, true)

	if root.bestMove != nil {
		return root.bestMove.coord, root.bestMove.value
	} else {
		return s.SVertex{X: -1, Y: -1}, 0
	}
}
