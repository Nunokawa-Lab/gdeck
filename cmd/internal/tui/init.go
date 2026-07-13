package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

func (m Model) Init() tea.Cmd {
	// 次の回転イベントを返す（予約）
	return tea.Batch(
		m.spinner.Tick,
		textinput.Blink,
	)
}

func InitialModel() (Model, error) {

	requests, err := store.List()
	if err != nil {
		return Model{}, err
	}

	// spinner new create
	s := spinner.New()
	s.Spinner = spinner.Dot

	// viewport new create
	vp := viewport.New(80, 20)

	// textinput new create
	ti := textinput.New()
	ti.CharLimit = 100
	ti.Width = 30
	ti.Placeholder = "get-comment"

	// textarea new create
	// 	ta := textarea.New()
	// 	ta.CharLimit = 0 // 制限なし
	// 	ta.SetWidth(60)  // TODO 要検討
	// 	ta.SetHeight(15)  // TODO 要検討
	// 	// ta.Prompt = ""
	// 	ta.ShowLineNumbers = false
	// 	ta.Placeholder = `GetSampleUsers GET https://api.example.com/v1/users
	// -H "Authorization: Bearer {{AUTH_TOKEN}}"
	// -d '{"role": "admin"}'`

	// name-textinput new create
	saveFormName := textinput.New()
	saveFormName.CharLimit = 100
	saveFormName.Width = 80
	saveFormName.Placeholder = "GetSampleUsers"

	// method-textinput new create
	saveFormMethod := textinput.New()
	saveFormMethod.CharLimit = 10
	saveFormMethod.Width = 20
	saveFormMethod.Placeholder = "GET"

	// url-textinput new create
	saveFormUrl := textinput.New()
	saveFormUrl.CharLimit = 0
	saveFormUrl.Width = 80
	saveFormUrl.Placeholder = "https://api.example.com/v1/users"

	// saveform init
	sf := saveForm{
		name:   saveFormName,
		method: saveFormMethod,
		url:    saveFormUrl,
		focus:  focusSaveFieldName,
	}

	// bubbleteaに渡すinterfaceは Init() Update() View() をレシーバーに持っている必要あり
	m := Model{
		requests:           requests,
		cursor:             0,
		spinner:            s,
		leftViewport:       vp,
		rightViewport:      vp,
		searchInput:        ti,
		rightPaneView:      RightPanePreview,
		saveForm:           sf,
		saveFormFieldCount: 3,
	}

	m.loadCurrentRequest()

	m.leftViewport.SetContent(m.requestListContent())
	m.rightViewport.SetContent(m.responseContent())

	return m, nil
}
