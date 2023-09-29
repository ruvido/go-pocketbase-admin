package admin

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Person struct {
    name   string
	email  string
	mobile string
	state  string
    age    string
	status string
}

type model struct {
	choices  []User           // items on the to-do list
	cursor   int                // which to-do list item our cursor is pointing at
	selected map[int]struct{}   // which to-do items are selected
}

func initialModel( people []User ) model {
	return model{
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

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	fmt.Println("Import pocketbase data")
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

		// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

			// The "up" and "k" keys move the cursor up
		case "a":
			m.choices[m.cursor].Status = "accepted"
			m.cursor++

		case "x":
			m.choices[m.cursor].Status = "rejected"
			m.cursor++

		case "w":
			m.choices[m.cursor].Status = "waiting"
			m.cursor++

			// update
		case "u":
			UpdateCollection(m.choices)

			// update & quit
		case "q":
			UpdateCollection(m.choices)
			return m, tea.Quit

			// Abort and exit the program.
		case "ctrl+c":
			return m, tea.Quit

			// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

			// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
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

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		// checked := " " // not selected
		// if _, ok := m.selected[i]; ok {
		// 	checked = "x" // selected!
		// }

		// Is this person accepted?
		accepted:= " " // not accepted
		if choice.Status == "accepted" {
			accepted = "*"
		}
		if choice.Status == "waiting" {
			accepted = "+"
		}

		// Render the row
		s += fmt.Sprintf("%s %2s %-15s %-25s\n", cursor, accepted, choice.Name, choice.Email)
	}

	// The footer
	s += "\n     a to accept."
	s += "\n     x to reject."
	s += "\n     w to waiting list."
	s += "\n     ------------------"
	s += "\n     u to update pocketbase."
	s += "\n     q to update and quit."
	s += "\n     ------------------"
	s += "\n     ctrl+c to abort.\n"

	// Send the UI for rendering
	return s
}

// func main() {
func BubbleList(people []User) {
	p := tea.NewProgram(initialModel(people))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

