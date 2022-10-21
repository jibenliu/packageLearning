package workflowPattern

// IntIterator is Iterator
type IntIterator interface {
	First()
	Next()
	IsDone() bool
	CurrentItem() int
}

// Numbers is ConcreteAggregate
type Numbers struct {
	Data []int
}

// GetIterator return Iterator
func (n Numbers) GetIterator() IntIterator {
	return &Iterator{n, 0}
}

// Iterator is ConcreteIterator
type Iterator struct {
	_numbers Numbers
	_index   int
}

// First positions the iterator to the first element
func (i *Iterator) First() {
	i._index = 0
}

// Next advances the current element
func (i *Iterator) Next() {
	i._index++
}

// IsDone checks whether the index refers to an element withinthe List
func (i *Iterator) IsDone() bool {
	return i._index >= len(i._numbers.Data)
}

// CurrentItem returns the item at the current index
func (i *Iterator) CurrentItem() int {
	return i._numbers.Data[i._index]
}

func IteratorDemo() {
	//Client
	numbers := Numbers{[]int{2, 3, 5, 7, 11}}
	iterator := numbers.GetIterator()
	sum := 0
	for iterator.First(); !iterator.IsDone(); iterator.Next() {
		sum += iterator.CurrentItem()
	}
	//sum is 28
}
