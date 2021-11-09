package algorithm

import (
	g "gomoku/game"
	h "gomoku/heuristic"
	s "gomoku/structures"
)

type playData struct {
	Vertex s.SVertex
	Heur   int32
}

type stockData2 struct {
	Heur   int32
	Vertex s.SVertex
	Depth  int8
	Alpha  int32
	Beta   int32
}

type stockData struct {
	Heur   int32
	Vertex s.SVertex
	Depth  int8
}

type heurData struct {
	Play stockData
	Heur int32
}

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

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func min(a int32, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func swapPlayer(ctx *s.SContext) {
	if ctx.CurrentPlayer == 1 {
		ctx.CurrentPlayer = 2
	} else {
		ctx.CurrentPlayer = 1
	}
}

func insert(a []playData, index int, value playData) []playData {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func insertNeighbors(plays []playData, play playData) []playData {

	size := len(plays)

	for i := 0; i < size; i++ {
		if plays[i].Heur <= play.Heur {
			return insert(plays, i, play)
		}
	}

	return append(plays, play)
}

func sortNeighbors(ctx s.SContext, neighbors []s.SVertex) []s.SVertex {
	var ret []s.SVertex
	var plays []playData

	for _, neighbor := range neighbors {
		placement := PlacementHeuristic(ctx, neighbor.X, neighbor.Y)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			heur := h.CalcHeuristic(ctx)
			plays = insertNeighbors(plays, playData{Heur: heur, Vertex: neighbor})

			ctx.Goban[neighbor.Y][neighbor.X] = 0
		}
	}

	for _, play := range plays {
		ret = append(ret, play.Vertex)
	}

	return ret
}

func PlacementHeuristic(ctx s.SContext, case_x int, case_y int) uint8 {
	// Retour 2 pour une obligation de capture, 1 pour un coup ok, 0 sinon
	capture := false
	if ctx.ActiveCapture && len(ctx.Capture) > 0 {
		for _, cap := range ctx.Capture {
			if case_x == cap.X && case_y == cap.Y {
				capture = true
				break
			}
		}
		if capture == false {
			return 2
		}
	}
	if ctx.ActiveDoubleThrees > 0 && !g.CheckDoubleThree(&ctx, case_x, case_y) {
		return 0
	}
	if case_y < 0 || case_y > int(ctx.NSize) {
		return 0
	}
	if case_x < 0 || case_x > int(ctx.NSize) {
		return 0
	}
	if ctx.Goban[int(case_y)][int(case_x)] == 0 {
		return 1
	} else {
		return 0
	}
}

func revertCapture(ctx *s.SContext, captureVertex []s.SVertex, captureP1 int, captureP2 int, player uint8) {
	ctx.NbCaptureP1 = captureP1
	ctx.NbCaptureP2 = captureP2

	swapPlayer := 1

	if player == 1 {
		swapPlayer = 2
	}

	for _, vertex := range captureVertex {
		ctx.Goban[vertex.Y][vertex.X] = s.Tnumber(swapPlayer)
	}
}
