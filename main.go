package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	s "gomoku/structures"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func display_message(visu *s.SVisu, size int32, line1 string, line2 string) {
	if visu.TextureMessage1 != nil {
		visu.TextureMessage1.Destroy()
		visu.TextureMessage2.Destroy()
	}
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: size + 4, Y: size - 100, W: (size / 4) - 10, H: 100})
	visu.Renderer.FillRect(&sdl.Rect{X: size + 4, Y: size - 100, W: (size / 4) - 10, H: 100})
	visu.Renderer.Present()
	if line1 != "" {
		color := sdl.Color{R: 212, G: 66, B: 62, A: 255}
		bmp, err := visu.FontMsg.RenderUTF8Solid(line1, color)
		bmp2, err2 := visu.FontMsg.RenderUTF8Solid(line2, color)
		if err != nil || err2 != nil {
			fmt.Fprintf(os.Stderr, "Failed to renderer font: %s\n", err)
			panic(err)
		}
		texture, err := visu.Renderer.CreateTextureFromSurface(bmp)
		texture2, err2 := visu.Renderer.CreateTextureFromSurface(bmp2)
		if err != nil || err2 != nil {
			fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err)
			panic(err)
		}
		visu.Renderer.Present()
		visu.Renderer.Copy(texture, nil, &sdl.Rect{X: size + 4, Y: size - 100, W: (size / 4) - 10, H: 50})
		visu.Renderer.Copy(texture2, nil, &sdl.Rect{X: size + 4, Y: size - 50, W: (size / 4) - 10, H: 50})
		visu.Renderer.Present()
	}
}

func display_player(ctx *s.SContext, visu *s.SVisu, size_case int, current bool) {
	var text string
	var color sdl.Color
	size := int32(int((ctx.NSize + 1)) * int(size_case))
	if ctx.CurrentPlayer == 1 && current == false {
		ctx.CurrentPlayer = 2
		color = sdl.Color{R: 35, G: 33, B: 33, A: 255}
		text = "Player 2"
	} else if ctx.CurrentPlayer == 2 && current == false {
		ctx.CurrentPlayer = 1
		color = sdl.Color{R: 240, G: 228, B: 229, A: 255}
		text = "Player 1"
	} else if ctx.CurrentPlayer == 1 && current == true {
		color = sdl.Color{R: 240, G: 228, B: 229, A: 255}
		text = "Player 1"
	} else if ctx.CurrentPlayer == 2 && current == true {
		color = sdl.Color{R: 35, G: 33, B: 33, A: 255}
		text = "Player 2"
	}
	bmp, err := visu.FontPlayer.RenderUTF8Solid(text, color)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to renderer font: %s\n", err)
		panic(err)
	}
	if visu.TexturePlayer != nil {
		visu.TexturePlayer.Destroy()
	}
	visu.TexturePlayer, err = visu.Renderer.CreateTextureFromSurface(bmp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err)
		panic(err)
	}
	bmp.Free()
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: size + 4, Y: 0, W: (size / 4) - 10, H: 50})
	visu.Renderer.FillRect(&sdl.Rect{X: size + 4, Y: 0, W: (size / 4) - 10, H: 50})
	visu.Renderer.Present()
	visu.Renderer.Copy(visu.TexturePlayer, nil, &sdl.Rect{X: size + 4, Y: 0, W: (size / 4) - 10, H: 50})
	visu.Renderer.Present()
	sdl.Delay(16)
}

func trace_stone(case_x float64, case_y float64, size_case int, nb_case int, ctx *s.SContext, visu *s.SVisu) {
	k := case_x * float64(size_case)
	h := case_y * float64(size_case)
	radius := 10.0

	if ctx.CurrentPlayer == 1 {
		visu.Renderer.SetDrawColor(240, 228, 229, 255)
	} else {
		visu.Renderer.SetDrawColor(35, 33, 33, 255)
	}
	visu.Renderer.DrawLine(int32(k-radius), int32(h), int32(k+radius), int32(h))
	// Draw circle
	visu.Renderer.Present()
	for dy := 1.0; dy < radius; dy += 1.0 {
		dx := math.Sqrt((2.0 * radius * dy) - (dy * dy))
		visu.Renderer.DrawLine(int32(k-dx), int32(h+dy-radius), int32(k+dx), int32(h+dy-radius))
		visu.Renderer.DrawLine(int32(k-dx), int32(h-dy+radius), int32(k+dx), int32(h-dy+radius))
	}
	visu.Renderer.Present()
	display_player(ctx, visu, size_case, false)
}

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
	ctx.CurrentPlayer = uint8((rand.Intn(3-1) + 1))
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
	size_case := (display.H - (int32(ctx.NSize * 3))) / (int32(ctx.NSize + 1))
	size := int32((int32(ctx.NSize + 1)) * size_case)
	visu.Window, err = sdl.CreateWindow("Gomoku", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		size+(size/4), size, sdl.WINDOW_SHOWN)
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
	if visu.FontPlayer, err = ttf.OpenFont("fonts/Quicksand-VariableFont_wght.ttf", 20); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	defer visu.FontPlayer.Close()
	if visu.FontMsg, err = ttf.OpenFont("fonts/Rubik-Regular.ttf", 12); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	defer visu.FontMsg.Close()
	visu.Renderer, err = sdl.CreateRenderer(visu.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer : %s\n", err)
		panic(err)
	}
	defer visu.Renderer.Destroy()
	// 0xe2c473
	// Initialize window with color and lines
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: 0, Y: 0, W: size + (size / 4), H: size})
	visu.Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: size + (size / 4), H: size})
	visu.Renderer.SetDrawColor(0, 0, 0, 200)
	for line := 0; uint8(line) < ctx.NSize+2; line++ {
		visu.Renderer.DrawLine(0, int32(line*int(size_case)), size, int32(line*int(size_case)))
		visu.Renderer.DrawLine(int32(line*int(size_case)), 0, int32(line*int(size_case)), size)
	}
	visu.Renderer.Present()
	display_player(&ctx, &visu, int(size_case), true)
	running := true
	// Loop de jeu
	for running {
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
				if t.State == sdl.PRESSED {
					if err != nil {
						panic(err)
					}
					// Trouver intersection la plus proche
					h_mouse := float64(t.Y - 5)
					k_mouse := float64(t.X - 5)
					case_x := math.Round(k_mouse / float64(size_case))
					case_y := math.Round(h_mouse / float64(size_case))
					if (case_x > 0 && uint8(case_x) <= ctx.NSize) && (case_y > 0 && uint8(case_y) <= ctx.NSize) {
						if ctx.Goban[int(case_y-1)][int(case_x-1)] == 0 {
							ctx.Goban[int(case_y-1)][int(case_x-1)] = s.Tnumber(ctx.CurrentPlayer)
							fmt.Println(ctx)
							display_message(&visu, size, "", "")
							trace_stone(case_x, case_y, int(size_case), int(ctx.NSize), &ctx, &visu)
						} else {
							display_message(&visu, size, "Il y a déjà", "une pierre")
							sdl.Log("Il y a déjà une pierre")
						}
					} else {
						display_message(&visu, size, "En dehors", "du terrain")
						sdl.Log("En dehors du terrain")
					}
				}
			}
		}
	}
}
