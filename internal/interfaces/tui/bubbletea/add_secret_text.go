package bubbletea

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type AddSecretTextModel struct {
	parent   tea.Model
	service  interfaces.ClientService
	input    textinput.Model
	textarea textarea.Model
	focused  int
	err      error
}

func NewAddSecretTextModel(parent tea.Model, service interfaces.ClientService) AddSecretTextModel {
	return AddSecretTextModel{
		parent:   parent,
		service:  service,
		input:    NewInput(inputText, true),
		textarea: NewTextarea(false),
		focused:  0,
		err:      nil,
	}
}

func (m AddSecretTextModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddSecretTextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 2)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == 1 {
				name := m.input.Value()
				text := m.textarea.Value()

				err := m.service.CreateUserSecret(context.Background(), &command.ClientCreateUserSecretCommand{
					SecretType: entities.UserSecretTextType,
					SecretName: name,
					SecretData: entities.NewUserSecretText(text),
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

		m.input.Blur()
		m.textarea.Blur()

		if m.focused == 0 {
			m.input.Focus()
		} else {
			m.textarea.Focus()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	if m.input.Err != nil {
		m.err = m.input.Err
	} else if m.textarea.Err != nil {
		m.err = m.textarea.Err
	}

	m.input, cmds[0] = m.input.Update(msg)
	m.textarea, cmds[1] = m.textarea.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m AddSecretTextModel) View() string {
	return docStyle.Render(fmt.Sprintf(
		`%s

%s
%s
%s
%s
%s
%s
`,
		titleStyle.Render("Добавить текст"),
		inputLabelStyle.Render("Название"),
		m.input.View(),
		inputLabelStyle.Render("Текст"),
		m.textarea.View(),
		errToString(m.err),
		continueStyle.Render("Enter сохранить, Esc вернуться"),
	) + "\n")
}

// nextInput focuses the next input field
func (m *AddSecretTextModel) nextInput() {
	m.focused = (m.focused + 1) % 2
}

// prevInput focuses the previous input field
func (m *AddSecretTextModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = 1
	}
}
