package algorithm

import (
	s "gomoku/structures"
)

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

func inNeighbors2(ctx s.SContext, vertex s.SVertex, ret_list []s.SVertex) []s.SVertex {
	if (vertex.Y < 0 || vertex.Y >= int(ctx.NSize)) || (vertex.X < 0 || vertex.X >= int(ctx.NSize)) {
		return ret_list
	}
	if ctx.Goban[vertex.Y][vertex.X] != 0 {
		return ret_list
	}
	ret_list = append(ret_list, vertex)
	return ret_list
}

func getNeighbors(ctx s.SContext, vertex s.SVertex) []s.SVertex {
	if ctx.CasesNonNull == nil {
		ctx.CasesNonNull = make([]s.SVertex, 0)
	}
	ret_list := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(ret_list, ctx.CasesNonNull)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X + 1, Y: vertex.Y + 1}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X - 1, Y: vertex.Y - 1}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X + 1, Y: vertex.Y - 1}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X - 1, Y: vertex.Y + 1}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X, Y: vertex.Y + 1}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X, Y: vertex.Y - 1}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X + 1, Y: vertex.Y}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X - 1, Y: vertex.Y}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X - 1, Y: vertex.Y}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X - 1, Y: vertex.Y}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X - 1, Y: vertex.Y}, ret_list)
	ret_list = inNeighbors2(ctx, s.SVertex{X: vertex.X - 1, Y: vertex.Y}, ret_list)
	ret_list = removeDuplicate2(ret_list, s.SVertex{X: vertex.X, Y: vertex.Y})
	return ret_list
}
