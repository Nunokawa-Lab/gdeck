package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
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
		textarea.Blink,
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

	// header-textarea new create
	saveFormHeader := textarea.New()
	saveFormHeader.CharLimit = 0 // 制限なし
	saveFormHeader.SetWidth(80)
	saveFormHeader.SetHeight(5)
	saveFormHeader.ShowLineNumbers = false
	saveFormHeader.Placeholder = `One header per line
Use {{TOKEN}} for env substitution

Content-Type: application/json
Authorization: Bearer {{TOKEN}}`

	// body-textarea new create
	saveFormBody := textarea.New()
	saveFormBody.CharLimit = 0 // 制限なし
	saveFormBody.SetWidth(80)
	saveFormBody.SetHeight(15)
	saveFormBody.ShowLineNumbers = false
	saveFormBody.Placeholder = `Use {{WEBHOOK_URL}} for env substitution

{
  "role": "admin",
  "webhook-url": "{{WEBHOOK_URL}}"
}`

	// saveform init
	sf := saveForm{
		name:   saveFormName,
		method: saveFormMethod,
		url:    saveFormUrl,
		header: saveFormHeader,
		body:   saveFormBody,
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
		saveFormFieldCount: 5,
	}

	m.loadCurrentRequest()

	m.leftViewport.SetContent(m.requestListContent())
	m.rightViewport.SetContent(m.responseContent())

	return m, nil
}
