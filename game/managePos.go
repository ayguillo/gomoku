package game

import s "gomoku/structures"

func coordinateOnGoban(coordinate s.SVertex) bool {
	if coordinate.Y < 0 || coordinate.Y > 18 ||
		coordinate.X < 0 || coordinate.X > 18 {
		return false
	}
	return true
}

func positionOccupiedByPlayer(coordinate s.SVertex, goban s.Tgoban, player uint8) bool {
	if coordinateOnGoban(coordinate) == true {
		if goban[coordinate.Y][coordinate.X] == s.Tnumber(player) {
			return true
		}
	}
	return false
}

func positionOccupiedByOpponent(coordinate s.SVertex, goban s.Tgoban, player uint8) bool {
	if coordinateOnGoban(coordinate) == true {
		if goban[coordinate.Y][coordinate.X] != 0 && goban[coordinate.Y][coordinate.X] != s.Tnumber(player) {
			return true
		}
	}
	return false
}

func positionUnoccupied(coordinate s.SVertex, goban s.Tgoban) bool {
	if coordinateOnGoban(coordinate) == true {
		if goban[coordinate.Y][coordinate.X] == 0 {
			return true
		}
	}
	return false
}

func findNeighbour(coordinate s.SVertex, y int8, x int8, multiple int8) s.SVertex {
	neighbour := coordinate
	neighbour.Y += int(y * multiple)
	neighbour.X += int(x * multiple)
	return neighbour
}
