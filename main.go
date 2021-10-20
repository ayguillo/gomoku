package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	a "gomoku/algorithm"
	d "gomoku/display"
	g "gomoku/game"
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
	ctx.Capture = s.SVertex{X: -1, Y: -1}
	c := 'A'
	for i := 1; i <= int(ctx.NSize); i++ {
		ctx.MapX[i] = string(c)
		c++
	}
	ctx.CurrentPlayer = 1
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

	if ctx.CurrentPlayer == 1 {
		color = [4]uint8{240, 228, 229, 255}
	} else {
		color = [4]uint8{35, 33, 33, 255}
	}
	a.FindNeighbors(ctx, int(vertex_next.X), int(vertex_next.Y))
	d.DisplayMessage(visu, ctx.Size, "", "", *ctx)
	d.TraceStone(float64(vertex_next.X), float64(vertex_next.Y), ctx, visu, color, false)
	g.Capture(ctx, visu, int(vertex_next.X), int(vertex_next.Y), true)
	d.DisplayCapture(*ctx, visu)
	if g.VictoryConditionAlign(ctx, int(vertex_next.X), int(vertex_next.Y), visu) == true || g.VictoryCapture(*ctx) {
		d.DisplayVictory(visu, *ctx)
		sdl.Log("VICTORY")
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
		depth := int8(6)
		now := time.Now()
		vertex_next, heuris := a.AlphaBetaPruning2(*ctx, depth)
		fmt.Println(vertex_next, heuris)
		delta := time.Since(now)
		fmt.Println(delta)
		ctx.Goban[int(vertex_next.Y)][int(vertex_next.X)] = s.Tnumber(ctx.CurrentPlayer)
		startgame, endgame = displayPlay(startgame, endgame, ctx, visu, vertex_next)
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
		if g.Placement(ctx, int(case_x), int(case_y)) == true {
			startgame, endgame = displayPlay(startgame, endgame, ctx, visu, s.SVertex{X: int(case_x), Y: int(case_y)})
		} else {
			d.DisplayMessage(visu, ctx.Size, "Il y a déjà", "une pierre", *ctx)
			sdl.Log("Il y a déjà une pierre")
		}
	} else {
		d.DisplayMessage(visu, ctx.Size, "En dehors", "du terrain", *ctx)
		sdl.Log("En dehors du terrain")
	}
	return startgame, endgame
}

func menu(visu *s.SVisu, ctx s.SContext) (int, bool, bool, int) {
	versus, double_threes, capture, time_limite := 0, true, true, 0
	// Window
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: 0, Y: 0, W: ctx.Size + ctx.Size/2, H: ctx.Size + ((ctx.SizeCase) / 2)})
	visu.Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: ctx.Size + ctx.Size/2, H: ctx.Size + ((ctx.SizeCase) / 2)})
	// Button quit
	visu.Renderer.SetDrawColor(212, 66, 62, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: ctx.Size + 4 + ctx.Size/4, Y: ctx.Size - 50, W: (ctx.Size / 4), H: 50})
	visu.Renderer.FillRect(&sdl.Rect{X: ctx.Size + 4 + ctx.Size/4, Y: ctx.Size - 50, W: (ctx.Size / 4), H: 50})
	// Buttons players
	visu.Renderer.SetDrawColor(0, 0, 0, 0)
	visu.Renderer.DrawRect(&sdl.Rect{X: 10, Y: 80, W: 50, H: 50})
	visu.Renderer.FillRect(&sdl.Rect{X: 10, Y: 80, W: 50, H: 50})
	visu.Renderer.DrawRect(&sdl.Rect{X: 70, Y: 80, W: 50, H: 50})
	visu.Renderer.FillRect(&sdl.Rect{X: 70, Y: 80, W: 50, H: 50})
	visu.Renderer.DrawRect(&sdl.Rect{X: 130, Y: 80, W: 50, H: 50})
	visu.Renderer.FillRect(&sdl.Rect{X: 130, Y: 80, W: 50, H: 50})

	color := sdl.Color{R: 240, G: 228, B: 229, A: 255}
	bmp, err := visu.FontPlayer.RenderUTF8Solid("Play >", color)
	bmp2, err2 := visu.FontPlayer.RenderUTF8Solid("Players :", color)
	bmp3, err2 := visu.FontPlayer.RenderUTF8Solid("Human vs Bot", color)
	bmp4, err2 := visu.FontPlayer.RenderUTF8Solid("Human vs Human", color)
	bmp5, err2 := visu.FontPlayer.RenderUTF8Solid("Bot vs Bot", color)

	if err != nil || err2 != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font: %s %s\n", err, err2)
		panic(err)
	}
	button, err := visu.Renderer.CreateTextureFromSurface(bmp)
	text_player, err2 := visu.Renderer.CreateTextureFromSurface(bmp2)
	hvb, err2 := visu.Renderer.CreateTextureFromSurface(bmp3)
	hvh, err2 := visu.Renderer.CreateTextureFromSurface(bmp4)
	bvb, err2 := visu.Renderer.CreateTextureFromSurface(bmp5)
	if err != nil || err2 != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font: %s %s\n", err, err2)
		panic(err)
	}
	bmp.Free()
	visu.Renderer.Copy(button, nil, &sdl.Rect{X: ctx.Size + 4 + ctx.Size/4, Y: ctx.Size - 50, W: (ctx.Size / 4), H: 50})
	visu.Renderer.Copy(text_player, nil, &sdl.Rect{X: 10, Y: 10, W: 130, H: 50})
	visu.Renderer.Present()
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			if versus == 0 {
				visu.Renderer.SetDrawColor(226, 196, 115, 255)
				visu.Renderer.DrawRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
				visu.Renderer.Copy(hvb, nil, &sdl.Rect{X: 150, Y: 10, W: 170, H: 50})
				visu.Renderer.SetDrawColor(83, 51, 237, 1)
				visu.Renderer.DrawRect(&sdl.Rect{X: 10, Y: 80, W: 50, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 10, Y: 80, W: 50, H: 50})
				visu.Renderer.SetDrawColor(0, 0, 0, 0)
				visu.Renderer.DrawRect(&sdl.Rect{X: 70, Y: 80, W: 50, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 70, Y: 80, W: 50, H: 50})
				visu.Renderer.DrawRect(&sdl.Rect{X: 130, Y: 80, W: 50, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 130, Y: 80, W: 50, H: 50})
				visu.Renderer.Present()
			} else if versus == 1 {
				visu.Renderer.SetDrawColor(226, 196, 115, 255)
				visu.Renderer.DrawRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
				visu.Renderer.Copy(hvh, nil, &sdl.Rect{X: 150, Y: 10, W: 220, H: 50})
				visu.Renderer.SetDrawColor(0, 0, 0, 0)
				visu.Renderer.DrawRect(&sdl.Rect{X: 10, Y: 80, W: 50, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 10, Y: 80, W: 50, H: 50})
				visu.Renderer.SetDrawColor(83, 51, 237, 1)
				visu.Renderer.DrawRect(&sdl.Rect{X: 70, Y: 80, W: 50, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 70, Y: 80, W: 50, H: 50})
				visu.Renderer.SetDrawColor(0, 0, 0, 0)
				visu.Renderer.DrawRect(&sdl.Rect{X: 130, Y: 80, W: 50, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 130, Y: 80, W: 50, H: 50})
				visu.Renderer.Present()
			} else {
				visu.Renderer.SetDrawColor(226, 196, 115, 255)
				visu.Renderer.DrawRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
				visu.Renderer.Copy(bvb, nil, &sdl.Rect{X: 150, Y: 10, W: 120, H: 50})
				visu.Renderer.SetDrawColor(0, 0, 0, 0)
				visu.Renderer.DrawRect(&sdl.Rect{X: 10, Y: 80, W: 50, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 10, Y: 80, W: 50, H: 50})
				visu.Renderer.DrawRect(&sdl.Rect{X: 70, Y: 80, W: 50, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 70, Y: 80, W: 50, H: 50})
				visu.Renderer.SetDrawColor(83, 51, 237, 1)
				visu.Renderer.DrawRect(&sdl.Rect{X: 130, Y: 80, W: 50, H: 50})
				visu.Renderer.FillRect(&sdl.Rect{X: 130, Y: 80, W: 50, H: 50})
				visu.Renderer.Present()
			}
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.State == sdl.PRESSED && t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED {
					if t.Y >= ctx.Size-50 && t.X >= ctx.Size+4+ctx.Size/4 {
						running = false
					}
					if (t.Y >= 80 && t.Y <= 80+50) && (t.X >= 10 && t.X <= 60) {
						versus = 0
					}
					if (t.Y >= 80 && t.Y <= 80+50) && (t.X >= 70 && t.X <= 120) {
						versus = 1
					}
					if (t.Y >= 80 && t.Y <= 80+50) && (t.X >= 130 && t.X <= 180) {
						versus = 2
					}
				}
			}

		}
	}
	return versus, double_threes, capture, time_limite
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

	versus, double_threes, capture, time_limite := menu(&visu, ctx)
	visu.Renderer.Clear()
	visu.Renderer.Present()
	fmt.Println(versus, double_threes, capture, time_limite)
	d.TraceGoban(&visu, ctx)
	d.DisplayPlayer(&ctx, &visu, true)
	d.DisplayCounter(ctx, &visu)
	ctx.Players = make(map[uint8]bool)
	ctx.Players[1] = true
	ctx.Players[2] = false
	// middle := math.Round(float64(ctx.NSize)/2) - 2
	// color = [4]uint8{35, 33, 33, 255}
	// d.TraceStone(middle, middle, &ctx, &visu, color, false)
	// ctx.Goban[int(middle)][int(middle)] = s.Tnumber(2)
	// a.FindNeighbors(&ctx, int(middle), int(middle))
	running := true
	endgame := false
	startgame := true
	// Loop de jeu
	for running {
		if ctx.Players[ctx.CurrentPlayer] == true && endgame != true {
			startgame, endgame = bot(startgame, endgame, &ctx, &visu)
		}
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
			case *sdl.KeyboardEvent:
				if t.State == sdl.PRESSED && t.Keysym.Sym == sdl.K_ESCAPE {
					fmt.Println("Quit")
					running = false
				}
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED && endgame == false {
					startgame, endgame = human(err, endgame, startgame, &ctx, &visu, t)
					if endgame == true {
						t.State = 0
					}
				}
				if t.State == sdl.PRESSED && endgame == true {
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
					d.DisplayPlayer(&ctx, &visu, false)
					d.DisplayCounter(ctx, &visu)
					ctx.NbCaptureP1 = 0
					ctx.NbCaptureP2 = 0
					ctx.CasesNonNull = nil
					endgame = false
					ctx.CurrentPlayer = 1
				}
			}
		}
	}
}
