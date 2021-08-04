package main

//工厂模式
type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	println("Inside Rectangle::draw() method")
}

type Square struct {
}

func (r *Square) Draw() {
	println("Inside Square ::draw() method.")
}

type Circle struct {
}

func (c *Circle) Draw() {
	println("Inside Circle  ::draw() method.")
}

type ShapeFactory struct {
}

//使用 getShape 方法获取形状类型的对象
func (s *ShapeFactory) getShape(shapeType string) Shape {
	if shapeType == "" {
		return nil
	}
	if shapeType == "CIRCLE" {
		return &Circle{}
	} else if shapeType == "RECTANGLE" {
		return &Rectangle{}
	} else if shapeType == "SQUARE" {
		return &Square{}
	}
	return nil
}

func main() {
	factory := ShapeFactory{}
	factory.getShape("CIRCLE").Draw()
	factory.getShape("RECTANGLE").Draw()
	factory.getShape("SQUARE").Draw()
}
