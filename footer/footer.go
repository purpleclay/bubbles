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

package footer

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/purpleclay/bubbles"
	theme "github.com/purpleclay/lipgloss-theme"
)

// ResizedMsg is a message that is sent when the footer has resized due to
// the help menu expanding and collapsing
type ResizedMsg struct{}

var (
	toggle = theme.H2.
		Copy().
		Width(8).
		AlignHorizontal(lipgloss.Center)

	bar = lipgloss.NewStyle().
		Background(lipgloss.Color("#262525")).
		Padding(0, 2).
		Faint(true)

	helpMargin = lipgloss.NewStyle().MarginTop(1)
)

// Model defines a footer TUI component
type Model struct {
	barWidth int
	barMsg   string
	expanded bool
	help     help.Model
	keymap   help.KeyMap
	width    int
}

// New creates a new model with a given [help.keymap]
func New(keymap help.KeyMap) *Model {
	help := help.New()
	help.Styles.ShortSeparator = lipgloss.NewStyle().Foreground(theme.S200).Faint(true)
	help.Styles.ShortKey = lipgloss.NewStyle().Foreground(theme.S100).Bold(true)

	return &Model{
		keymap:   keymap,
		help:     help,
		expanded: false,
	}
}

// Init initialises the footer
func (*Model) Init() tea.Cmd {
	return nil
}

// Update the current footer
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "?" {
			m.expanded = !m.expanded
			cmd = func() tea.Msg { return ResizedMsg{} }
		}
	}

	return m, cmd
}

// View displays the footer
func (m *Model) View() string {
	var helpMenu string
	toggleMsg := "? help"
	if m.expanded {
		toggleMsg = "? hide"
		helpMenu = helpMargin.Render(m.help.View(m.keymap))
	}

	var b strings.Builder

	b.WriteString(lipgloss.JoinHorizontal(
		lipgloss.Left,
		bar.Copy().Width(m.barWidth).Render(m.barMsg),
		toggle.Render(toggleMsg),
	))

	if helpMenu != "" {
		b.WriteString("\n")
		b.WriteString(helpMenu)
	}

	return b.String()
}

// Resize the dimensions of the footer
func (m *Model) Resize(width, _ int) bubbles.Model {
	m.width = width
	m.barWidth = width - toggle.GetWidth()
	return m
}

// Width returns the current width of the footer
func (m *Model) Width() int {
	return m.width
}

// Height returns the current height of the footer
func (m *Model) Height() int {
	return lipgloss.Height(m.View())
}

// Message sets the message within the status bar
func (m *Model) Message(msg string) {
	m.barMsg = msg
}
