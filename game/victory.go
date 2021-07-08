package game

import (
	d "gomoku/display"
	s "gomoku/structures"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

func horizontalAlign(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) bool {
	count_stone := 0
	var message string
	if CounterVertical(*ctx, case_x, case_y, capturePlayer) == true || CounterDiag(*ctx, case_x, case_y, capturePlayer) == true {
		message = "Capture in " + ctx.MapX[case_x+1] + " " + strconv.Itoa(int(ctx.NSize)-case_y)
		ctx.Capture = s.SVertex{X: case_x, Y: case_y}
	}
	for current_case := case_x + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(*ctx, current_case, case_y, capturePlayer) == true || CounterDiag(*ctx, current_case, case_y, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[current_case+1] + " " + strconv.Itoa(int(ctx.NSize)-case_y)
			}
			count_stone++
		} else {
			break
		}
	}
	for current_case := case_x - 1; current_case >= 0; current_case-- {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(*ctx, current_case, case_y, capturePlayer) == true || CounterDiag(*ctx, current_case, case_y, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[current_case+1] + " " + strconv.Itoa(int(ctx.NSize)-case_y)
			}
			count_stone++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" {
			sdl.Log(message)
			d.DisplayMessage(visu, int32((int32(ctx.NSize+1))*ctx.SizeCase), message, "", *ctx)
			return false
		}
		return true
	}
	return false
}

func verticalAlign(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) bool {
	count_stone := 0
	var message string
	if CounterHorizontal(*ctx, case_x, case_y, capturePlayer) == true || CounterDiag(*ctx, case_x, case_y, capturePlayer) == true {
		message = "Capture in " + ctx.MapX[case_x] + " " + strconv.Itoa(int(ctx.NSize)-case_y)
	}
	for current_case := case_y + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterHorizontal(*ctx, case_x, current_case, capturePlayer) == true || CounterDiag(*ctx, case_x, current_case, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[case_x+1] + " " + strconv.Itoa(int(ctx.NSize)-current_case)
			}
			count_stone++
		} else {
			break
		}
	}
	for current_case := case_y - 1; current_case >= 0; current_case-- {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterHorizontal(*ctx, case_x, current_case, capturePlayer) == true || CounterDiag(*ctx, case_x, current_case, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[case_x+1] + " " + strconv.Itoa(int(ctx.NSize)-current_case)
			}
			count_stone++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" {
			sdl.Log(message)
			d.DisplayMessage(visu, int32((int32(ctx.NSize+1))*ctx.SizeCase), message, "", *ctx)
			return false
		}
		return true
	}
	return false
}

func diagLeft(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) bool {
	count_stone := 0
	var message string
	if CounterVertical(*ctx, case_x, case_y, capturePlayer) == true || CounterHorizontal(*ctx, case_x, case_y, capturePlayer) ||
		CounterDiag(*ctx, case_x, case_y, capturePlayer) {
		message = "Capture in " + ctx.MapX[case_x] + " " + strconv.Itoa(int(ctx.NSize)-case_y+1)
	}

	for current_case_x, current_case_y := case_x+1, case_y+1; current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(*ctx, current_case_x, current_case_y, capturePlayer) == true || CounterHorizontal(*ctx, current_case_x, current_case_y, capturePlayer) ||
				CounterDiag(*ctx, current_case_x, current_case_y, capturePlayer) {
				message = "Capture in " + ctx.MapX[current_case_x+1] + " " + strconv.Itoa(int(ctx.NSize)-current_case_y)
			}
			count_stone++
			current_case_x++
			current_case_y++
		} else {
			break
		}
	}
	for current_case_x, current_case_y := case_x-1, case_y-1; current_case_x >= 0 && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(*ctx, current_case_x, current_case_y, capturePlayer) == true || CounterHorizontal(*ctx, current_case_x, current_case_y, capturePlayer) ||
				CounterDiag(*ctx, current_case_x, current_case_y, capturePlayer) {
				message = "Capture in " + ctx.MapX[current_case_x+1] + " " + strconv.Itoa(int(ctx.NSize)-current_case_y)
			}
			count_stone++
			current_case_x--
			current_case_y--
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" {
			sdl.Log(message)
			d.DisplayMessage(visu, int32((int32(ctx.NSize+1))*ctx.SizeCase), message, "", *ctx)
			return false
		}
		return true
	}
	return false
}

func diagRight(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) bool {
	count_stone := 0
	var message string
	if CounterVertical(*ctx, case_x, case_y, capturePlayer) == true || CounterHorizontal(*ctx, case_x, case_y, capturePlayer) ||
		CounterDiag(*ctx, case_x, case_y, capturePlayer) {
		message = "Capture in " + ctx.MapX[case_x] + " " + strconv.Itoa(int(ctx.NSize)-case_y+1)
	}
	for current_case_x, current_case_y := case_x+1, case_y-1; current_case_x < int(ctx.NSize) && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(*ctx, current_case_x, current_case_y, capturePlayer) == true || CounterHorizontal(*ctx, current_case_x, current_case_y, capturePlayer) ||
				CounterDiag(*ctx, current_case_x, current_case_y, capturePlayer) {
				message = "Capture in " + ctx.MapX[current_case_x+1] + " " + strconv.Itoa(int(ctx.NSize)-current_case_y+1)
			}
			count_stone++
			current_case_x++
			current_case_y--
		} else {
			break
		}
	}
	for current_case_x, current_case_y := case_x-1, case_y+1; current_case_x >= 0 && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(*ctx, current_case_x, current_case_y, capturePlayer) == true || CounterHorizontal(*ctx, current_case_x, current_case_y, capturePlayer) ||
				CounterDiag(*ctx, current_case_x, current_case_y, capturePlayer) {
				message = "Capture in " + ctx.MapX[current_case_x+1] + " " + strconv.Itoa(int(ctx.NSize)-current_case_y+1)
			}
			count_stone++
			current_case_x--
			current_case_y++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" {
			sdl.Log(message)
			d.DisplayMessage(visu, int32((int32(ctx.NSize+1))*ctx.SizeCase), message, "", *ctx)
			return false
		}
		return true
	}
	return false
}

func diagonalAlign(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) bool {
	if diagLeft(ctx, case_x, case_y, capturePlayer, nbCapture, visu) == true {
		return true
	}
	if diagRight(ctx, case_x, case_y, capturePlayer, nbCapture, visu) == true {
		return true
	}
	return false
}

func VictoryConditionAlign(ctx *s.SContext, case_x int, case_y int, visu *s.SVisu) bool {
	ret_value := false
	capturePlayer, nbCapture := 0, 0
	if ctx.CurrentPlayer == 1 {
		capturePlayer = 2
		nbCapture = ctx.NbCaptureP2
	} else {
		capturePlayer = 1
		nbCapture = ctx.NbCaptureP1
	}
	if horizontalAlign(ctx, case_x, case_y, capturePlayer, nbCapture, visu) == true {
		ret_value = true
	}
	if verticalAlign(ctx, case_x, case_y, capturePlayer, nbCapture, visu) == true {
		ret_value = true
	}
	if diagonalAlign(ctx, case_x, case_y, capturePlayer, nbCapture, visu) == true {
		ret_value = true
	}
	if ret_value == true {
		if ctx.CurrentPlayer == 1 {
			ctx.NbVictoryP1++
		} else {
			ctx.NbVictoryP2++
		}
	}
	return ret_value
}

func VictoryCapture(ctx s.SContext) bool {
	if ctx.NbCaptureP1 >= 5 {
		return true
	} else if ctx.NbCaptureP2 >= 5 {
		return true
	}
	return false
}
