package doubleThree

import (
	s "gomoku/structures"
)

func isInRange(x int, y int) bool {
	if x < 0 || x >= 19 || y < 0 || y >= 19 {
		return false
	}

	return true
}

func checkSideDoubleThree(goban s.Tgoban, minusOne s.SVertex, minusTwo s.SVertex, minusThree s.SVertex, minusFour s.SVertex, plusOne s.SVertex, player uint8) bool {
	if PositionUnoccupied(plusOne, goban) {
		if PositionOccupiedByPlayer(minusOne, goban, player) {
			if PositionOccupiedByPlayer(minusTwo, goban, player) && PositionUnoccupied(minusThree, goban) {
				return true
			}
			if PositionUnoccupied(minusTwo, goban) && PositionOccupiedByPlayer(minusThree, goban, player) && PositionUnoccupied(minusFour, goban) {
				return true
			}
		}
		if PositionUnoccupied(minusOne, goban) && PositionOccupiedByPlayer(minusTwo, goban, player) && PositionOccupiedByPlayer(minusThree, goban, player) && PositionUnoccupied(minusFour, goban) {
			return true
		}
	}

	return false
}

func checkMidDoubleThree(goban s.Tgoban, plusOne s.SVertex, plusTwo s.SVertex, plusThree s.SVertex, minusOne s.SVertex, minusTwo s.SVertex, minusThree s.SVertex, player uint8) bool {
	if PositionOccupiedByPlayer(minusOne, goban, player) && PositionUnoccupied(minusTwo, goban) {
		if PositionOccupiedByPlayer(plusOne, goban, player) && PositionUnoccupied(plusTwo, goban) {
			return true
		}
		if PositionUnoccupied(plusOne, goban) && PositionOccupiedByPlayer(plusTwo, goban, player) && PositionUnoccupied(plusThree, goban) {
			return true
		}
	}
	if PositionUnoccupied(minusOne, goban) && PositionOccupiedByPlayer(minusTwo, goban, player) && PositionUnoccupied(minusThree, goban) {
		if PositionOccupiedByPlayer(plusOne, goban, player) && PositionUnoccupied(plusTwo, goban) {
			return true
		}
	}

	return false
}

func checkDoubleThree(goban s.Tgoban, case_x int, case_y int, x int, y int, player uint8) bool {
	coordinate := s.SVertex{Y: case_y, X: case_x}

	minusOne := FindNeighbour(coordinate, y, x, -1)
	minusTwo := FindNeighbour(coordinate, y, x, -2)
	minusThree := FindNeighbour(coordinate, y, x, -3)
	minusFour := FindNeighbour(coordinate, y, x, -4)
	plusOne := FindNeighbour(coordinate, y, x, 1)
	plusTwo := FindNeighbour(coordinate, y, x, 2)
	plusThree := FindNeighbour(coordinate, y, x, 3)
	plusFour := FindNeighbour(coordinate, y, x, 4)

	if checkSideDoubleThree(goban, minusOne, minusTwo, minusThree, minusFour, plusOne, player) ||
		checkSideDoubleThree(goban, plusOne, plusTwo, plusThree, plusFour, minusOne, player) ||
		checkMidDoubleThree(goban, plusOne, plusTwo, plusThree, minusOne, minusTwo, minusThree, player) {
		return true
	}

	return false
}

func DoubleThree(coordinate s.SVertex, goban s.Tgoban, player uint8, capture bool) bool {
	if capture && isCapture(goban, coordinate.X, coordinate.Y, player) == true {
		return false
	}

	horizontal := checkDoubleThree(goban, coordinate.X, coordinate.Y, -1, 0, player)
	vertical := checkDoubleThree(goban, coordinate.X, coordinate.Y, 0, -1, player)
	diagLeft := checkDoubleThree(goban, coordinate.X, coordinate.Y, -1, -1, player)
	diagRight := checkDoubleThree(goban, coordinate.X, coordinate.Y, -1, 1, player)

	if horizontal && (vertical || diagLeft || diagRight) || 
		vertical && (horizontal || diagLeft || diagRight) ||
		diagLeft && (vertical || horizontal || diagRight) ||
		diagRight && (vertical || diagLeft || horizontal) {
		return true
	}
	return false
}
