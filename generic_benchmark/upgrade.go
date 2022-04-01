package generic_benchmark

// SubtractAble is a type constraint that defines subtractAble dataTypes to be used in generic functions
type SubtractAble interface {
	int | int64 | float32
}

// MoveAble is the interace for moving an Entity
type MoveAble[S SubtractAble] interface {
	Move(S)
}

// Car is a Generic Struct with the type S to be defined
type Car[S SubtractAble] struct {
	Name          string
	DistanceMoved S
}

// Person is a struct that accepts a type definition at initialization
// And uses that Type as the data type for meters as input
// Person is a Generic Struct with the type S to be defined
type Person[S SubtractAble] struct {
	Name          string
	DistanceMoved S
}

func (p *Person[S]) Move(meters S) {
	p.DistanceMoved += meters
}
func (c *Car[S]) Move(meters S) {
	c.DistanceMoved += meters
}

// Move is a generic function that takes in a Generic MoveAble and moves it
func Move[S SubtractAble, V MoveAble[S]](v V, meters S) {
	v.Move(meters)
}

// CarRegular Below is the Type casting based Solution
type CarRegular struct {
	Name          string
	DistanceMoved int
}

// PersonRegular personRegular
type PersonRegular struct {
	Name          string
	DistanceMoved float32
}

func (p *PersonRegular) Move(meters float32) {
	p.DistanceMoved += meters
}

func (c *CarRegular) Move(meters int) {
	c.DistanceMoved += meters
}

func MoveRegular(v interface{}, distance float32) {
	switch v.(type) {
	case *PersonRegular:
		v.(*PersonRegular).Move(distance)
	case *CarRegular:
		v.(*CarRegular).Move(int(distance))
	default:
		// Handle Unsupported types, not needed by Generic solution as Compiler does this for you
	}
}
