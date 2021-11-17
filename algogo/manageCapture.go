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

func checkDiagLeftUpCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x-1, case_y-1) && isInRange(case_x-2, case_y-2) && isInRange(case_x-3, case_y-3) && goban[case_y-3][case_x-3] == player && goban[case_y-2][case_x-2] == opp && goban[case_y-1][case_x-1] == opp {
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

func checkDiagRightUpCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x+1, case_y-1) && isInRange(case_x+2, case_y-2) && isInRange(case_x+3, case_y-3) && goban[case_y-3][case_x+3] == player && goban[case_y-2][case_x+2] == opp && goban[case_y-1][case_x+1] == opp {
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

func checkUpCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x, case_y-1) && isInRange(case_x, case_y-2) && isInRange(case_x, case_y-3) && goban[case_y-3][case_x] == player && goban[case_y-2][case_x] == opp && goban[case_y-1][case_x] == opp {
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

func checkDownCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x, case_y+1) && isInRange(case_x, case_y+2) && isInRange(case_x, case_y+3) && goban[case_y+3][case_x] == player && goban[case_y+2][case_x] == opp && goban[case_y+1][case_x] == opp {
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

func checkLeftCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x-1, case_y) && isInRange(case_x-2, case_y) && isInRange(case_x-3, case_y) && goban[case_y][case_x-3] == player && goban[case_y][case_x-2] == opp && goban[case_y][case_x-1] == opp {
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

func checkRightCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x+1, case_y) && isInRange(case_x+2, case_y) && isInRange(case_x+3, case_y) && goban[case_y][case_x+3] == player && goban[case_y][case_x+2] == opp && goban[case_y][case_x+1] == opp {
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

func checkDiagLeftDownCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x-1, case_y+1) && isInRange(case_x-2, case_y+2) && isInRange(case_x-3, case_y+3) && goban[case_y+3][case_x-3] == player && goban[case_y+2][case_x-2] == opp && goban[case_y+1][case_x-1] == opp {
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

func checkDiagRightDownCapture(goban s.Tgoban, case_x int, case_y int, player s.Tnumber, opp s.Tnumber) []s.SVertex {
	var ret []s.SVertex = nil

	if isInRange(case_x+1, case_y+1) && isInRange(case_x+2, case_y+2) && isInRange(case_x+3, case_y+3) && goban[case_y+3][case_x+3] == player && goban[case_y+2][case_x+2] == opp && goban[case_y+1][case_x+1] == opp {
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

func CaptureAlgoNode(node *node, case_x int, case_y int) []s.SVertex {
	var ret []s.SVertex = nil
	var opp uint8 = 1

	if node.player == 1 {
		opp = 2
	}

	leftUp := checkDiagLeftUpCapture(node.goban, case_x, case_y, s.Tnumber(node.player), s.Tnumber(opp))
	rightUp := checkDiagRightUpCapture(node.goban, case_x, case_y, s.Tnumber(node.player), s.Tnumber(opp))
	up := checkUpCapture(node.goban, case_x, case_y, s.Tnumber(node.player), s.Tnumber(opp))
	down := checkDownCapture(node.goban, case_x, case_y, s.Tnumber(node.player), s.Tnumber(opp))
	left := checkLeftCapture(node.goban, case_x, case_y, s.Tnumber(node.player), s.Tnumber(opp))
	right := checkRightCapture(node.goban, case_x, case_y, s.Tnumber(node.player), s.Tnumber(opp))
	leftDown := checkDiagLeftDownCapture(node.goban, case_x, case_y, s.Tnumber(node.player), s.Tnumber(opp))
	rightDown := checkDiagRightDownCapture(node.goban, case_x, case_y, s.Tnumber(node.player), s.Tnumber(opp))

	if leftUp != nil {
		if node.player == 1 {
			node.captures.Capture0++
		} else {
			node.captures.Capture1++
		}

		ret = append(ret, leftUp...)
	}

	if rightUp != nil {
		if node.player == 1 {
			node.captures.Capture0++
		} else {
			node.captures.Capture1++
		}

		ret = append(ret, rightUp...)
	}

	if up != nil {
		if node.player == 1 {
			node.captures.Capture0++
		} else {
			node.captures.Capture1++
		}

		ret = append(ret, up...)
	}

	if down != nil {
		if node.player == 1 {
			node.captures.Capture0++
		} else {
			node.captures.Capture1++
		}

		ret = append(ret, down...)
	}

	if left != nil {
		if node.player == 1 {
			node.captures.Capture0++
		} else {
			node.captures.Capture1++
		}

		ret = append(ret, left...)
	}

	if right != nil {
		if node.player == 1 {
			node.captures.Capture0++
		} else {
			node.captures.Capture1++
		}

		ret = append(ret, right...)
	}

	if leftDown != nil {
		if node.player == 1 {
			node.captures.Capture0++
		} else {
			node.captures.Capture1++
		}

		ret = append(ret, leftDown...)
	}

	if rightDown != nil {
		if node.player == 1 {
			node.captures.Capture0++
		} else {
			node.captures.Capture1++
		}

		ret = append(ret, rightDown...)
	}

	return ret
}

func CaptureAlgoCtx(ctx *s.SContext, case_x int, case_y int) []s.SVertex {
	var ret []s.SVertex = nil
	var opp uint8 = 1

	if ctx.CurrentPlayer == 1 {
		opp = 2
	}

	leftUp := checkDiagLeftUpCapture(ctx.Goban, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	rightUp := checkDiagRightUpCapture(ctx.Goban, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	up := checkUpCapture(ctx.Goban, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	down := checkDownCapture(ctx.Goban, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	left := checkLeftCapture(ctx.Goban, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	right := checkRightCapture(ctx.Goban, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	leftDown := checkDiagLeftDownCapture(ctx.Goban, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))
	rightDown := checkDiagRightDownCapture(ctx.Goban, case_x, case_y, s.Tnumber(ctx.CurrentPlayer), s.Tnumber(opp))

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
