package km

type tPersonDistanceCalculator struct {
}

var gMaxInt = 0x7fffffff_ffffffff

func newPersonDistanceCalculator() IDistanceCalculator {
	return &tPersonDistanceCalculator{}
}

func (c *tPersonDistanceCalculator) Calc(a, b IPoint) int {
	if a == b {
		return 0
	}

	p1, ok := a.(*tPerson)
	if !ok {
		return gMaxInt
	}

	p2, ok := b.(*tPerson)
	if !ok {
		return gMaxInt
	}

	dx := p1.x - p2.x
	dy := p1.y - p2.y

	d := dx*dx + dy*dy
	if d < 0 {
		panic(d)
	}
	return d
}

var PersonDistanceCalculator = newPersonDistanceCalculator()
