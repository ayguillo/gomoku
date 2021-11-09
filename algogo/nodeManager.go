package algogo

import (
	h "gomoku/heuristic"
	s "gomoku/structures"
)

var identity int

func createNode(id int, value int, newGoban s.Tgoban, coord s.SVertex, neighbors []s.SVertex, player uint8, maximizingPlayer bool, NbCaptureP1 uint8, NbCaptureP2 uint8, parent *node, depth uint8) *node {
	return &node{
		id:               id,
		value:            value,
		goban:            newGoban,
		coord:            coord,
		neighbors:        neighbors,
		player:           player,
		maximizingPlayer: maximizingPlayer,
		captures: captures{
			capture0: NbCaptureP1,
			capture1: NbCaptureP2,
		},
		parent: parent,
		depth:  depth,
	}
}

func generateBoard(current *node, coord s.SVertex, neighbors []s.SVertex, player uint8) {
	var value int
	identity++
	newGoban := copyGoban(current.goban, 19)
	newGoban[coord.Y][coord.X] = s.Tnumber(player)

	var ctx s.SContext

	ctx.Goban = newGoban
	ctx.NSize = 19
	ctx.CurrentPlayer = player
	ctx.NbCaptureP1 = int(current.captures.capture0)
	ctx.NbCaptureP2 = int(current.captures.capture1)
	ctx.ActiveCapture = false

	if player == 1 {
		player = 2
	} else {
		player = 1
	}

	if current.maximizingPlayer {
		value = current.value - int(h.CalcHeuristic(ctx))/int(current.depth)
		// value = current.value - evaluateMove(coord, newGoban, player, current.captures)/int(current.depth)

	} else {
		value = current.value + int(h.CalcHeuristic(ctx))/int(current.depth)
		// value = current.value + evaluateMove(coord, newGoban, player, current.captures)/int(current.depth)

	}

	newNeighbors := getNeighbors(current.goban, coord)

	child := createNode(identity, value, newGoban, coord, newNeighbors, player, !current.maximizingPlayer, current.captures.capture1, current.captures.capture1, current, current.depth+1)
	current.children = append(current.children, child)
}

func generateTree(current *node, coord s.SVertex, neighbors []s.SVertex) {
	player := uint8(1)

	if current.player == 1 {
		player = 2
	}

	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(current.goban, neighbor.X, neighbor.Y)
		if placement >= 1 {
			generateBoard(current, neighbor, neighbors, player)
		}
	}
}
