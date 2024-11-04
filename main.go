package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Post struct {
	Post struct {
		Author struct {
			Handle      string `json:"handle"`
			DisplayName string `json:"displayName"`
		} `json:"author"`
		Record struct {
			Text string `json:"text"`
		} `json:"record"`
	} `json:"post"`
}

type Response struct {
	Feed []Post `json:"feed"`
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
