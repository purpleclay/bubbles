/*
Copyright (c) 2023 Purple Clay

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/purpleclay/bubbles/footer"
)

var (
	red      = lipgloss.NewStyle().Background(lipgloss.Color("#cc2936"))
	redMsg   = "background :red:"
	green    = lipgloss.NewStyle().Background(lipgloss.Color("#678d58"))
	greenMsg = "background :green:"
	blue     = lipgloss.NewStyle().Background(lipgloss.Color("#08415c"))
	blueMsg  = "background :blue:"
	clear    = lipgloss.NewStyle().Background(lipgloss.NoColor{})
	clearMsg = "background :clear:"
)

type keys struct {
	bindings []key.Binding
}

func (k keys) ShortHelp() []key.Binding {
	return k.bindings
}

func (k keys) FullHelp() [][]key.Binding {
	return nil
}

type model struct {
	fill       lipgloss.Style
	fillHeight int
	height     int
	footer     *footer.Model
	width      int
}

func new() model {
	km := []key.Binding{
		key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "red")),
		key.NewBinding(key.WithKeys("g"), key.WithHelp("g", "green")),
		key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "blue")),
		key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "clear")),
		key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "quit")),
	}

	ftr := footer.New(keys{bindings: km})
	ftr.Message(clearMsg)

	return model{
		footer: ftr,
	}
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case footer.FooterResizedMsg:
		m.fillHeight = m.height - m.footer.Height()
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.footer = m.footer.Resize(m.width, m.height).(*footer.Model)
		m.fillHeight = m.height - m.footer.Height()
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			m.fill = red.Copy()
			m.footer.Message(redMsg)
		case "g":
			m.fill = green.Copy()
			m.footer.Message(greenMsg)
		case "b":
			m.fill = blue.Copy()
			m.footer.Message(blueMsg)
		case "c":
			m.fill = clear.Copy()
			m.footer.Message(clearMsg)
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	f, cmd := m.footer.Update(msg)
	m.footer = f.(*footer.Model)

	return m, cmd
}

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.fill.Height(m.fillHeight).Width(m.width).Render(),
		m.footer.View(),
	)
}

func main() {
	if _, err := tea.NewProgram(new(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("Could not start program: %v\n", err)
		os.Exit(1)
	}
}
