package structures

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type SVisu struct {
	Renderer           *sdl.Renderer
	FontMsg            *ttf.Font
	FontPlayer         *ttf.Font
	FontCounter        *ttf.Font
	TexturePlayer      *sdl.Texture
	TextureMessage1    *sdl.Texture
	TextureMessage2    *sdl.Texture
	TextureVictoryP1   *sdl.Texture
	TextureVictoryP2   *sdl.Texture
	TextureCaptureP1   *sdl.Texture
	TextureCaptureP2   *sdl.Texture
	TextureNotationX   *sdl.Texture
	TextureNotationY   *sdl.Texture
	TextureMessageTime *sdl.Texture
	Window             *sdl.Window
}

func (visu *SVisu) FillDefaults() {
	visu.TexturePlayer = nil
	visu.TextureMessage1 = nil
	visu.TextureMessage2 = nil
	visu.TextureVictoryP1 = nil
	visu.TextureVictoryP2 = nil
	visu.TextureCaptureP1 = nil
	visu.TextureCaptureP2 = nil
	visu.TextureNotationX = nil
	visu.TextureNotationY = nil
	visu.TextureMessageTime = nil
}
