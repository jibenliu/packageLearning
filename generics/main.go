package generics

type Addable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 | complex64 | complex128 | string
}

func Print[T Addable](value T, nums ...int) {
	print(value)
	print(nums)
}

func findFunc[T Addable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}

func main() {
	print(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
}
