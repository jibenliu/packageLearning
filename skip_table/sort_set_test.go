package skiptable

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type StItem struct {
	f float64
	k string
}

func (s *StItem) Key() string {
	return strconv.FormatFloat(s.f, 'f', -1, 64) + s.k
}

func (s *StItem) Score() float64 {
	return s.f
}

func (s *StItem) Compare(value ISkipListNode) int {
	_s := value.(*StItem)
	if s.Key() > _s.Key() {
		return 1
	}
	if s.Key() == _s.Key() {
		return 0
	}
	return -1
}

type SortS3 []ISortSet

func (s SortS3) Len() int {
	return len(s)
}

func (s SortS3) Less(i, j int) bool {
	if s[i].Score() < s[j].Score() {
		return true
	}
	if s[i].Score() == s[j].Score() {
		if s[i].Compare(s[j]) < 1 {
			return true
		}
	}
	return false
}

func (s SortS3) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//////////////////////////////////////////////////////////////////
func TestNewSortSet(t *testing.T) {
	set := NewDefaultSortSet()
	const N = 50
	s3 := insert3(set, N)
	for _, set := range s3 {
		fmt.Printf("%v  ", set.Key())
	}
	fmt.Println()
	result := set.Range(0, -1)
	for _, set := range result {
		fmt.Printf("%v  ", set.Key())
	}
	fmt.Println()

	//清空
	set.RemoveRangeByRank(0, -1)

	s3 = insert4(set, 50)
	for _, set := range s3 {
		fmt.Printf("%v  ", set.Key())
	}
	fmt.Println()
	result = set.Range(0, -1)
	for _, set := range result {
		fmt.Printf("%v  ", set.Key())
	}
	fmt.Println()
}

func TestSortSet_Rank(t *testing.T) {
	set := NewDefaultSortSet()
	const N = 100
	s4 := insert4(set, N)
	for i, set := range s4 {
		fmt.Printf("%d->%v  ", i, set.Key())
	}
	fmt.Println()
	n := rand.Int31n(N)
	fmt.Println("rand n:", n, s4[n].Key())
	rank := set.Rank(s4[n].Key())

	fmt.Println("rand n", n, " rank", rank)
}

func TestSortSet_RangeByScore(t *testing.T) {
	set := NewDefaultSortSet()
	const N = 100
	s3 := insert3(set, N)
	//测试随机范围
	n := rand.Int31n(10)
	x := rand.Int31n(N - n)
	fmt.Println(n, x)

	start := s3[x].Score()
	end := s3[n+x].Score()
	findRange := &SkipListFindRange{
		Min:    start,
		Max:    end,
		MinInf: false,
		MaxInf: false,
	}
	fmt.Printf("测试随机取范围 [%v,%v]\n", start, end)
	result := set.RangeByScore(findRange)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()

	//测试范围
	n = rand.Int31n(N)
	fmt.Printf("测试随机取范围 [-∞,+∞]\n")
	findRange.MinInf = true
	findRange.Min = 0
	findRange.MaxInf = true
	findRange.Max = 0
	result = set.RangeByScore(findRange)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()

	//测试范围
	n = rand.Int31n(N)
	fmt.Printf("测试随机取范围 [-∞,%v]\n", s3[n].Score())
	findRange.MinInf = true
	findRange.Min = 0
	findRange.MaxInf = false
	findRange.Max = s3[n].Score()
	result = set.RangeByScore(findRange)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()

	//测试范围
	n = rand.Int31n(N)
	fmt.Printf("测试随机取范围 [%v,+∞]\n", s3[n].Score())
	findRange.MinInf = false
	findRange.Min = s3[n].Score()
	findRange.MaxInf = true
	findRange.Max = 0
	result = set.RangeByScore(findRange)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
}
func TestSortSet_RevRangeByScore(t *testing.T) {
	set := NewDefaultSortSet()
	const N = 100
	s3 := insert3(set, N)
	//测试随机范围
	n := rand.Int31n(10)
	x := rand.Int31n(N - n)
	fmt.Println(n, x)

	start := s3[x].Score()
	end := s3[n+x].Score()
	findRange := &SkipListFindRange{
		Min:    end,
		Max:    start,
		MinInf: false,
		MaxInf: false,
	}
	fmt.Printf("测试随机取范围 [%v,%v]\n", end, start)
	result := set.RevRangeByScore(findRange)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()

	//测试范围
	n = rand.Int31n(N)
	fmt.Printf("测试随机取范围 [-∞,+∞]\n")
	findRange.MinInf = true
	findRange.Min = 0
	findRange.MaxInf = true
	findRange.Max = 0
	result = set.RevRangeByScore(findRange)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()

	//测试范围
	n = rand.Int31n(N)
	fmt.Printf("测试随机取范围 [-∞,%v]\n", s3[n].Score())
	findRange.MinInf = true
	findRange.Min = 0
	findRange.MaxInf = false
	findRange.Max = s3[n].Score()
	result = set.RevRangeByScore(findRange)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()

	//测试范围
	n = rand.Int31n(N)
	fmt.Printf("测试随机取范围 [%v,+∞]\n", s3[n].Score())
	findRange.MinInf = false
	findRange.Min = s3[n].Score()
	findRange.MaxInf = true
	findRange.Max = 0
	result = set.RevRangeByScore(findRange)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
}

func TestSortSet_Range(t *testing.T) {
	set := NewDefaultSortSet()
	const N = 100
	s3 := insert3(set, N)

	fmt.Println()

	//测试随机范围
	n := rand.Int63n(20)
	x := rand.Int63n(N - n)
	fmt.Printf("minRank:%v, maxRank:%v\n", x, n+x)
	fmt.Printf("min:%v, max:%v\n", s3[x].Score(), s3[x+n].Score())
	result := set.Range(x, x+n)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
	fmt.Println("-----------------------------------------------------")

	//测试[-rand, -rand]
	n = rand.Int63n(20)
	x = rand.Int63n(N)
	fmt.Printf("minRank:%v, maxRank:%v    minRank:%v, maxRank:%v \n", -(x + n), -x, N-(x+n), N-x)
	fmt.Printf("min:%v, max:%v\n", s3[N-(x+n)+1].Score(), s3[N-x+1].Score())
	result = set.Range(-(x + n), -x)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
	fmt.Println("-----------------------------------------------------")
}
func TestSortSet_RevRange(t *testing.T) {
	set := NewDefaultSortSet()
	const N = 100
	s3 := insert3(set, N)
	sort.Sort(sort.Reverse(s3))
	//for i, set := range s3 {
	//	fmt.Printf("%d->%v  ", i, set.Score())
	//}
	fmt.Println()

	//测试随机范围
	n := rand.Int63n(20)
	x := rand.Int63n(N - n)
	fmt.Printf("minRank:%v, maxRank:%v\n", x, n+x)
	fmt.Printf("min:%v, max:%v\n", s3[x].Score(), s3[x+n].Score())
	result := set.RevRange(x, x+n)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
	fmt.Println("-----------------------------------------------------")

	//测试[-rand, -rand]
	n = rand.Int63n(20)
	x = rand.Int63n(N)
	fmt.Printf("minRank:%v, maxRank:%v    minRank:%v, maxRank:%v \n", -(x + n), -x, N-(x+n), N-x)
	fmt.Printf("min:%v, max:%v\n", s3[N-(x+n)].Score(), s3[N-x].Score())
	result = set.RevRange(99, 100)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
	fmt.Println("-----------------------------------------------------")
}

func TestSortSet_RemoveRangeByRank(t *testing.T) {
	set := NewDefaultSortSet()
	const N = 100
	s3 := insert3(set, N)
	sort.Sort(sort.Reverse(s3))
	for i, set := range s3 {
		fmt.Printf("%d->%v  ", i, set.Score())
	}
	fmt.Println()
	result := set.Range(0, -1)
	fmt.Println("原始")
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
	fmt.Println("删除 10,20")
	set.RemoveRangeByRank(10, 20)
	result = set.Range(0, -1)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()

	fmt.Println("删除 30,-1")
	set.RemoveRangeByRank(30, -1)
	result = set.Range(0, -1)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()

	fmt.Println("删除 -10,-1")
	set.RemoveRangeByRank(-10, -1)
	result = set.Range(0, -1)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
	fmt.Println("删除 0,-1")
	set.RemoveRangeByRank(0, -1)
	result = set.Range(0, -1)
	if len(result) == 0 {
		fmt.Println("set is nil")
	}
	fmt.Println()

	fmt.Println("重新插入数据")
	insert3(set, 100)
	result = set.Range(0, -1)
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
}

func TestSortSet_RemoveRangeByScore(t *testing.T) {
	set := NewDefaultSortSet()
	const N = 100
	s3 := insert3(set, N)
	sort.Sort(sort.Reverse(s3))
	for i, set := range s3 {
		fmt.Printf("%d->%v  ", i, set.Score())
	}
	fmt.Println()
	result := set.Range(0, -1)
	fmt.Println("原始")
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println()
}

func TestSortSet_Big(t *testing.T) {
	set := NewDefaultSortSet()
	const N = 5000000

	arr := make(SortS3, 0, N)
	for _, x := range rand.Perm(N) {
		s := &StItem{
			f: float64(x),
			k: "A",
		}
		arr = append(arr, s)
	}
	t1 := time.Now()
	set.Add(arr...)
	t2 := time.Now()
	fmt.Println("插入用时:", t2.Sub(t1).Seconds())

	sort.Sort(arr)
	fmt.Println()
	t1 = time.Now()
	result := set.Range(N-10, -1)
	t2 = time.Now()
	fmt.Println("最后10个")
	for _, set := range result {
		fmt.Printf("%v  ", set.Score())
	}
	fmt.Println("查找用时:", t2.Sub(t1).Milliseconds())

	fmt.Println("get rank")
	n := rand.Int31n(N)
	fmt.Println("rand n:", n, "key", arr[n].Key())
	t1 = time.Now()
	rank := set.Rank(arr[n].Key())
	t2 = time.Now()
	fmt.Println("get rank:", rank, "查找用时:", t2.Sub(t1).Milliseconds())
}

func insert3(set *SortSet, size int) SortS3 {
	arrS := make(SortS3, 0, size)
	for _, x := range rand.Perm(size) {
		s := &StItem{
			f: float64(x),
			k: "A",
		}
		set.Add(s)
		arrS = append(arrS, s)
	}
	sort.Sort(arrS)
	return arrS
}

//包含重复项
func insert4(set *SortSet, size int) SortS3 {
	arrS := make(SortS3, 0, size)
	for _, x := range rand.Perm(size / 2) {
		s := &StItem{
			f: float64(x),
			k: "A",
		}
		set.Add(s)
		arrS = append(arrS, s)
	}
	for _, x := range rand.Perm(size / 2) {
		s := &StItem{
			f: float64(x),
			k: "B",
		}
		set.Add(s)
		arrS = append(arrS, s)
	}
	sort.Sort(arrS)
	return arrS
}
