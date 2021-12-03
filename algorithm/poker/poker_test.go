package poker

import (
	"testing"
)

func TestCreateNew(t *testing.T) {
	t.Logf("%v", CreateNew())
}

func TestShuffle(t *testing.T) {
	values := CreateNew()
	t.Logf("初始牌：%v", values)
	Shuffle(values)
	t.Logf("打乱后：%v", values)
}

func TestDispatcher(t *testing.T) {
	initValues := CreateNew()
	t.Logf("初始牌：%v", initValues)
	Shuffle(initValues)
	t.Logf("打乱后：%v", initValues)
	t.Logf("玩家1：%v", Dispatcher(0, initValues))
	t.Logf("玩家2：%v", Dispatcher(1, initValues))
	t.Logf("玩家3：%v", Dispatcher(2, initValues))
	t.Logf("底牌：%v", Dispatcher(3, initValues))
}

func TestParseCardsInSize(t *testing.T) {
	cardsA := []string{"A3", "B3", "C3", "A4", "B4", "C4", "A5", "B5", "A5", "A6", "B6", "A6", "A11", "A7", "B12", "B7"}
	aShowMode := ParseCardsInSize(cardsA)
	t.Logf("A玩家:%s", aShowMode.CardTypeStatus)

	cardsB := []string{"A3", "B3", "C3", "A4", "B4", "C4", "A5", "B5", "A5", "A6", "B6", "A6", "A11", "A7", "B12"}
	bShowMode := ParseCardsInSize(cardsB)
	t.Logf("A玩家:%s", bShowMode.CardTypeStatus)

	cardsC := []string{"A3", "B3", "C3", "A4", "B4", "C4", "A5", "B5", "A5", "A6", "B6", "A6"}
	cShowMode := ParseCardsInSize(cardsC)
	t.Logf("A玩家:%s", cShowMode.CardTypeStatus)


	cardsD := []string{"A3", "B4", "C5", "A6", "B7"}
	dShowMode := ParseCardsInSize(cardsD)
	t.Logf("A玩家:%s", dShowMode.CardTypeStatus)


	cardsA = []string{"A3", "B4", "C5", "A6"}
	cardsB = []string{"A3", "B4", "C5", "A8"}
	ashowMode := ParseCardsInSize(cardsA)
	bshowMode := ParseCardsInSize(cardsB)
	t.Logf("A玩家:%s", ashowMode.CardTypeStatus)
	t.Logf("B玩家:%s", bshowMode.CardTypeStatus)


	cardsA = []string{"A3", "B3", "C3", "D3"}
	ashowMode = ParseCardsInSize(cardsA)
	t.Logf("A玩家:%s", ashowMode.CardTypeStatus)


	cardsA = []string{"A3", "B3", "C3", "A4", "B4", "C4", "A5", "B5", "A5", "A6", "B6", "A6", "A11", "A7", "B12", "B7"}
	ashowMode = ParseCardsInSize(cardsA)
	cardsB = []string{"A4", "B4", "C4", "A5", "B5", "C5", "A6", "B6", "A6", "A7", "B7", "A7", "A11", "A10", "B12", "13"}
	bshowMode = ParseCardsInSize(cardsB)
	t.Logf("A玩家:%d", ashowMode.CompareValue)
	t.Logf("B玩家:%d", bshowMode.CompareValue)
}
