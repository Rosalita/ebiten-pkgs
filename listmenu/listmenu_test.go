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

	validInput := lm.Input{
		Width:      100,
		ItemHeight: 40,
		Items:      validItems,
	}

	missingWidth := lm.Input{
		ItemHeight: 30,
		Items:      validItems,
	}

	missingItemHeight := lm.Input{
		Width: 50,
		Items: validItems,
	}

	missingItems := lm.Input{
		ItemHeight: 30,
		Width:      50,
	}

	expectedMenu := lm.ListMenu{
		Tx:         0,
		Ty:         0,
		Width:      100,
		ItemHeight: 40,
		Height:     200,
		Offx:       0,
		Offy:       40,
		Items: []lm.Item{
			{
				Name: "playButton",
				Text: "PLAY",
				TxtX: 40,
				TxtY: 25,
			},
			{
				Name: "optionButton",
				Text: "OPTIONS",
				TxtX: 16,
				TxtY: 25,
			},
			{
				Name: "quitButton",
				Text: "QUIT",
				TxtX: 40,
				TxtY: 25,
			},
		},
	}

	tests := []struct {
		input        lm.Input
		expectedMenu lm.ListMenu
		err          error
	}{
		{validInput, expectedMenu, nil},
		{missingWidth, lm.ListMenu{}, errors.New("Mandatory input field Width is missing")},
		{missingItemHeight, lm.ListMenu{}, errors.New("Mandatory input field ItemHeight is missing")},
		{missingItems, lm.ListMenu{}, errors.New("Mandatory input field Items is missing")},
	}

	for _, test := range tests {
		result, err := lm.NewMenu(test.input)
		assert.Equal(t, test.expectedMenu.Width, result.Width)
		assert.Equal(t, test.expectedMenu.ItemHeight, result.ItemHeight)
		assert.Equal(t, test.expectedMenu.Height, result.Height)
		assert.Equal(t, test.expectedMenu.Offy, result.Offy)
		assert.Equal(t, len(test.expectedMenu.Items), len(result.Items))
		assert.Equal(t, test.err, err)
	}
}
