package circular_pointer

import (
	"fmt"
	"testing"
)

func TestPointIn(t *testing.T) {
	square := NewSquare(4, 4)
	pointsInside := [][]int{
		{2, 2},
		{3, 3},
		{4, 4},
		{0, 0},
		{0, 3},
		{3, 0},
	}
	pointsOutside := [][]int{
		{5, 5},
		{5, 1},
		{-1, 0},
		{-1, -1},
		{15, 18},
	}
	for _, p := range pointsInside {
		if !square.In(p[0], p[1]) {
			t.Errorf("(%d,%d) should be inside of %+v", p[0], p[1], square)
		}
	}
	for _, p := range pointsOutside {
		if square.In(p[0], p[1]) {
			t.Errorf("(%d,%d) should be outside of %+v", p[0], p[1], square)
		}
	}
}

var expectedPos = [][2]int{
	{0, 0},
	{1, 0},
	{2, 0},
	{2, 1},
	{2, 2},
	{1, 2},
	{0, 2},
	{0, 1},
	{1, 1},
}

func TestCircularPointer(t *testing.T) {
	testCursor := NewCursor(2, 2)
	fmt.Printf("%+v\n", testCursor)
	for n, exp := range expectedPos {
		x, y := exp[0], exp[1]
		if x != testCursor.X || y != testCursor.Y {
			t.Errorf("Expected (%d,%d) but the cursor was at (%d, %d)", x, y, testCursor.X, testCursor.Y)
		}
		if n == len(expectedPos)-1 {
			if testCursor.Next() != nil {
				t.Errorf("The cursor should finish now")
			}
		} else {
			if nil == testCursor.Next() {
				t.Errorf("The cursor finished before the expected navigation")
			}
		}
		fmt.Printf("%+v\n", testCursor)
	}
}
