package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func trace_stone(case_x float64, case_y float64, size_case int, nb_case int, renderer *sdl.Renderer) {
	k := case_x * float64(size_case)
	h := case_y * float64(size_case)

	radius := 10.0
	renderer.SetDrawColor(226, 228, 229, 255)
	renderer.DrawLine(int32(k-radius), int32(h), int32(k+radius), int32(h))
	// Draw circle
	renderer.Present()
	for dy := 1.0; dy < radius; dy += 1.0 {
		dx := math.Sqrt((2.0 * radius * dy) - (dy * dy))
		renderer.DrawLine(int32(k-dx), int32(h+dy-radius), int32(k+dx), int32(h+dy-radius))
		renderer.DrawLine(int32(k-dx), int32(h-dy+radius), int32(k+dx), int32(h-dy+radius))
		renderer.Present()
	}
	sdl.Delay(16)
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	// display, _ := sdl.GetDesktopDisplayMode(0)
	// fmt.Println(display.W)
	// fmt.Println(display.H)
	nb_case := 19
	size_case := 35
	size := int32((nb_case + 1) * size_case)
	window, err := sdl.CreateWindow("Gomoku", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		size, size, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	running := true

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	// 0xe2c473
	// Initialize window with color and lines
	renderer.SetDrawColor(226, 196, 115, 255)
	renderer.DrawRect(&sdl.Rect{X: 0, Y: 0, W: size, H: size})
	renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: size, H: size})
	renderer.SetDrawColor(0, 0, 0, 200)
	for line := 0; line < nb_case+1; line++ {
		renderer.DrawLine(0, int32(line*size_case), size, int32(line*size_case))
		renderer.DrawLine(int32(line*size_case), 0, int32(line*size_case), size)
	}
	renderer.Present()
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
					if case_x != 0 && case_y != 0 && int(case_x) != nb_case && int(case_y) != nb_case {
						trace_stone(case_x, case_y, size_case, nb_case, renderer)
					}
				}
			}
		}
	}
}
