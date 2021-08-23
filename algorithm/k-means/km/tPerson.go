package km

import "fmt"

type tPerson struct {
	x int
	y int
}

func NewPerson(x, y int) IPoint {
	return &tPerson{x, y}
}

func (p *tPerson) String() string {
	return fmt.Sprintf("p(%v,%v)", p.x, p.y)
}
