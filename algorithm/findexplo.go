package algorithm

import (
	s "gomoku/structures"
)

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

func FindLastNeighbors(ctx *s.SContext, vertex s.SVertex) {
	if ctx.LastMove.X != -1 {
		tmp_last := ctx.LastMove
		ctx.LastMove = vertex
		ctx.LastLastMove = tmp_last
		ctx.CasesNonNull = nil
		FindNeighbors(ctx, ctx.LastMove.X, ctx.LastMove.Y)
		FindNeighbors(ctx, ctx.LastLastMove.X, ctx.LastLastMove.Y)
	} else {
		FindNeighbors(ctx, vertex.X, vertex.Y)
		ctx.LastMove = vertex
	}
}
