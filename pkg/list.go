package admin

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
)

type Person struct {
    name   string
	email  string
	mobile string
	state  string
    age    string
	status string
}

type modelx struct {
	choices  []User           // items on the to-do list
	cursor   int                // which to-do list item our cursor is pointing at
	selected map[int]struct{}   // which to-do items are selected
	paginator paginator.Model
	pagCursor int             // cursor relative to pagination
}

var needUpdate = false

func newModel( people []User ) modelx {
	// var items []string
	// for i := 1; i < 101; i++ {
	// 	text := fmt.Sprintf("Item %d", i)
	// 	items = append(items, text)
	// }

	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 10
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(len(people))

	return modelx{
		paginator: p,
		// items:     items,
		choices: people,
	}
}

func initialModel( people []User ) modelx {
	return modelx{
		// Our to-do list is a grocery list
		// choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		// choices : []Person{
		// 	{name: "Alice Oliss",  email: "ali@ce.co",   mobile: "12343456", state: "Veneto", age: "23", status: ""},
		// 	{name: "Bob Ujik",     email: "bo@b.co",     mobile: "98653006", state: "Lazio",  age: "18", status: ""},
		// 	{name: "Mark Bombero", email: "mk@bombe.ro", mobile: "34273006", state: "Emilia", age: "29", status: ""},
		// },
		choices: people,


		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m modelx) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m modelx) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {

		// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

			// The "up" and "k" keys move the cursor up
		case "a":
			needUpdate = true
			m.choices[m.cursor].Status = "accepted"
			m.cursor++

		case "c":
			needUpdate = true
			m.choices[m.cursor].Status = "cancelled"
			m.cursor++

		case "x":
			needUpdate = true
			m.choices[m.cursor].Status = ""
			m.cursor++

		case "w":
			needUpdate = true
			m.choices[m.cursor].Status = "waiting"
			m.cursor++

			// update
		case "u":
			UpdateCollection(m.choices)
			needUpdate = false

			// update & quit
		case "q":
			if needUpdate {
				UpdateCollection(m.choices)
			}
			return m, tea.Quit

			// Abort and exit the program.
		case "ctrl+c":
			return m, tea.Quit

		case "left":
			// m.paginator.PrevPage()
			if m.paginator.Page > 0 {
				m.cursor = m.cursor - m.paginator.PerPage
			}

		case "right":
			// m.paginator.NextPage()
			m.cursor = m.cursor + m.paginator.PerPage
			if m.cursor >= len(m.choices)-1 {
				m.cursor = len(m.choices)-1
			}

			// The "up" and "k" keys move the cursor up
		case "up", "k":
			// if m.cursor > 0 {
			// 	m.cursor--
			// }
			if m.paginator.Page == 0 {
				if m.cursor > 0 {
					m.cursor--
				} 
				m.pagCursor=m.cursor
			} else {
				m.cursor--
				m.pagCursor = m.cursor - m.paginator.Page*m.paginator.PerPage
				if m.pagCursor < 0 {
					m.paginator.PrevPage()
					m.pagCursor = m.paginator.PerPage-1
				}
			}

			// The "down" and "j" keys move the cursor down
		case "down", "j":
			// if m.cursor < len(m.choices)-1 {
			// if m.cursor + m.paginator.PerPage*m.paginator.Page <  len(m.choices)-1  {
			if m.cursor  <  len(m.choices)-1  {
				m.cursor++
			}
			m.pagCursor = m.cursor - m.paginator.Page*m.paginator.PerPage
			if m.pagCursor  >= m.paginator.PerPage {
				m.paginator.NextPage()
				m.pagCursor = 0
			}

			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
		// case "enter", " ":
		// 	_, ok := m.selected[m.cursor]
		// 	if ok {
		// 		delete(m.selected, m.cursor)
		// 	} else {
		// 		m.selected[m.cursor] = struct{}{}
		// 	}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	// return m, nil

	// with paginator
	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m modelx) View() string {
	var b strings.Builder
	b.WriteString("\n  Pocketbase admin\n\n")
	start, end := m.paginator.GetSliceBounds(len(m.choices))
	for i, item := range m.choices[start:end] {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		m.pagCursor = m.cursor - m.paginator.Page*m.paginator.PerPage
		if m.pagCursor == i {
			cursor = ">" // cursor!
		}
		// Is this person accepted?
		accepted:= " " // not accepted
		if item.Status == "accepted" {
			accepted = "*"
		}
		if item.Status == "waiting" {
			accepted = "+"
		}
		if item.Status == "cancelled" {
			accepted = "C"
		}

		// Render the row
		s := fmt.Sprintf("%s %2s %-25s %-35s %-25s\n", cursor, accepted, item.Name, item.Email, item.Mobile)
		b.WriteString(s)
		// b.WriteString("  • " + item.Name + "\n\n")
	}
	b.WriteString("  " + m.paginator.View())
	b.WriteString("\n\n  a: accept, w: waiting, c: cancelled, x: reset • u: update • q: quit \n")
	c := fmt.Sprintf("%d %d", m.cursor, m.pagCursor )
	b.WriteString(c)
	return b.String()
}

// func main() {
func BubbleList(people []User) {
	// p := tea.NewProgram(initialModel(people))
	p := tea.NewProgram(newModel(people))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// func (m modelx) CAZZView() string {
// 	// The header
// 	s := "What should we buy at the market?\n\n"
//
// 	// Iterate over our choices
// 	for i, choice := range m.choices {
//
// 		// Is the cursor pointing at this choice?
// 		cursor := " " // no cursor
// 		// if m.cursor == i {
// 		if m.cursor == i {
// 			cursor = ">" // cursor!
// 		}
//
// 		// Is this person accepted?
// 		accepted:= " " // not accepted
// 		if choice.Status == "accepted" {
// 			accepted = "*"
// 		}
// 		if choice.Status == "waiting" {
// 			accepted = "+"
// 		}
// 		if choice.Status == "rejected" {
// 			accepted = "E"
// 		}
//
// 		// Render the row
// 		s += fmt.Sprintf("%s %2s %-15s %-25s\n", cursor, accepted, choice.Name, choice.Email)
// 	}
//
// 	// The footer
// 	s += "\n     a to accept."
// 	s += "\n     x to reject."
// 	s += "\n     w to waiting list."
// 	s += "\n     ------------------"
// 	s += "\n     u to update pocketbase."
// 	s += "\n     q to update and quit."
// 	s += "\n     ------------------"
// 	s += "\n     ctrl+c to abort.\n"
//
// 	// Send the UI for rendering
// 	return s
// }
//
