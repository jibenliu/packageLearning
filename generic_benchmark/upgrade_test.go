package generic_benchmark

import "testing"

func Benchmark_Structures(b *testing.B) {
	// Init the structs
	p := &Person[float32]{Name: "John"}
	c := &Car[int]{Name: "Ferrari"}

	pRegular := &PersonRegular{Name: "John"}
	cRegular := &CarRegular{Name: "Ferrari"}

	// Run the test
	b.Run("Person_Generic_Move", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// generic will try to use float64 if we dont tell it is a float32
			Move[float32](p, 10.2)
		}
	})

	b.Run("Car_Generic_Move", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Move[int](c, 10)
		}
	})

	b.Run("Person_Regular_Move", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MoveRegular(pRegular, 10.2)
		}
	})

	b.Run("Car_Regular_Move", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MoveRegular(cRegular, 10)
		}
	})
}

// go test -v -bench=Benchmark_Structures -benchtime=1000000000x -count 5
