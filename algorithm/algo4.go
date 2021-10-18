package algorithm

import (
	"fmt"
	s "gomoku/structures"

	"math"
)

func CopyGoban2(ctx s.SContext) s.Tgoban {
	newGoban := make([][]s.Tnumber, ctx.NSize)
	for Y, line := range ctx.Goban {
		newGoban[Y] = make([]s.Tnumber, ctx.NSize)
		for X, nb := range line {
			newGoban[Y][X] = nb
		}
	}
	return newGoban
}

// func CopyCases(ctx s.SContext) map[s.SVertex][]s.SVertex {
// 	cases := make(map[s.SVertex][]s.SVertex)
// 	for key, elem := range ctx.CasesNonNull {
// 		cases[key] = elem
// 	}
// 	return cases
// }

func AlphaBetaPruning2(ctx s.SContext, depth int8) (s.SVertex, int32) {
	goban := CopyGoban2(ctx)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  nil,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize}

	neighbors := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(neighbors, ctx.CasesNonNull)

	alpha := int32(-2147483648)
	beta := int32(2147483647)

	vertex := s.SVertex{X: -1, Y: -1}
	maxEval := int32(-2147483648)

	for _, neighbor := range neighbors {
		fmt.Println(neighbor)
		placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			newNeighbors := getNeighbors(tmp_ctx, neighbor) // fct getNeighbors a faire
			eval := minimax(tmp_ctx, newNeighbors, depth-1, alpha, beta, false)
			tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0

			if eval >= maxEval {
				maxEval = eval
				vertex = neighbor
			}
		}
	}

	return vertex, maxEval
}

func Heuristic2(ctx s.SContext, isMaximazingPlayer bool) int32 {
	value := 0
	player := s.Tnumber(2)

	if isMaximazingPlayer == true {
		player = 1
	}
	for x := range ctx.Goban {
		for y := range ctx.Goban[x] {
			if ctx.Goban[y][x] == player {
				value += Heuristic(ctx, x, y)
			}
		}
	}

	return int32(value)
}

func minimax(tmp_ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32, isMaximazingPlayer bool) int32 {
	if depth == 0 /* || gameIsOver(ctx.Goban)) */ {
		return Heuristic2(tmp_ctx, isMaximazingPlayer)
	}

	if isMaximazingPlayer {
		maxEval := int32(-2147483648)

		for _, neighbor := range neighbors {
			// fmt.Println("max", neighbor)
			placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(tmp_ctx.CurrentPlayer)
				newNeighbors := getNeighbors(tmp_ctx, neighbor) // fct getNeighbors a faire
				eval := minimax(tmp_ctx, newNeighbors, depth-1, alpha, beta, false)
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0
				maxEval = int32(math.Max(float64(maxEval), float64(eval)))
				alpha = int32(math.Max(float64(alpha), float64(eval)))

				if beta <= alpha {
					fmt.Println("Prunning")
					break
				}

			}
		}
		return maxEval
	} else {
		minEval := int32(2147483647)

		playerMin := s.Tnumber(0)
		if tmp_ctx.CurrentPlayer == 1 {
			playerMin = s.Tnumber(2)
		} else {
			playerMin = s.Tnumber(1)
		}

		for _, neighbor := range neighbors {
			// fmt.Println("min", neighbor)
			placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(playerMin)
				newNeighbors := getNeighbors(tmp_ctx, neighbor) // fct getNeighbors a faire
				eval := minimax(tmp_ctx, newNeighbors, depth-1, alpha, beta, true)
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0
				minEval = int32(math.Min(float64(minEval), float64(eval)))
				beta = int32(math.Min(float64(beta), float64(eval)))
				if beta <= alpha {
					fmt.Println("Prunning")
					break
				}

			}
		}
		return minEval
	}
}
