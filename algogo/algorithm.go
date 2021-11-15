package algogo

import (
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

func buildContext(node node) s.SContext {
	var ctx s.SContext

	ctx.Goban = node.goban
	ctx.NSize = 19
	ctx.CurrentPlayer = node.player
	ctx.NbCaptureP1 = int(node.captures.Capture0)
	ctx.NbCaptureP2 = int(node.captures.Capture1)
	ctx.ActiveCapture = isCapture

	return ctx
}

func minimaxRecursive(node *node, depth uint8, alpha int, beta int, maximizingPlayer bool) int {
	// fmt.Printf("%v\n", node.goban)
	// check, _ := victoryCondition(node.goban, int(node.captures.Capture0), int(node.captures.Capture1))
	if depth <= 0 || node.value >= 10000000 || node.value <= -10000000 {
		return node.value
	}

	generateTree(node, node.coord, node.neighbors)

	if maximizingPlayer {
		maxValue := minInt
		for _, child := range node.children {
			value := minimaxRecursive(child, depth-1, alpha, beta, false)
			if value > maxValue {
				node.bestMove = child
				maxValue = value
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
			value := minimaxRecursive(child, depth-1, alpha, beta, true)
			if value < minValue {
				node.bestMove = child
				minValue = value
			}
			beta = min(beta, minValue)
			if beta <= alpha {
				break
			}
		}
		return minValue
	}
}

func MinimaxTree(ctx s.SContext) (s.SVertex, int) {
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
		imp := CheckImpMove(ctx, ctx.CasesNonNull)
		if imp != nil {
			neighbors = imp
		} else {
			neighbors = make([]s.SVertex, len(ctx.CasesNonNull))
			copy(neighbors, ctx.CasesNonNull)
		}
	}

	opp := uint8(2)
	if ctx.CurrentPlayer == 2 {
		opp = 2
	}

	root := createNode(0, 0, copyGoban(ctx.Goban), emptyVertex, neighbors, opp, false, uint8(ctx.NbCaptureP1), uint8(ctx.NbCaptureP2), nil, 1)
	minimaxRecursive(root, ctx.Depth, alpha, beta, true)

	if root.bestMove != nil {
		return root.bestMove.coord, root.bestMove.value
	} else {
		return reMinimaxTree(ctx)
	}
}

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
		opp = 2
	}

	root := createNode(0, 0, copyGoban(ctx.Goban), emptyVertex, neighbors, opp, false, uint8(ctx.NbCaptureP1), uint8(ctx.NbCaptureP2), nil, 1)
	minimaxRecursive(root, ctx.Depth, alpha, beta, true)

	if root.bestMove != nil {
		return root.bestMove.coord, root.bestMove.value
	} else {
		return s.SVertex{X: -1, Y: -1}, 0
	}
}
