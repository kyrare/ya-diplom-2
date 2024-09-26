package bubbletea

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
)

type StartModel struct {
	list    list.Model
	service interfaces.ClientService
}

func New(service interfaces.ClientService) StartModel {
	list := list.New([]list.Item{
		ModelItem{title: "Авторизация", desc: "Авторизация по логину и паролю"},
		ModelItem{title: "Регистрация", desc: "Регистрация для новых пользователей"},
	}, list.NewDefaultDelegate(), 0, 0)

	list.SetFilteringEnabled(false)
	list.SetShowStatusBar(false)
	list.Title = "Необходимо авторизоваться"

	return StartModel{
		list:    list,
		service: service,
	}
}

func (m StartModel) Init() tea.Cmd {
	return nil
}

func (m StartModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.list.Cursor() {
			case 0:
				return NewLoginModel(m, m.service), nil
			case 1:
				return NewRegisterModel(m, m.service), nil
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

func (m StartModel) View() string {
	return docStyle.Render(m.list.View())
}
