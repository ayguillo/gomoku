package algorithm

import (
	g "gomoku/game"
	s "gomoku/structures"

	"math"
	"math/rand"
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

func Minimax(ctx s.SContext, depth int8, maximizingPlayer bool, child *s.SVertex, alpha int32, beta int32, max_depth int8, hit *s.SVertex) (int32, s.SVertex) {
	var next_player uint8
	var current_player uint8

	if depth == 0 {
		heuris := rand.Intn(10000)
		return int32(heuris), *hit
	}
	if child != nil {
		if g.VictoryConditionAlign(&ctx, int(child.X), int(child.Y), nil) || g.VictoryCapture(ctx) {
			return 1000000, *hit
		}
	}
	node := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(node, ctx.CasesNonNull)
	goban := CopyGoban(ctx)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  node,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize}
	if child != nil {
		FindNeighbors(&tmp_ctx, int(child.X), int(child.Y))
	}
	if tmp_ctx.CurrentPlayer == 1 {
		next_player = 2
		current_player = 1
	} else {
		next_player = 1
		current_player = 2
	}
	// fmt.Println(tmp_ctx.CasesNonNull, child, maximizingPlayer)
	if maximizingPlayer {
		value := int32(-2147483648)
		for _, neighbor := range tmp_ctx.CasesNonNull {
			if max_depth == depth {
				hit = &neighbor
			}
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(tmp_ctx.CurrentPlayer)
			// fmt.Println(tmp_ctx)
			capture, vertex_cap := g.Capture(&tmp_ctx, nil, int(neighbor.X), int(neighbor.Y), false)
			tmp_ctx.CurrentPlayer = next_player
			tmp_value, tmp_hit := Minimax(tmp_ctx, depth-1, false, &neighbor, alpha, beta, max_depth, hit)
			tmp_ctx.CurrentPlayer = current_player
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(0)
			if capture == true {
				if tmp_ctx.CurrentPlayer == 1 {
					tmp_ctx.NbCaptureP1--
				} else {
					tmp_ctx.NbCaptureP2--
				}
				for _, vertex := range vertex_cap {
					tmp_ctx.Goban[vertex.Y][vertex.X] = s.Tnumber(next_player)
				}
			}
			if tmp_value >= beta {
				return tmp_value, tmp_hit
			}
			if tmp_value > value {
				value = tmp_value
				hit = &tmp_hit
			}
			alpha = int32(math.Max(float64(alpha), float64(value)))
		}
		return value, *hit
	} else {
		value := int32(2147483647)
		for _, neighbor := range tmp_ctx.CasesNonNull {
			if max_depth == depth {
				hit = &neighbor
			}
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(tmp_ctx.CurrentPlayer)
			// fmt.Println(tmp_ctx)
			capture, vertex_cap := g.Capture(&tmp_ctx, nil, int(neighbor.X), int(neighbor.Y), false)
			tmp_ctx.CurrentPlayer = next_player
			tmp_value, tmp_hit := Minimax(tmp_ctx, depth-1, true, &neighbor, alpha, beta, max_depth, hit)
			tmp_ctx.CurrentPlayer = current_player
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(0)
			if capture == true {
				if tmp_ctx.CurrentPlayer == 1 {
					tmp_ctx.NbCaptureP1--
				} else {
					tmp_ctx.NbCaptureP2--
				}
				for _, vertex := range vertex_cap {
					tmp_ctx.Goban[vertex.Y][vertex.X] = s.Tnumber(next_player)
				}
			}
			if alpha >= tmp_value {
				return tmp_value, tmp_hit
			}
			if tmp_value < value {
				value = tmp_value
				hit = &tmp_hit
			}
			beta = int32(math.Min(float64(beta), float64(value)))
		}
		return value, *hit
	}
}
