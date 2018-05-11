package main

import (
  "bytes"
  "math/rand"
)

type Tile int

const (
  Explored_0 Tile = 0
  Explored_1 Tile = 1
  Explored_2 Tile = 2
  Explored_3 Tile = 3
  Explored_4 Tile = 4
  Explored_5 Tile = 5
  Explored_6 Tile = 6
  Explored_7 Tile = 7
  Explored_8 Tile = 8
  Unknown Tile = 9
  Bomb Tile = 10
  BombShown Tile = 11
)

func (tile Tile) String() string {
  names := [...]string{
    "  ",
    "1 ",
    "2 ",
    "3 ",
    "4 ",
    "5 ",
    "6 ",
    "7 ",
    "8 ",
    "██",
    "██",
    "╳ "}
  return names[tile]
}

type Board struct {
  WIDTH, HEIGHT int
  x, y int
  grid []Tile
  flags []bool
}

func newBlankBoard(width int, height int) *Board {
  b := new(Board)
  b.x = width / 2
  b.y = height / 2
  b.WIDTH = width
  b.HEIGHT = height
  b.grid = make([]Tile, width * height)
  b.flags = make([]bool, width * height)
  for y := 0; y < b.HEIGHT; y++ {
    for x := 0; x < b.WIDTH; x++ {
      b.SetFlag(x, y, false)
    }
  }
  return b
}

func newBoard(width int, height int, numBombs int) *Board {
  b := newBlankBoard(width, height)
  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      b.SetPoint(x, y, Unknown)
    }
  }
  for i := 0; i < numBombs; i++ {
    b.SetPoint(rand.Int() % width, rand.Int() % height, Bomb)
  }
  return b
}

type BoolBoard struct {
  b Board
  grid []bool
}

func newBoolBoard(board Board) *BoolBoard {
  b := new(BoolBoard)
  b.b = board
  b.grid = make([]bool, board.WIDTH * board.HEIGHT)
  for y := 0; y < board.HEIGHT; y++ {
    for x := 0; x < board.WIDTH; x++ {
      b.SetPoint(x, y, false)
    }
  }
  return b
}

func (b *Board) GetPoint(x int, y int) Tile {
  if OutOfBoard(*b, x, y) {
    return 0
  }
  return b.grid[y * b.WIDTH + x]
}

func (b *Board) GetFlag(x int, y int) bool {
  if OutOfBoard(*b, x, y) {
    return false
  }
  return b.flags[y * b.WIDTH + x]
}

func (b *Board) GetDrawString(x int, y int, gameOver bool) string {
  if b.GetFlag(x, y) {
    if gameOver && (b.GetPoint(x, y) == Bomb || b.GetPoint(x, y) == BombShown) {
      return Colorize("⚑ ", GREEN)
    }
    return Colorize("⚑ ", RED)
  } else {
    return b.GetPoint(x, y).String()
  }
}

func (b *Board) SetPoint(x int, y int, value Tile) {
  if OutOfBoard(*b, x, y) {
    return
  }
  b.grid[y * b.WIDTH + x] = value
}

func (b *BoolBoard) GetPoint(x int, y int) bool {
  if OutOfBoard(b.b, x, y) {
    return false
  }
  return b.grid[y * b.b.WIDTH + x]
}

func (b *BoolBoard) SetPoint(x int, y int, value bool) {
  if OutOfBoard(b.b, x, y) {
    return
  }
  b.grid[y * b.b.WIDTH + x] = value
}

func (b *Board) ToggleFlag() {
  b.flags[b.y * b.WIDTH + b.x] = !b.flags[b.y * b.WIDTH + b.x]
}

func (b *Board) SetFlag(x int, y int, value bool) {
  b.flags[y * b.WIDTH + x] = value
}

// returns are_you_alive?
func (b *Board) SearchRecursive(x int, y int, boolgrid *BoolBoard) bool {
  if OutOfBoard(*b, x, y) || boolgrid.GetPoint(x, y) {
    return true
  }
  point := b.GetPoint(x, y)
  if (point == Bomb) {
    return false
  } else {
    value := b.GetValueOfPoint(x, y)
    b.SetPoint(x, y, value)
    boolgrid.SetPoint(x, y, true)
    if value == 0 {
      b.SearchRecursive(x - 1, y - 1, boolgrid)
      b.SearchRecursive(x    , y - 1, boolgrid)
      b.SearchRecursive(x + 1, y - 1, boolgrid)
      b.SearchRecursive(x - 1, y    , boolgrid)
      b.SearchRecursive(x + 1, y    , boolgrid)
      b.SearchRecursive(x - 1, y + 1, boolgrid)
      b.SearchRecursive(x    , y + 1, boolgrid)
      b.SearchRecursive(x + 1, y + 1, boolgrid)
    }
  }
  return true
}

// returns are_you_alive?, did_you_win?
func (b *Board) Search() (bool, bool) {
  if b.SearchRecursive(b.x, b.y, newBoolBoard(*b)) {
    explored_all := true
    for _, tile := range b.grid {
      if tile == Unknown {
        explored_all = false
        break
      }
    }
    if explored_all {
      return false, true
    }
    return true, false
  } else {
    // updates board to show bombs
    for i, tile := range b.grid {
      if tile == Bomb {
        b.grid[i] = BombShown
      }
    }
    return false, false
  }
}

func (b *Board) GetValueOfPoint(x int, y int) Tile {
  if OutOfBoard(*b, x, y) {
    return Explored_0
  }
  points := [...]Point {
    Point{-1, -1},
    Point{0 , -1},
    Point{1 , -1},
    Point{-1,  0},
    Point{1 ,  0},
    Point{-1,  1},
    Point{0 ,  1},
    Point{1 ,  1},
  }
  value := Explored_0
  for _, p := range points {
    if b.GetPoint(p.x + x, p.y + y) == Bomb {
      value++
    }
  }
  return value;
}

func OutOfBoard(b Board, x int, y int) bool {
  return x < 0 || x > b.WIDTH - 1 || y < 0 || y > b.HEIGHT - 1
}

func (b *Board) String() string {
  var buffer bytes.Buffer
  for y := 0; y < b.HEIGHT; y++ {
    for x := 0; x < b.WIDTH; x++ {
      buffer.WriteString(b.GetPoint(x, y).String())
      buffer.WriteString(", ")
    }
    buffer.WriteString("\n")
  }
  return buffer.String()
}

type Point struct {
  x, y int
}
