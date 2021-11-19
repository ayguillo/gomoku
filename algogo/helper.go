package algogo

import (
	d "gomoku/doubleThree"
	s "gomoku/structures"
)

var isCapture bool
var isDoubleThree bool

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
	lastMove        s.SVertex
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
	if isDoubleThree && d.DoubleThree(s.SVertex{X: case_x, Y: case_y}, goban, player, isCapture) {
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
	oldNeighbors = inNeighbors2(goban, s.SVertex{X: vertex.X + 1, Y: vertex.Y + 1}, oldNeighbors)
	oldNeighbors = inNeighbors2(goban, s.SVertex{X: vertex.X - 1, Y: vertex.Y - 1}, oldNeighbors)
	oldNeighbors = inNeighbors2(goban, s.SVertex{X: vertex.X + 1, Y: vertex.Y - 1}, oldNeighbors)
	oldNeighbors = inNeighbors2(goban, s.SVertex{X: vertex.X - 1, Y: vertex.Y + 1}, oldNeighbors)
	oldNeighbors = inNeighbors2(goban, s.SVertex{X: vertex.X, Y: vertex.Y + 1}, oldNeighbors)
	oldNeighbors = inNeighbors2(goban, s.SVertex{X: vertex.X, Y: vertex.Y - 1}, oldNeighbors)
	oldNeighbors = inNeighbors2(goban, s.SVertex{X: vertex.X + 1, Y: vertex.Y}, oldNeighbors)
	oldNeighbors = inNeighbors2(goban, s.SVertex{X: vertex.X - 1, Y: vertex.Y}, oldNeighbors)
	oldNeighbors = removeDuplicate2(oldNeighbors, s.SVertex{X: vertex.X, Y: vertex.Y})
	return oldNeighbors
}

func removeDuplicate(ctx *s.SContext, vertex s.SVertex) {
	keys := make(map[s.SVertex]bool)
	list := []s.SVertex{}
	for _, entry := range ctx.CasesNonNull {
		if _, value := keys[entry]; !value && entry != vertex {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	ctx.CasesNonNull = list
}

func inNeighbors(ctx *s.SContext, vertex s.SVertex) {
	if (vertex.Y < 0 || vertex.Y >= int(ctx.NSize)) || (vertex.X < 0 || vertex.X >= int(ctx.NSize)) {
		return
	}
	if ctx.Goban[vertex.Y][vertex.X] != 0 {
		return
	}
	ctx.CasesNonNull = append(ctx.CasesNonNull, vertex)
}

func FindNeighbors(ctx *s.SContext, case_x int, case_y int) {
	if ctx.CasesNonNull == nil {
		ctx.CasesNonNull = make([]s.SVertex, 0)
	}
	inNeighbors(ctx, s.SVertex{X: case_x + 1, Y: case_y + 1})
	inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y - 1})
	inNeighbors(ctx, s.SVertex{X: case_x + 1, Y: case_y - 1})
	inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y + 1})
	inNeighbors(ctx, s.SVertex{X: case_x, Y: case_y + 1})
	inNeighbors(ctx, s.SVertex{X: case_x, Y: case_y - 1})
	inNeighbors(ctx, s.SVertex{X: case_x + 1, Y: case_y})
	inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y})
	inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y})
	inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y})
	inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y})
	inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y})
	removeDuplicate(ctx, s.SVertex{X: case_x, Y: case_y})
}
