package listmenu_test

import (
	"errors"
	"image/color"
	"testing"

	lm "github.com/Rosalita/ebiten-pkgs/listmenu"
	"github.com/stretchr/testify/assert"
)

var (
	white   = &color.NRGBA{0xff, 0xff, 0xff, 0xff}
	black   = &color.NRGBA{0x00, 0x00, 0x00, 0xff}
	magenta = &color.NRGBA{0xff, 0x00, 0xff, 0xff}
)

func TestNewMenu(t *testing.T) {
	// try create a new menu with missing values width, height, items

	validItems := []lm.Item{
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

	missingWidth := lm.Input{
		Height: 30,
		Items:  validItems,
	}

	missingHeight := lm.Input{
		Width: 50,
		Items: validItems,
	}

	missingItems := lm.Input{
		Height: 30,
		Width:  50,
	}

	tests := []struct {
		input lm.Input
		menu  lm.ListMenu
		err   error
	}{
		{missingWidth, lm.ListMenu{}, errors.New("Mandatory input field width is missing")},
		{missingHeight, lm.ListMenu{}, errors.New("Mandatory input field height is missing")},
		{missingItems, lm.ListMenu{}, errors.New("Mandatory input field MenuItems is missing")},
	}

	for _, test := range tests {
		result, err := lm.NewMenu(test.input)
		assert.Equal(t, test.menu, result)
		assert.Equal(t, test.err, err)
	}

}
