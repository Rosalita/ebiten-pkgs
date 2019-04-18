package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
	"github.com/hajimehoshi/ebiten/inpututil"  // required for isKeyJustPressed

	cm "github.com/Rosalita/my-ebiten/pkgs/charmenu"

)

var (
	charMenu cm.CharMenu
	displayString string
)

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	
	ebitenutil.DebugPrint(screen, displayString)

	charMenu.Draw(screen)

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		charMenu.DecRow()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		charMenu.IncRow()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		charMenu.IncCol()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		charMenu.DecCol()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter){
		displayString += charMenu.GetSelectedChar()
	}

	return nil
}

func main() {
	displayString = ""

	white := &color.NRGBA{0xff, 0xff, 0xff, 0xff}

	charMenuInput := cm.Input{
		Tx: 50,
		Ty: 50,
		DefaultBgColour: white,
	}

	charMenu, _ = cm.NewMenu(charMenuInput)


	if err := ebiten.Run(update, 320, 240, 2, "Character menu"); err != nil {
		panic(err)
	}
}
