package charmenu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharListToCharGrid(t *testing.T) {

	oneRow := [][]CharItem{
		[]CharItem{
			{Char: "a"}, {Char: "b"}, {Char: "c"},
		},
	}

	twoRows := [][]CharItem{
		[]CharItem{
			{Char: "a"}, {Char: "b"}, {Char: "c"},
		},
		[]CharItem{
			{Char: "d"}, {Char: "e"},
		},
	}

	tests := []struct {
		charList   string
		lineLength int
		result     [][]CharItem
	}{
		{"abc", 3, oneRow},
		{"abcde", 3, twoRows},
	}

	for _, test := range tests {
		result := charListToCharGrid(test.charList, test.lineLength)
		assert.Equal(t, test.result, result)
	}
}

func TestSplitIntoLines(t *testing.T) {

	tests := []struct {
		charList   string
		lineLength int
		result     []string
	}{
		{"abc", 3, []string{"abc"}},
		{"abcdef", 3, []string{"abc", "def"}},
		{"12345", 2, []string{"12", "34", "5"}},
	}

	for _, test := range tests {
		result := splitIntoLines(test.charList, test.lineLength)
		assert.Equal(t, test.result, result)
	}

}

func TestInitGrid(t *testing.T) {

	grid := [][]CharItem{
		[]CharItem{
			{Char: "a"}, {Char: "b"}, {Char: "c"},
		},
		[]CharItem{
			{Char: "d"}, {Char: "e"},
		},
	}

	expected := [][]CharItem{
		[]CharItem{
			{Char: "a"}, {Char: "b"}, {Char: "c"},
		},
		[]CharItem{
			{Char: "d"}, {Char: "e"},
		},
	}

	tests := []struct {
		grid   [][]CharItem
		width  int
		height int
		result [][]CharItem
	}{
		{grid, 10, 20, expected},
	}

	for _, test := range tests {
		result := initGrid(test.grid, test.width, test.height)
		for y, row := range result {
			for x := range row {
				assert.NotNil(t, result[y][x].image)
				assert.NotNil(t, result[y][x].TxtX)
				assert.NotNil(t, result[y][x].TxtY)
			}
		}
	}
}
