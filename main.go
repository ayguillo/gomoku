package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	a "gomoku/algogo"
	d "gomoku/display"
	g "gomoku/game"
	m "gomoku/menu"
	s "gomoku/structures"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func initialize() (s.SVisu, s.SContext, error) {
	rand.Seed(time.Now().UnixNano())
	display, _ := sdl.GetDesktopDisplayMode(0)
	// Déclaration de la stricture context
	ctx := s.SContext{}
	ctx.NSize = 19
	ctx.Goban = make([][]s.Tnumber, ctx.NSize)
	ctx.MapX = make(map[int]string)
	ctx.Capture = make([]s.SVertex, 0)
	c := 'A'
	for i := 1; i <= int(ctx.NSize); i++ {
		ctx.MapX[i] = string(c)
		c++
	}
	ctx.CurrentPlayer = 1
	ctx.LastMove = s.SVertex{X: -1, Y: -1}
	ctx.LastLastMove = s.SVertex{X: -1, Y: -1}
	ctx.NbVictoryP1, ctx.NbVictoryP2, ctx.NbCaptureP1, ctx.NbCaptureP2 = 0, 0, 0, 0
	index := 0
	for index < int(ctx.NSize) {
		ctx.Goban[index] = make([]s.Tnumber, ctx.NSize)
		index++
	}
	// Création du plateau + Déclaration de la structure visu
	visu := s.SVisu{}
	visu.FillDefaults()
	size_case := (display.H - (int32(ctx.NSize * 3))) / (int32(ctx.NSize) + 2)
	ctx.SizeCase = size_case
	size := int32((int32(ctx.NSize + 1)) * ctx.SizeCase)
	ctx.Size = size
	err := error(nil)
	visu.Window, err = sdl.CreateWindow("Gomoku", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		size+(size/2), size, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize window: %s\n", err)
		panic(err)
	}
	if err = ttf.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize TTF: %s\n", err)
		panic(err)
	}
	if visu.FontPlayer, err = ttf.OpenFont("fonts/Quicksand-VariableFont_wght.ttf", int(size)/4); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	if visu.FontMsg, err = ttf.OpenFont("fonts/Rubik-Regular.ttf", int(size)/4); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	if visu.FontCounter, err = ttf.OpenFont("fonts/Rubik-Regular.ttf", int(size)/4); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	visu.Renderer, err = sdl.CreateRenderer(visu.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer : %s\n", err)
		panic(err)
	}
	return visu, ctx, err
}

func displayPlay(startgame bool, endgame bool, ctx *s.SContext, visu *s.SVisu, vertex_next s.SVertex) (bool, bool) {
	var color [4]uint8

	if vertex_next.X < 0 || vertex_next.X >= int(ctx.NSize) || vertex_next.Y < 0 || vertex_next.Y >= int(ctx.NSize) {
		d.DisplayEquality(visu, *ctx)
		sdl.Log("Equality")
		d.DisplayMessage(visu, ctx.Size, "Cliquez pour", "relancer", *ctx)
		return true, true
	}

	if ctx.CurrentPlayer == 1 {
		color = [4]uint8{240, 228, 229, 255}
	} else {
		color = [4]uint8{35, 33, 33, 255}
	}
	a.FindNeighbors(ctx, vertex_next.X, vertex_next.Y)
	if ctx.LastMove.X != -1 {
		tmp_last := ctx.LastMove
		ctx.LastMove = vertex_next
		ctx.LastLastMove = tmp_last
	} else {
		ctx.LastMove = vertex_next
	}
	d.DisplayMessage(visu, ctx.Size, "", "", *ctx)
	d.TraceStone(float64(vertex_next.X), float64(vertex_next.Y), ctx, visu, color, false)
	if ctx.ActiveCapture {
		g.Capture(ctx, visu, int(vertex_next.X), int(vertex_next.Y), true)
		d.DisplayCapture(*ctx, visu)
	}
	if g.VictoryGoban(ctx, visu) == true || g.VictoryCapture(*ctx) {
		d.DisplayVictory(visu, *ctx)
		sdl.Log("VICTORY")
		d.DisplayMessage(visu, ctx.Size, "Cliquez pour", "relancer", *ctx)
		return true, true
	} else if g.CheckDraw(*ctx) {
		d.DisplayEquality(visu, *ctx)
		sdl.Log("Equality")
		d.DisplayMessage(visu, ctx.Size, "Cliquez pour", "relancer", *ctx)
		return true, true
	} else {
		d.DisplayPlayer(ctx, visu, false)
	}
	return startgame, endgame
}

func bot(startgame bool, endgame bool, ctx *s.SContext, visu *s.SVisu) (bool, bool) {
	var color [4]uint8
	if startgame == true {
		if ctx.CurrentPlayer == 1 {
			color = [4]uint8{240, 228, 229, 255}
		} else {
			color = [4]uint8{35, 33, 33, 255}
		}
		middle := math.Round(float64(ctx.NSize)/2) - 2
		if ctx.Goban[int(middle)][int(middle)] != 0 {
			middle++
		}
		d.TraceStone(middle, middle, ctx, visu, color, false)
		ctx.Goban[int(middle)][int(middle)] = s.Tnumber(ctx.CurrentPlayer)
		a.FindNeighbors(ctx, int(middle), int(middle))
		startgame = false
		d.DisplayPlayer(ctx, visu, false)
	} else {
		player := ctx.CurrentPlayer
		now := time.Now()
		vertex_next, _ := a.MinimaxTree(*ctx, ctx.Depth)
		delta := time.Since(now)
		d.DisplayTime(visu, delta.String(), ctx.Size, player)
		sdl.Log(delta.String())
		if vertex_next.X != -1 && vertex_next.Y != -1 {
			ctx.Goban[int(vertex_next.Y)][int(vertex_next.X)] = s.Tnumber(ctx.CurrentPlayer)
		}
		startgame, endgame = displayPlay(startgame, endgame, ctx, visu, vertex_next)
	}
	if ctx.ActiveHelp {
		current := ctx.CurrentPlayer
		if ctx.CurrentPlayer == 1 {
			ctx.CurrentPlayer = 2
		} else {
			ctx.CurrentPlayer = 1
		}
		if ctx.VertexHelp.X != -1 && ctx.Goban[ctx.VertexHelp.Y][ctx.VertexHelp.X] == 0 {
			color = [4]uint8{226, 196, 115, 255}
			d.TraceStone(float64(ctx.VertexHelp.X), float64(ctx.VertexHelp.Y), ctx, visu, color, true)
		}

		ctx.CurrentPlayer = current
	}
	return startgame, endgame
}

func human(err error, startgame bool, endgame bool, ctx *s.SContext, visu *s.SVisu, t *sdl.MouseButtonEvent) (bool, bool) {
	if err != nil {
		panic(err)
	}
	// Trouver intersection la plus proche
	h_mouse := float64(t.Y - 5)
	k_mouse := float64(t.X - 5)
	// Traduit la coordonnee sur le tableau
	case_x := math.Round((k_mouse-float64(ctx.Size/4))/float64(ctx.SizeCase)) - 1
	case_y := math.Round(h_mouse/float64(ctx.SizeCase)) - 1
	if (case_x >= 0 && uint8(case_x) < ctx.NSize) && (case_y >= 0 && uint8(case_y) < ctx.NSize) {
		placement := g.Placement(ctx, int(case_x), int(case_y))
		if placement == 0 {
			ctx.Goban[int(case_y)][int(case_x)] = s.Tnumber(ctx.CurrentPlayer)
			startgame, endgame = displayPlay(startgame, endgame, ctx, visu, s.SVertex{X: int(case_x), Y: int(case_y)})
		} else if placement < 0 {
			d.DisplayMessage(visu, ctx.Size, "Il y a déjà", "une pierre", *ctx)
			sdl.Log("Il y a déjà une pierre")
		} else if placement == 1 {
			d.DisplayMessage(visu, ctx.Size, "Double Three", "", *ctx)
			sdl.Log("Double three")
		}
	} else {
		d.DisplayMessage(visu, ctx.Size, "En dehors", "du terrain", *ctx)
		sdl.Log("En dehors du terrain")
	}
	startgame = false
	return startgame, endgame
}

func main() {

	// 0xe2c473
	// Initialize window with color and lines
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize sdl: %s\n", err)
		panic(err)
	}
	defer sdl.Quit()
	visu, ctx, err := initialize()
	defer visu.TexturePlayer.Destroy()
	defer visu.TextureMessage1.Destroy()
	defer visu.TextureMessage2.Destroy()
	defer visu.TextureVictoryP1.Destroy()
	defer visu.TextureVictoryP2.Destroy()
	defer visu.TextureCaptureP1.Destroy()
	defer visu.TextureCaptureP2.Destroy()
	defer visu.TextureNotationX.Destroy()
	defer visu.TextureNotationY.Destroy()
	defer visu.Window.Destroy()
	defer ttf.Quit()
	defer visu.FontPlayer.Close()
	defer visu.FontMsg.Close()
	defer visu.FontCounter.Close()
	defer visu.Renderer.Destroy()

	versus, double_threes, capture, help, end, difficulty, thinking := m.Menu(&visu, ctx)
	visu.Renderer.Clear()
	visu.Renderer.Present()
	if !end {
		d.TraceGoban(&visu, ctx)
		d.DisplayPlayer(&ctx, &visu, true)
		d.DisplayCounter(ctx, &visu)
		ctx.Players = make(map[uint8]bool)
		if versus == 0 {
			ctx.Players[1] = true
			ctx.Players[2] = false
		} else if versus == 1 {
			ctx.Players[1] = false
			ctx.Players[2] = false
		} else {
			ctx.Players[1] = true
			ctx.Players[2] = true
		}

		ctx.ActiveCapture = capture
		ctx.ActiveHelp = help
		ctx.ActiveDoubleThrees = double_threes
		ctx.TimeToThink = thinking
		if help {
			ctx.VertexHelp = s.SVertex{X: -1, Y: -1}
		}
		if difficulty == 0 {
			ctx.Depth = 2
		} else if difficulty == 1 {
			ctx.Depth = 5
		} else {
			ctx.Depth = 7
		}
	}
	running := true
	endgame := false
	startgame := true
	displayHelp := true
	// Loop de jeu
	for running && !end {
		if ctx.Players[ctx.CurrentPlayer] == true && endgame != true {
			startgame, endgame = bot(startgame, endgame, &ctx, &visu)
			displayHelp = true
		}
		if ctx.Players[ctx.CurrentPlayer] == false && ctx.ActiveHelp == true && endgame != true && displayHelp == true {
			color := [4]uint8{83, 51, 237, 1}
			if startgame == true {
				middle := math.Round(float64(ctx.NSize)/2) - 2
				ctx.VertexHelp = s.SVertex{X: int(middle), Y: int(middle)}
				d.TraceStone(middle, middle, &ctx, &visu, color, false)
			} else {
				vertex_help, _ := a.MinimaxTree(ctx, 2)
				ctx.VertexHelp = vertex_help
				d.TraceStone(float64(vertex_help.X), float64(vertex_help.Y), &ctx, &visu, color, false)
			}
			displayHelp = false
		}
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				sdl.Log("Quit")
				running = false
			case *sdl.KeyboardEvent:
				if t.State == sdl.PRESSED && t.Keysym.Sym == sdl.K_ESCAPE {
					sdl.Log("Quit")
					running = false
				}
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED && endgame == false && ctx.Players[ctx.CurrentPlayer] == false {
					if ctx.VertexHelp.X != -1 {
						color := [4]uint8{226, 196, 115, 255}
						d.TraceStone(float64(ctx.VertexHelp.X), float64(ctx.VertexHelp.Y), &ctx, &visu, color, true)
					}
					startgame, endgame = human(err, startgame, endgame, &ctx, &visu, t)
					displayHelp = true
					if endgame == true {
						t.State = 0
					}
				}
				if t.State == sdl.PRESSED && endgame == true {
					displayHelp = true
					visu.Renderer.Clear()
					visu.Renderer.Present()
					index := 0
					for index < int(ctx.NSize) {
						ctx.Goban[index] = nil
						index++
					}
					ctx.Goban = nil
					ctx.Goban = make([][]s.Tnumber, ctx.NSize)
					index = 0
					for index < int(ctx.NSize) {
						ctx.Goban[index] = make([]s.Tnumber, ctx.NSize)
						index++
					}
					d.TraceGoban(&visu, ctx)
					d.DisplayCounter(ctx, &visu)
					ctx.NbCaptureP1 = 0
					ctx.NbCaptureP2 = 0
					ctx.CasesNonNull = nil
					endgame = false
					startgame = true
					d.DisplayPlayer(&ctx, &visu, false)
				}
			}
		}
	}
}
