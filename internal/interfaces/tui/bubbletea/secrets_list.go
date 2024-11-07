package bubbletea

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type SecretsListModel struct {
	parent         tea.Model
	service        interfaces.ClientService
	list           list.Model
	deleteOnSelect bool
}

func NewSecretsList(
	parent tea.Model,
	secrets []*entities.UserSecret,
	service interfaces.ClientService,
	deleteOnSelect bool,
) SecretsListModel {
	delegate := list.NewDefaultDelegate()

	items := make([]list.Item, 0)
	for _, secret := range secrets {
		items = append(items, SecretListItem{*secret})
	}

	w, h := getListSizes()

	l := list.New(items, delegate, w, h)
	l.Title = "Все ваши секреты"

	return SecretsListModel{
		parent:         parent,
		service:        service,
		list:           l,
		deleteOnSelect: deleteOnSelect,
	}
}

func (m SecretsListModel) Init() tea.Cmd {
	return nil
}

func (m SecretsListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			return m.parent, nil
		case tea.KeyEnter:
			if m.deleteOnSelect {
				item := m.list.SelectedItem().(SecretListItem)

				return NewConfirmDelete(m.parent, m.service, item), nil
			}

			return NewSelectAction(m.service), nil
		}

	case tea.WindowSizeMsg:
		h, v := docForListStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m SecretsListModel) View() string {
	return docForListStyle.Render(m.list.View())
}
