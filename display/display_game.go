package display

import (
	s "gomoku/structures"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func TraceStone(case_x float64, case_y float64, ctx *s.SContext, visu *s.SVisu, cover bool) {
	case_x++
	case_y++
	k := case_x * float64(ctx.SizeCase)
	h := case_y * float64(ctx.SizeCase)
	radius := 10.0
	if ctx.CurrentPlayer == 1 && cover == false {
		visu.Renderer.SetDrawColor(240, 228, 229, 255)
	} else if ctx.CurrentPlayer == 2 && cover == false {
		visu.Renderer.SetDrawColor(35, 33, 33, 255)
	} else if cover == true {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
	}
	// Draw circle
	if cover == false {
		visu.Renderer.DrawLine(int32(k-radius), int32(h), int32(k+radius+1), int32(h))
	}
	visu.Renderer.Present()
	for dy := 0.0; dy < radius; dy += 1.0 {
		dx := math.Round(math.Sqrt((2.0 * radius * dy) - (dy * dy)))
		visu.Renderer.DrawLine(int32(k-dx), int32(h+dy-radius), int32(k+dx), int32(h+dy-radius))
		visu.Renderer.DrawLine(int32(k-dx), int32(h-dy+radius), int32(k+dx), int32(h-dy+radius))
	}
	if cover == true {
		visu.Renderer.SetDrawColor(0, 0, 0, 200)
		visu.Renderer.DrawLine(int32(k-radius), int32(h), int32(k+radius+1), int32(h))
		visu.Renderer.DrawLine(int32(k), int32(h-radius), int32(k), int32(h+radius))
	}
	visu.Renderer.Present()
}

func TraceGoban(visu *s.SVisu, ctx s.SContext) {
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: 0, Y: 0, W: ctx.Size + (ctx.Size / 4), H: ctx.Size})
	visu.Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: ctx.Size + (ctx.Size / 4), H: ctx.Size})
	visu.Renderer.SetDrawColor(0, 0, 0, 200)
	for line := 0; uint8(line) < ctx.NSize+2; line++ {
		visu.Renderer.DrawLine(0, int32(line*int(ctx.SizeCase)), ctx.Size, int32(line*int(ctx.SizeCase)))
		visu.Renderer.DrawLine(int32(line*int(ctx.SizeCase)), 0, int32(line*int(ctx.SizeCase)), ctx.Size)
	}
	visu.Renderer.Present()
}
