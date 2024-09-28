package bubbletea

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
	"github.com/kyrare/ya-diplom-2/internal/domain/entities"
)

type SecretsListModel struct {
	parent  tea.Model
	service interfaces.ClientService
	list    list.Model
}

func NewSecretsList(
	parent tea.Model,
	secrets []*entities.UserSecret,
	service interfaces.ClientService,
) SecretsListModel {
	delegate := list.NewDefaultDelegate()

	items := make([]list.Item, 0)
	for _, secret := range secrets {
		items = append(items, SecretListItem{*secret})
	}

	w, h := getListSizes()

	l := list.New(items, delegate, w, h)

	//l.SetFilteringEnabled(false)
	//l.SetShowStatusBar(false)
	l.Title = "Все ваши секреты"

	return SecretsListModel{
		parent:  parent,
		service: service,
		list:    l,
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
			return NewSelectAction(m.service), nil
			//switch m.list.Cursor() {
			//case 0:
			//	return NewAddSecretPasswordModel(m.parent, m.service), nil
			//case 1:
			//	return NewAddSecretCartModel(m.parent, m.service), nil
			//case 2:
			//	return NewAddSecretTextModel(m.parent, m.service), nil
			//default:
			//	return m.parent, nil
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

func (m SecretsListModel) View() string {
	return docStyle.Render(m.list.View())
}
