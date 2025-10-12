package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// Структура графического интерфейса
type Terminal struct {
	windowTitle  string
	theme        *material.Theme
	inputField   widget.Editor
	sendButton   widget.Clickable
	outputField  widget.Editor
	ops          op.Ops
	history      []string
	historyIndex int
	buffer       chan string
}

// Создать новый терминал
func NewTerminal(title string) *Terminal {
	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	t := &Terminal{
		windowTitle: title,
		theme:       theme,
	}

	t.history = make([]string, 0)
	t.buffer = make(chan string)

	t.inputField.SingleLine = true
	t.inputField.Submit = true
	t.outputField.ReadOnly = true

	return t
}

// Вывести текст в поле вывода
func (t *Terminal) Print(a any) {
	t.outputField.SetText(t.outputField.Text() + fmt.Sprint(a))
	t.outputField.MoveCaret(len(t.outputField.Text()), len(t.outputField.Text()))
}

// Вывести текст в поле вывода
func (t *Terminal) Println(a any) {
	t.Print(fmt.Sprint(a) + "\n")
}

// Основной цикл программы
func (t *Terminal) Main() {
	w := new(app.Window)
	w.Option(app.Title(t.windowTitle),
		app.Size(unit.Dp(800), unit.Dp(600)))

	for {
		e := w.Event()
		switch e := e.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&t.ops, e)

			// Обрабатываем нажатие кнопки
			if t.sendButton.Clicked(gtx) {
				t.writeToBuffer()
				continue
			}

			// Обрабатываем SubmitEvent
			for {
				inputEvent, ok := t.inputField.Update(gtx)
				if !ok {
					break
				}
				if _, ok := inputEvent.(widget.SubmitEvent); ok {
					t.writeToBuffer()
					break
				}
			}

			// Обрабатываем клавиатуру
			for {
				event, ok := gtx.Event(key.Filter{})
				if !ok {
					break
				}

				if keyEvent, ok := event.(key.Event); ok && keyEvent.State == key.Press {
					switch keyEvent.Name {
					case key.NameUpArrow:
						if t.historyIndex < len(t.history) {
							t.inputField.SetText(t.history[t.historyIndex])
							t.historyIndex++
						}
					}
					break
				}
			}

			// Рисуем интерфейс
			t.draw(gtx, e)
		case app.DestroyEvent:
			exit()
		}
	}
}

// Получить ввод пользователя
func (t *Terminal) Read() (string, error) {
	return <-t.buffer, nil
}

// Записывает данные из поля ввода в буфер
func (t *Terminal) writeToBuffer() {
	text := t.inputField.Text()
	if text == "" {
		return
	}

	t.inputField.SetText("")

	// Добавляем в историю
	history := make([]string, 1)
	history[0] = text
	history = append(history, t.history...)
	t.history = history
	t.historyIndex = 0

	t.buffer <- text
}

// Отрисовать элементы интерфейса
func (t *Terminal) draw(gtx layout.Context, e app.FrameEvent) {
	// Поле ввода
	inputField := func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			editor := material.Editor(t.theme, &t.inputField, "Type here...")
			return editor.Layout(gtx)
		})
	}

	// Кнопка ввода
	sendButton := func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(t.theme, &t.sendButton, "Execute")
		return btn.Layout(gtx)
	}

	// Поле вывода
	outputField := func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			editor := material.Editor(t.theme, &t.outputField, "")
			return editor.Layout(gtx)
		})
	}

	layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		// Верхняя часть - текстовое поле вывода
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				border := widget.Border{
					Color:        t.theme.Palette.Fg,
					CornerRadius: unit.Dp(4),
					Width:        unit.Dp(1),
				}
				return border.Layout(gtx, outputField)
			})
		}),
		// Нижняя часть - поле ввода и кнопка
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Spacing:   layout.SpaceBetween,
					Alignment: layout.Middle,
				}.Layout(gtx, layout.Flexed(1, inputField), layout.Rigid(sendButton))
			})
		}),
	)

	e.Frame(gtx.Ops)
}
