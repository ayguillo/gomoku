package display

import (
	s "gomoku/structures"

	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func fillCircle(k float64, h float64, visu *s.SVisu, radius float64) {
	for dy := 0.0; dy < radius; dy += 0.075 {
		dx := math.Round(math.Sqrt((2.0 * radius * dy) - (dy * dy)))
		visu.Renderer.DrawLine(int32(k-dx), int32(h+dy-radius), int32(k+dx), int32(h+dy-radius))
		visu.Renderer.DrawLine(int32(k-dx), int32(h-dy+radius), int32(k+dx), int32(h-dy+radius))
	}
	visu.Renderer.Present()
}

func drawCircle(k float64, h float64, visu *s.SVisu, radius int) {
	delta := 0.01 - float64(radius)
	x := float64(radius) - 0.5
	y := 0.5
	cx := k - 0.5
	cy := h - 0.5

	for x >= y {
		visu.Renderer.DrawPoint(int32(cx+x), int32(cy+y))
		visu.Renderer.DrawPoint(int32(cx+y), int32(cy+x))
		if x != 0 {
			visu.Renderer.DrawPoint(int32(cx-x), int32(cy+y))
			visu.Renderer.DrawPoint(int32(cx-y), int32(cy-x))
		}

		if y != 0 {
			visu.Renderer.DrawPoint(int32(cx+x), int32(cy-y))
			visu.Renderer.DrawPoint(int32(cx-y), int32(cy+x))
		}

		if x != 0 && y != 0 {
			visu.Renderer.DrawPoint(int32(cx-x), int32(cy-y))
			visu.Renderer.DrawPoint(int32(cx-y), int32(cy-x))
		}
		delta += y
		y++
		delta += y

		if delta >= 0 {
			x--
			delta -= x
			delta -= x
		}
		visu.Renderer.DrawPoint(int32(x), int32(y))
	}
	visu.Renderer.Present()
}

func createCircle(k float64, h float64, visu *s.SVisu, radius int, r int, g int, b int, a int) {
	visu.Renderer.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
	drawCircle(k, h, visu, radius)
	fillCircle(k, h, visu, float64(radius))
}

func TraceStone(case_x float64, case_y float64, ctx *s.SContext, visu *s.SVisu, color [4]uint8, cover bool) {
	case_x++
	case_y++
	k := case_x*float64(ctx.SizeCase) + float64(ctx.Size/4)
	h := case_y * float64(ctx.SizeCase)
	radius := 10.0
	visu.Renderer.SetDrawColor(color[0], color[1], color[2], color[3])
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
		if (case_x == 4 && case_y == 4) || (case_x == float64(ctx.NSize)-3.0 && case_y == 4) ||
			(case_x == 4 && case_y == float64(ctx.NSize)-3.0) ||
			(case_x == float64(ctx.NSize)-3.0 && case_y == float64(ctx.NSize)-3.0) {
			createCircle(k, h, visu, 5, 0, 0, 0, 200)
		}
	}
	visu.Renderer.Present()
}

func TraceGoban(visu *s.SVisu, ctx s.SContext) {
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: 0, Y: 0, W: ctx.Size + ctx.Size/2, H: ctx.Size + ((ctx.SizeCase) / 2)})
	visu.Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: ctx.Size + ctx.Size/2, H: ctx.Size + ((ctx.SizeCase) / 2)})
	visu.Renderer.SetDrawColor(0, 0, 0, 200)
	color := sdl.Color{R: 0, G: 0, B: 0, A: 200}
	for line := 0; uint8(line) < ctx.NSize+2; line++ {
		if line > 0 && line <= int(ctx.NSize) {
			if visu.TextureNotationX != nil {
				visu.TextureNotationX.Destroy()
			}
			if visu.TextureNotationY != nil {
				visu.TextureNotationY.Destroy()
			}
			visu.Renderer.DrawLine((ctx.Size/4 + ctx.SizeCase), int32(line*int(ctx.SizeCase)), ctx.Size+ctx.Size/4-ctx.SizeCase, int32(line*int(ctx.SizeCase)))
			visu.Renderer.DrawLine(int32(line*int(ctx.SizeCase))+ctx.Size/4, ctx.SizeCase, int32(line*int(ctx.SizeCase))+ctx.Size/4, ctx.Size-ctx.SizeCase)
			visu.Renderer.Present()
			bmpX, errX := visu.FontMsg.RenderUTF8Solid(ctx.MapX[line], color)
			bmpY, errY := visu.FontMsg.RenderUTF8Solid(strconv.Itoa(int(ctx.NSize)-line+1), color)
			time.Sleep(50 * time.Millisecond)
			if errX != nil || errY != nil {
				fmt.Fprintf(os.Stderr, "Failed to renderer font: %s\n", errX)
				panic(errX)
				panic(errY)
			}
			visu.TextureNotationX, errX = visu.Renderer.CreateTextureFromSurface(bmpX)
			visu.TextureNotationY, errY = visu.Renderer.CreateTextureFromSurface(bmpY)
			if errX != nil || errY != nil {
				fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", errX)
				panic(errX)
				panic(errY)
			}
			bmpX.Free()
			bmpY.Free()
			visu.Renderer.Copy(visu.TextureNotationX, nil, &sdl.Rect{X: int32(line*int(ctx.SizeCase)) + ctx.Size/4 - (ctx.SizeCase / 12), Y: ctx.Size - ctx.SizeCase, W: ctx.SizeCase / 5, H: ctx.SizeCase / 2})
			visu.Renderer.Copy(visu.TextureNotationX, nil, &sdl.Rect{X: int32(line*int(ctx.SizeCase)) + ctx.Size/4 - (ctx.SizeCase / 12), Y: ctx.SizeCase / 2, W: ctx.SizeCase / 5, H: ctx.SizeCase / 2})
			visu.Renderer.Copy(visu.TextureNotationY, nil, &sdl.Rect{X: ctx.Size/4 + ctx.SizeCase/2, Y: int32(line*int(ctx.SizeCase)) - ctx.SizeCase/4, W: ctx.SizeCase / 5, H: ctx.SizeCase / 2})
			visu.Renderer.Copy(visu.TextureNotationY, nil, &sdl.Rect{X: ctx.Size + ctx.Size/4 - ctx.SizeCase/2, Y: int32(line*int(ctx.SizeCase)) - ctx.SizeCase/4, W: ctx.SizeCase / 5, H: ctx.SizeCase / 2})
		}
	}
	if ctx.NSize == 19 {
		x, y := 4.0, 4.0
		k := x*float64(ctx.SizeCase) + float64(ctx.Size/4)
		h := y * float64(ctx.SizeCase)
		createCircle(k, h, visu, 5, 0, 0, 0, 200)
		x, y = float64(ctx.NSize-3.0), 4.0
		k = x*float64(ctx.SizeCase) + float64(ctx.Size/4)
		h = y * float64(ctx.SizeCase)
		createCircle(k, h, visu, 5, 0, 0, 0, 200)
		x, y = 4.0, float64(ctx.NSize)-3.0
		k = x*float64(ctx.SizeCase) + float64(ctx.Size/4)
		h = y * float64(ctx.SizeCase)
		createCircle(k, h, visu, 5, 0, 0, 0, 200)
		x, y = float64(ctx.NSize)-3.0, float64(ctx.NSize)-3.0
		k = x*float64(ctx.SizeCase) + float64(ctx.Size/4)
		h = y * float64(ctx.SizeCase)
		createCircle(k, h, visu, 5, 0, 0, 0, 200)
	}
	visu.Renderer.Present()
}
