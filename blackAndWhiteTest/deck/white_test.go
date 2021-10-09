package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeckShouldBeShuffledOnce(t *testing.T) {
	var num uint8 = 5

	d := NewDeck(num)
	assert.Equal(t, len(d.cards), int(num))
	assert.Equal(t, d.shuffled, false, "Deck should init as not shuffled")
	orderBefore := fmt.Sprint(d.cards)

	d.shuffle()
	assert.Equal(t, d.shuffled, true, "Deck has not been marked as shuffled")
	orderAfter := fmt.Sprint(d.cards)

	assert.NotEqual(t, orderBefore, orderAfter, "Deck once shuffled should have new card order")
}
