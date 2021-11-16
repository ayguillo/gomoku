package algogo

import (
	s "gomoku/structures"
	"time"
)

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

var ALLITERATION int
var initDepth = uint8(0)

var startTime time.Time
var endTime time.Time

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
	check, _ := victoryCondition(node.goban, int(node.captures.Capture0), int(node.captures.Capture1))
	if depth <= 0 || (check && depth != initDepth) {
		return node.value
	}

	generateTree(node, node.neighbors)

	if maximizingPlayer {
		maxValue := minInt
		for _, child := range node.children {
			value := minimaxRecursive(child, depth-1, alpha, beta, false)
			if value > maxValue {
				node.bestMove = child
				maxValue = value
			}
			alpha = max(alpha, value)
			if beta <= alpha {
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
			beta = min(beta, value)
			if beta <= alpha {
				break
			}
		}
		return minValue
	}
}

func getMinimaxValue(node *node, depth uint8, alpha int, beta int, ch chan playData) {
	value := minimaxRecursive(node, depth-1, alpha, beta, false)
	ch <- playData{Heur: int32(value), Vertex: node.coord}
}

func startRoutine(node *node, depth uint8) (s.SVertex, int) {
	var ch = make(chan playData)

	generateTree(node, node.neighbors)

	stockPlays := make([]playData, len(node.children))

	alpha := minInt
	beta := maxInt

	for _, child := range node.children {
		go getMinimaxValue(child, depth, alpha, beta, ch)
	}
	for i := range stockPlays {
		stockPlays[i] = <-ch
	}

	maxValue := minInt
	stockMove := s.SVertex{X: -1, Y: -1}
	for _, play := range stockPlays {
		if int(play.Heur) >= minInt {
			stockMove = play.Vertex
			maxValue = int(play.Heur)
		}
	}

	return stockMove, maxValue
}

func MinimaxTree(ctx s.SContext, depth uint8) (s.SVertex, int) {
	ALLITERATION = 0

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
		// imp := CheckImpMove(ctx, ctx.CasesNonNull)
		// if imp != nil {
		// 	neighbors = imp
		// } else {
		neighbors = make([]s.SVertex, len(ctx.CasesNonNull))
		copy(neighbors, ctx.CasesNonNull)
		// }
	}

	opp := uint8(2)
	if ctx.CurrentPlayer == 2 {
		opp = 1
	}

	startTime = time.Now()
	endTime = startTime.Add(time.Millisecond * 500)
	root := createNode(0, 0, copyGoban(ctx.Goban), emptyVertex, sortNeighbors(ctx, neighbors), opp, false, uint8(ctx.NbCaptureP1), uint8(ctx.NbCaptureP2), nil, 1)
	return startRoutine(root, depth)
	// minimaxRecursive(root, depth, alpha, beta, true)

	// for _, children := range root.children {
	// 	fmt.Printf("%v %v\n", children.coord, children.value)
	// }

	// if true {
	// 	return plays.Vertex, int(plays.Heur)
	// }

	// if root.bestMove != nil {
	// 	return root.bestMove.coord, root.bestMove.value
	// } else {
	// 	println("INFO: Reprunning minimax")
	// 	return reMinimaxTree(ctx)
	// }
}

//
//
//
//
//
//
//
//
//
///
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
