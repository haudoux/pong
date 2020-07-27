package pong

const winWidth, winHeigth int = 800, 600

type gameState int

const (
	start gameState = iota
	play
)

var state = start

func getCenter() Pos {
	return NewPos(float32(winWidth)/2, float32(winHeigth)/2)
}

func drawMiddleLine(color Color, width int, height int, pixels *Pixels) {
	startX := winWidth / 2
	startY := 0

	for startY < winHeigth {
		for i, v := range line {
			if v == 1 {
				for y := startY; y < startY+width; y++ {
					for x := startX; x < startX+height; x++ {
						pixels.setPixel(x, y, color)
					}
				}
			}
			if (i+1)%3 == 0 {
				startY += height
			}
		}
	}
}
