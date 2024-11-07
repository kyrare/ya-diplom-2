package bubbletea

import (
	"context"
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
)

type ConfirmModel struct {
	parent  tea.Model
	service interfaces.ClientService
	item    SecretListItem
	err     error
}

func NewConfirmDelete(
	parent tea.Model,
	service interfaces.ClientService,
	item SecretListItem,
) *ConfirmModel {
	return &ConfirmModel{
		parent:  parent,
		service: service,
		item:    item,
	}
}

func (m ConfirmModel) Init() tea.Cmd {
	return nil
}

func (m ConfirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.err = errors.New("test1")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			m.err = errors.New("test2")
			err := m.service.DeleteUserSecret(context.Background(), m.item.Id)
			if err != nil {
				m.err = err
				return m, nil
			}

			return m.parent, nil
		case tea.KeyEsc:
			return m.parent, nil
		}
	}

	return m, nil
}

func (m ConfirmModel) View() string {
	return docStyle.Render(fmt.Sprintf(
		`%s

%s
%s
`,
		titleStyle.Render("Вы уверены, что хотите удалить этот секрет?"),
		errToString(m.err),
		continueStyle.Render("Enter удалить, Esc вернуться"),
	) + "\n")
}
