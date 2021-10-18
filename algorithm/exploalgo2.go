package algorithm

import (
	s "gomoku/structures"
)

func allNeighbors(ctx s.SContext, vertex s.SVertex, array_neigbors []s.SVertex, neigbors map[s.SVertex][]s.SVertex) []s.SVertex {
	if (vertex.Y < 0 || vertex.Y >= int(ctx.NSize)) || (vertex.X < 0 || vertex.X >= int(ctx.NSize)) {
		return array_neigbors
	}
	if ctx.Goban[vertex.Y][vertex.X] != 0 {
		return array_neigbors
	}
	for _, elem := range neigbors {
		for _, onecase := range elem {
			if onecase == vertex {
				return array_neigbors
			}
		}
	}
	array_neigbors = append(array_neigbors, vertex)
	return array_neigbors
}

func FindNeighborsExplo(ctx s.SContext, case_x int, case_y int, neigbors map[s.SVertex][]s.SVertex) map[s.SVertex][]s.SVertex {
	cases_vertex := s.SVertex{X: case_x, Y: case_y}
	array := make([]s.SVertex, 0)
	cases := make(map[s.SVertex][]s.SVertex)

	array = allNeighbors(ctx, s.SVertex{X: case_x + 1, Y: case_y + 1}, array, neigbors)
	array = allNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y - 1}, array, neigbors)
	array = allNeighbors(ctx, s.SVertex{X: case_x + 1, Y: case_y - 1}, array, neigbors)
	array = allNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y + 1}, array, neigbors)
	array = allNeighbors(ctx, s.SVertex{X: case_x, Y: case_y + 1}, array, neigbors)
	array = allNeighbors(ctx, s.SVertex{X: case_x, Y: case_y - 1}, array, neigbors)
	array = allNeighbors(ctx, s.SVertex{X: case_x + 1, Y: case_y}, array, neigbors)
	array = allNeighbors(ctx, s.SVertex{X: case_x - 1, Y: case_y}, array, neigbors)
	cases[cases_vertex] = array
	for key, elem := range neigbors {
		len_elem, index := len(elem), 0
		for index < len_elem {
			if ctx.Goban[elem[index].Y][elem[index].X] == 0 {
				cases[key] = append(cases[key], elem[index])
			}
			index++
		}
	}
	return cases
}
