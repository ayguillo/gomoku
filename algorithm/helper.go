package algorithm

import s "gomoku/structures"

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

type stockData struct {
	Heur   int32
	Vertex s.SVertex
	Depth  int8
}

type heurData struct {
	Play stockData
	Heur int32
}
