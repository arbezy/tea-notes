package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	store	  *Store
	state	  uint
	textarea  textarea.Model
	textinput textinput.Model
	currNote  Note
	notes	  []Note
	listIndex int
}

func NewModel(store *Store) model {
	notes, err := store.GetNotes()
	if err != nil {
		log.Fatal("unable to get notes: %v", err)
	}

	return model{
		store: store,
		state: listView,
		textarea: textarea.New(),
		textinput: textinput.New(),
		notes: notes,
	}
}

func (m model) Init() tea.cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.textarea, _ = m.textarea.Update(msg)
	m.textinput, _ = m.textinput.Update(msg)

	switch msg := msg.(type) {
	// handle key strokes
	case tea.KeyMsg:
		key := msg.String()
		
		switch m.state {
		case listView:
			switch key {
			case "q":
				return m, tea.Quit
			case "n":
				// ...
			case "up":
				// ...
			case "down":
				// ...
			case "enter":
				// ...
			}

		case titleView:
			switch key {
			case "enter":
				//...
			case "esc":
				// ...
			}

		case bodyView:
			switch key {
			case "ctrl+s":
				// ...
			case "esc":
				// ...
			}
		}
	}

	return m, nil
}
