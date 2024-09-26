package bubbletea

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kyrare/ya-diplom-2/internal/app/interfaces"
)

type SuccessModel struct {
	service interfaces.ClientService
}

func NewSuccess(service interfaces.ClientService) SuccessModel {
	return SuccessModel{service}
}

func (m SuccessModel) Init() tea.Cmd {
	return nil
}

func (m SuccessModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter, tea.KeyEsc:
			return NewSelectAction(m.service), nil
		}
	}

	return m, nil
}

func (m SuccessModel) View() string {

	return docStyle.Render(fmt.Sprintf(
		`%s

%s
`,
		titleStyle.Render("Запись успешно добавлена"),
		continueStyle.Render("Enter продолжить"),
	) + "\n")
}
