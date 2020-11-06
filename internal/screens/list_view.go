package screens

import (
	"github.com/4thel00z/down/pkg/libdown"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"strings"
)

type ListItem struct {
	Text    string
	Checked bool
}

var (
	p = termenv.ColorProfile()
	DefaultStandard  = termenv.Style{}.Foreground(p.Color("15"))
	DefaultHighlight = termenv.Style{}.Foreground(p.Color("87"))
)

func (l ListItem) String() string {
	if l.Checked {
		return libdown.S("[•] %s", l.Text)
	}
	return libdown.S("[ ] %s", l.Text)
}

type ListView struct {
	Items     []ListItem
	Current   int
	Standard  termenv.Style
	Highlight termenv.Style
}

func toListItems(items ...string) []ListItem {
	n := len(items)
	listItems := make([]ListItem, n)
	for i := 0; i < n; i++ {
		listItems[i] = ListItem{
			Text: items[i],
		}
	}
	return listItems
}

func (lv *ListView) Format() []string {
	n := len(lv.Items)
	items := make([]string, n)
	for i := 0; i < n; i++ {
		listItem := lv.Items[i]
		if lv.Current == i {
			items[i] = lv.Highlight.Styled(listItem.String())
		} else {
			items[i] = lv.Standard.Styled(listItem.String())
		}
	}
	return items
}

func NewListView(standard, highlight termenv.Style, items ...string) *ListView {
	return &ListView{
		Items:     toListItems(items...),
		Current:   0,
		Standard:  standard,
		Highlight: highlight,
	}
}

func (lv *ListView) Next() {
	lv.Current = (lv.Current + 1) % len(lv.Items)
}

func (lv *ListView) Previous() {
	if lv.Current-1 < 0 {
		lv.Current = len(lv.Items) - 1
		return
	}
	lv.Current = lv.Current - 1
}

func (lv *ListView) Init() tea.Cmd {
	return nil
}

func (lv *ListView) Update(m tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := m.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "space":
		case "tab":
			lv.Items[lv.Current].Checked = !lv.Items[lv.Current].Checked
			return lv, nil

		case "down":
			lv.Next()
			return lv, nil

		case "up":
			lv.Previous()
			return lv, nil
		}

	}
	return lv, nil
}

func (lv *ListView) View() string {
	return strings.Join(lv.Format(), "\n")
}
