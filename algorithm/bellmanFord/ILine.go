package bellmanFord

type tLine struct {
	from   string
	to     string
	weight int
}

func NewLine(from string, to string, weight int) ILine {
	return &tLine{
		from, to, weight,
	}
}

func (t *tLine) From() string {
	return t.from
}

func (t *tLine) To() string {
	return t.to
}

func (t *tLine) Weight() int {
	return t.weight
}
