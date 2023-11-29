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

package header

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
	"github.com/purpleclay/bubbles"
	theme "github.com/purpleclay/lipgloss-theme"
)

// Option defines options when creating a new model
type Option func(*Model)

// WithVersion provides an option for setting and displaying a version
// alongside the header title. An empty string is ignored
func WithVersion(ver string) Option {
	return func(m *Model) {
		if strings.TrimSpace(ver) != "" {
			m.version = ver
			m.showVersion = true
		}
	}
}

// WithDesc provides an option for setting and displaying a description
// underneath the header title. An empty string is ignored
func WithDesc(desc string) Option {
	return func(m *Model) {
		if strings.TrimSpace(desc) != "" {
			m.desc = desc
			m.showDesc = true
		}
	}
}

// WithBorder provides an option for setting and displaying a border that
// underlines the header. This border will always be the width of the terminal
func WithBorder() Option {
	return func(m *Model) {
		m.showBorder = true
	}
}

var (
	borderBottom = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(theme.S600).
			MarginBottom(1)

	faint = lipgloss.NewStyle().Faint(true)
)

// Model defines a header TUI component
type Model struct {
	desc        string
	height      int
	title       string
	version     string
	width       int
	showBorder  bool
	showDesc    bool
	showVersion bool
}

// New creates a new model with a given title. Further customization can
// be achieved by providing additional options:
//
//	hdr := New("bubbles", WithDesc("a collection of TUI components"))
func New(title string, opts ...Option) *Model {
	m := &Model{title: title}

	for _, o := range opts {
		o(m)
	}
	// The height of the header will be static, once initialised
	m.height = lipgloss.Height(m.View())
	return m
}

// Init initialises the header
func (*Model) Init() tea.Cmd {
	return nil
}

// Update the current header
func (m *Model) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View displays the header
func (m *Model) View() string {
	var b strings.Builder
	b.WriteString(theme.H2.Render(m.title))

	if m.showVersion {
		b.WriteString(theme.H4.Render(m.version))
	}

	if m.showDesc {
		b.WriteString("\n\n")
		b.WriteString(faint.Render(truncate.StringWithTail(m.desc, uint(m.width), "...")))
	}

	if m.showBorder {
		return borderBottom.Width(m.width).Render(b.String())
	}

	return b.String()
}

// Resize the dimensions of the header
func (m *Model) Resize(width, _ int) bubbles.Model {
	m.width = width
	return m
}

// Width returns the current width of the header
func (m *Model) Width() int {
	return m.width
}

// Height returns the current height of the header
func (m *Model) Height() int {
	return m.height
}
