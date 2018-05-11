package main

import (
  "testing"
  "reflect"
)

func TestSetAndGetPoint(t *testing.T) {
  b := newBlankBoard(2, 2)
  b.SetPoint(1, 1, Bomb)
  if b.GetPoint(1, 1) != Bomb {
    t.Errorf("Point{1, 1} was incorect, got: %v, want: %v.", b.GetPoint(1, 1), Bomb)
  }
  b.SetPoint(0, 0, Bomb)
  if b.GetPoint(0, 0) != Bomb {
    t.Errorf("Point{0, 0} was incorect, got: %v, want: %v.", b.GetPoint(0, 0), Bomb)
  }
}

func TestSearchCorner(t *testing.T) {
  b := newBlankBoard(3, 3)
  b.SetPoint(2, 2, Bomb)
  b.x = 0
  b.y = 0
  b.Search()
  wantedBoard := newBlankBoard(3, 3)
  wantedBoard.x = 0
  wantedBoard.y = 0
  a := [...]Tile {Explored_0, Explored_0, Explored_0,
                  Explored_0, Explored_1, Explored_1,
                  Explored_0, Explored_1, Bomb}
  for i, e := range a {
    wantedBoard.grid[i] = e
  }
  t.Logf("%v", reflect.DeepEqual(*wantedBoard, *b))
  if !reflect.DeepEqual(*wantedBoard, *b) {
    t.Errorf("Search was incorrect. Got: \n%#v, want: \n%#v.", b, wantedBoard)
  }
}

func TestSearchMiddle(t *testing.T) {
  b := newBlankBoard(3, 3)
  b.SetPoint(1, 1, Bomb)
  b.x = 0
  b.y = 0
  b.Search()
  wantedBoard := newBlankBoard(3, 3)
  a := [...]Tile {Explored_1, Unknown   , Unknown   ,
                  Unknown   , Bomb      , Unknown   ,
                  Unknown   , Unknown   , Unknown   }
  for i, e := range a {
    wantedBoard.grid[i] = e
  }
  if !reflect.DeepEqual(*wantedBoard, *b) {
    t.Errorf("Search was incorrect. Got: \n%#v, want: \n%#v.", b, wantedBoard)
  }
}

func TestGetValueOfPoint(t *testing.T) {
  b := newBlankBoard(3, 3) //  B: Bomb, S: Search location, U: Unknown
  b.SetPoint(2, 2, Bomb)
  b.SetPoint(1, 0, Bomb) // S B U
  b.SetPoint(0, 1, Bomb) // B U U
  b.SetPoint(1, 2, Bomb) // U B B  to make sure it gets the right area and doesn't check negatives
  if b.GetValueOfPoint(0, 0) != 2 {
    t.Errorf("GetValueOfPoint was incorrect. Got: %d, want: %v.", b.GetValueOfPoint(1, 1), 2)
  }
}

func TestOutOfBoard(t *testing.T) {
  b := newBlankBoard(3, 3)
  if OutOfBoard(*b, 3, 2) != true { // true is outside board
    t.Errorf("OutOfBoard was incorrect. Got: %v, want: %v.", false, true)
  }
  if OutOfBoard(*b, 2, 2) != false {  // false is inside board
    t.Errorf("OutOfBoard was incorrect. Got: %v, want: %v.", true, false)
  }
}
