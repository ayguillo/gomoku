package algogo

import (
	g "gomoku/game"
	s "gomoku/structures"
)

func measureOpponent(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) (int8, bool) {
	var multiple int8
	var length int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := g.FindNeighbour(coordinate, y, x, multiple)
		if g.PositionOccupiedByOpponent(neighbour, goban, player) == true {
			length++
		} else if g.PositionUnoccupied(neighbour, goban) == false {
			return length, true
		} else {
			break
		}
	}
	return length, false
}

func measurePlayer(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) (int8, bool) {
	var length int8
	var multiple int8
	for multiple = 1; multiple < 5; multiple++ {
		neighbour := g.FindNeighbour(coordinate, y, x, multiple)
		if g.PositionOccupiedByPlayer(neighbour, goban, player) == true {
			length++
		} else if g.PositionUnoccupied(neighbour, goban) == false {
			return length, true
		} else {
			break
		}
	}
	return length, false
}

func lengthPlayerChain(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) (int8, bool, bool) {
	a, flanked1 := measurePlayer(coordinate, goban, y, x, player)
	b, flanked2 := measurePlayer(coordinate, goban, -y, -x, player)
	if a+b+1 > 5 {
		return 5, flanked1, flanked2
	}
	if a+b+1 == 3 && flanked1 != flanked2 {
		if !(a == 2 && flanked1 || b == 2 && flanked2) {
			return 0, flanked1, flanked2
		}
	}
	return a + b + 1, flanked1, flanked2
}

func lengthOpponentChain(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) (int8, bool, bool) {
	a, flanked1 := measureOpponent(coordinate, goban, y, x, player)
	b, flanked2 := measureOpponent(coordinate, goban, -y, -x, player)
	if a+b > 4 {
		return 4, flanked1, flanked2
	}
	return a + b, flanked1, flanked2
}

func chainAttackDefend(coordinate s.SVertex, goban s.Tgoban, y, x int8, player uint8) int {
	var attack int
	var defend int

	opponentChain, flanked1, flanked2 := lengthOpponentChain(coordinate, goban, y, x, player)
	switch opponentChain {
	case 4:
		defend = blockWin
	case 3:
		if flanked1 == false && flanked2 == false {
			defend = block3Free
		}
	case 2:
		if flanked1 == false && flanked2 == false {
			defend = block2
		}
	}

	playerChain, flanked1, flanked2 := lengthPlayerChain(coordinate, goban, y, x, player)
	switch playerChain {
	case 5:
		attack = align5Win
	case 4:
		if flanked1 == false && flanked2 == false {
			attack = align4Free
		} else if flanked1 == true && flanked2 == true {
			attack = 0
		} else {
			attack = align4FLanked
		}
	case 3:
		if flanked1 == false && flanked2 == false {
			attack = align3Free
		} else if flanked1 == true && flanked2 == true {
			attack = 0
		} else {
			attack = align3Flanked
		}
	case 2:
		if flanked1 == false && flanked2 == false {
			attack = align2Free
		}
	}
	println("ad", x, y, attack, defend)
	return attack + defend
}

func EvaluateMove(coord s.SVertex, goban s.Tgoban, player uint8, captures Captures) int {
	var x int8
	var y int8

	eval := 0
	for y = -1; y <= 1; y++ {
		for x = -1; x <= 1; x++ {
			if x != 0 && y != 0 {
				if isCapture {
					eval += captureAttackDefend(coord, goban, y, x, player, captures)
				}
				eval += chainAttackDefend(coord, goban, y, x, player)
			}
		}
	}

	return eval
}
