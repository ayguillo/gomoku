package algogo

import (
	s "gomoku/structures"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

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

func minimaxRecursive(node *node, depth uint8, alpha int, beta int, maximizingPlayer bool) int {
	if depth == 0 || node.value >= align5Win || node.value <= -align5Win {
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

	var emptyVertex s.SVertex = s.SVertex{X: -1, Y: -1}
	neighbors := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(neighbors, ctx.CasesNonNull)

	root := createNode(0, 0, copyGoban(ctx.Goban, int(ctx.Size)), emptyVertex, neighbors, ctx.CurrentPlayer, false, uint8(ctx.NbCaptureP1), uint8(ctx.NbCaptureP2), nil, 1)
	minimaxRecursive(root, ctx.Depth, alpha, beta, true)

	return root.bestMove.coord, root.bestMove.value
}
