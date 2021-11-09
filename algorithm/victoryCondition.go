package algorithm

import (
	s "gomoku/structures"
)

func checkCaptureVictory(ctx s.SContext) (bool, s.Tnumber) {
	if ctx.NbCaptureP1 >= 5 {
		return true, 1
	} else if ctx.NbCaptureP2 >= 5 {
		return true, 2
	}
	return false, 0
}

func checkVictory(ctx s.SContext, player s.Tnumber, x int, y int, i int, j int) int {
	count := 0
	for k := 0; k < 5 && y >= 0 && y < int(ctx.NSize) && x >= 0 && x < int(ctx.NSize) && ctx.Goban[y][x] == player; k++ {
		x += i
		y += j
		count += 1
	}

	return count
}

func checkHorizontalVictory(ctx s.SContext, x int, y int) bool {
	if checkVictory(ctx, ctx.Goban[y][x], x, y, 1, 0)+checkVictory(ctx, ctx.Goban[y][x], x, y, -1, 0) > 5 {
		return true
	}

	return false
}

func checkVerticalVictory(ctx s.SContext, x int, y int) bool {
	if checkVictory(ctx, ctx.Goban[y][x], x, y, 0, 1)+checkVictory(ctx, ctx.Goban[y][x], x, y, 0, -1) > 5 {
		return true
	}

	return false
}

func checkDiagVictory(ctx s.SContext, x int, y int) bool {
	if checkVictory(ctx, ctx.Goban[y][x], x, y, 1, 1)+checkVictory(ctx, ctx.Goban[y][x], x, y, -1, -1) > 5 || checkVictory(ctx, ctx.Goban[y][x], x, y, -1, 1)+checkVictory(ctx, ctx.Goban[y][x], x, y, 1, -1) > 5 {
		return true
	}

	return false
}

func CheckAlignVictory(ctx s.SContext, x int, y int) bool {
	if checkHorizontalVictory(ctx, x, y) || checkVerticalVictory(ctx, x, y) || checkDiagVictory(ctx, x, y) {
		return true
	}
	return false
}

func VictoryCondition(ctx s.SContext) (bool, s.Tnumber) {
	if ctx.ActiveCapture {
		check, player := checkCaptureVictory(ctx)

		if check {
			return true, player
		}
	}

	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] != 0 {
				if CheckAlignVictory(ctx, x, y) {
					return true, ctx.Goban[y][x]
				}
			}
		}
	}

	return false, 0
}
