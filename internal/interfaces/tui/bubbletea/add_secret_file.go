package bubbletea

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type AddSecretFileModel struct {
	parent    tea.Model
	service   interfaces.ClientService
	inputName textinput.Model
	inputFile textinput.Model
	focused   int
	err       error
}

func NewAddSecretFileModel(parent tea.Model, service interfaces.ClientService) AddSecretFileModel {
	return AddSecretFileModel{
		parent:    parent,
		service:   service,
		inputName: NewInput(inputText, true),
		inputFile: NewInput(inputText, false),
		focused:   0,
		err:       nil,
	}
}

func (m AddSecretFileModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddSecretFileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 2)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == 1 {
				name := m.inputName.Value()
				file := m.inputFile.Value()

				if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
					m.err = errors.New("file does not exist")
					return m, nil
				}

				err := m.service.CreateUserSecret(context.Background(), &command.ClientCreateUserSecretCommand{
					SecretType: entities.UserSecretFileType,
					SecretName: name,
					SecretData: entities.NewUserSecretFile(file),
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

		m.inputName.Blur()
		m.inputFile.Blur()

		if m.focused == 0 {
			m.inputName.Focus()
		} else {
			m.inputFile.Focus()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	if m.inputName.Err != nil {
		m.err = m.inputName.Err
	} else if m.inputFile.Err != nil {
		m.err = m.inputFile.Err
	}

	m.inputName, cmds[0] = m.inputName.Update(msg)
	m.inputFile, cmds[1] = m.inputFile.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m AddSecretFileModel) View() string {
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
		m.inputName.View(),
		inputLabelStyle.Render("Путь к файлу"),
		m.inputFile.View(),
		errToString(m.err),
		continueStyle.Render("Enter сохранить, Esc вернуться"),
	) + "\n")
}

// nextInput focuses the next input field
func (m *AddSecretFileModel) nextInput() {
	m.focused = (m.focused + 1) % 2
}

// prevInput focuses the previous input field
func (m *AddSecretFileModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = 1
	}
}
