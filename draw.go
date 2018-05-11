package main

import (
  "fmt"
  "bytes"
  "strings"
)

func DrawBoard(board Board, gameOver bool) {
  //Clear()
  MoveCursor(0, 0)
  fmt.Print(GetDrawString(board, gameOver))
}

func Draw(board Board) {
  points := [...]Point {
    Point{-1, 0 },
    Point{1 , 0 },
    Point{0 , -1},
    Point{0 , 1 }}
  for _, p := range points {
    if (OutOfBoard(board, board.x + p.x, board.y + p.y)) {
      continue
    }
    drawOverLines(board.x + p.x, board.y + p.y, board, GetAroundCursorString(
      board.x + p.x,
      board.y + p.y,
      board.WIDTH,
      board.HEIGHT,
      board.GetDrawString(board.x + p.x, board.y + p.y, false)))
  }
  drawOverLines(board.x, board.y, board, Colorize(GetCursorString(board.GetDrawString(board.x, board.y, false)), BLUE))
}

func drawOverLines(x int, y int, board Board, str string) {
  array := strings.Split(str, "\n")
  for i, s := range array {
    MoveCursor(x * 4, y * 2 + i)
    fmt.Print(s)
  }
}

func GetCursorString(tile string) string {
  var buffer bytes.Buffer
  buffer.WriteString("┏ ━ ┓\n┃ ")
  buffer.WriteString(tile)
  buffer.WriteString(Colorize("┃\n┗ ━ ┛", BLUE))
  return buffer.String()
}

func GetAroundCursorString(x int, y int, width int, height int, tile string) string {
  var buffer bytes.Buffer
  if y == 0 { // no bottom; this is only the top row
    if x == 0 {
      buffer.WriteString("┏ ━ ┳") // top left
    } else if x == width - 1 {
      buffer.WriteString("┳ ━ ┓") // top right
    } else {
      buffer.WriteString("┳ ━ ┳") // top middle
    }
  } else {
    if x == 0 {
      buffer.WriteString("┣ ━ ╋") // far left middle
    } else if x == width - 1 {
      buffer.WriteString("╋ ━ ┫") // far right middle
    } else {
      buffer.WriteString("╋ ━ ╋") // middle
    }
  }
  buffer.WriteString("\n┃ ")
  buffer.WriteString(tile)
  buffer.WriteString("┃\n")
  if y == 0 { // bottom row
    if x == 0 {
      buffer.WriteString("┗ ━ ┻") // bottom left
    } else if x == width - 1 {
      buffer.WriteString("┻ ━ ┛") // bottom right
    } else {
      buffer.WriteString("┻ ━ ┻") // bottom middle
    }
  } else {
    if x == 0 {
      buffer.WriteString("┣ ━ ╋") // far left middle
    } else if x == width - 1 {
      buffer.WriteString("╋ ━ ┫") // far right middle
    } else {
      buffer.WriteString("╋ ━ ╋") // middle
    }
  }
  return Colorize(buffer.String(), WHITE)
}

func GetDrawString(board Board, gameOver bool) string {
  var buffer bytes.Buffer
  for y := 0; y < board.HEIGHT; y++ {
    for x := 0; x < board.WIDTH; x++ {
      if y == 0 {
        if x == 0 {
          buffer.WriteString("┏ ━ ")
        } else {
          buffer.WriteString("┳ ━ ")
        }
      } else {
        if x == 0 {
          buffer.WriteString("┣ ━ ")
        } else {
          buffer.WriteString("╋ ━ ")
        }
      }
    }
    if y == 0 {
      buffer.WriteString("┓")
    } else {
      buffer.WriteString("┫")
    }
    buffer.WriteString("\n")
    for x := 0; x < board.WIDTH; x++ {
      buffer.WriteString("┃ ")
      buffer.WriteString(board.GetDrawString(x, y, gameOver))
      buffer.WriteString(Colorize("", WHITE))
      if x == board.WIDTH - 1 {
        buffer.WriteString("┃")
      }
    }
    buffer.WriteString("\n")
  }
  for x := 0; x < board.WIDTH; x++ {
    if x == 0 {
      buffer.WriteString("┗ ━ ")
    } else {
      buffer.WriteString("┻ ━ ")
    }
  }
  buffer.WriteString("┛\n")
  return Colorize(buffer.String(), WHITE)
}
