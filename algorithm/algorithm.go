package algorithm

import (
	"fmt"
	"math"
	"math/rand"

	g "gomoku/game"
	s "gomoku/structures"
)

// func CopyGoban(ctx s.SContext) s.Tgoban {
// 	newGoban := make([][]s.Tnumber, ctx.NSize)
// 	for Y, line := range ctx.Goban {
// 		newGoban[Y] = make([]s.Tnumber, ctx.NSize)
// 		for X, nb := range line {
// 			newGoban[Y][X] = nb
// 		}
// 	}
// 	return newGoban
// }

func AlphaBetaPruning(ctx s.SContext, max_explor int) (s.SVertex, int32) {
	alpha := int32(-2147483648)
	beta := int32(2147483647)
	vertex := s.SVertex{X: -1, Y: -1}
	u := int32(-1 << 31)
	for _, stone := range ctx.CasesNonNull {
		// fmt.Println(stone, ctx.CasesNonNull)
		tmp_u := max_player(ctx, alpha, beta, max_explor, &stone)
		if tmp_u > u {
			u = tmp_u
			vertex = stone
		}
	}
	return vertex, u
}

// 1st player
func max_player(ctx s.SContext, alpha int32, beta int32, explor int, new_vertex *s.SVertex) int32 {
	// fmt.Println("alpha beta max", alpha, beta)
	u := int32(-2147483648)
	goban := CopyGoban(ctx)
	casesnonnull := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(casesnonnull, ctx.CasesNonNull)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  casesnonnull,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize}
	if explor == 0 || g.VictoryConditionAlign(&ctx, int(new_vertex.X), int(new_vertex.Y), nil) == true {
		// tmp_heuris := Heuristic(tmp_ctx, int(new_vertex.X), int(new_vertex.Y))
		// fmt.Println("Ret max", tmp_heuris, explor)
		// fmt.Println(ctx)
		tmp_heuris := rand.Intn(10000)
		return int32(tmp_heuris)
	}
	if new_vertex != nil {
		FindNeighbors(&tmp_ctx, int(new_vertex.X), int(new_vertex.Y))
	}
	// // fmt.Println(alpha, beta)
	for _, neighbor := range tmp_ctx.CasesNonNull {
		placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(tmp_ctx.CurrentPlayer)
			tmp_u := min_player(tmp_ctx, alpha, beta, explor-1, &neighbor)
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(0)
			if tmp_u >= beta {
				return tmp_u
			} else {
				alpha = int32(math.Max(float64(alpha), float64(u)))
			}
			if tmp_u > u {
				u = tmp_u
			}
		}
	}
	return u
}

// // 2nd player
func min_player(ctx s.SContext, alpha int32, beta int32, explor int, new_vertex *s.SVertex) int32 {
	// fmt.Println("alpha beta min", alpha, beta)
	playerMin := s.Tnumber(0)
	if ctx.CurrentPlayer == 1 {
		playerMin = s.Tnumber(2)
	} else {
		playerMin = s.Tnumber(1)
	}
	u := int32(2147483647)
	goban := CopyGoban(ctx)
	casesnonnull := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(casesnonnull, ctx.CasesNonNull)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  casesnonnull,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize}
	if explor == 0 || g.VictoryConditionAlign(&ctx, int(new_vertex.X), int(new_vertex.Y), nil) == true {
		tmp_heuris := Heuristic(tmp_ctx, int(new_vertex.X), int(new_vertex.Y))
		// fmt.Println("Ret min", tmp_heuris, explor)
		return int32(-tmp_heuris)
	}
	if new_vertex != nil {
		FindNeighbors(&tmp_ctx, int(new_vertex.X), int(new_vertex.Y))
	}
	for _, neighbor := range tmp_ctx.CasesNonNull {
		placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(playerMin)
			tmp_u := max_player(tmp_ctx, alpha, beta, explor-1, &neighbor)
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(0)
			if tmp_u <= alpha {
				fmt.Println("Prunning alpha")
				return tmp_u
			} else {
				beta = int32(math.Min(float64(beta), float64(u)))
			}
			if tmp_u < u {
				u = tmp_u
			}
		}
	}
	return u
}
