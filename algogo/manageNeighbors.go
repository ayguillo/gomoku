package algogo

import s "gomoku/structures"

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
		placement := PlacementHeuristic(ctx.Goban, neighbor.X, neighbor.Y, ctx.CurrentPlayer)
		if placement >= 1 {
			ctx.Goban[neighbor.Y][neighbor.X] = s.Tnumber(ctx.CurrentPlayer)
			heur := EvaluateGoban(ctx)
			plays = insertNeighbors(plays, playData{Heur: heur, Vertex: neighbor})

			ctx.Goban[neighbor.Y][neighbor.X] = 0
		}
	}

	for _, play := range plays {
		ret = append(ret, play.Vertex)
	}

	return ret
}
