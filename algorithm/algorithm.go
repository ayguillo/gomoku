package algorithm

import (
	s "gomoku/structures"
	"math"
)

func CopyGoban(ctx s.SContext) s.Tgoban {
	newGoban := make([][]s.Tnumber, ctx.NSize)
	for Y, line := range ctx.Goban {
		newGoban[Y] = make([]s.Tnumber, ctx.NSize)
		for X, nb := range line {
			newGoban[Y][X] = nb
		}
	}
	return newGoban
}

func CopyCases(ctx s.SContext) map[s.SVertex][]s.SVertex {
	cases := make(map[s.SVertex][]s.SVertex)
	for key, elem := range ctx.CasesNonNull {
		cases[key] = elem
	}
	return cases
}

func AlphaBetaPruning(ctx s.SContext, max_explor int) (s.SVertex, int32) {
	alpha := int32(-1 << 31)
	beta := int32(1<<31 - 1)
	vertex := s.SVertex{X: -1, Y: -1}
	u := int32(-1000000)
	for stone := range ctx.CasesNonNull {
		for _, neighbor := range ctx.CasesNonNull[stone] {
			tmp_u := max_player(ctx, alpha, beta, max_explor, nil)
			if tmp_u >= u {
				u = tmp_u
				vertex = neighbor
			}
		}
	}
	return vertex, u
}

// 1st player
func max_player(ctx s.SContext, alpha int32, beta int32, explor int, new_vertex *s.SVertex) int32 {
	// fmt.Println("alpha beta max", alpha, beta)
	u := int32(-1 << 31)
	goban := CopyGoban(ctx)
	casesnonnull := CopyCases(ctx)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  casesnonnull,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize}
	if explor == 0 {
		tmp_heuris := Heuristic(tmp_ctx, int(new_vertex.X), int(new_vertex.Y))
		return int32(tmp_heuris)
	}
	if new_vertex != nil {
		FindNeighbors(&tmp_ctx, int(new_vertex.X), int(new_vertex.Y), nil)
	}
	// fmt.Println(alpha, beta)
	for stone := range tmp_ctx.CasesNonNull {
		for _, neighbor := range tmp_ctx.CasesNonNull[stone] {
			placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				if placement == 2 {
					return int32(50000)
				}
				// fmt.Println("Max")
				// tmp_heuris := Heuristic(tmp_ctx, int(neighbor.X), int(neighbor.Y))
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(tmp_ctx.CurrentPlayer)
				if tmp_ctx.CurrentPlayer == 1 {
					tmp_ctx.CurrentPlayer = 2
				} else {
					tmp_ctx.CurrentPlayer = 1
				}
				tmp_u := min_player(tmp_ctx, alpha, beta, explor-1, &neighbor)
				// fmt.Println("u max", u)
				// Alpha beta prunning a ajouter
				if u >= beta {
					return u
				} else {
					alpha = int32(math.Max(float64(alpha), float64(u)))
				}
				u = int32(math.Max(float64(tmp_u), float64(u)))
			}
			
		}
	}
	return u
}

// 2nd player
func min_player(ctx s.SContext, alpha int32, beta int32, explor int, new_vertex *s.SVertex) int32 {
	// fmt.Println("alpha beta min", alpha, beta)
	playerMin := s.Tnumber(0)
	if ctx.CurrentPlayer == 1 {
		playerMin = s.Tnumber(2)
	} else {
		playerMin = s.Tnumber(1)
	}
	goban := CopyGoban(ctx)
	casesnonnull := CopyCases(ctx)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  casesnonnull,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize}
	if new_vertex != nil {
		FindNeighbors(&tmp_ctx, int(new_vertex.X), int(new_vertex.Y), nil)
	}
	u := int32(1<<31 - 1)
	for stone := range tmp_ctx.CasesNonNull {
		for _, neighbor := range tmp_ctx.CasesNonNull[stone] {
			placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				if placement == 2 {
					return int32(-50000)
				}
				// fmt.Println("Min")
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(playerMin)
				tmp_u := max_player(tmp_ctx, alpha, beta, explor-1, &neighbor)
				// fmt.Println("u min", u)
				// Alpha beta prunning a ajouter
				if u <= alpha {
					return u
				} else {
					beta = int32(math.Min(float64(beta), float64(u)))
				}
				u = int32(math.Min(float64(tmp_u), float64(u)))
			}
		}
	}
	return u
}
