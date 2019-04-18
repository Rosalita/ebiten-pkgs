package imagemenu

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

// Item represents an item in a image menu
type Item struct {
	Name  string        // a name to describe each menu item
	Bytes []byte        // bytes used to generate images
	image *ebiten.Image // used to store the image of each menu item
}

// ImageMenu is a navigatable graphical menu where each item is an image
type ImageMenu struct {
	Tx            float64 // x translation of the menu
	Ty            float64 // y translation of the menu
	ImgWidth      int     // width of all image menu items
	ImgHeight     int     // height of all image menu items
	Items         []Item  // menu items
	SelectedIndex *int
}

// Input is an object used to create an image menu
type Input struct {
	Tx        float64 // x translation of the menu
	Ty        float64 // x translation of the menu
	ImgWidth  int     // mandatory, width of all image menu items
	ImgHeight int     // mandatory, height of all image menu items
	Items     []Item  // menu items
}

var (
	rightArrow *ebiten.Image
	leftArrow  *ebiten.Image
)

func init() {
	rightArrow, _ = ebiten.NewImage(36, 36, ebiten.FilterDefault)
	leftArrow, _ = ebiten.NewImage(36, 36, ebiten.FilterDefault)
}

//NewMenu constructs a new image menu from a Input
func NewMenu(input Input) (ImageMenu, error) {

	if input.ImgWidth == 0 {
		return ImageMenu{}, errors.New("Mandatory input ImgWidth is missing")
	}
	if input.ImgHeight == 0 {
		return ImageMenu{}, errors.New("Mandatory input ImgHeight is missing")
	}
	if len(input.Items) < 1 {
		return ImageMenu{}, errors.New("Mandatory input field Items is missing")
	}

	defaultSelectedIndex := 0

	m := ImageMenu{
		Tx:            input.Tx,
		Ty:            input.Ty,
		ImgWidth:      input.ImgWidth,
		ImgHeight:     input.ImgHeight,
		Items:         input.Items,
		SelectedIndex: &defaultSelectedIndex,
	}

	// initialise images for each []bytes in items
	for i := range m.Items {
		if m.Items[i].Bytes != nil {
			img, _, err := image.Decode(bytes.NewReader(m.Items[i].Bytes))
			if err != nil {
				log.Fatal(err)
			}
			newImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
			m.Items[i].image = newImage
		}
	}

	return m, nil
}

//Draw draws the image menu to the screen
func (m *ImageMenu) Draw(screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(m.Tx, m.Ty)

	screen.DrawImage(m.Items[*m.SelectedIndex].image, opts)

	drawArrows(m.Tx, m.Ty, m.ImgWidth, m.ImgHeight, screen)

}

func drawArrows(tx, ty float64, imgWidth int, imgHeight int, screen *ebiten.Image) {

	imgMidY := float64(imgHeight) / 2

	lv1 := ebiten.Vertex{
		DstX:   float32(tx) - 10,
		DstY:   float32(ty) + float32(imgMidY),
		SrcX:   0,
		SrcY:   0,
		ColorR: 1,
		ColorG: 1,
		ColorB: 1,
		ColorA: 1,
	}

	lv2 := ebiten.Vertex{
		DstX:   float32(tx),
		DstY:   float32(ty) + float32(imgMidY) - 10,
		SrcX:   0,
		SrcY:   0,
		ColorR: 1,
		ColorG: 1,
		ColorB: 1,
		ColorA: 1,
	}

	lv3 := ebiten.Vertex{
		DstX:   float32(tx),
		DstY:   float32(ty) + float32(imgMidY) + 10,
		SrcX:   0,
		SrcY:   0,
		ColorR: 1,
		ColorG: 1,
		ColorB: 1,
		ColorA: 1,
	}

	rv1 := ebiten.Vertex{
		DstX:   float32(tx) + float32(imgWidth) + 10,
		DstY:   float32(ty) + float32(imgMidY),
		SrcX:   0,
		SrcY:   0,
		ColorR: 1,
		ColorG: 1,
		ColorB: 1,
		ColorA: 1,
	}

	rv2 := ebiten.Vertex{
		DstX:   float32(tx) + float32(imgWidth),
		DstY:   float32(ty) + float32(imgMidY) - 10,
		SrcX:   0,
		SrcY:   0,
		ColorR: 1,
		ColorG: 1,
		ColorB: 1,
		ColorA: 1,
	}

	rv3 := ebiten.Vertex{
		DstX:   float32(tx) + float32(imgWidth),
		DstY:   float32(ty) + float32(imgMidY) + 10,
		SrcX:   0,
		SrcY:   0,
		ColorR: 1,
		ColorG: 1,
		ColorB: 1,
		ColorA: 1,
	}

	leftvs := []ebiten.Vertex{lv1, lv2, lv3}
	rightvs := []ebiten.Vertex{rv1, rv2, rv3}

	indices := []uint16{0, 1, 2}

	op := &ebiten.DrawTrianglesOptions{}

	rightArrow.Fill(color.White)
	leftArrow.Fill(color.White)

	screen.DrawTriangles(leftvs, indices, leftArrow, op)
	screen.DrawTriangles(rightvs, indices, rightArrow, op)

}

//IncrementSelected increments the selected index provided it is not already at maximum
func (m *ImageMenu) IncrementSelected() {
	maxIndex := len(m.Items) - 1
	if *m.SelectedIndex < maxIndex {
		*m.SelectedIndex++
	}
}

//DecrementSelected decrements the selected index provided it is not already at minimum
func (m *ImageMenu) DecrementSelected() {
	minIndex := 0
	if *m.SelectedIndex > minIndex {
		*m.SelectedIndex--
	}
}
