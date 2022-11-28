package rangeAndForBench

func CreateABigSlice(count int) [][4096]int {
	ret := make([][4096]int, count)
	for i := 0; i < count; i++ {
		ret[i] = [4096]int{}
	}
	return ret
}
