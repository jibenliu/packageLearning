package k_means

import (
	"packageLearning/algorithm/k-means/km"
	"strings"
	"testing"
)

func Test_KMeans(t *testing.T) {
	// 创建样本点
	samples := []km.IPoint{
		km.NewPerson(2, 11),
		km.NewPerson(2, 8),
		km.NewPerson(2, 6),

		km.NewPerson(3, 12),
		km.NewPerson(3, 10),

		km.NewPerson(4, 7),
		km.NewPerson(4, 3),

		km.NewPerson(5, 11),
		km.NewPerson(5, 9),
		km.NewPerson(5, 2),

		km.NewPerson(7, 9),
		km.NewPerson(7, 6),
		km.NewPerson(7, 3),

		km.NewPerson(8, 12),

		km.NewPerson(9, 3),
		km.NewPerson(9, 5),
		km.NewPerson(9, 10),

		km.NewPerson(10, 3),
		km.NewPerson(10, 6),
		km.NewPerson(10, 12),

		km.NewPerson(11, 9),
	}

	fnPoints2String := func(points []km.IPoint) string {
		items := make([]string, len(points))
		for i, it := range points {
			items[i] = it.String()
		}
		return strings.Join(items, " ")
	}

	for k := 1; k <= 3; k++ {
		centers := km.KMeansClassifier.Classify(samples, km.PersonDistanceCalculator, k)
		t.Log(fnPoints2String(centers))
	}
}
