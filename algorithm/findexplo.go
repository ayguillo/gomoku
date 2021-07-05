package algorithm

import (
	"fmt"
	d "gomoku/display"
	s "gomoku/structures"
)

func inNeighbors(ctx *s.SContext, vertex s.SVertex, array_neigbors []s.SVertex) []s.SVertex {
	if (vertex.Y < 0 || vertex.Y >= 19) || (vertex.X < 0 || vertex.X >= 19) {
		return array_neigbors
	}
	fmt.Println(ctx.Goban[vertex.Y][vertex.X], vertex.Y, vertex.X)
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
	color := [4]uint8{83, 51, 237, 1}
	for key, elem := range ctx.CasesNonNull {
		fmt.Println("Key", key)
		len_elem, index := len(elem), 0
		for index < len_elem {
			fmt.Println("\tarray", elem[index])
			if ctx.Goban[elem[index].Y][elem[index].X] != 0 {
				elem[index] = elem[len(elem)-1]
				elem[len(elem)-1] = s.SVertex{X: 0, Y: 0}
				elem = elem[:len(elem)-1]
				len_elem--
			} else {
				d.TraceStone(float64(elem[index].X), float64(elem[index].Y), ctx, visu, color, false)
				index++
			}
		}
	}
	fmt.Println("\n")
}
