package workflowPattern

// Point is State
type Point struct {
	X, Y int
}

// Memento contains a State field
type Memento struct {
	_state Point
}

func (m Memento) getState() Point {
	return m._state
}

// Shape is Originator
type Shape struct {
	Position Point
}

// Move Shape
func (s *Shape) Move(left, top int) {
	s.Position.X += left
	s.Position.Y += top
}

func (s *Shape) getMemento() Memento {
	state := Point{s.Position.X, s.Position.Y}
	return Memento{state}
}

func (s *Shape) setMemento(memento Memento) {
	s.Position = memento.getState()
}

// ShapeHelper is Caretaker
type ShapeHelper struct {
	_shape *Shape
	_stack []Memento
}

// NewShapeHelper creates a new ShapeHelper
func NewShapeHelper(shape *Shape) ShapeHelper {
	return ShapeHelper{shape, []Memento{}}
}

// Move Shape and save prior state
func (sh *ShapeHelper) Move(left, top int) {
	sh._stack = append(sh._stack, sh._shape.getMemento())
	sh._shape.Move(left, top)
}

// Undo move shape to previous position
func (sh *ShapeHelper) Undo() {
	l := len(sh._stack)
	if l > 0 {
		memento := sh._stack[l-1]
		sh._stack = sh._stack[:l-1]
		sh._shape.setMemento(memento)
	}
}

func MementoDemo() {
	//Client
	shape := &Shape{}
	helper := NewShapeHelper(shape)

	helper.Move(2, 3)
	//shape.Position is (2, 3)
	helper.Move(-5, 4)
	//shape.Position is (-3, 7)

	helper.Undo()
	//shape.Position is (2, 3)
	helper.Undo()
	//shape.Position is (0, 0)
}
