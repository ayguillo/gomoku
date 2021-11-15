package game

import (
	d "gomoku/display"
	s "gomoku/structures"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

func CheckDraw(ctx s.SContext) bool {
	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] == 0 {
				return false
			}
		}
	}

	return true
}

func horizontalAlign(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) uint8 {
	count_stone := 0
	var message string
	if CounterVertical(ctx, case_x, case_y, capturePlayer) == true || CounterDiag(ctx, case_x, case_y, capturePlayer) == true {
		message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
	}
	for current_case := case_x + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(ctx, current_case, case_y, capturePlayer) == true || CounterDiag(ctx, current_case, case_y, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			count_stone++
		} else {
			break
		}
	}
	for current_case := case_x - 1; current_case >= 0; current_case-- {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(ctx, current_case, case_y, capturePlayer) == true || CounterDiag(ctx, current_case, case_y, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			count_stone++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" && ctx.ActiveCapture {
			if visu != nil {
				sdl.Log(message)
				d.DisplayMessage(visu, int32((int32(ctx.NSize+1))*ctx.SizeCase), message, "", *ctx)
			}
			return 2
		}
		return 1
	}
	return 0
}

func verticalAlign(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) uint8 {
	count_stone := 0
	var message string
	if CounterHorizontal(ctx, case_x, case_y, capturePlayer) == true || CounterDiag(ctx, case_x, case_y, capturePlayer) == true {
		message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)

	}
	for current_case := case_y + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterHorizontal(ctx, case_x, current_case, capturePlayer) == true || CounterDiag(ctx, case_x, current_case, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			count_stone++
		} else {
			break
		}
	}
	for current_case := case_y - 1; current_case >= 0; current_case-- {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterHorizontal(ctx, case_x, current_case, capturePlayer) == true || CounterDiag(ctx, case_x, current_case, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			count_stone++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" && ctx.ActiveCapture {
			if visu != nil {
				sdl.Log(message)
				d.DisplayMessage(visu, int32((int32(ctx.NSize+1))*ctx.SizeCase), message, "", *ctx)
			}
			return 2
		}
		ctx.Capture = nil
		return 1
	}
	return 0
}

func diagLeft(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) uint8 {
	count_stone := 0
	var message string
	if CounterVertical(ctx, case_x, case_y, capturePlayer) == true || CounterHorizontal(ctx, case_x, case_y, capturePlayer) ||
		CounterDiag(ctx, case_x, case_y, capturePlayer) {
		message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
	}

	for current_case_x, current_case_y := case_x+1, case_y+1; current_case_x < int(ctx.NSize) && current_case_y < int(ctx.NSize); {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(ctx, current_case_x, current_case_y, capturePlayer) == true || CounterHorizontal(ctx, current_case_x, current_case_y, capturePlayer) ||
				CounterDiag(ctx, current_case_x, current_case_y, capturePlayer) {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
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
			if CounterVertical(ctx, current_case_x, current_case_y, capturePlayer) == true || CounterHorizontal(ctx, current_case_x, current_case_y, capturePlayer) ||
				CounterDiag(ctx, current_case_x, current_case_y, capturePlayer) {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			count_stone++
			current_case_x--
			current_case_y--
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" && ctx.ActiveCapture {
			if visu != nil {
				sdl.Log(message)
				d.DisplayMessage(visu, int32((int32(ctx.NSize+1))*ctx.SizeCase), message, "", *ctx)
			}
			return 2
		}
		return 1
	}
	return 0
}

func diagRight(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) uint8 {
	count_stone := 0
	var message string
	if CounterVertical(ctx, case_x, case_y, capturePlayer) == true || CounterHorizontal(ctx, case_x, case_y, capturePlayer) ||
		CounterDiag(ctx, case_x, case_y, capturePlayer) {
		message = "Capture in " + ctx.MapX[case_x] + " " + strconv.Itoa(int(ctx.NSize)-case_y+1)
	}
	for current_case_x, current_case_y := case_x+1, case_y-1; current_case_x < int(ctx.NSize) && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(ctx, current_case_x, current_case_y, capturePlayer) == true || CounterHorizontal(ctx, current_case_x, current_case_y, capturePlayer) ||
				CounterDiag(ctx, current_case_x, current_case_y, capturePlayer) {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
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
			if CounterVertical(ctx, current_case_x, current_case_y, capturePlayer) == true || CounterHorizontal(ctx, current_case_x, current_case_y, capturePlayer) ||
				CounterDiag(ctx, current_case_x, current_case_y, capturePlayer) {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			count_stone++
			current_case_x--
			current_case_y++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" && ctx.ActiveCapture {
			if visu != nil {
				sdl.Log(message)
				d.DisplayMessage(visu, int32((int32(ctx.NSize+1))*ctx.SizeCase), message, "", *ctx)
			}
			return 2
		}
		return 1
	}
	return 0
}

func diagonalAlign(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int, visu *s.SVisu) uint8 {
	left := diagLeft(ctx, case_x, case_y, capturePlayer, nbCapture, visu)
	right := diagRight(ctx, case_x, case_y, capturePlayer, nbCapture, visu)
	if left != 0 {
		return left
	}
	if right != 0 {
		return right
	}
	return 0
}

func VictoryConditionAlign(ctx *s.SContext, case_x int, case_y int, visu *s.SVisu) bool {
	tmp_ret, ret_value := uint8(0), false
	capturePlayer, nbCapture := 0, 0
	if ctx.CurrentPlayer == 1 {
		capturePlayer = 2
		nbCapture = ctx.NbCaptureP2
	} else {
		capturePlayer = 1
		nbCapture = ctx.NbCaptureP1
	}
	horiz := horizontalAlign(ctx, case_x, case_y, capturePlayer, nbCapture, visu)
	vert := verticalAlign(ctx, case_x, case_y, capturePlayer, nbCapture, visu)
	diag := diagonalAlign(ctx, case_x, case_y, capturePlayer, nbCapture, visu)
	if horiz >= tmp_ret {
		tmp_ret = horiz
	}
	if vert >= tmp_ret {
		tmp_ret = vert
	}
	if diag >= tmp_ret {
		tmp_ret = diag
	}
	if tmp_ret == 1 || tmp_ret == 0 {
		ctx.Capture = nil
		if tmp_ret == 1 {
			ret_value = true
		}
	}
	return ret_value
}

func VictoryGoban(ctx *s.SContext, visu *s.SVisu) bool {
	tmp_player := ctx.CurrentPlayer
	victory := false
	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] != 0 {
				victory_condition := VictoryConditionAlign(ctx, x, y, visu)
				if ctx.Capture != nil {
					return false
				}
				if victory_condition {
					victory = true
				}
			}
		}
	}
	if victory {
		if visu != nil {
			if ctx.CurrentPlayer == 1 {
				ctx.NbVictoryP1++
			} else {
				ctx.NbVictoryP2++
			}
		}
		return true
	}
	if ctx.CurrentPlayer == 2 {
		ctx.CurrentPlayer = 1
	} else {
		ctx.CurrentPlayer = 2
	}
	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] != 0 {
				victory_condition := VictoryConditionAlign(ctx, x, y, visu)
				if ctx.Capture != nil {
					return false
				}
				if victory_condition {
					victory = true
				}
			}
		}
	}
	if victory {
		if visu != nil {
			if ctx.CurrentPlayer == 1 {
				ctx.NbVictoryP1++
			} else {
				ctx.NbVictoryP2++
			}
		}
		return true
	}
	ctx.CurrentPlayer = tmp_player
	return victory
}

func VictoryCapture(ctx s.SContext) bool {
	if !ctx.ActiveCapture {
		return false
	}
	if ctx.NbCaptureP1 >= 5 {
		return true
	} else if ctx.NbCaptureP2 >= 5 {
		return true
	}
	return false
}
