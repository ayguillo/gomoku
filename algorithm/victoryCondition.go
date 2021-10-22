package algorithm

import (
	s "gomoku/structures"
)

func checkCaptureVictory(ctx s.SContext) bool {
	return false
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
