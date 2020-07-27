package pong

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

//Pixels each pixel of the screen
type Pixels struct {
	screen []byte
}

func firstInit() {
	rand.Seed(time.Now().UnixNano())
}

func initScreen() *Pixels {
	pixels := Pixels{}
	pixels.screen = make([]byte, winWidth*winHeigth*4)
	pixels.resetScreen()
	return &pixels
}
func (pixels *Pixels) resetScreen() {
	for y := 0; y < winHeigth; y++ {
		for x := 0; x < winWidth; x++ {
			pixels.setPixel(x, y, Black())
		}
	}
	/*for i := range pixels.screen {
		pixels.screen[i] = 0
	}*/
}

/*
func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}*/
func (pixels *Pixels) setPixel(x, y int, c Color) {
	index := (y*winWidth + x) * 4
	if index < len(pixels.screen)-4 && index >= 0 {
		pixels.screen[index] = c.Red
		pixels.screen[index+1] = c.Blue
		pixels.screen[index+2] = c.Green
	}
}

func startSDL2() (*sdl.Window, *sdl.Renderer, *sdl.Texture) {
	/*err := sdl.Init(sdl.INIT_GAMECONTROLLER)
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil
	}*/
	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeigth), sdl.WINDOW_SHOWN)

	if err != nil {
		fmt.Println(err)
		return window, nil, nil
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return window, renderer, nil
	}
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeigth))
	if err != nil {
		fmt.Println(err)
	}
	return window, renderer, texture
}

func initControlHandler() []*sdl.GameController {
	var controlHandlers []*sdl.GameController
	for i := 0; i < sdl.NumJoysticks(); i++ {
		controlHandlers = append(controlHandlers, sdl.GameControllerOpen(i))
		defer controlHandlers[i].Close()
	}
	return controlHandlers
}

//Run Start the sdl
func Run() {
	firstInit()
	window, renderer, texture := startSDL2()
	defer sdl.Quit()
	defer window.Destroy()
	if renderer != nil {
		defer renderer.Destroy()
	}
	if texture != nil {
		defer texture.Destroy()
	}
	mainLoop(renderer, texture)

}
func mainLoop(renderer *sdl.Renderer, texture *sdl.Texture) {
	//controlHandlers := initControlHandler()
	keyState := sdl.GetKeyboardState()
	var frameStart time.Time
	var elapsedTime float32
	var controllerAxis int16
	player1 := paddle{NewPos(100, 100), 15, 50, 300, 300, 0, White()}
	player2 := paddle{NewPos(float32(winWidth)-50, 100), 15, 50, 300, 300, 0, White()}
	ball := ball{NewPos(300, 300), 20, 100, 100, rand.Float32() * 800, White()}
	pixels := initScreen()

	for {
		frameStart = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		/*for _, controller := range controlHandlers {
			if controller != nil {
				controllerAxis = controller.Axis(sdl.CONTROLLER_AXIS_LEFTY)
			}
		}*/
		if state == play {
			player1.update(keyState, controllerAxis, elapsedTime)
			player2.aiUpdate(&ball, elapsedTime)
			ball.update(&player1, &player2, elapsedTime)
		} else if state == start {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				if player1.score == 3 || player2.score == 3 {
					player1.score = 0
					player2.score = 0
				}
				state = play
			}
		}
		pixels.resetScreen()
		drawMiddleLine(White(), 1, 4, pixels)
		player1.draw(1, 5, pixels)
		player2.draw(1, 5, pixels)
		ball.draw(pixels)

		err := texture.Update(nil, pixels.screen, winWidth*4)
		if err != nil {
			fmt.Println(sdl.GetError())
		}
		err = renderer.Copy(texture, nil, nil)
		if err != nil {
			fmt.Println(sdl.GetError())
		}
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())

		//144 FPS
		if elapsedTime < 0.0069 {
			sdl.Delay(5 - uint32(elapsedTime)*1000.0)
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}

}
