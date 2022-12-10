package handheld

// Keeps track of device screen with drawing cycles and sprites
type Screen struct {
	width   int
	height  int
	display [][]rune
}

func NewScreen(width int, height int) *Screen {
	screen := Screen{width: width, height: height}
	screen.display = [][]rune{}
	for y := 0; y < screen.height; y++ {
		screen.display = append(screen.display, []rune{})
		for x := 0; x < screen.width; x++ {
			screen.display[y] = append(screen.display[y], '.')
		}
	}
	return &screen
}

func (screen *Screen) drawPixel(cycle int, register int) {
	row := int((cycle - 1) / screen.width)
	pixel := (cycle - 1) % screen.width
	if pixel >= register-1 && pixel <= register+1 {
		screen.display[row][pixel] = '#'
	} else {
		screen.display[row][pixel] = '.'
	}
}

func (screen *Screen) getImage() []string {
	image := []string{}
	for _, row := range screen.display {
		image = append(image, string(row))
	}
	return image
}
