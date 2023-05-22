package title

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)


const icon = "ᕙ(˘ ₃˘)ᕗ"
const title = "manly"

// Static title
// TODO make title look better
func View() string {
	title := fmt.Sprintf(
		"%s%s%s%s%s",
		"\n",
		"\n",
		icon+" "+title+" "+icon,
		"\n",
		"\n",
	)
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#F25D94")).
		PaddingLeft(4).
		Width(32)
	return titleStyle.Render(title)
}
