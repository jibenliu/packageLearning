package deck_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"blackAndWhiteTest/deck"
)

func TestDeckCanDrawCards(t *testing.T) {
	var num uint8 = 10

	d := deck.NewDeck(num)
	for i := uint8(0); i < num; i++ {
		_, err := d.Draw()
		assert.Nil(t, err)
	}
	_, err := d.Draw()
	assert.Equal(t, err, deck.Empty)
}
