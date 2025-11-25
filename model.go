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

func (m model) Init() tea.Cmd {
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
				// new note
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.currNote = Note{}
				m.state = titleView
			case "up", "j": // vim... ¬‿¬
			// TODO: decr model listindex...
				if m.listIndex > 0 {
					m.listIndex--
				}
			case "down", "k":
				if m.listIndex < len(m.notes) - 1 {
					m.listIndex++
				}
			case "enter":
				// open note
				m.currNote = m.notes[m.listIndex]
				m.state = bodyView
				m.textarea.SetValue(m.currNote.Body)
				m.textarea.Focus()
				m.textarea.CursorEnd()
			}

		case titleView:
			switch key {
			case "enter":
				title := m.textinput.Value()
				if title != "" {
					m.currNote.Title = title 
					m.state = bodyView
					m.textarea.SetValue("")
					m.textarea.Focus()
					m.textarea.CursorEnd()
				}
			case "esc":
				m.state = listView
			}

		case bodyView:
			switch key {
			case "ctrl+s":
				m.currNote.Body = m.textarea.Value()

				m.store.SaveNote(m.currNote)
				
				tea.Quit()
			case "esc":
				m.state = listView
			}
		}
	}

	return m, nil
}
