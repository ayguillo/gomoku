package display

import (
	s "gomoku/structures"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func TraceStone(case_x float64, case_y float64, size_case int, ctx *s.SContext, visu *s.SVisu) {
	k := case_x * float64(size_case)
	h := case_y * float64(size_case)
	radius := 10.0

	if ctx.CurrentPlayer == 1 {
		visu.Renderer.SetDrawColor(240, 228, 229, 255)
	} else {
		visu.Renderer.SetDrawColor(35, 33, 33, 255)
	}
	// Draw circle
	visu.Renderer.DrawLine(int32(k-radius), int32(h), int32(k+radius), int32(h))
	visu.Renderer.Present()
	for dy := 0.0; dy < radius; dy += 1.0 {
		dx := math.Round(math.Sqrt((2.0 * radius * dy) - (dy * dy)))
		visu.Renderer.DrawLine(int32(k-dx), int32(h+dy-radius), int32(k+dx), int32(h+dy-radius))
		visu.Renderer.DrawLine(int32(k-dx), int32(h-dy+radius), int32(k+dx), int32(h-dy+radius))
	}
	visu.Renderer.Present()
}

func TraceGoban(visu *s.SVisu, ctx s.SContext, size int32, size_case int) {
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: 0, Y: 0, W: size + (size / 4), H: size})
	visu.Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: size + (size / 4), H: size})
	visu.Renderer.SetDrawColor(0, 0, 0, 200)
	for line := 0; uint8(line) < ctx.NSize+2; line++ {
		visu.Renderer.DrawLine(0, int32(line*size_case), size, int32(line*size_case))
		visu.Renderer.DrawLine(int32(line*size_case), 0, int32(line*size_case), size)
	}
	visu.Renderer.Present()
}
