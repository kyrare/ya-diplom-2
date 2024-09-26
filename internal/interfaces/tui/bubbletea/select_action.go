package bubbletea

import (
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
	list := list.New([]list.Item{
		ModelItem{title: "Добавить новый секрет", desc: "Сохраните новый секрет!"},
		ModelItem{title: "Показать все секреты", desc: "Показать все секреты, которые уже были сохранили"},
	}, list.NewDefaultDelegate(), 60, 15)

	list.SetFilteringEnabled(false)
	list.SetShowStatusBar(false)
	list.Title = "Выберете доступное действие"

	return SelectActionModel{
		list:    list,
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
			//switch m.list.Cursor() {
			//case 0:
			//	return NewLoginModel(m), nil
			//case 1:
			//	return NewRegisterModel(m), nil
			//}
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
