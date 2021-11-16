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
	// if capture && isCapture(coordinate, goban, player) == true {
	// 	return false
	// }
	// var freeThree bool
	// var y int
	// var x int
	// for y = -1; y <= 1; y++ {
	// 	for x = -1; x <= 1; x++ {
	// 		if !(x == 0 && y == 0) {
	// 			foundThree := checkDoubleThree(goban, coordinate.X, coordinate.Y, x, y, player)
	// 			if foundThree == true {
	// 				return true
	// 				if freeThree == true {
	// 					return true
	// 				} else {
	// 					freeThree = true
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	horizontal := checkDoubleThree(goban, coordinate.X, coordinate.Y, -1, 0, player)
	if horizontal {
		return true
	}
	return false
}
