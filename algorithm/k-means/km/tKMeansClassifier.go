package km

import (
	"math/rand"
	"time"
)

type tKMeansClassifier struct {
}

type tPointEntry struct {
	point    IPoint
	distance int
	index    int
}

func newPointEntry(p IPoint, d int, i int) *tPointEntry {
	return &tPointEntry{
		p, d, i,
	}
}

func newKMeansClassifier() IClassifier {
	return &tKMeansClassifier{}
}

// Classify 将samples聚类成k个, 并返回k个中心点
func (c *tKMeansClassifier) Classify(samples []IPoint, calc IDistanceCalculator, k int) []IPoint {
	sampleCount := len(samples)
	if sampleCount <= k {
		return samples
	}

	// 初始化, 随机选择k个中心点
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	centers := make([]IPoint, k)
	for selected, i := make(map[int]bool, 0), 0; i < k; {
		n := rnd.Intn(sampleCount)
		_, ok := selected[n]

		if !ok {
			selected[n] = true
			centers[i] = samples[n]
			i++
		}
	}

	// 根据到中心点的距离, 划分samples
	for {
		groups := c.split(samples, centers, calc)

		newCenters := make([]IPoint, k)
		for i, g := range groups {
			newCenters[i] = c.centerOf(g, calc)
		}

		if c.groupEquals(centers, newCenters) {
			return centers
		}
		centers = newCenters
	}
}

// 将样本点距离中心点的距离进行分簇
func (c *tKMeansClassifier) split(samples []IPoint, centers []IPoint, calc IDistanceCalculator) [][]IPoint {
	k := len(centers)
	result := make([][]IPoint, k)
	for i := 0; i < k; i++ {
		result[i] = make([]IPoint, 0)
	}

	entries := make([]*tPointEntry, k)
	for i, c := range centers {
		entries[i] = newPointEntry(c, 0, i)
	}

	for _, p := range samples {
		for _, e := range entries {
			e.distance = calc.Calc(p, e.point)
		}

		center := c.min(entries)
		result[center.index] = append(result[center.index], p)
	}

	return result
}

// 计算一簇样本的重心. 重心就是距离各点的总和最小的点
func (c *tKMeansClassifier) centerOf(samples []IPoint, calc IDistanceCalculator) IPoint {
	entries := make([]*tPointEntry, len(samples))
	for i, src := range samples {
		distance := 0
		for _, it := range samples {
			distance += calc.Calc(src, it)
		}
		entries[i] = newPointEntry(src, distance, i)
	}

	return c.min(entries).point
}

// 判断两组点是否相同
func (c *tKMeansClassifier) groupEquals(g1, g2 []IPoint) bool {
	if len(g1) != len(g2) {
		return false
	}

	for i, v := range g1 {
		if g2[i] != v {
			return false
		}
	}

	return true
}

// 查找距离最小的点
func (c *tKMeansClassifier) min(entries []*tPointEntry) *tPointEntry {
	minI := 0
	minD := gMaxInt
	for i, it := range entries {
		if it.distance < minD {
			minI = i
			minD = it.distance
		}
	}

	return entries[minI]
}

var KMeansClassifier = newKMeansClassifier()
