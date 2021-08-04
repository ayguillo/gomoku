package algorithm

import (
	"fmt"
	s "gomoku/structures"
)

func PlayOne(ctx s.SContext) s.SVertex {
	vertex_next := s.SVertex{X: -1, Y: -1}
	heuris := -50000
	for stone, _ := range ctx.CasesNonNull {
		for _, neighbor := range ctx.CasesNonNull[stone] {
			placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
			if placement == 1 {
				tmp_heuris := Heuristic(ctx, int(neighbor.X), int(neighbor.Y))
				if tmp_heuris >= heuris {
					heuris = tmp_heuris
					vertex_next = neighbor
				}
			} else if placement == 2 {
				heuris = 50000
				vertex_next = neighbor
			}
		}
	}
	fmt.Println("Heuristic = ", heuris)
	return vertex_next
}
