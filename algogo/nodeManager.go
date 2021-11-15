package algogo

import (
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
		captures: Captures{
			Capture0: NbCaptureP1,
			Capture1: NbCaptureP2,
		},
		parent: parent,
		depth:  depth,
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
	newNeighbors := getNeighbors(current.goban, coord)

	var ctx s.SContext

	ctx.Goban = newGoban
	ctx.NSize = 19
	ctx.CurrentPlayer = opp
	ctx.NbCaptureP1 = int(current.captures.Capture0)
	ctx.NbCaptureP2 = int(current.captures.Capture1)
	ctx.ActiveDoubleThrees = isDoubleThree
	ctx.ActiveCapture = isCapture

	if isCapture {
		capturesVertex := CaptureAlgo(&ctx, coord.X, coord.Y)

		for _, capture := range capturesVertex {
			ctx.Goban[capture.Y][capture.X] = 0
			newNeighbors = append(newNeighbors, s.SVertex{Y: capture.Y, X: capture.X})
		}
	}

	if current.maximizingPlayer {
		value = current.value - int(EvaluateMove(ctx, coord.X, coord.Y))/int(current.depth)
	} else {
		value = current.value + int(EvaluateMove(ctx, coord.X, coord.Y))/int(current.depth)
	}
	// value = current.value + int(EvaluateGoban(ctx))/int(current.depth)

	child := createNode(identity, value, ctx.Goban, coord, newNeighbors, opp, !current.maximizingPlayer, uint8(ctx.NbCaptureP1), uint8(ctx.NbCaptureP2), current, current.depth+1)
	current.children = append(current.children, child)
}

func generateTree(current *node, neighbors []s.SVertex) {
	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(current.goban, neighbor.X, neighbor.Y, current.player)
		if placement >= 1 {
			generateBoard(current, neighbor, neighbors)
		}
	}
}
