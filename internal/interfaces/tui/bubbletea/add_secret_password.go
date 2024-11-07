package bubbletea

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type AddSecretPasswordModel struct {
	parent  tea.Model
	service interfaces.ClientService
	inputs  []textinput.Model
	focused int
	err     error
}

const (
	secretPasswordName  = iota
	secretPasswordLogin = iota
	secretPasswordPassword
)

func NewAddSecretPasswordModel(parent tea.Model, service interfaces.ClientService) AddSecretPasswordModel {
	inputs := make([]textinput.Model, 3)

	inputs[secretPasswordName] = NewInput(inputText, true)
	inputs[secretPasswordLogin] = NewInput(inputText, false)
	inputs[secretPasswordPassword] = NewInput(inputPassword, false)

	return AddSecretPasswordModel{
		parent:  parent,
		service: service,
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m AddSecretPasswordModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddSecretPasswordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				name := m.inputs[secretPasswordName].Value()
				login := m.inputs[secretPasswordLogin].Value()
				password := m.inputs[secretPasswordPassword].Value()

				err := m.service.CreateUserSecret(context.Background(), &command.ClientCreateUserSecretCommand{
					SecretType: entities.UserSecretPasswordType,
					SecretName: name,
					SecretData: entities.NewUserSecretPassword(login, password),
				})

				if err != nil {
					m.err = err
					return m, nil
				}

				return NewSuccess(m.service), nil
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

func (m AddSecretPasswordModel) View() string {
	return docStyle.Render(fmt.Sprintf(
		`%s

%s
%s
%s
%s
%s
%s
%s
%s
`,
		titleStyle.Render("Добавить логин/пароль"),
		inputLabelStyle.Render("Название"),
		m.inputs[secretPasswordName].View(),
		inputLabelStyle.Render("Логин"),
		m.inputs[secretPasswordLogin].View(),
		inputLabelStyle.Render("Пароль"),
		m.inputs[secretPasswordPassword].View(),
		errToString(m.err),
		continueStyle.Render("Enter сохранить, Esc вернуться"),
	) + "\n")
}

// nextInput focuses the next input field
func (m *AddSecretPasswordModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *AddSecretPasswordModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
