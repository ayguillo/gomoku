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
	newGoban := copyGoban(current.goban)
	newGoban[coord.Y][coord.X] = s.Tnumber(opp)

	tmp_last := current.lastMoves[0]
	current.lastMoves = [2]s.SVertex{coord, tmp_last}

	newNeighbors := getNeighbors(current.goban, neighbors, coord)

	var capturesVertex []s.SVertex

	if isCapture {
		capturesVertex = CaptureAlgoNode(current, coord.X, coord.Y, opp)
		for _, capture := range capturesVertex {
			newGoban[capture.Y][capture.X] = 0
			newNeighbors = append(newNeighbors, s.SVertex{Y: capture.Y, X: capture.X})
		}
	}

	child := createNode(identity, value, newGoban, coord, newNeighbors, opp, !current.maximizingPlayer, current.captures.Capture0, current.captures.Capture1, current, current.depth+1, current.lastMoves[0], current.lastMoves[1])
	current.children = append(current.children, child)

	if isCapture && capturesVertex != nil {
		child.capturesVertex = capturesVertex
	}

}

func generateTree(current *node, neighbors []s.SVertex) {
	if current.depth <= 2 {
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(current.goban, neighbor.X, neighbor.Y, current.player)
			if placement >= 1 {
				generateBoard(current, neighbor, neighbors)
			}
		}
	} else {
		lastMoves := current.lastMoves[0]
		lastlastMoves := current.lastMoves[1]

		var y int
		var x int
		var threatSpace int = 1

		for y = lastMoves.Y - threatSpace; y <= lastMoves.Y+threatSpace; y++ {
			for x = lastMoves.X - threatSpace; x <= lastMoves.X+threatSpace; x++ {

				placement := PlacementHeuristic(current.goban, x, y, current.player)
				if placement >= 1 {
					generateBoard(current, s.SVertex{X: x, Y: y}, neighbors)
				}
			}
		}

		for y = lastlastMoves.Y - threatSpace; y <= lastlastMoves.Y+threatSpace; y++ {
			for x = lastlastMoves.X - threatSpace; x <= lastlastMoves.X+threatSpace; x++ {
				placement := PlacementHeuristic(current.goban, x, y, current.player)
				if placement >= 1 {
					generateBoard(current, s.SVertex{X: x, Y: y}, neighbors)
				}
			}
		}
	}
}
