package algogo

import (
	s "gomoku/structures"
)

var identity int

func createNode(id int, value int, newGoban s.Tgoban, coord s.SVertex, neighbors []s.SVertex, player uint8, maximizingPlayer bool, NbCaptureP1 uint8, NbCaptureP2 uint8, parent *node, depth uint8, lastMove s.SVertex, lastlastMove s.SVertex) *node {

	return &node{
		id:               id,
		value:            value,
		goban:            newGoban,
		coord:            coord,
		neighbors:        neighbors,
		player:           player,
		maximizingPlayer: maximizingPlayer,
		captures: Captures{
			Capture0: NbCaptureP1,
			Capture1: NbCaptureP2,
		},
		parent:    parent,
		depth:     depth,
		lastMoves: [2]s.SVertex{lastMove, lastlastMove},
	}
}

func generateBoard(current *node, coord s.SVertex, neighbors []s.SVertex) {
	var value int

	opp := uint8(2)
	if current.player == 2 {
		opp = 1
	}

	identity++
	newGoban := current.goban
	newGoban[coord.Y][coord.X] = s.Tnumber(opp)

	tmp_last := current.lastMoves[0]
	current.lastMoves = [2]s.SVertex{{X: coord.X, Y: coord.Y}, tmp_last}

	newNeighbors := getNeighbors(current.goban, neighbors, coord)

	// var ctx s.SContext

	// ctx.Goban = newGoban
	// ctx.NSize = 19
	// ctx.CurrentPlayer = opp
	// ctx.NbCaptureP1 = int(current.captures.Capture0)
	// ctx.NbCaptureP2 = int(current.captures.Capture1)
	// ctx.ActiveDoubleThrees = isDoubleThree
	// ctx.ActiveCapture = isCapture
	var capturesVertex []s.SVertex

	if isCapture {
		capturesVertex = CaptureAlgoNode(current, coord.X, coord.Y)
		for _, capture := range capturesVertex {
			newNeighbors = append(newNeighbors, s.SVertex{Y: capture.Y, X: capture.X})
		}
	}

	newGoban[coord.Y][coord.X] = 0

	// if current.maximizingPlayer {
	// 	value = -int(EvaluateGoban(ctx)) / int(current.depth)
	// } else {
	// 	value = int(EvaluateGoban(ctx)) / int(current.depth)
	// }
	ctx := s.SContext{}
	ctx.Goban = current.goban
	ctx.CurrentPlayer = uint8(current.player)
	var child *node
	child = createNode(identity, value, newGoban, coord, newNeighbors, opp, !current.maximizingPlayer, current.captures.Capture0, current.captures.Capture1, current, current.depth+1, current.lastMoves[0], current.lastMoves[1])
	current.children = append(current.children, child)

	if isCapture && capturesVertex != nil {
		child.capturesVertex = capturesVertex
	}

}

func generateTree(current *node, neighbors []s.SVertex) {
	if current.depth == 1 {
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(current.goban, neighbor.X, neighbor.Y, current.player)
			if placement >= 1 {
				generateBoard(current, neighbor, neighbors)
			}
		}
	} else {
		for _, neighbor := range current.lastMoves {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if y == 0 && x == 0 {
						continue
					}
					placement := PlacementHeuristic(current.goban, neighbor.X, neighbor.Y, current.player)
					if placement >= 1 {
						generateBoard(current, neighbor, neighbors)
					}
				}
			}
		}
	}
}
