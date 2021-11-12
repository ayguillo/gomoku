package algogo

import (
	"fmt"
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
	// check, _ := victoryCondition(node.goban, 0, 0)
	if depth <= 0 || node.value >= 10000000 || node.value <= -10000000 {
		return node.value
		// if node.player == 1 {
		// 	return node.value * int(node.captures.capture0)
		// } else {
		// 	return node.value * int(node.captures.capture1)
		// }
	}

	println("before tree")
	generateTree(node, node.coord, node.neighbors)

	println("after tree")

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

	isCapture = ctx.ActiveCapture
	isDoubleThree = ctx.ActiveDoubleThrees

	var emptyVertex s.SVertex = s.SVertex{X: -1, Y: -1}

	var neighbors []s.SVertex

	if isCapture && len(ctx.Capture) > 0 {
		println("ON EST ALLL")
		fmt.Printf("%v\n", ctx.Capture)
		neighbors = make([]s.SVertex, len(ctx.Capture))
		copy(neighbors, ctx.Capture)
	} else {
		neighbors = make([]s.SVertex, len(ctx.CasesNonNull))
		copy(neighbors, ctx.CasesNonNull)
	}

	// if isCatpure == false && ctx.Capture != nil {
	// 	neighbors = make([]s.SVertex, len(ctx.Capture))
	// 	copy(neighbors, ctx.Capture)
	// } else {
	// 	neighbors = make([]s.SVertex, len(ctx.CasesNonNull))
	// 	copy(neighbors, ctx.CasesNonNull)
	// }

	root := createNode(0, 0, copyGoban(ctx.Goban), emptyVertex, neighbors, ctx.CurrentPlayer, false, uint8(ctx.NbCaptureP1), uint8(ctx.NbCaptureP2), nil, 1)
	minimaxRecursive(root, ctx.Depth, alpha, beta, true)

	if root.bestMove != nil {
		return root.bestMove.coord, root.bestMove.value
	} else {
		return s.SVertex{X: -1, Y: -1}, 0
	}
}
