package charmenu

import (
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

var mplusNormalFont font.Face

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

// CharItem holds everything needed to display a single character
type CharItem struct {
	Char  string        // The character displayed on the box
	TxtX  int           // X location to draw text
	TxtY  int           // Y location to draw text
	image *ebiten.Image // used to store the image for the char box
}

// CharMenu is a navigatable, selectable menu that displays chars
type CharMenu struct {
	Tx                  float64      // x translation of the menu
	Ty                  float64      // y translation of the menu
	Offx                float64      // X offset for each character box
	Offy                float64      // Y offset for each character box
	DefaultBgColour     *color.NRGBA // default background colour
	DefaultTxtColour    *color.NRGBA // default text colour
	DefaultSelBgColour  *color.NRGBA // default selected background colour
	DefaultSelTxtColour *color.NRGBA // default selected text colour
	CharList            string       // a string containing all the characters to display
	CharsPerRow         int          // the number of chars in a row
	SelectedRow         *int         // row index of the selected item
	SelectedCol         *int         // column index of the selected item
	CharGrid            [][]CharItem
}

// Input is an object used to create an alpha list
type Input struct {
	Tx                  float64      // optional, x translation of the menu, if not provided will be 0
	Ty                  float64      // optional, y translation of the menu, if not provided will be 0
	DefaultBgColour     *color.NRGBA // optional, default background colour of menu, if not provided will be cyan
	DefaultTxtColour    *color.NRGBA // optional, default text colour, if not provided will be black
	DefaultSelBGColour  *color.NRGBA // optional, default selected background colour of menu, if not provided will be magenta
	DefaultSelTxtColour *color.NRGBA //optional, default selected text colour of menu, if not provided it will be white
}

//NewMenu constructs a new alpha menu from a Input
func NewMenu(input Input) (CharMenu, error) {

	if input.DefaultBgColour == nil {
		input.DefaultBgColour = &color.NRGBA{0x00, 0xff, 0xff, 0xff}
	}

	if input.DefaultTxtColour == nil {
		input.DefaultTxtColour = &color.NRGBA{0x00, 0x00, 0x00, 0xff}
	}

	if input.DefaultSelBGColour == nil {
		input.DefaultSelBGColour = &color.NRGBA{0xff, 0x00, 0xff, 0xff}
	}

	if input.DefaultSelTxtColour == nil {
		input.DefaultSelTxtColour = &color.NRGBA{0xff, 0xff, 0xff, 0xff}
	}

	defaultSelectedRow := 0
	defaultSelectedCol := 0
	defaultOffx := 20.0
	defaultOffy := 20.0
	defaultWidth := 18
	defaultHeight := 18
	defaultLineLength := 13

	charList := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ. "

	charGrid := charListToCharGrid(charList, defaultLineLength)

	charGrid = initGrid(charGrid, defaultWidth, defaultHeight)

	m := CharMenu{
		Tx:                  input.Tx,
		Ty:                  input.Ty,
		Offx:                defaultOffx,
		Offy:                defaultOffy,
		DefaultBgColour:     input.DefaultBgColour,
		DefaultTxtColour:    input.DefaultTxtColour,
		DefaultSelBgColour:  input.DefaultSelBGColour,
		DefaultSelTxtColour: input.DefaultSelTxtColour,
		CharList:            charList,
		CharsPerRow:         defaultLineLength,
		SelectedRow:         &defaultSelectedRow,
		SelectedCol:         &defaultSelectedCol,
		CharGrid:            charGrid,
	}

	return m, nil
}

//Draw draws the list menu to the screen
func (m *CharMenu) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(m.Tx, m.Ty)

	for y, row := range m.CharGrid {
		for x := range row {

			if y != 0 && x == 0 { // if  not first row and is first item in the row
				// translate back to start of row
				// translate down to start drawing new row
				opts.GeoM.Translate(-(float64(m.CharsPerRow) * m.Offx), m.Offy)
			}

			if *m.SelectedRow == y && *m.SelectedCol == x {
				m.CharGrid[y][x].image.Fill(m.DefaultSelBgColour)
			} else {
				m.CharGrid[y][x].image.Fill(m.DefaultBgColour)
			}

			text.Draw(m.CharGrid[y][x].image, m.CharGrid[y][x].Char, mplusNormalFont, int(m.CharGrid[y][x].TxtX), int(m.CharGrid[y][x].TxtY), m.DefaultTxtColour)

			screen.DrawImage(m.CharGrid[y][x].image, opts)
			opts.GeoM.Translate(m.Offx, 0.0)

		}
	}
}

//GetSelectedChar returns the char selected in the menu
func (m *CharMenu) GetSelectedChar() string {
	return m.CharGrid[*m.SelectedRow][*m.SelectedCol].Char
}

//IncRow increments the selectedRow index if able
func (m *CharMenu) IncRow() {
	lastRowIndex := len(m.CharList) / m.CharsPerRow

	if *m.SelectedRow < lastRowIndex-1 {
		*m.SelectedRow++
	} else if *m.SelectedRow == lastRowIndex-1 {
		numCharsOnLastRow := len(m.CharList) % m.CharsPerRow
		if *m.SelectedCol <= (numCharsOnLastRow - 1) {
			*m.SelectedRow++
		}
	}
}

//DecRow decrements the selectedRow index if able
func (m *CharMenu) DecRow() {
	minIndex := 0
	if *m.SelectedRow > minIndex {
		*m.SelectedRow--
	}
}

//IncCol increments the selectedCol index if able
func (m *CharMenu) IncCol() {
	maxIndex := m.CharsPerRow - 1
	lastRowIndex := len(m.CharList) / m.CharsPerRow

	if *m.SelectedCol < maxIndex && *m.SelectedRow < lastRowIndex {
		*m.SelectedCol++
	} else if *m.SelectedRow == lastRowIndex {

		numCharsOnLastRow := len(m.CharList) % m.CharsPerRow

		if *m.SelectedCol < numCharsOnLastRow-1 {
			*m.SelectedCol++
		}

	}
}

//DecCol decrements the selectedCol index if able
func (m *CharMenu) DecCol() {
	minIndex := 0
	if *m.SelectedCol > minIndex {
		*m.SelectedCol--
	}
}

func charListToCharGrid(charList string, lineLength int) (foo [][]CharItem) {
	charGrid := [][]CharItem{}
	lines := splitIntoLines(charList, lineLength)
	for _, line := range lines {

		row := []CharItem{}

		for _, char := range line {

			charItem := CharItem{
				Char: string(char),
			}
			row = append(row, charItem)
		}
		charGrid = append(charGrid, row)
	}
	return charGrid
}

func initGrid(grid [][]CharItem, width, height int) [][]CharItem {
	for y, row := range grid {
		for x := range row {
			img, _ := ebiten.NewImage(width, height, ebiten.FilterNearest)
			grid[y][x].image = img

			if grid[y][x].Char == "w" || grid[y][x].Char == "M" {
				grid[y][x].TxtX = 3
				grid[y][x].TxtY = 14
			} else if  grid[y][x].Char == "W"{
				grid[y][x].TxtX = 2
				grid[y][x].TxtY = 14
			}else if  grid[y][x].Char == "N" || grid[y][x].Char == "Q" || grid[y][x].Char == "O" ||  grid[y][x].Char == "m"{
					grid[y][x].TxtX = 4
					grid[y][x].TxtY = 14
			} else {
				grid[y][x].TxtX = 5
				grid[y][x].TxtY = 14
			}

		}
	}
	return grid
}

func splitIntoLines(s string, lineLength int) []string {

	runeList := []rune(s)
	lines := []string{}
	line := ""

	for i, r := range runeList {
		line = line + string(r)

		if i > 0 && (i+1)%lineLength == 0 || i == len(s)-1 {
			lines = append(lines, line)
			line = ""
		}
	}
	return lines
}
