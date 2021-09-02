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

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize sdl: %s\n", err)
		panic(err)
	}
	defer sdl.Quit()
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
	defer visu.TexturePlayer.Destroy()
	defer visu.TextureMessage1.Destroy()
	defer visu.TextureMessage2.Destroy()
	defer visu.TextureVictoryP1.Destroy()
	defer visu.TextureVictoryP2.Destroy()
	defer visu.TextureCaptureP1.Destroy()
	defer visu.TextureCaptureP2.Destroy()
	defer visu.TextureNotationX.Destroy()
	defer visu.TextureNotationY.Destroy()
	size_case := (display.H - (int32(ctx.NSize * 3))) / (int32(ctx.NSize) + 2)
	ctx.SizeCase = size_case
	size := int32((int32(ctx.NSize + 1)) * ctx.SizeCase)
	ctx.Size = size
	visu.Window, err = sdl.CreateWindow("Gomoku", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		size+(size/2), size, sdl.WINDOW_SHOWN)
	defer visu.Window.Destroy()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize window: %s\n", err)
		panic(err)
	}
	defer visu.Window.Destroy()
	if err = ttf.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize TTF: %s\n", err)
		panic(err)
	}
	defer ttf.Quit()
	if visu.FontPlayer, err = ttf.OpenFont("fonts/Quicksand-VariableFont_wght.ttf", int(size)/4); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	defer visu.FontPlayer.Close()
	if visu.FontMsg, err = ttf.OpenFont("fonts/Rubik-Regular.ttf", int(size)/4); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	defer visu.FontMsg.Close()
	if visu.FontCounter, err = ttf.OpenFont("fonts/Rubik-Regular.ttf", int(size)/4); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	defer visu.FontCounter.Close()
	visu.Renderer, err = sdl.CreateRenderer(visu.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer : %s\n", err)
		panic(err)
	}
	defer visu.Renderer.Destroy()
	// 0xe2c473
	// Initialize window with color and lines
	d.TraceGoban(&visu, ctx)
	d.DisplayPlayer(&ctx, &visu, true)
	d.DisplayCounter(ctx, &visu)
	middle := math.Round(float64(ctx.NSize)/2) - 2
	var color [4]uint8
	color = [4]uint8{35, 33, 33, 255}
	d.TraceStone(middle, middle, &ctx, &visu, color, false)
	ctx.Goban[int(middle)][int(middle)] = s.Tnumber(2)
	a.FindNeighbors(&ctx, int(middle), int(middle), &visu)
	running := true
	endgame := false
	// Loop de jeu
	for running {
		if ctx.CurrentPlayer == 2 {
			// time.Sleep(1 * time.Second)
			now := time.Now()
			vertex_next, heuris := a.AlphaBetaPruning(ctx, 3)
			delta := time.Since(now)
			fmt.Println(delta)
			fmt.Println("Heuristic", heuris, "vertex", vertex_next)
			color := [4]uint8{35, 33, 33, 255}
			ctx.Goban[int(vertex_next.Y)][int(vertex_next.X)] = s.Tnumber(2)
			a.FindNeighbors(&ctx, int(vertex_next.X), int(vertex_next.Y), &visu)
			d.TraceStone(float64(vertex_next.X), float64(vertex_next.Y), &ctx, &visu, color, false)
			g.Capture(&ctx, &visu, int(vertex_next.X), int(vertex_next.Y), true)
			d.DisplayCapture(ctx, &visu)
			if g.VictoryConditionAlign(&ctx, int(vertex_next.X), int(vertex_next.Y), &visu) == true || g.VictoryCapture(ctx) {
				d.DisplayVictory(&visu, ctx)
				sdl.Log("VICTORY")
				d.DisplayMessage(&visu, size, "Cliquez pour", "relancer", ctx)
				ctx.CurrentPlayer = 1
				endgame = true
				continue
			} else {
				d.DisplayPlayer(&ctx, &visu, false)
			}
			fmt.Println(ctx)
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
						if g.Placement(&ctx, int(case_x), int(case_y)) == true {
							a.FindNeighbors(&ctx, int(case_x), int(case_y), &visu)
							d.DisplayMessage(&visu, size, "", "", ctx)
							// heuris := a.Heuristic(ctx, int(case_x), int(case_y))
							if ctx.CurrentPlayer == 1 {
								color = [4]uint8{240, 228, 229, 255}
							} else {
								color = [4]uint8{35, 33, 33, 255}
							}
							d.TraceStone(case_x, case_y, &ctx, &visu, color, false)
							g.Capture(&ctx, &visu, int(case_x), int(case_y), true)
							d.DisplayCapture(ctx, &visu)
							fmt.Println(ctx)
							if g.VictoryConditionAlign(&ctx, int(case_x), int(case_y), &visu) == true || g.VictoryCapture(ctx) {
								d.DisplayVictory(&visu, ctx)
								sdl.Log("VICTORY")
								d.DisplayMessage(&visu, size, "Cliquez pour", "relancer", ctx)
								endgame = true
								continue
							} else {
								d.DisplayPlayer(&ctx, &visu, false)
							}
						} else {
							d.DisplayMessage(&visu, size, "Il y a déjà", "une pierre", ctx)
							sdl.Log("Il y a déjà une pierre")
						}
					} else {
						d.DisplayMessage(&visu, size, "En dehors", "du terrain", ctx)
						sdl.Log("En dehors du terrain")
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
					color = [4]uint8{35, 33, 33, 255}
					d.TraceStone(middle, middle, &ctx, &visu, color, false)
					ctx.Goban[int(middle)][int(middle)] = s.Tnumber(2)
					a.FindNeighbors(&ctx, int(middle), int(middle), &visu)
					ctx.CurrentPlayer = 1
				}
			}
		}
	}
}
