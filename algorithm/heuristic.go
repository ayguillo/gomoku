package algorithm

import (
	g "gomoku/game"
	s "gomoku/structures"
)

func conditionAlignHeuristic(ctx s.SContext, case_x int, case_y int) int {
	capturePlayer, nbCapture := 0, 0
	if ctx.CurrentPlayer == 1 {
		capturePlayer = 2
		nbCapture = ctx.NbCaptureP2
	} else {
		capturePlayer = 1
		nbCapture = ctx.NbCaptureP1
	}
	horizontal := horizontalHeuristicAlign(ctx, case_x, case_y, capturePlayer, nbCapture)
	vertical := verticalHeuristicAlign(ctx, case_x, case_y, capturePlayer, nbCapture)
	diagonal := diagonalHeuristicAlign(ctx, case_x, case_y, capturePlayer, nbCapture)
	return horizontal + vertical + diagonal
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
	if ctx.ActiveDoubleThrees && !g.CheckDoubleThree(&ctx, case_x, case_y) {
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

func captureHeuristic(ctx s.SContext, case_x int, case_y int) int {
	var capture uint8
	score_capture := 0
	if ctx.CurrentPlayer == 1 {
		capture = 2
		score_capture = 10 * ctx.NbCaptureP1
	} else {
		capture = 1
		score_capture = 10 * ctx.NbCaptureP2
	}
	horizontal := horizontalCaptureHeuristic(ctx, case_x, case_y, capture, score_capture)
	vertical := verticalCaptureHeuristic(ctx, case_x, case_y, capture, score_capture)
	diagonal := diagonalCaptureHeuristic(ctx, case_x, case_y, capture, score_capture)
	return horizontal + vertical + diagonal
}

func blockHeuristic(ctx s.SContext, case_x int, case_y int) int {
	var opponent s.Tnumber
	if ctx.CurrentPlayer == 1 {
		opponent = 2
	} else {
		opponent = 1
	}
	horizontal := blockHorizontal(ctx, case_x, case_y, opponent)
	vertical := blockVertical(ctx, case_x, case_y, opponent)
	diagLeft := blockDiagLeft(ctx, case_x, case_y, opponent)
	diagRight := blockDiagRight(ctx, case_x, case_y, opponent)
	return horizontal + vertical + diagLeft + diagRight
}

func Heuristic(ctx s.SContext, case_x int, case_y int) int {
	align := conditionAlignHeuristic(ctx, case_x, case_y)
	capture := captureHeuristic(ctx, case_x, case_y)
	block := blockHeuristic(ctx, case_x, case_y)
	// fmt.Println("align", align, "capture", capture, "block", block, "total", align+capture+block)
	return align + capture + block
}
