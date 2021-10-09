package deck1

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"math/rand"
)

var Empty = errors.New("Empty deck")

type Deck struct {
	cards    []uint8
	shuffled uint
}

func NewDeck(numbers uint8) *Deck {
	cards := make([]uint8, 0, numbers)
	for i := uint8(0); i < numbers; i++ {
		cards = append(cards, i+1)
	}

	d := Deck{cards: cards}

	var n uint32
	binary.Read(crand.Reader, binary.LittleEndian, &n)
	for n = 1 + n%20; n > 0; n-- {
		d.shuffle()
	}

	return &d
}

func (d *Deck) Draw() (card uint8, err error) {
	if len(d.cards) == 0 {
		return 0, Empty
	}
	card, d.cards = d.cards[0], d.cards[1:]

	return card, nil
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	d.shuffled++
}
