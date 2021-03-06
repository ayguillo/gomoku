package display

import (
	"fmt"
	"os"
	"strconv"

	s "gomoku/structures"

	"github.com/veandco/go-sdl2/sdl"
)

func DisplayTime(visu *s.SVisu, time string, size int32, player uint8) {
	if visu.TextureMessageTime != nil {
		visu.TextureMessageTime.Destroy()
	}
	color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	bmp, err := visu.FontMsg.RenderUTF8Solid(time, color)
	if err == nil {
		visu.TextureMessageTime, err = visu.Renderer.CreateTextureFromSurface(bmp)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to renderer font: %s\n", err)
		panic(err)
	}
	bmp.Free()
	if player == 1 {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: 4, Y: size/4 + 50, W: (size / 4) - 10, H: 30})
		visu.Renderer.FillRect(&sdl.Rect{X: 4, Y: size/4 + 50, W: (size / 4) - 10, H: 30})
		visu.Renderer.Present()
		visu.Renderer.Copy(visu.TextureMessageTime, nil, &sdl.Rect{X: 4, Y: size/4 + 50, W: (size / 4) - 10, H: 30})
	} else {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: size + 4 + size/4, Y: size/4 + 50, W: (size / 4) - 10, H: 30})
		visu.Renderer.FillRect(&sdl.Rect{X: size + 4 + size/4, Y: size/4 + 50, W: (size / 4) - 10, H: 30})
		visu.Renderer.Present()
		visu.Renderer.Copy(visu.TextureMessageTime, nil, &sdl.Rect{X: size + 4 + size/4, Y: size/4 + 50, W: (size / 4) - 10, H: 30})
		visu.Renderer.Present()
	}
	visu.Renderer.Present()
}

func DisplayMessage(visu *s.SVisu, size int32, line1 string, line2 string, ctx s.SContext) {
	if visu.TextureMessage1 != nil {
		visu.TextureMessage1.Destroy()
		visu.TextureMessage2.Destroy()
	}
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: 4, Y: size - 100, W: (size / 4) - 10, H: 100})
	visu.Renderer.FillRect(&sdl.Rect{X: 4, Y: size - 100, W: (size / 4) - 10, H: 100})
	visu.Renderer.DrawRect(&sdl.Rect{X: size + 4 + size/4, Y: size - 100, W: (size / 4) - 10, H: 100})
	visu.Renderer.FillRect(&sdl.Rect{X: size + 4 + size/4, Y: size - 100, W: (size / 4) - 10, H: 100})
	visu.Renderer.Present()
	if line1 != "" && line2 != "" {
		color := sdl.Color{R: 212, G: 66, B: 62, A: 255}
		bmp, err := visu.FontMsg.RenderUTF8Solid(line1, color)
		bmp2, err2 := visu.FontMsg.RenderUTF8Solid(line2, color)
		if err != nil || err2 != nil {
			fmt.Fprintf(os.Stderr, "Failed to renderer font: %s %s\n", err, err2)
			panic(err)
		}
		visu.TextureMessage1, err = visu.Renderer.CreateTextureFromSurface(bmp)
		visu.TextureMessage2, err2 = visu.Renderer.CreateTextureFromSurface(bmp2)
		bmp.Free()
		bmp2.Free()
		if err != nil || err2 != nil {
			fmt.Fprintf(os.Stderr, "Failed to create texture font: %s %s\n", err, err2)
			panic(err)
		}
		visu.Renderer.Present()
		if ctx.CurrentPlayer == 1 {
			visu.Renderer.Copy(visu.TextureMessage1, nil, &sdl.Rect{X: 4, Y: size - 100, W: (size / 4) - 10, H: 50})
			visu.Renderer.Copy(visu.TextureMessage2, nil, &sdl.Rect{X: 4, Y: size - 50, W: (size / 4) - 10, H: 50})
		} else if ctx.CurrentPlayer == 2 {
			visu.Renderer.Copy(visu.TextureMessage1, nil, &sdl.Rect{X: size + 4 + size/4, Y: size - 100, W: (size / 4) - 10, H: 50})
			visu.Renderer.Copy(visu.TextureMessage2, nil, &sdl.Rect{X: size + 4 + size/4, Y: size - 50, W: (size / 4) - 10, H: 50})
		}
		visu.Renderer.Present()
	} else if line1 != "" && line2 == "" {
		color := sdl.Color{R: 212, G: 66, B: 62, A: 255}
		bmp, err := visu.FontMsg.RenderUTF8Solid(line1, color)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to renderer font: %s\n", err)
			panic(err)
		}
		visu.TextureMessage1, err = visu.Renderer.CreateTextureFromSurface(bmp)
		bmp.Free()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err)
			panic(err)
		}
		visu.Renderer.Present()
		if ctx.CurrentPlayer == 1 {
			visu.Renderer.Copy(visu.TextureMessage1, nil, &sdl.Rect{X: 4, Y: size - 100, W: (size / 4) - 10, H: 50})
		} else if ctx.CurrentPlayer == 2 {
			visu.Renderer.Copy(visu.TextureMessage1, nil, &sdl.Rect{X: size + 4 + size/4, Y: size - 100, W: (size / 4) - 10, H: 50})
		}
		visu.Renderer.Present()
	}
}

func DisplayPlayer(ctx *s.SContext, visu *s.SVisu, current bool) {
	var text string
	var color sdl.Color
	size := ctx.Size
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
	if ctx.CurrentPlayer == 2 {
		visu.Renderer.DrawRect(&sdl.Rect{X: size + 4 + size/4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: size + 4 + size/4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.Present()
		visu.Renderer.Copy(visu.TexturePlayer, nil, &sdl.Rect{X: size + 4 + size/4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: 4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: 4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.Present()
	} else if ctx.CurrentPlayer == 1 {
		visu.Renderer.DrawRect(&sdl.Rect{X: 4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: 4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.Present()
		visu.Renderer.Copy(visu.TexturePlayer, nil, &sdl.Rect{X: 4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: size + 4 + size/4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: size + 4 + size/4, Y: 0, W: (size / 4) - 10, H: 50})
		visu.Renderer.Present()
	}
}

func DisplayEquality(visu *s.SVisu, ctx s.SContext) {
	size := ctx.Size
	color := sdl.Color{R: 35, G: 33, B: 33, A: 255}

	if visu.TexturePlayer != nil {
		visu.TexturePlayer.Destroy()
	}
	bmp, err := visu.FontPlayer.RenderUTF8Solid("Equality", color)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err)
		panic(err)
	}
	visu.TexturePlayer, err = visu.Renderer.CreateTextureFromSurface(bmp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err)
		panic(err)
	}
	bmp.Free()
	visu.Renderer.Copy(visu.TexturePlayer, nil, &sdl.Rect{X: 4, Y: 54, W: (size / 4) - 10, H: 50})
	visu.Renderer.Present()
}

func DisplayVictory(visu *s.SVisu, ctx s.SContext) {
	var color sdl.Color
	size := ctx.Size
	if ctx.CurrentPlayer == 1 {
		color = sdl.Color{R: 240, G: 228, B: 229, A: 255}
	} else {
		color = sdl.Color{R: 35, G: 33, B: 33, A: 255}
	}
	if visu.TexturePlayer != nil {
		visu.TexturePlayer.Destroy()
	}
	bmp, err := visu.FontPlayer.RenderUTF8Solid("Victory", color)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err)
		panic(err)
	}
	visu.TexturePlayer, err = visu.Renderer.CreateTextureFromSurface(bmp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err)
		panic(err)
	}
	bmp.Free()
	if ctx.CurrentPlayer == 1 {
		visu.Renderer.Copy(visu.TexturePlayer, nil, &sdl.Rect{X: 4, Y: 54, W: (size / 4) - 10, H: 50})
	} else if ctx.CurrentPlayer == 2 {
		visu.Renderer.Copy(visu.TexturePlayer, nil, &sdl.Rect{X: size + 4 + size/4, Y: 54, W: (size / 4) - 10, H: 50})
	}
	visu.Renderer.Present()
}

func DisplayCounter(ctx s.SContext, visu *s.SVisu) {
	size := ctx.Size
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	// Player 1
	visu.Renderer.DrawRect(&sdl.Rect{X: 4, Y: size / 2, W: (size / 4) - 10, H: 50})
	visu.Renderer.FillRect(&sdl.Rect{X: 4, Y: size / 2, W: (size / 4) - 10, H: 50})
	// Player 2
	visu.Renderer.DrawRect(&sdl.Rect{X: size + 4 + size/4, Y: size / 2, W: (size / 4) - 10, H: 50})
	visu.Renderer.FillRect(&sdl.Rect{X: size + 4 + size/4, Y: size / 2, W: (size / 4) - 10, H: 50})
	visu.Renderer.Present()
	if visu.TextureVictoryP1 != nil {
		visu.TextureVictoryP1.Destroy()
		visu.TextureVictoryP2.Destroy()
	}
	victoryP1 := "Nb win Player 1 : " + strconv.Itoa(ctx.NbVictoryP1)
	victoryP2 := "Nb win Player 2 : " + strconv.Itoa(ctx.NbVictoryP2)
	colorP1 := sdl.Color{R: 240, G: 228, B: 229, A: 255}
	colorP2 := sdl.Color{R: 35, G: 33, B: 33, A: 255}
	bmp_p1, err := visu.FontMsg.RenderUTF8Solid(victoryP1, colorP1)
	bmp_p2, err2 := visu.FontMsg.RenderUTF8Solid(victoryP2, colorP2)
	if err != nil || err2 != nil {
		fmt.Fprintf(os.Stderr, "Failed to renderer font: %s\n", err)
		panic(err)
	}
	visu.TextureVictoryP1, err = visu.Renderer.CreateTextureFromSurface(bmp_p1)
	visu.TextureVictoryP2, err2 = visu.Renderer.CreateTextureFromSurface(bmp_p2)
	if err != nil || err2 != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err)
		panic(err)
	}
	bmp_p1.Free()
	bmp_p2.Free()

	visu.Renderer.Copy(visu.TextureVictoryP1, nil, &sdl.Rect{X: 4, Y: size / 2, W: (size / 4) - 10, H: 50})
	visu.Renderer.Copy(visu.TextureVictoryP2, nil, &sdl.Rect{X: size + 4 + size/4, Y: size / 2, W: (size / 4) - 10, H: 50})

	visu.Renderer.Present()
}

func DisplayCapture(ctx s.SContext, visu *s.SVisu) {
	colorP1 := sdl.Color{R: 240, G: 228, B: 229, A: 255}
	colorP2 := sdl.Color{R: 35, G: 33, B: 33, A: 255}
	size := ctx.Size
	if ctx.NbCaptureP1 != 0 {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: 4, Y: size / 4, W: (size / 4) - 10, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: 4, Y: size / 4, W: (size / 4) - 10, H: 50})
		captureP1 := "Nb Capture Player 1 : " + strconv.Itoa(ctx.NbCaptureP1)
		bmp_p1, err := visu.FontMsg.RenderUTF8Solid(captureP1, colorP1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to renderer font: %s\n", err)
			panic(err)
		}
		visu.TextureCaptureP1, err = visu.Renderer.CreateTextureFromSurface(bmp_p1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err)
			panic(err)
		}
		bmp_p1.Free()
		visu.Renderer.Copy(visu.TextureCaptureP1, nil, &sdl.Rect{X: 4, Y: size / 4, W: (size / 4) - 10, H: 50})
	}
	if ctx.NbCaptureP2 != 0 {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: size + 4 + size/4, Y: size / 4, W: (size / 4) - 10, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: size + 4 + size/4, Y: size / 4, W: (size / 4) - 10, H: 50})
		captureP2 := "Nb Capture Player 2 : " + strconv.Itoa(ctx.NbCaptureP2)
		bmp_p2, err2 := visu.FontMsg.RenderUTF8Solid(captureP2, colorP2)
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "Failed to renderer font: %s\n", err2)
			panic(err2)
		}
		visu.TextureCaptureP2, err2 = visu.Renderer.CreateTextureFromSurface(bmp_p2)
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "Failed to create texture font: %s\n", err2)
			panic(err2)
		}
		bmp_p2.Free()
		visu.Renderer.Copy(visu.TextureCaptureP2, nil, &sdl.Rect{X: size + 4 + size/4, Y: size / 4, W: (size / 4) - 10, H: 50})
	}
	if ctx.NbCaptureP1 != 0 || ctx.NbCaptureP2 != 0 {
		visu.Renderer.Present()
	}
}
