package bubbletea

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
)

type SelectSecretTypeModel struct {
	parent  tea.Model
	service interfaces.ClientService
	list    list.Model
}

func NewSecretType(parent tea.Model, service interfaces.ClientService) SelectSecretTypeModel {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false

	list := list.New([]list.Item{
		ModelItem{title: "Добавить логин/пароль", desc: ""},
		ModelItem{title: "Добавить банковскую карту", desc: ""},
		ModelItem{title: "Добавить файл", desc: ""},
		ModelItem{title: "Добавить производный текст", desc: ""},
	}, delegate, 60, 15)

	list.SetFilteringEnabled(false)
	list.SetShowStatusBar(false)
	list.Title = "Выберете тип секрета"

	return SelectSecretTypeModel{
		parent:  parent,
		service: service,
		list:    list,
	}
}

func (m SelectSecretTypeModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m SelectSecretTypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return m.parent, nil
		case tea.KeyEnter:
			switch m.list.Cursor() {
			case 0:
				return NewAddSecretPassportModel(m.parent, m.service), nil
			case 1:
				return m.parent, nil
			}
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m SelectSecretTypeModel) View() string {
	return docStyle.Render(m.list.View())
}
