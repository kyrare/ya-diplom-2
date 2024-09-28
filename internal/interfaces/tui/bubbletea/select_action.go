package bubbletea

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
)

type SelectActionModel struct {
	service interfaces.ClientService
	list    list.Model
}

func NewSelectAction(service interfaces.ClientService) SelectActionModel {
	w, h := getListSizes()

	l := list.New([]list.Item{
		ModelItem{title: "Добавить новый секрет", desc: "Сохраните новый секрет!"},
		ModelItem{title: "Показать все секреты", desc: "Показать все секреты, которые уже были сохранили"},
		ModelItem{title: "Удалить секрет", desc: "Удалить один из секретов"},
	}, list.NewDefaultDelegate(), w, h)

	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.Title = "Выберете доступное действие"

	return SelectActionModel{
		list:    l,
		service: service,
	}
}

func (m SelectActionModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m SelectActionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.list.Cursor() {
			case 0:
				return NewSecretType(m, m.service), nil
			case 1, 2:
				secrets, _ := m.service.GetUserSecrets(context.Background())
				// todo tui error
				return NewSecretsList(m, secrets, m.service, m.list.Cursor() == 2), nil
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

func (m SelectActionModel) View() string {
	return docStyle.Render(m.list.View())
}
