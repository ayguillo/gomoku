package algogo

import (
	s "gomoku/structures"
)

var identity int

func createNode(id int, value int, newGoban s.Tgoban, coord s.SVertex, neighbors []s.SVertex, player uint8, maximizingPlayer bool, NbCaptureP1 uint8, NbCaptureP2 uint8, parent *node, depth uint8, lastMove s.SVertex) *node {

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
		lastMove: lastMove,
	}
}

func generateBoard(current *node, coord s.SVertex, lastMove s.SVertex, neighbors []s.SVertex) {
	var value int

	opp := uint8(2)
	if current.player == 2 {
		opp = 1
	}

	identity++
	newGoban := copyGoban(current.goban)
	newGoban[coord.Y][coord.X] = s.Tnumber(opp)
	newNeighbors := getNeighbors(current.goban, neighbors, coord)
	var capturesVertex []s.SVertex

	if isCapture {
		capturesVertex = CaptureAlgoNode(current, coord.X, coord.Y, opp)
		for _, capture := range capturesVertex {
			newGoban[capture.Y][capture.X] = 0
			newNeighbors = append(newNeighbors, s.SVertex{Y: capture.Y, X: capture.X})
		}
	}

	child := createNode(identity, value, newGoban, coord, newNeighbors, opp, !current.maximizingPlayer, current.captures.Capture0, current.captures.Capture1, current, current.depth+1, lastMove)
	current.children = append(current.children, child)

	if isCapture && capturesVertex != nil {
		child.capturesVertex = capturesVertex
	}

}

func generateTree(current *node, neighbors []s.SVertex) {
	if current.depth <= 3 {
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(current.goban, neighbor.X, neighbor.Y, current.player)
			if placement >= 1 {
				generateBoard(current, neighbor, current.lastMove, neighbors)
			}
		}
	} else {
		lastMove := current.coord
		lastlastMove := current.lastMove

		var y int
		var x int
		var threatSpace int = 1

		for y = lastMove.Y - threatSpace; y <= lastMove.Y+threatSpace; y++ {
			for x = lastMove.X - threatSpace; x <= lastMove.X+threatSpace; x++ {

				placement := PlacementHeuristic(current.goban, x, y, current.player)
				if placement >= 1 {
					generateBoard(current, s.SVertex{X: x, Y: y}, lastMove, neighbors)
				}
			}
		}

		for y = lastlastMove.Y - threatSpace; y <= lastlastMove.Y+threatSpace; y++ {
			for x = lastlastMove.X - threatSpace; x <= lastlastMove.X+threatSpace; x++ {
					placement := PlacementHeuristic(current.goban, x, y, current.player)
					if placement >= 1 {
						generateBoard(current, s.SVertex{X: x, Y: y}, lastlastMove, neighbors)
				}
			}
		}
	}
}
