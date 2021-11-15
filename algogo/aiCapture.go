package algogo

import (
	g "gomoku/game"
	s "gomoku/structures"
)

func measureChain(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) int8 {
	var length int8
	var multiple int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := g.FindNeighbour(coordinate, y, x, multiple)
		if g.PositionOccupiedByPlayer(neighbour, goban, player) == true {
			length++
		} else {
			break
		}
	}
	return length
}

func checkVertexAlignFive(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) bool {
	a := measureChain(coordinate, goban, y, x, player)
	b := measureChain(coordinate, goban, -y, -x, player)
	if a+b >= 4 {
		return true
	}
	return false
}

func breakFive(coordinate s.SVertex, goban s.Tgoban, player uint8) bool {
	var x int8
	var y int8
	for y = -1; y <= 0; y++ {
		for x = -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				return false
			}
			if checkVertexAlignFive(coordinate, goban, y, x, player) == true {
				return true
			}
		}
	}
	return false
}

func breakFiveDirection(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8, opp uint8) bool {
	one := g.FindNeighbour(coordinate, y, x, 1)
	two := g.FindNeighbour(coordinate, y, x, 2)
	if (g.PositionOccupiedByOpponent(one, goban, player) == true && breakFive(one, goban, opp) == true) ||
		(g.PositionOccupiedByOpponent(two, goban, player) == true && breakFive(two, goban, opp) == true) {
		return true
	}
	return false
}

func willBreak5Align(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8, opp uint8) bool {
	if breakFiveDirection(coordinate, goban, y, x, player, opp) == true || breakFiveDirection(coordinate, goban, -y, -x, player, opp) == true {
		return true
	}
	return false
}

func canBeCapturedVertex(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) bool {
	minusOne := g.FindNeighbour(coordinate, y, x, -1)
	one := g.FindNeighbour(coordinate, y, x, 1)
	two := g.FindNeighbour(coordinate, y, x, 2)
	if g.PositionOccupiedByPlayer(one, goban, player) {
		if (g.PositionOccupiedByOpponent(minusOne, goban, player) && g.PositionUnoccupied(two, goban)) ||
			(g.PositionOccupiedByOpponent(two, goban, player) && g.PositionUnoccupied(minusOne, goban)) {
			return true
		}
	}
	return false
}

func willCaptureDirection(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) bool {
	one := g.FindNeighbour(coordinate, y, x, 1)
	two := g.FindNeighbour(coordinate, y, x, 2)
	three := g.FindNeighbour(coordinate, y, x, 3)
	if g.PositionOccupiedByOpponent(one, goban, player) == true &&
		g.PositionOccupiedByOpponent(two, goban, player) == true &&
		g.PositionOccupiedByPlayer(three, goban, player) == true {
		return true
	}
	return false
}

func willCapture(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) uint8 {
	var cap uint8
	if willCaptureDirection(coordinate, goban, y, x, player) == true {
		cap++
	}
	if willCaptureDirection(coordinate, goban, -y, -x, player) == true {
		cap++
	}
	return cap
}

func willBeCaptured(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) bool {
	if canBeCapturedVertex(coordinate, goban, y, x, player) == true || canBeCapturedVertex(coordinate, goban, -y, -x, player) == true {
		return true
	}
	return false
}

func capturedEight(player uint8, capture0 uint8, capture1 uint8) bool {
	if (player == 1 && capture0 >= 8) ||
		(player == 2 && capture1 >= 8) {
		return true
	}
	return false
}

func captureAttackDefend(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8, captures Captures) int {
	opp := uint8(1)
	if player == 1 {
		opp = 2
	}

	cap := willCapture(coordinate, goban, y, x, player)
	if cap != 0 {
		if capturedEight(player, captures.Capture0, captures.Capture1) == true {
			return capture10
		} else if willBreak5Align(coordinate, goban, y, x, player, opp) == true {
			return break5Align
		}
		if cap == 2 {
			return capture2 * 2
		}
		return capture2
	} else if willBeCaptured(coordinate, goban, y, x, player) == true {
		if capturedEight(opp, captures.Capture0, captures.Capture1) == true {
			return willBeCaptured8
		}
		return willBeCaptured2
	}
	return 0
}
