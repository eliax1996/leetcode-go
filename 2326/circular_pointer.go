package circular_pointer

const (
	Right = 0
	Left  = 1
	Down  = 2
	Up    = 3
)

type Square struct {
	// y axis low
	xl int
	// y axis high
	xh int
	yl int
	yh int
}

func NewSquare(x, y int) Square {
	return Square{
		0,
		x,
		0,
		y,
	}
}

func (s *Square) In(x, y int) bool {
	if x < s.xl || x > s.xh || y < s.yl || y > s.yh {
		return false
	}
	return true
}

type Cursor struct {
	X         int
	Y         int
	limits    Square
	direction int
}

func NewCursor(bottomLimit, rightLimit int) Cursor {
	return Cursor{
		0,
		0,
		NewSquare(rightLimit, bottomLimit),
		Right,
	}
}

func derivative(direction int) (int, int) {
	switch direction {
	case Right:
		return 1, 0
	case Left:
		return -1, 0
	case Down:
		return 0, 1
	case Up:
		return 0, -1
	}
	panic("unexpected state")
}

func (c *Cursor) InBoundaries(dx, dy int) bool {
	return c.limits.In(c.X+dx, c.Y+dy)
}

func (c *Cursor) Next() *Cursor {
	dx, dy := derivative(c.direction)
	if !c.InBoundaries(dx, dy) {
		switch c.direction {
		case Left:
			c.direction = Up
			// since from left we are going to move up
			// we are never going to see a x value as low
			// as now
			c.limits.yh -= 1
		case Right:
			c.direction = Down
			// since we are going to move down we will never
			// reach again an y axis value as low as now.
			c.limits.yl += 1
		case Down:
			c.direction = Left
			// since we are going to move left we
			// are never going to experience a value y value
			// as high as now.
			c.limits.xh -= 1
		case Up:
			c.direction = Right
			// since we are going to move right
			// we are never going to se a x as low as now.
			c.limits.xl += 1
		}
		dx, dy = derivative(c.direction)
	}
	if !c.InBoundaries(dx, dy) {
		return nil
	}
	c.X += dx
	c.Y += dy
	return c
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	// fill with -1
	matrix := make([][]int, m)
	for i := range m {
		matrix[i] = make([]int, n)
		for j := range n {
			matrix[i][j] = -1
		}
	}

	c := NewCursor(m-1, n-1)
	matrix[c.Y][c.X] = head.Val

	for head = head.Next; head != nil && c.Next() != nil; head = head.Next {
		matrix[c.Y][c.X] = head.Val
	}

	return matrix
}
