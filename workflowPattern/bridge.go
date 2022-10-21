package workflowPattern

import (
	"fmt"
	"strings"
)

// AText is Abstraction
type AText interface {
	getText() string
	addLine(value string)
}

// ITextImp is abstract Implementator
type ITextImp interface {
	getString() string
	appendLine(value string)
}

// TextImp is Implementator
type TextImp struct {
	rows []string
}

func (ti TextImp) getString() string {
	return strings.Join(ti.rows, "\n")
}

// BTextMaker RefinedAbstraction
type BTextMaker struct {
	textImp ITextImp
}

func (tm BTextMaker) getText() string {
	return tm.textImp.getString()
}

func (tm BTextMaker) addLine(value string) {
	tm.textImp.appendLine(value)
}

// BbTextBuilder is ConcreteImplementator1
type BbTextBuilder struct {
	TextImp
}

func (tb *BbTextBuilder) appendLine(value string) {
	tb.rows = append(tb.rows, value)
}

// BHTMLBuilder is ConcreteImplementator2
type BHTMLBuilder struct {
	TextImp
}

func (hb *BHTMLBuilder) appendLine(value string) {
	hb.rows = append(hb.rows,
		"<span>"+value+"</span><br/>")
}

func fillBbTextBuilder(textImp ITextImp) AText {
	BTextMaker := BTextMaker{textImp}
	BTextMaker.addLine("line 1")
	BTextMaker.addLine("line 2")
	return BTextMaker
}

func BridgeDemo() {
	//Client
	BTextMaker := fillBbTextBuilder(&BbTextBuilder{})
	text := BTextMaker.getText()
	//test: line 1
	//      line 2

	htmlMaker := fillBbTextBuilder(&BHTMLBuilder{})
	html := htmlMaker.getText()
	//html: <span>line 1</span><br/>
	//      <span>line 2</span><br/>

	fmt.Println(text)
	fmt.Println(html)
}
