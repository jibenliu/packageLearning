package km

type IDistanceCalculator interface {
	Calc(a, b IPoint) int
}
