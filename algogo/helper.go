package algogo

import (
	g "gomoku/game"
	s "gomoku/structures"
)

const capture10 = 100000
const break5Align = 12000
const willBeCaptured8 = -10000
const align5Win = 10000
const blockWin = 3000
const willBeCaptured2 = -1500
const align4Free = 1490
const block3Free = 1300
const capture2 = 1000
const align3Free = 1000
const align4FLanked = 500
const align3Flanked = 400
const block2 = 1
const align2Free = 1

var isCapture bool
var isDoubleThree bool
var initPlayer uint8

type playData struct {
	Vertex s.SVertex
	Heur   int32
}

type Captures struct {
	Capture0 uint8
	Capture1 uint8
}

type node struct {
	id               int
	value            int
	goban            s.Tgoban
	coord            s.SVertex
	neighbors        []s.SVertex
	player           uint8
	maximizingPlayer bool
	captures         Captures
	parent           *node
	children         []*node
	bestMove         *node
	depth            uint8
}

func copyGoban(goban s.Tgoban) s.Tgoban {
	newGoban := make([][]s.Tnumber, 19)
	for Y, line := range goban {
		newGoban[Y] = make([]s.Tnumber, 19)
		for X, nb := range line {
			newGoban[Y][X] = nb
		}
	}
	return newGoban
}

func PlacementHeuristic(goban s.Tgoban, case_x int, case_y int, player uint8) uint8 {
	// if ActiveCapture && len(ctx.Capture) > 0 {
	// 	for _, cap := range ctx.Capture {
	// 		if case_x == cap.X && case_y == cap.Y {
	// 			capture = true
	// 			break
	// 		}
	// 	}
	// 	if capture == false {
	// 		return 2
	// 	}
	// }
	if isDoubleThree && g.DoubleThree(s.SVertex{X: case_x, Y: case_y}, goban, player, isCapture) {
		return 0
	}
	if case_y < 0 || case_y >= 19 {
		return 0
	}
	if case_x < 0 || case_x >= 19 {
		return 0
	}
	if goban[int(case_y)][int(case_x)] == 0 {
		return 1
	} else {
		return 0
	}
}

func removeDuplicate2(ret_list []s.SVertex, vertex s.SVertex) []s.SVertex {
	keys := make(map[s.SVertex]bool)
	list := []s.SVertex{}
	for _, entry := range ret_list {
		if _, value := keys[entry]; !value && entry != vertex {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func inNeighbors2(goban s.Tgoban, vertex s.SVertex, ret_list []s.SVertex) []s.SVertex {
	if (vertex.Y < 0 || vertex.Y >= 19) || (vertex.X < 0 || vertex.X >= 19) {
		return ret_list
	}
	if goban[vertex.Y][vertex.X] != 0 {
		return ret_list
	}
	ret_list = append(ret_list, vertex)
	return ret_list
}

func getNeighbors(goban s.Tgoban, oldNeighbors []s.SVertex, vertex s.SVertex) []s.SVertex {
	ret_list := make([]s.SVertex, len(oldNeighbors))
	copy(ret_list, oldNeighbors)
	ret_list = inNeighbors2(goban, s.SVertex{X: vertex.X + 1, Y: vertex.Y + 1}, ret_list)
	ret_list = inNeighbors2(goban, s.SVertex{X: vertex.X - 1, Y: vertex.Y - 1}, ret_list)
	ret_list = inNeighbors2(goban, s.SVertex{X: vertex.X + 1, Y: vertex.Y - 1}, ret_list)
	ret_list = inNeighbors2(goban, s.SVertex{X: vertex.X - 1, Y: vertex.Y + 1}, ret_list)
	ret_list = inNeighbors2(goban, s.SVertex{X: vertex.X, Y: vertex.Y + 1}, ret_list)
	ret_list = inNeighbors2(goban, s.SVertex{X: vertex.X, Y: vertex.Y - 1}, ret_list)
	ret_list = inNeighbors2(goban, s.SVertex{X: vertex.X + 1, Y: vertex.Y}, ret_list)
	ret_list = inNeighbors2(goban, s.SVertex{X: vertex.X - 1, Y: vertex.Y}, ret_list)
	ret_list = removeDuplicate2(ret_list, s.SVertex{X: vertex.X, Y: vertex.Y})
	return ret_list
}
