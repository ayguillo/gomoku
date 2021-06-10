package structures

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type SVisu struct {
	Renderer         *sdl.Renderer
	FontMsg          *ttf.Font
	FontPlayer       *ttf.Font
	FontCounter      *ttf.Font
	TexturePlayer    *sdl.Texture
	TextureMessage1  *sdl.Texture
	TextureMessage2  *sdl.Texture
	TextureVictoryP1 *sdl.Texture
	TextureVictoryP2 *sdl.Texture
	Window           *sdl.Window
}

func (visu *SVisu) FillDefaults() {
	visu.TexturePlayer = nil
	visu.TextureMessage1 = nil
	visu.TextureMessage2 = nil
}
