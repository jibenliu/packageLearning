package generic_benchmark

// SubtractInt will subtract the second value from the first
func SubtractInt(a, b int) int {
	return a - b
}

// SubtractInt64 will subtract the second value from the first
func SubtractInt64(a, b int64) int64 {
	return a - b
}

// SubtractFloat32 will subtract the second value from the first
func SubtractFloat32(a, b float32) float32 {
	return a - b
}

// SubtractTypeSwitch is used to subtract using interfaces
func SubtractTypeSwitch(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) - b.(int)
	case int64:
		return a.(int64) - b.(int64)
	case float32:
		return a.(float32) - b.(float32)
	default:
		return nil
	}
}

// Subtract will subtract the second value from the first
func Subtract[V int64 | int | float32](a, b V) V {
	return a - b
}
