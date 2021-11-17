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

type S1 struct {
	f float64
}

func (s *S1) Key() string {
	return strconv.FormatFloat(s.f, 'f', -1, 64)
}

func (s *S1) Score() float64 {
	return s.f
}

func (s *S1) Compare(value ISkipListNode) int {
	_s := value.(*S1)
	if s.f > _s.f {
		return 1
	}
	if s.f == _s.f {
		return 0
	}
	return -1
}

type SortS1 []*S1

func (s SortS1) Len() int {
	return len(s)
}

func (s SortS1) Less(i, j int) bool {
	return s[i].f < s[j].f
}

func (s SortS1) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SortS2 []*SkipListNode

func (s SortS2) Len() int {
	return len(s)
}

func (s SortS2) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s SortS2) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestNewSkipTable(t *testing.T) {
	st := NewDefaultSkipTable()

	s1 := insert(st, 10000)

	size := st.Size()
	fmt.Println(size)

	for _, s := range s1 {
		fmt.Printf("%f  ", s.f)
	}
	fmt.Println()
	list := st.GetNodeByRank(1, 10000)
	for _, node := range list {
		fmt.Printf("%f  ", node.score)
	}

	fmt.Println()
}

func TestSkipList_GetNodeByRank(t *testing.T) {
	st := NewDefaultSkipTable()

	s1 := insert(st, 10000)

	size := st.Size()
	fmt.Println(size)

	for _, s := range s1 {
		fmt.Printf("%f  ", s.f)

	}
	fmt.Println()
	list := st.GetNodeByRank(1, 10000)
	for _, node := range list {
		fmt.Printf("%f  ", node.score)
	}

	fmt.Println()
}

func TestSkipList_GetNodeByScore(t *testing.T) {
	st := NewDefaultSkipTable()

	const N = 100
	s1 := insert(st, N)

	size := st.Size()
	fmt.Println(size)

	//测试随机范围
	n := rand.Int31n(10)
	x := rand.Int31n(N - n)
	fmt.Println(n, x)

	start := s1[x].f
	end := s1[n+x].f

	fmt.Println("获取范围:", start, end)

	r := &SkipListFindRange{
		Min:    start,
		Max:    end,
		MinInf: true,
		MaxInf: false,
	}
	list := st.GetNodeByScore(r)
	for _, node := range list {
		fmt.Printf("%v  ", node.score)
	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------")

	//测试获取第一个
	fmt.Println("测试获取第一个")
	r.Min = s1[0].Score()
	r.Max = s1[0].Score()
	list = st.GetNodeByScore(r)
	fmt.Print(s1[0].Score(), "  ")
	for _, node := range list {
		fmt.Printf("%v  ", node.score)
	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------")

	//测试[-∞,+∞]
	fmt.Printf("测试 [-∞,+∞] [%v,%v]\n", s1[0].Score(), s1[N-1].Score())
	r.MinInf = true
	r.MaxInf = true
	list = st.GetNodeByScore(r)
	for _, node := range list {
		fmt.Printf("%v  ", node.score)
	}
	fmt.Println()
	fmt.Println("正常", N, "实际", len(list), "个")
	fmt.Println("---------------------------------------------------------")

	//测试[-∞,n]
	n = rand.Int31n(N)
	fmt.Printf("测试 [-∞,%d] [%v,%v]\n", n, s1[0].Score(), s1[n-1].Score())
	r.MaxInf = false
	r.MinInf = true
	r.Max = s1[n-1].Score()
	r.Min = 0
	list = st.GetNodeByScore(r)
	for _, node := range list {
		fmt.Printf("%v  ", node.score)
	}
	fmt.Println()
	fmt.Println("正常", n, "实际", len(list), "个")
	fmt.Println("---------------------------------------------------------")

	//测试 [n,+∞]
	n = rand.Int31n(N)
	fmt.Printf("测试 [%d,+∞] [%v,%v]\n", n, s1[n-1].Score(), s1[N-1].Score())
	r.MaxInf = true
	r.MinInf = false
	r.Max = 0
	r.Min = s1[n-1].Score()

	list = st.GetNodeByScore(r)
	for _, node := range list {
		fmt.Printf("%v  ", node.score)
	}
	fmt.Println()
	fmt.Println("正常", N-n+1, "实际", len(list), "个")
}

func TestSkipList_UpdateScore(t *testing.T) {
	st := NewDefaultSkipTable()

	const N = 10
	insert(st, N)
	fmt.Println()

	list := st.GetNodeByRank(1, N)
	for _, node := range list {
		fmt.Printf("%v  ", node.score)
	}
	fmt.Println()
	nodes := st.GetNodeByRank(1, 1)
	if len(nodes) == 0 {
		panic("sl GetNodeByRank 1 not find")
	}

	st.UpdateScore(nodes[0], -nodes[0].score)

	fmt.Println()
	list = st.GetNodeByRank(1, N)
	for _, node := range list {
		fmt.Printf("%v  ", node.score)
	}

	fmt.Println()
}

func TestSkipList_GetNodeRank(t *testing.T) {
	st := NewDefaultSkipTable()

	const N = 1000
	s1 := insert2(st, N)
	fmt.Println()

	n := int64(rand.Int31n(N))
	fmt.Println("n:", n+1)

	fmt.Println("rank", st.GetNodeRank(s1[n]))
}

func TestSkipList_Delete(t *testing.T) {
	st := NewDefaultSkipTable()

	const N = 100
	s1 := insert2(st, N)

	for _, node := range s1 {
		st.Delete(node, st.GetUpdateList(node))
	}

	s1 = insert2(st, N)
	fmt.Println(s1[0].score, st.GetNodeByRank(1, 1)[0].score)
	st.Delete(s1[0], st.GetUpdateList(s1[0]))
	fmt.Println(s1[1].score, st.GetNodeByRank(1, 1)[0].score)

	st.Delete(s1[N-1], st.GetUpdateList(s1[N-1]))
	fmt.Println(s1[N-2].score, st.GetNodeByRank(st.Size(), st.Size())[0].score)

	st.Delete(s1[N-1], st.GetUpdateList(s1[N-1]))
	fmt.Println(s1[N-2].score, st.GetNodeByRank(st.Size(), st.Size())[0].score)
}

/////////////////////////////////////////////////////////////////////
func insert(st *SkipList, size int) SortS1 {
	arrS := make(SortS1, 0, size)
	for i := 0; i < size; i++ {
		f := rand.Float64()
		s := &S1{f: f}
		st.InsertByScore(f, s)
		arrS = append(arrS, s)
	}
	sort.Sort(arrS)
	return arrS
}

func insert2(st *SkipList, size int) SortS2 {
	arrS := make(SortS2, 0, size)
	for i := 0; i < size; i++ {
		f := rand.Float64()
		s := &S1{f: f}
		arrS = append(arrS, st.InsertByScore(f, s))
	}
	sort.Sort(arrS)
	return arrS
}

func skipListDump(st *SkipList) {
	fmt.Println("*************SKIP LIST DUMP START*************")
	for i := st.level - 1; i >= 0; i-- {
		fmt.Printf("level:--------%v--------\n", i)
		x := st.head
		for x != nil {
			if x == st.head {
				fmt.Printf("Head span: %d\n", x.Span(i))
			} else {
				fmt.Printf("span: %d value : %v\n", x.Span(i), x)
			}
			x = x.Next(i)
		}
	}
	fmt.Println("*************SKIP LIST DUMP END*************")
}
