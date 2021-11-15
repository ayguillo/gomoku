package game

import s "gomoku/structures"

func threeBlocked(end1 s.SVertex, end2 s.SVertex, goban s.Tgoban) bool {
	if PositionUnoccupied(end1, goban) == true &&
		PositionUnoccupied(end2, goban) == true {
		return false
	}
	return true
}

func checkVertexForThree(coordinate s.SVertex, goban s.Tgoban, y int8, x int8, player uint8) bool {
	minusTwo := FindNeighbour(coordinate, y, x, -2)
	minusOne := FindNeighbour(coordinate, y, x, -1)
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	four := FindNeighbour(coordinate, y, x, 4)
	if PositionOccupiedByPlayer(one, goban, player) == true {
		if PositionOccupiedByPlayer(two, goban, player) == true &&
			threeBlocked(minusOne, three, goban) == false {
			return true
		}
		if PositionOccupiedByPlayer(three, goban, player) == true &&
			threeBlocked(minusOne, four, goban) == false &&
			PositionOccupiedByOpponent(two, goban, player) == false {
			return true
		}
		if y < 0 || (y == 0 && x == -1) {
			if PositionOccupiedByPlayer(minusOne, goban, player) == true &&
				threeBlocked(minusTwo, two, goban) == false {
				return true
			}
		}
	}
	if PositionOccupiedByPlayer(two, goban, player) == true {
		if PositionOccupiedByPlayer(three, goban, player) == true &&
			threeBlocked(minusOne, four, goban) == false &&
			PositionOccupiedByOpponent(one, goban, player) == false {
			return true
		}
		if PositionOccupiedByPlayer(minusOne, goban, player) == true &&
			threeBlocked(minusTwo, three, goban) == false &&
			PositionOccupiedByOpponent(one, goban, player) == false {
			return true
		}
	}
	return false
}

func isCaptureDirection(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) bool {
	one := FindNeighbour(coordinate, y, x, 1)
	two := FindNeighbour(coordinate, y, x, 2)
	three := FindNeighbour(coordinate, y, x, 3)
	if PositionOccupiedByOpponent(one, goban, player) == true &&
		PositionOccupiedByOpponent(two, goban, player) == true &&
		PositionOccupiedByPlayer(three, goban, player) == true {
		return true
	}
	return false
}

func isCapture(coordinate s.SVertex, goban s.Tgoban, player uint8) bool {
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				if isCaptureDirection(coordinate, goban, y, x, player) == true {
					return true
				}
			}
		}
	}
	return false
}

func DoubleThree(coordinate s.SVertex, goban s.Tgoban, player uint8, capture bool) bool {
	if capture && isCapture(coordinate, goban, player) == true {
		return false
	}
	var freeThree bool
	var y int8
	var x int8
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if !(x == 0 && y == 0) {
				foundThree := checkVertexForThree(coordinate, goban, y, x, player)
				if foundThree == true {
					if freeThree == true {
						return true
					} else {
						freeThree = true
					}
				}
			}
		}
	}
	return false
}
