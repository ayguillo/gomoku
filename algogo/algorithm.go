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
				child.value = maxValue
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
				child.value = minValue
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
	root := createNode(0, 0, copyGoban(ctx.Goban), emptyVertex, sortNeighbors(ctx, neighbors, true), opp, false, uint8(ctx.NbCaptureP1), uint8(ctx.NbCaptureP2), nil, 1, ctx.LastMove, ctx.LastLastMove)

	minimaxRecursive(root, depth, alpha, beta, true)
	if root.bestMove != nil {
		return root.bestMove.coord, root.bestMove.value
	} else {
		println("INFO: Reprunning minimax")
		return s.SVertex{X: -1, Y: -1}, 0
	}
}
