package main

import(
  "fmt"
  "bytes"
  "strconv"
)

const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

var DEFAULT int;

func getColor(code int) string {
	return fmt.Sprintf("\033[3%dm", code)
}

func MoveCursor(x int, y int) {
	fmt.Printf("\033[%d;%dH", y + 1, x + 1)
}

func Clear() {
	fmt.Printf("\033[2J")
}

func Colorize(str string, color int) string {
  var buffer bytes.Buffer
  buffer.WriteString("\033[")
  buffer.WriteString(strconv.Itoa(30 + color))
  buffer.WriteString("m")
  buffer.WriteString(str)
  buffer.WriteString("\033[")
  buffer.WriteString(strconv.Itoa(30 + DEFAULT))
  buffer.WriteString("m")
  return buffer.String()
}

func ResetColor() {
  fmt.Printf("\033[0m")
}
