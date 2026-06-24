package navigation

const (
	MobileSheetTriggerID = "ui8kit-mobile-sheet-trigger"
	MobileSheetPanelID   = "ui8kit-mobile-sheet-panel"
	MobileSheetTitleID   = "ui8kit-mobile-sheet-title"
)

// Item describes a single navigation link.
type Item struct {
	Path  string
	Label string
	Icon  string
}

// NavProps configures horizontal or vertical navigation links.
type NavProps struct {
	Items     []Item
	Active    string
	Vertical  bool
	AriaLabel string
	Class     string
}

// MobileSheetProps configures the mobile navigation sheet panel.
type MobileSheetProps struct {
	ID         string
	TitleID    string
	Title      string
	AriaLabel  string
	CloseLabel string
	Items      []Item
	Active     string
}

// MobileSheetTriggerProps configures the header menu trigger button.
type MobileSheetTriggerProps struct {
	ID        string
	For       string
	AriaLabel string
	Class     string
}
