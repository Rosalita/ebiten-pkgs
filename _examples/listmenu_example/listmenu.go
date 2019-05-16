package main

import (
	"fmt"
	"image/color"

	lm "github.com/Rosalita/ebiten-pkgs/listmenu"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
	"github.com/hajimehoshi/ebiten/inpututil"  // required for isKeyJustPressed
)

var (
	listMenu lm.ListMenu
)

func update(screen *ebiten.Image) error {

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		listMenu.DecrementSelected()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		listMenu.IncrementSelected()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch listMenu.GetSelectedItem() {
		case "playButton":
			fmt.Println("Play was selected")
		case "optionButton":
			fmt.Println("Options was selected")
		case "quitButton":
			fmt.Println("Quit was selected")
		}
		return nil
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
	ebitenutil.DebugPrint(screen, "List Menu - use up and down arrow keys")
	listMenu.Draw(screen)

	return nil
}

func main() {

	white := &color.NRGBA{0xff, 0xff, 0xff, 0xff}
	pink := &color.NRGBA{0xff, 0x69, 0xb4, 0xff}

	menuItems := []lm.Item{
		{Name: "playButton",
			Text:     "PLAY",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
		{Name: "optionButton",
			Text:     "OPTIONS",
			TxtX:     16,
			TxtY:     25,
			BgColour: white},
		{Name: "quitButton",
			Text:     "QUIT",
			TxtX:     40,
			TxtY:     25,
			BgColour: white},
	}

	menuInput := lm.Input{
		Width:              140,
		ItemHeight:         36,
		Tx:                 24,
		Ty:                 24,
		Offy:               40,
		DefaultSelBgColour: pink,
		Items:              menuItems,
	}

	listMenu, _ = lm.NewMenu(menuInput)

	if err := ebiten.Run(update, 320, 240, 1, "List menu"); err != nil {
		panic(err)
	}
}
