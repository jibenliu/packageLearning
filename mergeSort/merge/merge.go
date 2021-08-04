package merge

func update(finalArr []int, arr []int, index *int, incrementIndex *int) {
	finalArr[*index] = arr[*incrementIndex]
	*incrementIndex++
	*index++
}

func Merge(arr1 []int, arr2 []int) []int {
	size1 := len(arr1)
	size2 := len(arr2)
	finalArr := make([]int, size1+size2)
	i := 0
	j := 0
	index := 0
	for ; i < size1 && j < size2; {
		if arr1[i] < arr2[j] {
			update(finalArr, arr1, &index, &i)
		} else {
			update(finalArr, arr2, &index, &j)
		}
	}

	for ; i < size1; {
		update(finalArr, arr1, &index, &i)
	}
	for ; j < size2; {
		update(finalArr, arr2, &index, &j)
	}
	return finalArr
}
