package km

type IClassifier interface {
	// Classify 将samples聚类成k个, 并返回k个中心点
	Classify(samples []IPoint, calc IDistanceCalculator, k int) []IPoint
}
