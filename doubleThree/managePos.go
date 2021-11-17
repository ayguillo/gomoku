package doubleThree

import s "gomoku/structures"

func coordinateOnGoban(coordinate s.SVertex) bool {
	if coordinate.Y < 0 || coordinate.Y > 18 ||
		coordinate.X < 0 || coordinate.X > 18 {
		return false
	}
	return true
}

func PositionOccupiedByPlayer(coordinate s.SVertex, goban s.Tgoban, player uint8) bool {
	if coordinateOnGoban(coordinate) == true {
		if goban[coordinate.Y][coordinate.X] == s.Tnumber(player) {
			return true
		}
	}
	return false
}

func PositionOccupiedByOpponent(coordinate s.SVertex, goban s.Tgoban, player uint8) bool {
	if coordinateOnGoban(coordinate) == true {
		if goban[coordinate.Y][coordinate.X] != 0 && goban[coordinate.Y][coordinate.X] != s.Tnumber(player) {
			return true
		}
	}
	return false
}

func PositionUnoccupied(coordinate s.SVertex, goban s.Tgoban) bool {
	if coordinateOnGoban(coordinate) == true {
		if goban[coordinate.Y][coordinate.X] == 0 {
			return true
		}
	}
	return false
}

func FindNeighbour(coordinate s.SVertex, y int, x int, multiple int) s.SVertex {
	neighbour := coordinate
	neighbour.Y += int(y * multiple)
	neighbour.X += int(x * multiple)
	return neighbour
}
