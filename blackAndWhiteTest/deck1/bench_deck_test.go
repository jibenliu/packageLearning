package deck1

import (
	"testing"
)

func BenchmarkGC(b *testing.B) {
	b.ReportAllocs()
	shuffled := 0

	for i := 0; i < b.N; i++ {
		d := NewDeck(100)
		_, _ = d.Draw()
		shuffled += int(d.shuffled)
	}

	b.ReportMetric(float64(shuffled)/float64(b.N), "shuffle/op")
}
