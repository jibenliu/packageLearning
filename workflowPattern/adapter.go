package workflowPattern

import (
	"fmt"
	"strings"
)

// StringList is Adaptee
type StringList struct {
	rows []string
}

func (sl StringList) getString() string {
	return strings.Join(sl.rows, "\n")
}

func (sl *StringList) add(value string) {
	sl.rows = append(sl.rows, value)
}

// TextAdapter is Adapter
type TextAdapter struct {
	RowList StringList
}

func (ta TextAdapter) getText() string {
	return ta.RowList.getString()
}

func getTextAdapter() TextAdapter {
	rowList := StringList{}
	rowList.add("line 1")
	rowList.add("line 2")
	adapter := TextAdapter{rowList}
	return adapter
}

func AdapterDemo() {
	//Client
	adapter := getTextAdapter()
	text := adapter.getText()
	//text: line 1
	//      line 2

	fmt.Println(text)
}
