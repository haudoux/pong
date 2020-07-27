package pong

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type paddle struct {
	Pos
	w          float32
	h          float32
	speed      float32
	startSpeed float32
	score      int
	color      Color
}

//16
//2
func (paddle *paddle) draw(width, height int, pixels *Pixels) {
	startX := int(paddle.X - paddle.w/2)
	startY := int(paddle.Y - paddle.h/2)

	for i, v := range pad {
		if v == 1 {
			for y := startY; y < startY+width; y++ {
				for x := startX; x < startX+height; x++ {
					pixels.setPixel(x, y, White())
				}
			}
		}
		if (i+1)%8 == 0 {
			startY += height
		}
	}

	paddle.drawNumbers(paddle.color, 3, 1, paddle.score, pixels)
}

func (paddle *paddle) drawNumbers(color Color, width, heigth, num int, pixels *Pixels) {
	startX := 0
	startY := winHeigth / 20
	if int(paddle.X) < winWidth/2 {
		startX = winWidth / 4
	} else {
		startX = winWidth / 4 * 3
	}

	for i, v := range nums[num] {
		if v == 1 {
			for y := startY; y < startY+heigth; y++ {
				for x := startX; x < startX+width; x++ {
					pixels.setPixel(x, y, color)
				}
			}
		}
		startX += width
		if (i+1)%10 == 0 {
			startY += heigth + 1
			startX -= width * 10
		}
	}
}
func (paddle *paddle) update(keyState []uint8, controllerAxis int16, elapsedTime float32) {
	if keyState[sdl.SCANCODE_UP] != 0 && paddle.Y-paddle.h > 0 {
		paddle.Y -= paddle.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 && int(paddle.Y+paddle.h) < winHeigth {
		paddle.Y += paddle.speed * elapsedTime
	}
	if math.Abs(float64(controllerAxis)) > 2800 {
		pct := float32(controllerAxis) / 32767.0
		paddle.Y += paddle.speed * pct * elapsedTime
	}
}
func (paddle *paddle) aiUpdate(ball *ball, elapsedTime float32) {
	paddle.Y = ball.Y //* elapsedTime
}
