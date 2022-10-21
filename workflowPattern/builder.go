package workflowPattern

import "fmt"

// TextWorker is AbstractBuilder
type TextWorker interface {
	AddText(text string)
	AddNewLine(text string)
}

// TextBuilder is ConcreteBuilder 1
type TextBuilder struct {
	Text string
}

// AddText adds text to the current line
func (tb *TextBuilder) AddText(text string) {
	tb.Text += text
}

// AddNewLine adds new line
func (tb *TextBuilder) AddNewLine(text string) {
	tb.Text += ("\n" + text)
}

// HTMLBuilder is ConcreteBuilder 2
type HTMLBuilder struct {
	HTML string
}

// AddText adds span to the current line
func (tb *HTMLBuilder) AddText(text string) {
	tb.HTML += ("<span>" + text + "</span>")
}

// AddNewLine adds new line
func (tb *HTMLBuilder) AddNewLine(text string) {
	tb.HTML += "<br/>\n"
	tb.AddText(text)
}

// TextMaker is Director
type TextMaker struct{}

// MakeText fills the text
func (tm TextMaker) MakeText(textBuilder TextWorker) {
	textBuilder.AddText("line 1")
	textBuilder.AddNewLine("line 2")
}

func BuilderDemo() {
	//Client
	textMaker := TextMaker{}

	textBuilder := TextBuilder{}
	textMaker.MakeText(&textBuilder)
	text := textBuilder.Text
	//text: line 1
	//      line 2

	htmlBuilder := HTMLBuilder{}
	textMaker.MakeText(&htmlBuilder)
	html := htmlBuilder.HTML
	//html: <span>line 1</span><br/>
	//      <span>line 2</span>

	fmt.Println(text)
	fmt.Println(html)
}
