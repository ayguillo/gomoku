package algorithm

import (
	"fmt"
	g "gomoku/game"
	h "gomoku/heuristic"
	s "gomoku/structures"

	"math"
)

var depthData []int8
var depthStart int8

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

func swapPlayer(ctx *s.SContext) {
	if ctx.CurrentPlayer == 1 {
		ctx.CurrentPlayer = 2
	} else {
		ctx.CurrentPlayer = 1
	}
}

type stockData struct {
	Heur   int32
	Vertex s.SVertex
	Depth  int8
}

func AlphaBetaPruning(ctx s.SContext, depth int8) (s.SVertex, int32) {
	neighbors := make([]s.SVertex, len(ctx.CasesNonNull))
	copy(neighbors, ctx.CasesNonNull)

	alpha := int32(-2147483648)
	beta := int32(2147483647)

	vertex := s.SVertex{X: -1, Y: -1}
	maxEval := int32(-2147483648)

	ch := make(chan stockData)
	i := len(neighbors)
	depthData = make([]int8, i)
	depthStart = depth

	k := 0
	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			go initMax(ctx, depth, neighbor, alpha, beta, ch, k)
			k++
		}
	}

	var data []stockData

	for i > 0 {
		ret := <-ch
		data = append(data, ret)
		i--
	}

	fmt.Printf("%v\n", data)
	close(ch)

	for _, value := range data {
		eval := value.Heur
		if eval >= maxEval {
			maxEval = eval
			vertex = value.Vertex
		}
	}

	fmt.Printf("%v %v\n", vertex, maxEval)

	return vertex, maxEval
}

func initMax(ctx s.SContext, depth int8, neighbor s.SVertex, alpha int32, beta int32, ch chan stockData, i int) {
	goban := CopyGoban(ctx)
	tmp_ctx := s.SContext{
		Goban:         goban,
		CurrentPlayer: ctx.CurrentPlayer,
		CasesNonNull:  nil,
		Capture:       ctx.Capture,
		NbCaptureP1:   ctx.NbCaptureP1,
		NbCaptureP2:   ctx.NbCaptureP2,
		NSize:         ctx.NSize}

	tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
	newNeighbors := getNeighbors(tmp_ctx, neighbor) // fct getNeighbors a faire
	tmp := tmp_ctx.CurrentPlayer
	swapPlayer(&tmp_ctx)
	eval := minimax(tmp_ctx, newNeighbors, depth-1, alpha, beta, false, i)
	tmp_ctx.CurrentPlayer = tmp
	tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0

	ret := stockData{
		Heur:   eval,
		Vertex: neighbor,
		Depth:  depthStart - depthData[i],
	}

	ch <- ret
}

func Heuristic2(ctx s.SContext, isMaximazingPlayer bool) int32 {
	value := 0

	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
				value += Heuristic(ctx, x, y)
			}
		}
	}

	return int32(value)
}

func victoryCondition(ctx s.SContext) bool {
	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] != 0 {
				if g.VictoryConditionAlign(&ctx, x, y, nil) {
					println("true")
					return true
				}
			}
		}
	}

	return false
}

func minimax(tmp_ctx s.SContext, neighbors []s.SVertex, depth int8, alpha int32, beta int32, isMaximazingPlayer bool, i int) int32 {
	if depth == 0 || victoryCondition(tmp_ctx) {
		depthData[i] = depth
		return h.CalcHeuristic(tmp_ctx)
	}

	if isMaximazingPlayer {
		maxEval := int32(-2147483648)

		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(tmp_ctx.CurrentPlayer)
				newNeighbors := getNeighbors(tmp_ctx, neighbor)

				tmp := tmp_ctx.CurrentPlayer
				swapPlayer(&tmp_ctx)

				eval := minimax(tmp_ctx, newNeighbors, depth-1, alpha, beta, false, i)

				tmp_ctx.CurrentPlayer = tmp
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0

				maxEval = int32(math.Max(float64(maxEval), float64(eval)))
				alpha = int32(math.Max(float64(alpha), float64(maxEval)))

				if alpha >= beta {
					break
				}
			}
		}
		return maxEval
	} else {
		minEval := int32(2147483647)

		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(tmp_ctx, neighbor.X, neighbor.Y)
			if placement >= 1 {
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(tmp_ctx.CurrentPlayer)
				newNeighbors := getNeighbors(tmp_ctx, neighbor)

				tmp := tmp_ctx.CurrentPlayer
				swapPlayer(&tmp_ctx)

				eval := minimax(tmp_ctx, newNeighbors, depth-1, alpha, beta, true, i)

				tmp_ctx.CurrentPlayer = tmp
				tmp_ctx.Goban[neighbor.Y][neighbor.X] = 0

				minEval = int32(math.Min(float64(minEval), float64(eval)))
				beta = int32(math.Min(float64(beta), float64(minEval)))

				if alpha >= beta {
					break
				}
			}
		}
		return minEval
	}
}
