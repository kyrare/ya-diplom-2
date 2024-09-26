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
	loginLogin = iota
	loginPassword
)

type LoginModel struct {
	parent  tea.Model
	service interfaces.ClientService
	inputs  []textinput.Model
	focused int
	err     error
}

func NewLoginModel(parent tea.Model, service interfaces.ClientService) LoginModel {
	inputs := make([]textinput.Model, 2)

	inputs[loginLogin] = textinput.New()
	inputs[loginLogin].Placeholder = ""
	inputs[loginLogin].Focus()
	inputs[loginLogin].CharLimit = 255
	inputs[loginLogin].Width = 50
	inputs[loginLogin].Prompt = ""
	inputs[loginLogin].SetValue("test")

	inputs[loginPassword] = textinput.New()
	inputs[loginPassword].Placeholder = "******"
	inputs[loginPassword].CharLimit = 255
	inputs[loginPassword].Width = 50
	inputs[loginPassword].Prompt = ""
	inputs[loginPassword].EchoMode = textinput.EchoPassword
	inputs[loginPassword].EchoCharacter = '•'
	inputs[loginPassword].SetValue("123123")

	return LoginModel{
		parent:  parent,
		service: service,
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

				err := m.service.Login(context.Background(), &command.LoginCommand{
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

func (m LoginModel) View() string {
	title := "Авторизация"
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
		m.inputs[loginLogin].View(),
		inputLabelStyle.Render("Пароль"),
		m.inputs[loginPassword].View(),
		errMessage,
		continueStyle.Render("Enter авторизоваться, Esc вернуться"),
	) + "\n")
}

// nextInput focuses the next input field
func (m *LoginModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *LoginModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
