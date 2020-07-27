package pong

//Color type
type Color struct {
	Red, Green, Blue byte
}

//NewColor Create and return a color struct
func NewColor(r byte, g byte, b byte) Color {
	return Color{r, g, b}
}

//White  return white color
func White() Color {
	return Color{255, 255, 255}
}

//Black  return black color
func Black() Color {
	return Color{0, 0, 0}
}
