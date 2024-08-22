package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	originalString string
	modifiedString string
	idx int
	lastKey rune 
}

func initialModel() model {
	challenge := "Hello, World!"
	initialState := make([]bool, len(challenge))
	for i := 0; i < len(challenge); i++ {
	  initialState = append(initialState, true)
	}
return model{originalString: challenge, modifiedString: challenge, idx: 0}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) GetOriginalStringByIdx() rune {
		return []rune(m.originalString)[m.idx]
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle key presses
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Check if the key pressed is "esc"
		if msg.String() == "esc" {
			// If the pressed key is "esc"
			return m, tea.Quit
		} else if msg.String() == "backspace" {
			// If the pressed key is "backspace"
			if m.idx > 0 {
				m.idx--
			originalStringAsRunes := m.GetOriginalStringByIdx() 
			m.modifiedString = m.ReplaceAtIndex(originalStringAsRunes)
			}
		}	else {
			keyAsString := msg.String()
			m.lastKey = []rune(keyAsString)[0] 
			m.modifiedString = m.ReplaceAtIndex(m.lastKey)
			if m.idx < len(m.originalString) {
				m.idx++
			}
		}
	}
	return m, nil
}

func (m model) ReplaceAtIndex(lastKey rune) string {
	runes := []rune(m.modifiedString)
	if m.idx >= len(runes) {
		m.idx = len(runes)
  		return m.modifiedString
	}

	runes[m.idx] = lastKey 
	return string(runes)
}

func (m model) View() string {
	orgStringAsRunes := []rune(m.originalString)
	modStringAsRunes := []rune(m.modifiedString)
	var output string

	for i, modRune := range modStringAsRunes {
		if i < len(orgStringAsRunes) {
			if modRune == orgStringAsRunes[i] && i < m.idx {
				// Correct, color green
				output += "\033[32m" + string(modRune) + "\033[0m"
			} else if i < m.idx {
				// Incorrect, color red
				output += "\033[31m" + string(modRune) + "\033[0m"
			} else {
			output += string(modRune)
			}
		}
	}
	return output
}


func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "could not start program: %v\n", err)
		os.Exit(1)
	}
}
