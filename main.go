package main

import (
  "strconv"
  "fmt"
  "time"
  "math/rand"
  "os"
  "os/exec"
)

func main() {
  DEFAULT = WHITE
  // for a unique board each time; this is not automatic
  rand.Seed(time.Now().UTC().UnixNano())
  // disable input buffering
  exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
  // do not display entered characters on the screen
  exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
  Clear()

  var b []byte = make([]byte, 1)
  bombs := 32
  width := 16
  height := 16
  var program_args map[string]string
  if len(os.Args) > 1 {
    raw_args := os.Args[1:]
    program_args = make(map[string]string, len(raw_args) / 2)
    for i, arg := range raw_args {
      if i % 2 == 0 && i < len(raw_args) - 1 {
        program_args[arg] = raw_args[i + 1]
      }
    }
  }
  if val, ok := program_args["-b"]; ok {
    var err error
    bombs, err = strconv.Atoi(val)
    if (err != nil) {
      panic(err)
    }
  }
  if val, ok := program_args["-w"]; ok {
    var err error
    width, err = strconv.Atoi(val)
    if (err != nil) {
      panic(err)
    }
  }
  if val, ok := program_args["-h"]; ok {
    var err error
    height, err = strconv.Atoi(val)
    if (err != nil) {
      panic(err)
    }
  }
  board := newBoard(width, height, bombs)
  fmt.Println(board.HEIGHT)
  DrawBoard(*board, false)
  for {
    os.Stdin.Read(b)
    needs_redraw, reset := Update(board, string(b[0]))
    if reset {
      board = newBoard(width, height, bombs)
      Clear()
    }
    if needs_redraw {
      DrawBoard(*board, false)
    }
    Draw(*board)
  }
}

// whenever this returns, the board needs to be reset
func exit_replay(board *Board, quiting bool, won bool) {
  DrawBoard(*board, true)
  if (!quiting) {
    if (won) {
      fmt.Printf(Colorize("YOU WON!!! GOOD JOB!!!\n", GREEN))
    }
    fmt.Printf("Replay? (space for yes, anything else for no): ")
    var b []byte = make([]byte, 1)
    os.Stdin.Read(b)
    if (string(b[0]) == " ") {
      DEFAULT = WHITE
      return
    }
  }
  ResetColor()
  fmt.Printf("\n")
  exec.Command("stty", "-f", "/dev/tty", "echo").Run()
  os.Exit(0)
}

// returns need_redraw?, needs_board_reset?
func Update(board *Board, key string) (bool, bool) {
  switch key {
  case "w":
    board.y--
  case "a":
    board.x--
  case "s":
    board.y++
  case "d":
    board.x++
  case "f":
    board.ToggleFlag()
  case " ":
    alive, won := board.Search()
    if alive {
      return true, false
    } else {
      exit_replay(board, false, won)
      return true, true
    }
  case "q":
    MoveCursor(0, board.HEIGHT * 2 + 1)
    exit_replay(board, true, false)
    return true, true
  }
  if board.y < 0 {
    board.y = 0
  }
  if board.y >= board.HEIGHT {
    board.y = board.HEIGHT - 1
  }
  if board.x < 0 {
    board.x = 0
  }
  if board.x >= board.WIDTH {
    board.x = board.WIDTH - 1
  }
  return false, false
}
