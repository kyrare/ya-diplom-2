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

	w, h := getListSizes()

	l := list.New([]list.Item{
		ModelItem{title: "Добавить логин/пароль", desc: ""},
		ModelItem{title: "Добавить банковскую карту", desc: ""},
		ModelItem{title: "Добавить производный текст", desc: ""},
		ModelItem{title: "Добавить файл", desc: ""},
	}, delegate, w, h)

	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.Title = "Выберете тип секрета"

	return SelectSecretTypeModel{
		parent:  parent,
		service: service,
		list:    l,
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
				return NewAddSecretPasswordModel(m.parent, m.service), nil
			case 1:
				return NewAddSecretCartModel(m.parent, m.service), nil
			case 2:
				return NewAddSecretTextModel(m.parent, m.service), nil
			case 3:
				return NewAddSecretFileModel(m.parent, m.service), nil
			default:
				return m.parent, nil
			}
		}

	case tea.WindowSizeMsg:
		h, v := docForListStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m SelectSecretTypeModel) View() string {
	return docForListStyle.Render(m.list.View())
}
