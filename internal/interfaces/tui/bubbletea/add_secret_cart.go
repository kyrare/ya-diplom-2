package bubbletea

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/command"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
	"github.com/kyrare/ya-diplom-2/internal/interfaces/tui/bubbletea/validators"
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
	parent  tea.Model
	service interfaces.ClientService
	inputs  []textinput.Model
	focused int
	err     error
}

func NewAddSecretCartModel(parent tea.Model, service interfaces.ClientService) AddSecretCartModel {
	inputs := make([]textinput.Model, 3)

	inputs[ccn] = NewInput(inputText, true)
	inputs[exp] = NewInput(inputText, false)
	inputs[cvv] = NewInput(inputText, false)

	inputs[ccn].Placeholder = "4505 **** **** 1234"
	inputs[ccn].CharLimit = 20
	inputs[ccn].Width = 30
	//inputs[ccn].SetValue("1234567890123456")

	inputs[exp].Placeholder = "MM/YY"
	inputs[exp].CharLimit = 5
	inputs[exp].Width = 5

	inputs[cvv].Placeholder = "XXX"
	inputs[cvv].CharLimit = 3
	inputs[cvv].Width = 5

	return AddSecretCartModel{
		parent:  parent,
		service: service,
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
				ccnV := m.inputs[ccn].Value()
				expV := m.inputs[exp].Value()
				cvvV := m.inputs[cvv].Value()

				if err := validators.CcnValidator(ccnV); err != nil {
					m.err = err
					return m, nil
				}

				if err := validators.ExpValidator(expV); err != nil {
					m.err = err
					return m, nil
				}

				if err := validators.CvvValidator(cvvV); err != nil {
					m.err = err
					return m, nil
				}

				parts := strings.Split(expV, "/")
				month, err := strconv.ParseInt(parts[0], 10, 64)
				if err != nil {
					m.err = fmt.Errorf("EXP is invalid")
					return m, nil
				}
				year, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					m.err = fmt.Errorf("EXP is invalid")
					return m, nil
				}

				year += 2000

				cvvI, err := strconv.ParseInt(cvvV, 10, 64)
				if err != nil {
					m.err = fmt.Errorf("CVV is invalid")
					return m, nil
				}

				err = m.service.CreateUserSecret(context.Background(), &command.ClientCreateUserSecretCommand{
					SecretType: entities.UserSecretBankCardType,
					SecretName: ccnV,
					SecretData: entities.NewUserSecretBankCard(ccnV, month, year, cvvI),
				})

				if err != nil {
					m.err = err
					return m, nil
				}

				return NewSuccess(m.service), nil
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
 %s
`,
		inputLabelStyle.Width(30).Render("Card Number"),
		m.inputs[ccn].View(),
		inputLabelStyle.Width(6).Render("EXP"),
		inputLabelStyle.Width(6).Render("CVV"),
		m.inputs[exp].View(),
		m.inputs[cvv].View(),
		errToString(m.err),
		continueStyle.Render("Enter сохранить, Esc вернуться"),
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
