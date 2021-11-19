package algogo

import (
	s "gomoku/structures"
)

func insert(a []playData, index int, value playData) []playData {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func insertNeighborsMax(plays []playData, play playData) []playData {

	size := len(plays)

	for i := 0; i < size; i++ {
		if plays[i].Heur <= play.Heur {
			return insert(plays, i, play)
		}
	}

	return append(plays, play)
}

func insertNeighborsMin(plays []playData, play playData) []playData {

	size := len(plays)

	for i := 0; i < size; i++ {
		if plays[i].Heur >= play.Heur {
			return insert(plays, i, play)
		}
	}

	return append(plays, play)
}

func sortNeighbors(ctx s.SContext, neighbors []s.SVertex, isMaximizing bool) []s.SVertex {
	var ret []s.SVertex
	var plays []playData

	if isMaximizing {
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx.Goban, neighbor.X, neighbor.Y, ctx.CurrentPlayer)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
				heur := EvaluateGoban(ctx)
				plays = insertNeighborsMax(plays, playData{Heur: heur, Vertex: neighbor})
				ctx.Goban[neighbor.Y][neighbor.X] = 0

			}
		}
	} else {
		for _, neighbor := range neighbors {
			placement := PlacementHeuristic(ctx.Goban, neighbor.X, neighbor.Y, ctx.CurrentPlayer)
			if placement >= 1 {
				ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
				heur := EvaluateGoban(ctx)
				plays = insertNeighborsMin(plays, playData{Heur: heur, Vertex: neighbor})
				ctx.Goban[neighbor.Y][neighbor.X] = 0

			}
		}
	}
	for _, play := range plays {
		ret = append(ret, play.Vertex)
	}
	return ret
}
