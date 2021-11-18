package menu

import (
	"fmt"
	s "gomoku/structures"
	"math"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func fillCircle(k float64, h float64, visu *s.SVisu, radius float64) {
	for dy := 0.0; dy <= radius; dy += 0.075 {
		dx := math.Round(math.Sqrt((2.0 * radius * dy) - (dy * dy)))
		visu.Renderer.DrawLine(int32(k-dx), int32(h+dy-radius), int32(k+dx), int32(h+dy-radius))
		visu.Renderer.DrawLine(int32(k-dx), int32(h-dy+radius), int32(k+dx), int32(h-dy+radius))
	}
	visu.Renderer.Present()
}

func createCircle(k float64, h float64, visu *s.SVisu, radius int, r int, g int, b int, a int) {
	visu.Renderer.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
	fillCircle(k, h, visu, float64(radius))
}

func changeBot(versus int, hvb *sdl.Texture, hvh *sdl.Texture, bvb *sdl.Texture, visu *s.SVisu) {
	if versus == 0 {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
		visu.Renderer.Copy(hvb, nil, &sdl.Rect{X: 150, Y: 10, W: 170, H: 50})
		createCircle(35, 105, visu, 25, 83, 51, 237, 1)
		createCircle(35+60, 105, visu, 25, 0, 0, 0, 0)
		createCircle(35+120, 105, visu, 25, 0, 0, 0, 0)
	} else if versus == 1 {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
		visu.Renderer.Copy(hvh, nil, &sdl.Rect{X: 150, Y: 10, W: 220, H: 50})
		createCircle(35, 105, visu, 25, 0, 0, 0, 0)
		createCircle(35+60, 105, visu, 25, 83, 51, 237, 1)
		createCircle(35+120, 105, visu, 25, 0, 0, 0, 0)
	} else {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: 150, Y: 10, W: 230, H: 50})
		visu.Renderer.Copy(bvb, nil, &sdl.Rect{X: 150, Y: 10, W: 120, H: 50})
		createCircle(35, 105, visu, 25, 0, 0, 0, 0)
		createCircle(35+60, 105, visu, 25, 0, 0, 0, 0)
		createCircle(35+120, 105, visu, 25, 83, 51, 237, 1)
	}
	visu.Renderer.Present()
}

func changeThink(thinking int, visu *s.SVisu, thinkt1 *sdl.Texture, thinkt2 *sdl.Texture, thinkt3 *sdl.Texture, ctx s.SContext) {
	if thinking == 0 {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50 + 200, Y: ctx.Size/2 + 100, W: 50, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50 + 200, Y: ctx.Size/2 + 100, W: 50, H: 50})
		visu.Renderer.Copy(thinkt1, nil, &sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50 + 200, Y: ctx.Size/2 + 100, W: 50, H: 50})
		createCircle(float64(ctx.Size/2+ctx.Size/4-25), float64(ctx.Size/2+100)+95, visu, 25, 83, 51, 237, 1)
		createCircle(float64(ctx.Size/2+ctx.Size/4+35), float64(ctx.Size/2+100)+95, visu, 25, 0, 0, 0, 0)
		createCircle(float64(ctx.Size/2+ctx.Size/4+95), float64(ctx.Size/2+100)+95, visu, 25, 0, 0, 0, 0)
	} else if thinking == 1 {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50 + 200, Y: ctx.Size/2 + 100, W: 50, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50 + 200, Y: ctx.Size/2 + 100, W: 50, H: 50})
		visu.Renderer.Copy(thinkt2, nil, &sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50 + 200, Y: ctx.Size/2 + 100, W: 50, H: 50})
		createCircle(float64(ctx.Size/2+ctx.Size/4-25), float64(ctx.Size/2+100)+95, visu, 25, 0, 0, 0, 0)
		createCircle(float64(ctx.Size/2+ctx.Size/4+35), float64(ctx.Size/2+100)+95, visu, 25, 83, 51, 237, 1)
		createCircle(float64(ctx.Size/2+ctx.Size/4+95), float64(ctx.Size/2+100)+95, visu, 25, 0, 0, 0, 0)
	} else {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50 + 200, Y: ctx.Size/2 + 100, W: 50, H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50 + 200, Y: ctx.Size/2 + 100, W: 50, H: 50})
		visu.Renderer.Copy(thinkt3, nil, &sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50 + 200, Y: ctx.Size/2 + 100, W: 50, H: 50})
		createCircle(float64(ctx.Size/2+ctx.Size/4-25), float64(ctx.Size/2+100)+95, visu, 25, 0, 0, 0, 0)
		createCircle(float64(ctx.Size/2+ctx.Size/4+35), float64(ctx.Size/2+100)+95, visu, 25, 0, 0, 0, 0)
		createCircle(float64(ctx.Size/2+ctx.Size/4+95), float64(ctx.Size/2+100)+95, visu, 25, 83, 51, 237, 1)
	}
	visu.Renderer.Present()
}

func changeHelp(help bool, visu *s.SVisu, ctx s.SContext, yes *sdl.Texture, no *sdl.Texture) {
	createCircle(float64(ctx.Size/2+ctx.Size/4-30), float64(ctx.Size/2-ctx.Size/6)+90, visu, 25, 50, 205, 50, 255)
	createCircle(float64(ctx.Size/2+ctx.Size/4+30), float64(ctx.Size/2-ctx.Size/6)+90, visu, 25, 178, 34, 34, 255)
	if help == true {
		visu.Renderer.Copy(yes, nil, &sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 55, Y: ctx.Size/2 - ctx.Size/6 + 75, W: 50, H: 30})
	} else {
		visu.Renderer.Copy(no, nil, &sdl.Rect{X: ctx.Size/2 + ctx.Size/4 + 5, Y: ctx.Size/2 - ctx.Size/6 + 75, W: 50, H: 30})
	}
	visu.Renderer.Present()
}

func changeCapture(capture bool, visu *s.SVisu, ctx s.SContext, yes *sdl.Texture, no *sdl.Texture) {
	createCircle(float64(ctx.Size+ctx.Size/4), float64(ctx.Size/2-ctx.Size/6)+90, visu, 25, 50, 205, 50, 255)
	createCircle(float64(ctx.Size+ctx.Size/4+70), float64(ctx.Size/2-ctx.Size/6)+90, visu, 25, 178, 34, 34, 255)
	if capture == true {
		visu.Renderer.Copy(yes, nil, &sdl.Rect{X: ctx.Size + ctx.Size/4 - 25, Y: ctx.Size/2 - ctx.Size/6 + 75, W: 50, H: 30})
	} else {
		visu.Renderer.Copy(no, nil, &sdl.Rect{X: ctx.Size + ctx.Size/4 + 45, Y: ctx.Size/2 - ctx.Size/6 + 75, W: 50, H: 30})
	}
	visu.Renderer.Present()
}

func changeDifficulty(difficulty int, visu *s.SVisu, ctx s.SContext, ez *sdl.Texture, med *sdl.Texture, hard *sdl.Texture) {
	if difficulty == 0 {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: ctx.Size + ctx.Size/11 + (ctx.Size / 6) + 50, Y: 10, W: (ctx.Size / 6), H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: ctx.Size + ctx.Size/11 + (ctx.Size / 6) + 50, Y: 10, W: (ctx.Size / 6), H: 50})
		visu.Renderer.Copy(ez, nil, &sdl.Rect{X: ctx.Size + ctx.Size/11 + (ctx.Size / 6) + 50, Y: 10, W: (ctx.Size / 6), H: 50})
		createCircle(float64(ctx.Size+80), 105, visu, 25, 83, 51, 237, 1)
		createCircle(float64(ctx.Size+80)+60, 105, visu, 25, 0, 0, 0, 0)
		createCircle(float64(ctx.Size+80)+120, 105, visu, 25, 0, 0, 0, 0)
	} else if difficulty == 1 {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: ctx.Size + ctx.Size/11 + (ctx.Size / 6) + 50, Y: 10, W: (ctx.Size / 6), H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: ctx.Size + ctx.Size/11 + (ctx.Size / 6) + 50, Y: 10, W: (ctx.Size / 6), H: 50})
		visu.Renderer.Copy(med, nil, &sdl.Rect{X: ctx.Size + ctx.Size/11 + (ctx.Size / 6) + 50, Y: 10, W: (ctx.Size / 6), H: 50})
		createCircle(float64(ctx.Size+80), 105, visu, 25, 0, 0, 0, 0)
		createCircle(float64(ctx.Size+80)+60, 105, visu, 25, 83, 51, 237, 1)
		createCircle(float64(ctx.Size+80)+120, 105, visu, 25, 0, 0, 0, 0)
	} else {
		visu.Renderer.SetDrawColor(226, 196, 115, 255)
		visu.Renderer.DrawRect(&sdl.Rect{X: ctx.Size + ctx.Size/11 + (ctx.Size / 6) + 50, Y: 10, W: (ctx.Size / 6), H: 50})
		visu.Renderer.FillRect(&sdl.Rect{X: ctx.Size + ctx.Size/11 + (ctx.Size / 6) + 50, Y: 10, W: (ctx.Size / 6), H: 50})
		visu.Renderer.Copy(hard, nil, &sdl.Rect{X: ctx.Size + ctx.Size/11 + (ctx.Size / 6) + 50, Y: 10, W: (ctx.Size / 6), H: 50})
		createCircle(float64(ctx.Size+80), 105, visu, 25, 0, 0, 0, 0)
		createCircle(float64(ctx.Size+80)+60, 105, visu, 25, 0, 0, 0, 0)
		createCircle(float64(ctx.Size+80)+120, 105, visu, 25, 83, 51, 237, 1)
	}
	visu.Renderer.Present()
}

func changeDoubleThrees(double_threes bool, visu *s.SVisu, ctx s.SContext, yes *sdl.Texture, no *sdl.Texture) {
	// X: 10, Y: ctx.Size/2 - ctx.Size/6
	createCircle(35, float64(ctx.Size/2-ctx.Size/6)+90, visu, 25, 50, 205, 50, 255)
	createCircle(35+70, float64(ctx.Size/2-ctx.Size/6)+90, visu, 25, 178, 34, 34, 255)
	if double_threes == true {
		visu.Renderer.Copy(yes, nil, &sdl.Rect{X: 10, Y: ctx.Size/2 - ctx.Size/6 + 75, W: 50, H: 30})
	} else {
		visu.Renderer.Copy(no, nil, &sdl.Rect{X: 80, Y: ctx.Size/2 - ctx.Size/6 + 75, W: 50, H: 30})
	}
	visu.Renderer.Present()
}

func Menu(visu *s.SVisu, ctx s.SContext) (int, bool, bool, bool, bool, int, int) {
	versus, double_threes, capture, help, difficulty, thinking := 0, true, true, true, 1, 0
	// Window
	visu.Renderer.SetDrawColor(226, 196, 115, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: 0, Y: 0, W: ctx.Size + ctx.Size/2, H: ctx.Size + ((ctx.SizeCase) / 2)})
	visu.Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: ctx.Size + ctx.Size/2, H: ctx.Size + ((ctx.SizeCase) / 2)})
	// Button quit
	visu.Renderer.SetDrawColor(212, 66, 62, 255)
	visu.Renderer.DrawRect(&sdl.Rect{X: ctx.Size + 4 + ctx.Size/4, Y: ctx.Size - 50, W: (ctx.Size / 4), H: 50})
	visu.Renderer.FillRect(&sdl.Rect{X: ctx.Size + 4 + ctx.Size/4, Y: ctx.Size - 50, W: (ctx.Size / 4), H: 50})
	// Str buttons
	color := sdl.Color{R: 240, G: 228, B: 229, A: 255}
	bmp, err := visu.FontPlayer.RenderUTF8Solid("Play >", color)
	bmp2, err1 := visu.FontPlayer.RenderUTF8Solid("Players :", color)
	bmp3, err2 := visu.FontPlayer.RenderUTF8Solid("Human vs Bot", color)
	bmp4, err3 := visu.FontPlayer.RenderUTF8Solid("Human vs Human", color)
	bmp5, err4 := visu.FontPlayer.RenderUTF8Solid("Bot vs Bot", color)
	bmp6, err5 := visu.FontPlayer.RenderUTF8Solid("Help ?", color)
	bmp7, err6 := visu.FontPlayer.RenderUTF8Solid("Yes", color)
	bmp8, err7 := visu.FontPlayer.RenderUTF8Solid("No", color)
	bmp9, err8 := visu.FontPlayer.RenderUTF8Solid("Capture ?", color)
	bmp10, err9 := visu.FontPlayer.RenderUTF8Solid("Difficulty :", color)
	bmp11, err10 := visu.FontPlayer.RenderUTF8Solid("Easy", color)
	bmp12, err11 := visu.FontPlayer.RenderUTF8Solid("Medium", color)
	bmp13, err12 := visu.FontPlayer.RenderUTF8Solid("Hard", color)
	bmp14, err13 := visu.FontPlayer.RenderUTF8Solid("Double Threes ?", color)
	bmp15, err14 := visu.FontPlayer.RenderUTF8Solid("Thinking time", color)
	bmp16, err15 := visu.FontPlayer.RenderUTF8Solid("0.5s", color)
	bmp17, err16 := visu.FontPlayer.RenderUTF8Solid("1s", color)
	bmp18, err17 := visu.FontPlayer.RenderUTF8Solid("2s", color)
	if err != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil ||
		err7 != nil || err8 != nil || err9 != nil || err10 != nil || err11 != nil || err12 != nil || err13 != nil ||
		err14 != nil || err15 != nil || err16 != nil || err17 != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font\n")
		panic(err)
	}
	button, err := visu.Renderer.CreateTextureFromSurface(bmp)
	text_player, err1 := visu.Renderer.CreateTextureFromSurface(bmp2)
	hvb, err2 := visu.Renderer.CreateTextureFromSurface(bmp3)
	hvh, err3 := visu.Renderer.CreateTextureFromSurface(bmp4)
	bvb, err4 := visu.Renderer.CreateTextureFromSurface(bmp5)
	help_text, err5 := visu.Renderer.CreateTextureFromSurface(bmp6)
	yes, err6 := visu.Renderer.CreateTextureFromSurface(bmp7)
	no, err7 := visu.Renderer.CreateTextureFromSurface(bmp8)
	cap, err8 := visu.Renderer.CreateTextureFromSurface(bmp9)
	diff, err9 := visu.Renderer.CreateTextureFromSurface(bmp10)
	ez, err9 := visu.Renderer.CreateTextureFromSurface(bmp11)
	med, err9 := visu.Renderer.CreateTextureFromSurface(bmp12)
	hard, err9 := visu.Renderer.CreateTextureFromSurface(bmp13)
	dt, err10 := visu.Renderer.CreateTextureFromSurface(bmp14)
	thinkt, err11 := visu.Renderer.CreateTextureFromSurface(bmp15)
	thinkt1, err12 := visu.Renderer.CreateTextureFromSurface(bmp16)
	thinkt2, err13 := visu.Renderer.CreateTextureFromSurface(bmp17)
	thinkt3, err14 := visu.Renderer.CreateTextureFromSurface(bmp18)
	if err != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil ||
		err5 != nil || err6 != nil || err7 != nil || err8 != nil || err9 != nil ||
		err10 != nil || err11 != nil || err12 != nil || err13 != nil || err14 != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture font\n")
		panic(err)
	}
	defer button.Destroy()
	defer text_player.Destroy()
	defer hvb.Destroy()
	defer hvh.Destroy()
	defer bvb.Destroy()
	defer yes.Destroy()
	defer no.Destroy()
	defer cap.Destroy()
	defer diff.Destroy()
	defer dt.Destroy()
	defer thinkt.Destroy()
	defer thinkt1.Destroy()
	defer thinkt2.Destroy()
	defer thinkt3.Destroy()
	bmp.Free()
	bmp2.Free()
	bmp3.Free()
	bmp4.Free()
	bmp5.Free()
	bmp6.Free()
	bmp7.Free()
	bmp8.Free()
	bmp9.Free()
	bmp10.Free()
	bmp11.Free()
	bmp12.Free()
	bmp13.Free()
	bmp14.Free()
	bmp15.Free()
	bmp16.Free()
	bmp17.Free()
	bmp18.Free()
	visu.Renderer.Copy(button, nil, &sdl.Rect{X: ctx.Size + 4 + ctx.Size/4, Y: ctx.Size - 50, W: (ctx.Size / 4), H: 50})
	visu.Renderer.Copy(text_player, nil, &sdl.Rect{X: 10, Y: 10, W: 130, H: 50})
	visu.Renderer.Copy(help_text, nil, &sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50, Y: ctx.Size/2 - ctx.Size/6, W: 130, H: 50})
	visu.Renderer.Copy(cap, nil, &sdl.Rect{X: ctx.Size + ctx.Size/4 - 25, Y: ctx.Size/2 - ctx.Size/6, W: 190, H: 50})
	visu.Renderer.Copy(diff, nil, &sdl.Rect{X: ctx.Size + 50, Y: 10, W: (ctx.Size / 4), H: 50})
	visu.Renderer.Copy(dt, nil, &sdl.Rect{X: 10, Y: ctx.Size/2 - ctx.Size/6, W: 190, H: 50})
	visu.Renderer.Copy(thinkt, nil, &sdl.Rect{X: ctx.Size/2 + ctx.Size/4 - 50, Y: ctx.Size/2 + 100, W: 190, H: 50})
	visu.Renderer.Present()
	running, end := true, false
	changeBot(versus, hvb, hvh, bvb, visu)
	changeHelp(help, visu, ctx, yes, no)
	changeCapture(capture, visu, ctx, yes, no)
	changeDifficulty(difficulty, visu, ctx, ez, med, hard)
	changeDoubleThrees(double_threes, visu, ctx, yes, no)
	changeThink(thinking, visu, thinkt1, thinkt2, thinkt3, ctx)
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				end = true
			case *sdl.KeyboardEvent:
				if t.State == sdl.PRESSED && t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
					end = true
				}
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED {
					if t.Y >= ctx.Size-50 && t.X >= ctx.Size+4+ctx.Size/4 {
						running = false
					}
					if (t.Y >= 80 && t.Y <= 80+50) && (t.X >= 10 && t.X <= 60) {
						versus = 0
						changeBot(versus, hvb, hvh, bvb, visu)
					}
					if (t.Y >= 80 && t.Y <= 80+50) && (t.X >= 70 && t.X <= 120) {
						versus = 1
						changeBot(versus, hvb, hvh, bvb, visu)
					}
					if (t.Y >= 80 && t.Y <= 80+50) && (t.X >= 130 && t.X <= 180) {
						versus = 2
						changeBot(versus, hvb, hvh, bvb, visu)
					}
					if (t.Y >= ctx.Size/2+ctx.Size/4-50 && t.Y <= ctx.Size/2+ctx.Size/4) && (t.X >= ctx.Size/2+170 && t.X <= ctx.Size/2+220) {
						thinking = 0
						changeThink(thinking, visu, thinkt1, thinkt2, thinkt3, ctx)
					}
					if (t.Y >= ctx.Size/2+ctx.Size/4-50 && t.Y <= ctx.Size/2+ctx.Size/4) && (t.X >= ctx.Size/2+ctx.Size/4+15 && t.X <= ctx.Size/2+ctx.Size/4+65) {
						thinking = 1
						changeThink(thinking, visu, thinkt1, thinkt2, thinkt3, ctx)
					}
					if (t.Y >= ctx.Size/2+ctx.Size/4-50 && t.Y <= ctx.Size/2+ctx.Size/4) && (t.X >= ctx.Size/2+ctx.Size/4+70 && t.X <= ctx.Size/2+ctx.Size/4+120) {
						thinking = 2
						changeThink(thinking, visu, thinkt1, thinkt2, thinkt3, ctx)
					}
					if (t.Y >= ctx.Size/2-ctx.Size/6+65 && t.Y <= ctx.Size/2-ctx.Size/6+115) && (t.X >= ctx.Size/2+ctx.Size/4-55 && t.X <= ctx.Size/2+ctx.Size/4-5) {
						help = true
						changeHelp(help, visu, ctx, yes, no)
					}
					if (t.Y >= ctx.Size/2-ctx.Size/6+65 && t.Y <= ctx.Size/2-ctx.Size/6+115) && (t.X >= ctx.Size/2+ctx.Size/4+5 && t.X <= ctx.Size/2+ctx.Size/4+55) {
						help = false
						changeHelp(help, visu, ctx, yes, no)
					}
					if (t.Y >= ctx.Size/2-ctx.Size/6+65 && t.Y <= ctx.Size/2-ctx.Size/6+115) && (t.X >= ctx.Size+ctx.Size/4-25 && t.X <= ctx.Size+ctx.Size/4+25) {
						capture = true
						changeCapture(capture, visu, ctx, yes, no)
					}
					if (t.Y >= ctx.Size/2-ctx.Size/6+65 && t.Y <= ctx.Size/2-ctx.Size/6+115) && (t.X >= ctx.Size+ctx.Size/4+45 && t.X <= ctx.Size+ctx.Size/4+90) {
						capture = false
						changeCapture(capture, visu, ctx, yes, no)
					}
					if (t.Y >= 80 && t.Y <= 80+50) && (t.X >= ctx.Size+80-25 && t.X <= ctx.Size+80+25) {
						difficulty = 0
						changeDifficulty(difficulty, visu, ctx, ez, med, hard)
					}
					if (t.Y >= 80 && t.Y <= 80+50) && (t.X >= ctx.Size+80+35 && t.X <= ctx.Size+80+85) {
						difficulty = 1
						changeDifficulty(difficulty, visu, ctx, ez, med, hard)
					}
					if (t.Y >= 80 && t.Y <= 80+50) && (t.X >= ctx.Size+80+95 && t.X <= ctx.Size+80+145) {
						difficulty = 2
						changeDifficulty(difficulty, visu, ctx, ez, med, hard)
					}
					if (t.Y >= ctx.Size/2-ctx.Size/6+65 && t.Y <= ctx.Size/2-ctx.Size/6+115) && (t.X >= 10 && t.X <= 60) {
						double_threes = true
						changeDoubleThrees(double_threes, visu, ctx, yes, no)
					}
					if (t.Y >= ctx.Size/2-ctx.Size/6+65 && t.Y <= ctx.Size/2-ctx.Size/6+115) && (t.X >= 80 && t.X <= 130) {
						double_threes = false
						changeDoubleThrees(double_threes, visu, ctx, yes, no)
					}
				}
			}

		}
	}
	return versus, double_threes, capture, help, end, difficulty, thinking
}
