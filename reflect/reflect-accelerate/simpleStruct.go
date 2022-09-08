package reflect_accelerate

type SimpleStruct struct {
	A int
	B int
}

func populateStruct(in *SimpleStruct) {
	in.B = 42
}
