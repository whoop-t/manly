package main

import "fmt"

func render(m model) string {
	return fmt.Sprintf(
		"Search Man Pages\n\n%s\n\n%s\n\n%s",
		m.searchInput.View(),
		m.resultsList.View(),
		"(esc to quit)",
	) + "\n"
}
