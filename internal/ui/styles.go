package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors - Blue/Cyan theme
	primaryColor   = lipgloss.Color("#00CED1") // Dark Cyan
	secondaryColor = lipgloss.Color("#20B2AA") // Light Sea Green
	accentColor    = lipgloss.Color("#5F9EA0") // Cadet Blue
	highlightColor = lipgloss.Color("#00FFFF") // Cyan
	mutedColor     = lipgloss.Color("#708090") // Slate Gray
	errorColor     = lipgloss.Color("#FF6B6B") // Red
	successColor   = lipgloss.Color("#98FB98") // Pale Green
	warningColor   = lipgloss.Color("#FFD700") // Gold

	// Base styles
	BaseStyle = lipgloss.NewStyle()

	// Title
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			Padding(0, 1)

	// Tabs
	ActiveTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(primaryColor).
			Padding(0, 2)

	InactiveTabStyle = lipgloss.NewStyle().
				Foreground(mutedColor).
				Padding(0, 2)

	TabGapStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	// List items
	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(highlightColor).
				Bold(true)

	NormalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	CompletedItemStyle = lipgloss.NewStyle().
				Foreground(mutedColor).
				Strikethrough(true)

	// Checkbox
	CheckboxEmpty    = "[ ] "
	CheckboxChecked  = "[x] "
	CheckboxSelected = " > "
	CheckboxNormal   = "   "

	// Panels
	ListPanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(accentColor).
			Padding(1)

	DetailPanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(accentColor).
				Padding(1)

	ActivePanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(primaryColor).
				Padding(1)

	// Detail view
	DetailTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(primaryColor).
				MarginBottom(1)

	DetailLabelStyle = lipgloss.NewStyle().
				Foreground(secondaryColor).
				Bold(true)

	DetailValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF"))

	// Status bar
	StatusBarStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Padding(0, 1)

	StatusMessageStyle = lipgloss.NewStyle().
				Foreground(successColor)

	StatusErrorStyle = lipgloss.NewStyle().
				Foreground(errorColor)

	// Help
	HelpKeyStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	HelpDescStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	HelpSeparatorStyle = lipgloss.NewStyle().
				Foreground(mutedColor)

	// Modal
	ModalStyle = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2).
			Background(lipgloss.Color("#1a1a2e"))

	ModalTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlightColor).
			MarginBottom(1)

	// Input
	InputLabelStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Bold(true)

	InputStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(accentColor).
			Padding(0, 1)

	InputFocusedStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(primaryColor).
				Padding(0, 1)

	// Projects badge
	ProjectBadgeStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#1a1a2e")).
				Background(secondaryColor).
				Padding(0, 1)

	// Project names in task list
	ProjectNamesStyle = lipgloss.NewStyle().
				Foreground(mutedColor)

	// Category badges
	CategoryTodayStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#FF6B6B")).
				Padding(0, 1)

	CategoryWeekStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#FFD700")).
				Padding(0, 1)

	CategoryNotUrgentStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#98FB98")).
				Padding(0, 1)

	CategoryGeneralStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(mutedColor).
				Padding(0, 1)
)
