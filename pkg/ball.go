package pong

import (
	"fmt"
	"math/rand"
)

type ball struct {
	Pos
	radius  float32
	xv      float32
	startXV float32
	yv      float32
	color   Color
}

func (ball *ball) draw(pixels *Pixels) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				pixels.setPixel(int(ball.X+x), int(ball.Y+y), ball.color)
			}
		}
	}
}

func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle, elapsedTime float32) {
	//Set velocity
	ball.X += ball.xv * elapsedTime
	ball.Y += ball.yv * elapsedTime

	//Bounce top and bot
	if ball.Y-ball.radius < 0 {
		ball.Y = 0 + ball.radius
		ball.yv = -ball.yv
	} else if ball.Y+ball.radius > float32(winHeigth) {
		ball.Y = float32(winHeigth) - ball.radius
		ball.yv = -ball.yv
	}
	//Reverse velocity if contact
	if ball.X-ball.radius < leftPaddle.X+leftPaddle.w/2 {
		if ball.Y > leftPaddle.Y-leftPaddle.h/2 && ball.Y < leftPaddle.Y+leftPaddle.h/2 {
			leftPaddle.speed *= 1.1
			rightPaddle.speed *= 1.1
			ball.xv = -ball.xv * 1.25
			ball.yv = getRandomYv()
			ball.X = leftPaddle.X + leftPaddle.w/2 + ball.radius
		}
	}
	if ball.X+ball.radius > rightPaddle.X-rightPaddle.w/2 {
		if ball.Y > rightPaddle.Y-rightPaddle.h/2 && ball.Y < rightPaddle.Y+rightPaddle.h/2 {
			leftPaddle.speed *= 1.1
			rightPaddle.speed *= 1.1
			ball.xv = -ball.xv * 1.25
			ball.yv = getRandomYv()
			ball.X = rightPaddle.X - rightPaddle.w/2.0 - ball.radius
		}
	}
	//Score if behind paddle
	if ball.X-ball.radius < leftPaddle.X-leftPaddle.w-(ball.radius*2) {
		rightPaddle.score++
		ball.Pos = getCenter()
		ball.yv = getRandomYv()
		ball.xv = ball.startXV
		leftPaddle.speed = leftPaddle.startSpeed
		rightPaddle.speed = rightPaddle.startSpeed
		rightPaddle.Y = getCenter().Y
		state = start
	} else if ball.X+ball.radius > rightPaddle.X+rightPaddle.w+(ball.radius*2) {
		leftPaddle.score++
		ball.Pos = getCenter()
		ball.xv = ball.startXV
		leftPaddle.Y = getCenter().Y
		ball.yv = getRandomYv()
		state = start
	}
	fmt.Println(ball.xv)
}

func getRandomYv() float32 {
	var nb float32
	if rand.Float32() > 0.5 {
		nb = rand.Float32() * 800
	}
	if rand.Float32() < 0.5 {
		nb = rand.Float32() * -800
	}
	return nb
}
