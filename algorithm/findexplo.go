package algorithm

import (
	s "gomoku/structures"
)

func inNeighbors(ctx *s.SContext, vertex s.SVertex, array_neigbors []s.SVertex) []s.SVertex {
	if ctx.Goban[vertex.Y][vertex.X] != 0 {
		return array_neigbors
	}
	for _, elem := range ctx.CasesNonNull {
		for _, onecase := range elem {
			if onecase == vertex {
				return array_neigbors
			}
		}
	}
	array_neigbors = append(array_neigbors, vertex)
	return array_neigbors
}

func FindNeighbors(ctx *s.SContext, case_x int, case_y int, visu *s.SVisu) {
	if ctx.CasesNonNull == nil {
		ctx.CasesNonNull = make(map[s.SVertex][]s.SVertex)
	}
	cases_vertex := s.SVertex{X: case_x, Y: case_y}
	array := make([]s.SVertex, 0)

	array = inNeighbors(ctx, s.SVertex{X: case_x + 1, Y: case_y + 1}, array)
	array = inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y - 1}, array)
	array = inNeighbors(ctx, s.SVertex{X: case_x + 1, Y: case_y - 1}, array)
	array = inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y + 1}, array)
	array = inNeighbors(ctx, s.SVertex{X: case_x, Y: case_y + 1}, array)
	array = inNeighbors(ctx, s.SVertex{X: case_x, Y: case_y - 1}, array)
	array = inNeighbors(ctx, s.SVertex{X: case_x + 1, Y: case_y}, array)
	array = inNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y}, array)
	ctx.CasesNonNull[cases_vertex] = array
	// color := [4]uint8{83, 51, 237, 1}
	// for _, elem := range ctx.CasesNonNull {
	// 	// fmt.Println("Key", key)
	// 	for _, onecase := range elem {
	// 		// fmt.Println("\tarray", onecase)
	// 		d.TraceStone(float64(onecase.X), float64(onecase.Y), ctx, visu, color, false)
	// 	}
	// }
}
