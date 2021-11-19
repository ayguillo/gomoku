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

func horizontalAlign(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) (uint8, []s.SVertex, string) {
	count_stone := 0
	var message string
	stones := make([]s.SVertex, 0)
	if CounterVertical(ctx, case_x, case_y, capturePlayer) == true || CounterDiag(ctx, case_x, case_y, capturePlayer) == true {
		message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
	}
	for current_case := case_x + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[case_y][current_case] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(ctx, current_case, case_y, capturePlayer) == true || CounterDiag(ctx, current_case, case_y, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			count_stone++
			stones = append(stones, s.SVertex{X: current_case, Y: case_y})
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
			stones = append(stones, s.SVertex{X: current_case, Y: case_y})
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" && ctx.ActiveCapture {
			return 2, stones, message
		}
		return 1, stones, message
	}
	return 0, stones, message
}

func verticalAlign(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) (uint8, []s.SVertex, string) {
	count_stone := 0
	var message string
	stones := make([]s.SVertex, 0)
	if CounterHorizontal(ctx, case_x, case_y, capturePlayer) == true || CounterDiag(ctx, case_x, case_y, capturePlayer) == true {
		message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)

	}
	for current_case := case_y + 1; current_case < int(ctx.NSize); current_case++ {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterHorizontal(ctx, case_x, current_case, capturePlayer) == true || CounterDiag(ctx, case_x, current_case, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			count_stone++
			stones = append(stones, s.SVertex{X: case_x, Y: current_case})
		} else {
			break
		}
	}
	for current_case := case_y - 1; current_case >= 0; current_case-- {
		if ctx.Goban[current_case][case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterHorizontal(ctx, case_x, current_case, capturePlayer) == true || CounterDiag(ctx, case_x, current_case, capturePlayer) == true {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			stones = append(stones, s.SVertex{X: case_x, Y: current_case})
			count_stone++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" && ctx.ActiveCapture {
			return 2, stones, message
		}
		ctx.Capture = nil
		return 1, stones, message
	}
	return 0, stones, message
}

func diagLeft(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) (uint8, []s.SVertex, string) {
	count_stone := 0
	var message string
	stones := make([]s.SVertex, 0)
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
			stones = append(stones, s.SVertex{X: current_case_x, Y: current_case_y})
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
			stones = append(stones, s.SVertex{X: current_case_x, Y: current_case_y})
			current_case_x--
			current_case_y--
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" && ctx.ActiveCapture {
			return 2, stones, message
		}
		return 1, stones, message
	}
	return 0, stones, message
}

func diagRight(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) (uint8, []s.SVertex, string) {
	count_stone := 0
	var message string
	stones := make([]s.SVertex, 0)
	if CounterVertical(ctx, case_x, case_y, capturePlayer) == true || CounterHorizontal(ctx, case_x, case_y, capturePlayer) ||
		CounterDiag(ctx, case_x, case_y, capturePlayer) {
		message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
	}
	for current_case_x, current_case_y := case_x+1, case_y-1; current_case_x < int(ctx.NSize) && current_case_y >= 0; {
		if ctx.Goban[current_case_y][current_case_x] == s.Tnumber(ctx.CurrentPlayer) {
			if CounterVertical(ctx, current_case_x, current_case_y, capturePlayer) == true || CounterHorizontal(ctx, current_case_x, current_case_y, capturePlayer) ||
				CounterDiag(ctx, current_case_x, current_case_y, capturePlayer) {
				message = "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			}
			count_stone++
			stones = append(stones, s.SVertex{X: current_case_x, Y: current_case_y})
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
			stones = append(stones, s.SVertex{X: current_case_x, Y: current_case_y})
			current_case_x--
			current_case_y++
		} else {
			break
		}
	}
	if count_stone >= 4 {
		if message != "" && ctx.ActiveCapture {
			return 2, stones, message
		}
		return 1, stones, message
	}
	return 0, stones, message
}

func diagonalAlign(ctx *s.SContext, case_x int, case_y int, capturePlayer int, nbCapture int) (uint8, []s.SVertex, string) {
	left, stonesleft, messageL := diagLeft(ctx, case_x, case_y, capturePlayer, nbCapture)
	right, stonesright, messageD := diagRight(ctx, case_x, case_y, capturePlayer, nbCapture)
	if left != 0 {
		return left, stonesleft, messageL
	}
	if right != 0 {
		return right, stonesright, messageD
	}
	return 0, nil, ""
}

func removeDuplicateCaptures(ctx *s.SContext) {
	keys := make(map[s.SVertex]bool)
	list := []s.SVertex{}
	for _, entry := range ctx.Capture {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	ctx.Capture = list
}

func VictoryConditionAlign(ctx *s.SContext, case_x int, case_y int) (bool, []s.SVertex, string) {
	tmp_ret, ret_value := uint8(0), false
	capturePlayer, nbCapture := 0, 0
	message := ""
	var ret_stones []s.SVertex

	if ctx.CurrentPlayer == 1 {
		capturePlayer = 2
		nbCapture = ctx.NbCaptureP2
	} else {
		capturePlayer = 1
		nbCapture = ctx.NbCaptureP1
	}
	horiz, stonesHoriz, messageH := horizontalAlign(ctx, case_x, case_y, capturePlayer, nbCapture)
	vert, stonesVert, messageV := verticalAlign(ctx, case_x, case_y, capturePlayer, nbCapture)
	diag, stonesDiag, messageD := diagonalAlign(ctx, case_x, case_y, capturePlayer, nbCapture)
	removeDuplicateCaptures(ctx)
	if horiz >= tmp_ret {
		tmp_ret = horiz
		ret_stones = stonesHoriz
		if horiz == 2 {
			message = messageH
		}
	}
	if vert >= tmp_ret {
		tmp_ret = vert
		ret_stones = stonesVert
		if vert == 2 {
			message = messageV
		}
	}
	if diag >= tmp_ret {
		tmp_ret = diag
		ret_stones = stonesDiag
		if diag == 2 {
			message = messageD
		}
	}
	if tmp_ret == 1 || tmp_ret == 0 {
		ctx.Capture = nil
		message = ""
	}
	if tmp_ret == 1 || tmp_ret == 2 {
		ret_value = true
	}
	return ret_value, ret_stones, message
}

func verifStones(stoneGoban s.SVertex, stones []s.SVertex) bool {
	for stone := range stones {
		if stoneGoban.X == stones[stone].X && stoneGoban.Y == stones[stone].Y {
			return true
		}
	}
	return false
}

func copyGoban(goban s.Tgoban) s.Tgoban {
	newGoban := make([][]s.Tnumber, 19)
	for Y, line := range goban {
		newGoban[Y] = make([]s.Tnumber, 19)
		for X, nb := range line {
			newGoban[Y][X] = nb
		}
	}
	return newGoban
}

func verifyCaptures(tmp_ctx *s.SContext, ctx *s.SContext, visu *s.SVisu) bool {
	counterVictory := 0
	cap := make([]s.SVertex, 0)
	stones := make([]s.SVertex, 0)
	capture := make([]s.SVertex, len(tmp_ctx.Capture))
	copy(capture, tmp_ctx.Capture)
	for i := range capture {
		opp := 1
		tmp := tmp_ctx.CurrentPlayer
		if tmp_ctx.CurrentPlayer == 1 {
			opp = 2
		}
		vertex_cap := capture[i]
		tmp_ctx.CurrentPlayer = uint8(opp)
		tmp_ctx.Goban[vertex_cap.Y][vertex_cap.X] = s.Tnumber(opp)
		save_cap := tmp_ctx.Goban[vertex_cap.Y][vertex_cap.X]
		_, ret_vertex := Capture(tmp_ctx, nil, vertex_cap.X, vertex_cap.Y, false)
		tmp_ctx.CurrentPlayer = tmp
		for y := range tmp_ctx.Goban {
			for x := range tmp_ctx.Goban[y] {
				if tmp_ctx.Goban[y][x] != 0 && tmp_ctx.Goban[y][x] == s.Tnumber(tmp_ctx.CurrentPlayer) && verifStones(s.SVertex{X: x, Y: y}, stones) == false {
					victory_condition, tmp_stones, _ := VictoryConditionAlign(tmp_ctx, x, y)
					for stone := range tmp_stones {
						stones = append(stones, tmp_stones[stone])
					}
					if victory_condition {
						counterVictory++
					} else {
						cap = append(cap, vertex_cap)
					}
				}
			}
		}
		opp = 1
		if save_cap == 1 {
			opp = 2
		}
		for vert := range ret_vertex {
			tmp_ctx.Goban[ret_vertex[vert].Y][ret_vertex[vert].X] = s.Tnumber(opp)
		}
		tmp_ctx.Goban[vertex_cap.Y][vertex_cap.X] = 0
	}
	if counterVictory == 0 {
		ctx.Capture = cap
		if visu != nil {
			message := "Capture in " + ctx.MapX[ctx.Capture[0].X+1] + " " + strconv.Itoa(int(ctx.NSize)-ctx.Capture[0].Y)
			sdl.Log(message)
			d.DisplayMessage(visu, int32((int32(ctx.NSize+1))*ctx.SizeCase), message, "", *ctx)
		}
		return false
	} else {
		if visu != nil {
			if ctx.CurrentPlayer == 1 {
				ctx.NbVictoryP1++
			} else {
				ctx.NbVictoryP2++
			}
		}
		return true
	}
}

func VictoryGoban(ctx *s.SContext, visu *s.SVisu) bool {
	tmp_player := ctx.CurrentPlayer
	victory := false
	counterVictory := 0
	stones := make([]s.SVertex, 0)
	message := ""
	var tmp_cap []s.SVertex
	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] != 0 && ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) && verifStones(s.SVertex{X: x, Y: y}, stones) == false {
				victory_condition, tmp_stones, tmp_message := VictoryConditionAlign(ctx, x, y)
				for stone := range tmp_stones {
					stones = append(stones, tmp_stones[stone])
				}
				if victory_condition {
					counterVictory++
					victory = true
				}
				if tmp_message != "" {
					message = tmp_message
					tmp_cap = append(tmp_cap, ctx.Capture...)
				}
			}
		}
	}
	if (victory && counterVictory > 1) || (victory && counterVictory == 1 && message == "") {
		if visu != nil {
			if ctx.CurrentPlayer == 1 {
				ctx.NbVictoryP1++
			} else {
				ctx.NbVictoryP2++
			}
		}
		return true
	} else if victory && counterVictory == 1 && message != "" {
		tmp_ctx := s.SContext{
			CurrentPlayer: ctx.CurrentPlayer,
			Goban:         copyGoban(ctx.Goban),
			Capture:       tmp_cap,
			MapX:          ctx.MapX,
			NSize:         ctx.NSize,
			ActiveCapture: ctx.ActiveCapture,
		}
		victory = verifyCaptures(&tmp_ctx, ctx, visu)
		if victory {
			return true
		}
		ctx.Capture = tmp_cap
	}
	if ctx.CurrentPlayer == 2 {
		ctx.CurrentPlayer = 1
	} else {
		ctx.CurrentPlayer = 2
	}
	victory = false
	counterVictory = 0
	stones = make([]s.SVertex, 0)
	message = ""
	for y := range ctx.Goban {
		for x := range ctx.Goban[y] {
			if ctx.Goban[y][x] != 0 && ctx.Goban[y][x] == s.Tnumber(ctx.CurrentPlayer) {
				victory_condition, tmp_stones, tmp_message := VictoryConditionAlign(ctx, x, y)
				for stone := range tmp_stones {
					stones = append(stones, tmp_stones[stone])
				}
				if victory_condition {
					victory = true
				}
				if tmp_message != "" {
					message = tmp_message
					tmp_cap = append(tmp_cap, ctx.Capture...)
				}
			}
		}
	}
	if (victory && counterVictory > 1) || (victory && counterVictory == 1 && message == "") {
		if visu != nil {
			if ctx.CurrentPlayer == 1 {
				ctx.NbVictoryP1++
			} else {
				ctx.NbVictoryP2++
			}
		}
		ctx.CurrentPlayer = tmp_player
		return true
	}
	ctx.CurrentPlayer = tmp_player
	ctx.Capture = tmp_cap
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
