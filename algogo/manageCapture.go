package algogo

import (
	s "gomoku/structures"
)

func isInRange(x int, y int) bool {
	if x < 0 || x >= 19 || y < 0 || y >= 19 {
		return false
	}

	return true
}

func checkDiagLeftUpCapture(ctx *s.SContext, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x-1, case_y-1) && isInRange(case_x-2, case_y-2) && isInRange(case_x-3, case_y-3) && ctx.Goban[case_y-3][case_x-3] == player && ctx.Goban[case_y-2][case_x-2] == opp && ctx.Goban[case_y-1][case_x-1] == opp {
		vertex1 := s.SVertex{
			X: case_x - 1,
			Y: case_y - 1,
		}
		vertex2 := s.SVertex{
			X: case_x - 2,
			Y: case_y - 2,
		}

		ret = append(ret, vertex1, vertex2)
	}

	return ret
}

func checkDiagRightUpCapture(ctx *s.SContext, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x+1, case_y-1) && isInRange(case_x+2, case_y-2) && isInRange(case_x+3, case_y-3) && ctx.Goban[case_y-3][case_x+3] == player && ctx.Goban[case_y-2][case_x+2] == opp && ctx.Goban[case_y-1][case_x+1] == opp {
		vertex1 := s.SVertex{
			X: case_x + 2,
			Y: case_y - 2,
		}
		vertex2 := s.SVertex{
			X: case_x + 1,
			Y: case_y - 1,
		}

		ret = append(ret, vertex1, vertex2)
	}

	return ret
}

func checkUpCapture(ctx *s.SContext, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x, case_y-1) && isInRange(case_x, case_y-2) && isInRange(case_x, case_y-3) && ctx.Goban[case_y-3][case_x] == player && ctx.Goban[case_y-2][case_x] == opp && ctx.Goban[case_y-1][case_x] == opp {
		vertex1 := s.SVertex{
			X: case_x,
			Y: case_y - 2,
		}
		vertex2 := s.SVertex{
			X: case_x,
			Y: case_y - 1,
		}

		ret = append(ret, vertex1, vertex2)
	}

	return ret
}

func checkDownCapture(ctx *s.SContext, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x, case_y+1) && isInRange(case_x, case_y+2) && isInRange(case_x, case_y+3) && ctx.Goban[case_y+3][case_x] == player && ctx.Goban[case_y+2][case_x] == opp && ctx.Goban[case_y+1][case_x] == opp {
		vertex1 := s.SVertex{
			X: case_x,
			Y: case_y + 2,
		}
		vertex2 := s.SVertex{
			X: case_x,
			Y: case_y + 1,
		}

		ret = append(ret, vertex1, vertex2)
	}

	return ret
}

func checkLeftCapture(ctx *s.SContext, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x-1, case_y) && isInRange(case_x-2, case_y) && isInRange(case_x-3, case_y) && ctx.Goban[case_y][case_x-3] == player && ctx.Goban[case_y][case_x-2] == opp && ctx.Goban[case_y][case_x-1] == opp {
		vertex1 := s.SVertex{
			X: case_x - 1,
			Y: case_y,
		}
		vertex2 := s.SVertex{
			X: case_x - 2,
			Y: case_y,
		}

		ret = append(ret, vertex1, vertex2)
	}

	return ret
}

func checkRightCapture(ctx *s.SContext, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x+1, case_y) && isInRange(case_x+2, case_y) && isInRange(case_x+3, case_y) && ctx.Goban[case_y][case_x+3] == player && ctx.Goban[case_y][case_x+2] == opp && ctx.Goban[case_y][case_x+1] == opp {
		vertex1 := s.SVertex{
			X: case_x + 1,
			Y: case_y,
		}
		vertex2 := s.SVertex{
			X: case_x + 2,
			Y: case_y,
		}

		ret = append(ret, vertex1, vertex2)
	}

	return ret
}

func checkDiagLeftDownCapture(ctx *s.SContext, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x-1, case_y+1) && isInRange(case_x-2, case_y+2) && isInRange(case_x-3, case_y+3) && ctx.Goban[case_y+3][case_x-3] == player && ctx.Goban[case_y+2][case_x-2] == opp && ctx.Goban[case_y+1][case_x-1] == opp {
		vertex1 := s.SVertex{
			X: case_x - 1,
			Y: case_y + 1,
		}
		vertex2 := s.SVertex{
			X: case_x - 2,
			Y: case_y + 2,
		}

		ret = append(ret, vertex1, vertex2)
	}

	return ret
}

func checkDiagRightDownCapture(ctx *s.SContext, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x+1, case_y+1) && isInRange(case_x+2, case_y+2) && isInRange(case_x+3, case_y+3) && ctx.Goban[case_y+3][case_x+3] == player && ctx.Goban[case_y+2][case_x+2] == opp && ctx.Goban[case_y+1][case_x+1] == opp {
		vertex1 := s.SVertex{
			X: case_x + 1,
			Y: case_y + 1,
		}
		vertex2 := s.SVertex{
			X: case_x + 2,
			Y: case_y + 2,
		}

		ret = append(ret, vertex1, vertex2)
	}

	return ret
}

func CaptureAlgo(ctx *s.SContext, case_x int, case_y int) []s.SVertex {
	var ret []s.SVertex = nil
	var opp uint8 = 1

	if ctx.CurrentPlayer == 1 {
		opp = 2
	}

	leftUp := checkDiagLeftUpCapture(ctx, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	rightUp := checkDiagRightUpCapture(ctx, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	up := checkUpCapture(ctx, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	down := checkDownCapture(ctx, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	left := checkLeftCapture(ctx, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	right := checkRightCapture(ctx, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	leftDown := checkDiagLeftDownCapture(ctx, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	rightDown := checkDiagRightDownCapture(ctx, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))

	if leftUp != nil {
		if ctx.CurrentPlayer == 1 {
			ctx.NbCaptureP1++
		} else {
			ctx.NbCaptureP2++
		}

		ret = append(ret, leftUp...)
	}

	if rightUp != nil {
		if ctx.CurrentPlayer == 1 {
			ctx.NbCaptureP1++
		} else {
			ctx.NbCaptureP2++
		}

		ret = append(ret, rightUp...)
	}

	if up != nil {
		if ctx.CurrentPlayer == 1 {
			ctx.NbCaptureP1++
		} else {
			ctx.NbCaptureP2++
		}

		ret = append(ret, up...)
	}

	if down != nil {
		if ctx.CurrentPlayer == 1 {
			ctx.NbCaptureP1++
		} else {
			ctx.NbCaptureP2++
		}

		ret = append(ret, down...)
	}

	if left != nil {
		if ctx.CurrentPlayer == 1 {
			ctx.NbCaptureP1++
		} else {
			ctx.NbCaptureP2++
		}

		ret = append(ret, left...)
	}

	if right != nil {
		if ctx.CurrentPlayer == 1 {
			ctx.NbCaptureP1++
		} else {
			ctx.NbCaptureP2++
		}

		ret = append(ret, right...)
	}

	if leftDown != nil {
		if ctx.CurrentPlayer == 1 {
			ctx.NbCaptureP1++
		} else {
			ctx.NbCaptureP2++
		}

		ret = append(ret, leftDown...)
	}

	if rightDown != nil {
		if ctx.CurrentPlayer == 1 {
			ctx.NbCaptureP1++
		} else {
			ctx.NbCaptureP2++
		}

		ret = append(ret, rightDown...)
	}

	return ret
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
