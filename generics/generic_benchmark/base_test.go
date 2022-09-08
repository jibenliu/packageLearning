package generic_benchmark

import (
	"math/rand"
	"testing"
	"time"
)

// Benchmark_Subtract is used to determine the most performant solution to subtraction
func Benchmark_Subtract(b *testing.B) {
	// Create a slice of random numbers based on the number of iterations set
	// to test the performance of the function
	// Default iterations for me is 1000000000
	// b.N is always 1 so we can use that to set the number of iterations
	numbers := make([]int, 1000000001)
	floatNumbers := make([]float32, 1000000001)
	// Create a random seed

	seed := rand.NewSource(time.Now().UnixNano())
	// Give the seed to the random package
	randomizer := rand.New(seed)
	for i := 0; i < b.N; i++ {
		// randomize numbers between 0-100
		numbers[i] = randomizer.Intn(100)
		floatNumbers[i] = float32(randomizer.Intn(100))
	}
	// run a benchmark for regular Ints
	b.Run("SubtractInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SubtractInt(numbers[i], numbers[i+1])
		}
	})
	// run a benchmark for regular Floats
	b.Run("SubtractFloat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SubtractFloat32(floatNumbers[i], floatNumbers[i+1])
		}
	})
	// run a benchmark for TypeSwitched Ints
	b.Run("Type_Subtraction_int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SubtractTypeSwitch(numbers[i], numbers[i+1])
		}
	})
	// run a benchmark for TypeSwitched Floats
	b.Run("Type_Subtraction_float", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SubtractTypeSwitch(floatNumbers[i], floatNumbers[i+1])
		}
	})

	// run a benchmark for Generic Ints
	b.Run("Generic_Subtraction_int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Subtract[int](numbers[i], numbers[i+1])
		}
	})
	// run a benchmark for Generic Floats
	b.Run("Generic_Subtraction_float", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Subtract[float32](floatNumbers[i], floatNumbers[i+1])
		}
	})
	// run a benchmark where generic type is infered
	b.Run("Generic_Inferred_int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Subtract(numbers[i], numbers[i+1])
		}
	})
}

// go test -v -bench=Benchmark -benchtime=1000000000x -count 5
