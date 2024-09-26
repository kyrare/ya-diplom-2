package bubbletea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

const (
	ccn = iota
	exp
	cvv
)

type AddSecretCartModel struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func NewAddSecretCartModel() AddSecretCartModel {
	inputs := make([]textinput.Model, 3)

	inputs[ccn] = NewInput(inputText, true)
	inputs[exp] = NewInput(inputText, false)
	inputs[cvv] = NewInput(inputText, false)

	inputs[ccn].Placeholder = "4505 **** **** 1234"
	inputs[ccn].CharLimit = 20
	inputs[ccn].Width = 30

	inputs[exp].Placeholder = "MM/YY"
	inputs[exp].CharLimit = 5
	inputs[exp].Width = 5

	inputs[cvv].Placeholder = "XXX"
	inputs[cvv].CharLimit = 3
	inputs[cvv].Width = 5

	return AddSecretCartModel{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m AddSecretCartModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddSecretCartModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}

		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		if m.inputs[i].Err != nil {
			m.err = m.inputs[i].Err
		}
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m AddSecretCartModel) View() string {
	return fmt.Sprintf(
		` Добавление новой карты:

 %s
 %s

 %s  %s
 %s  %s

 %s
`,
		inputLabelStyle.Width(30).Render("Card Number"),
		m.inputs[ccn].View(),
		inputLabelStyle.Width(6).Render("EXP"),
		inputLabelStyle.Width(6).Render("CVV"),
		m.inputs[exp].View(),
		m.inputs[cvv].View(),
		continueStyle.Render("Сохранить ->"),
	) + "\n"
}

// nextInput focuses the next input field
func (m *AddSecretCartModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *AddSecretCartModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
