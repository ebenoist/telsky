package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	resp     Response
	cursor   int
	selected map[int]struct{} // which to-do items are selected
}

func (m model) View() string {
	post := template.New("post")
	post = template.Must(post.Parse(`
	{{ .Current.Post.Author.DisplayName }} (@{{.Current.Post.Author.Handle }})

	{{ .Current.Post.Record.Text }}

	{{.Cursor}}{{.Checked}}Like * Repost
	------------`))

	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.resp.Feed {

		// Is the cursor pointing at this choice?
		cursor := "" // no cursor
		if m.cursor == i {
			cursor = "*" // cursor!
		}

		// Is this choice selected?
		checked := "" // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		var row bytes.Buffer
		err := post.Execute(&row, struct {
			Cursor  string
			Checked string
			Current Post
		}{
			Cursor:  cursor,
			Checked: checked,
			Current: choice,
		})

		if err != nil {
			panic(err)
		}

		s += fmt.Sprintf("%s\n", row.String())
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func initialModel() model {
	feed, err := getAuthorFeed("@benoist.dev")
	if err != nil {
		fmt.Printf("got an error %s", err)
		panic(err)
	}

	var result Response
	err = json.Unmarshal(feed, &result)
	if err != nil {
		fmt.Printf("got an error %s", err)
		panic(err)
	}

	return model{
		resp:     result,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.resp.Feed)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
