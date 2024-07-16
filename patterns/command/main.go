// Пример с использовванием паттерна для создания команд "Копировать" и "Вставить" в текстовом редакторе.

package main

import "fmt"

// Command интерфейс для всех команд
type Command interface {
	Execute()
	Undo()
}

// Receiver - получатель команды
type TextEditor struct {
	content string
}

func (e *TextEditor) Append(text string) {
	e.content += text
}

func (e *TextEditor) DeleteLast(n int) {
	if len(e.content) >= n {
		e.content = e.content[:len(e.content)-n]
	}
}

func (e *TextEditor) ShowContent() {
	fmt.Println("Content:", e.content)
}

// Concrete Command - команда копирования
type CopyCommand struct {
	editor *TextEditor
	text   string
}

func NewCopyCommand(editor *TextEditor, text string) *CopyCommand {
	return &CopyCommand{editor: editor, text: text}
}

func (c *CopyCommand) Execute() {
	c.editor.Append(c.text)
}

func (c *CopyCommand) Undo() {
	c.editor.DeleteLast(len(c.text))
}

// Concrete Command - команда вставки
type PasteCommand struct {
	editor *TextEditor
	text   string
}

func NewPasteCommand(editor *TextEditor, text string) *PasteCommand {
	return &PasteCommand{editor: editor, text: text}
}

func (p *PasteCommand) Execute() {
	p.editor.Append(p.text)
}

func (p *PasteCommand) Undo() {
	p.editor.DeleteLast(len(p.text))
}

func main() {
	editor := &TextEditor{}
	copyCmd := NewCopyCommand(editor, "Hello, ")
	pasteCmd := NewPasteCommand(editor, "World!")

	copyCmd.Execute()
	editor.ShowContent()

	pasteCmd.Execute()
	editor.ShowContent()

	pasteCmd.Undo()
	editor.ShowContent()

	copyCmd.Undo()
	editor.ShowContent()
}
