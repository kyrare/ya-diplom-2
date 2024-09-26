package bubbletea

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/interfaces/tui/bubbletea/validators"
)

const (
	registerLogin = iota
	registerPassword
)

type RegisterModel struct {
	parent  tea.Model
	service interfaces.ClientService
	inputs  []textinput.Model
	focused int
	err     error
}

func NewRegisterModel(parent tea.Model, service interfaces.ClientService) RegisterModel {
	inputs := make([]textinput.Model, 2)

	inputs[registerLogin] = NewInput(inputText, true)
	inputs[registerPassword] = NewInput(inputPassword, false)

	return RegisterModel{
		parent:  parent,
		service: service,
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m RegisterModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m RegisterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				login := m.inputs[loginLogin].Value()
				password := m.inputs[loginPassword].Value()
				if err := validators.LoginValidator(login); err != nil {
					m.err = err
					return m, nil
				}
				if err := validators.PasswordValidator(password); err != nil {
					m.err = err
					return m, nil
				}

				err := m.service.Register(context.Background(), &command.RegisterCommand{
					Login:    login,
					Password: password,
				})
				if err != nil {
					m.err = err
					return m, nil
				}

				return NewSelectAction(m.service), nil
			}
			m.nextInput()
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return m.parent, nil
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}

		m.err = nil
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
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m RegisterModel) View() string {
	title := "Регистрация"
	errMessage := ""
	if m.err != nil {
		errMessage = "\n" + errorStyle.Render(m.err.Error()) + "\n"
	}

	return docStyle.Render(fmt.Sprintf(
		`%s

%s
%s

%s
%s
%s
%s
`,
		titleStyle.Render(title),
		inputLabelStyle.Render("Логин"),
		m.inputs[registerLogin].View(),
		inputLabelStyle.Render("Пароль"),
		m.inputs[registerPassword].View(),
		errMessage,
		continueStyle.Render("Enter зарегистрироваться, Esc вернуться"),
	) + "\n")
}

// nextInput focuses the next input field
func (m *RegisterModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *RegisterModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
