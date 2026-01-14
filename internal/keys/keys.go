package keys

import (
	"t7t/internal/i18n"

	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	// Navigation
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	NextTab  key.Binding
	PrevTab  key.Binding
	Projects key.Binding

	// Task operations
	NewTask        key.Binding
	NewTaskGeneral key.Binding
	EditTask       key.Binding
	DeleteTask     key.Binding
	CompleteTask   key.Binding
	DeleteDone     key.Binding
	AssocProjects  key.Binding

	// Move task
	MoveToday     key.Binding
	MoveWeek      key.Binding
	MoveNotUrgent key.Binding
	MoveGeneral   key.Binding

	// Project operations
	NewProject      key.Binding
	EditProject     key.Binding
	DeleteProject   key.Binding
	CompleteProject key.Binding

	// General
	Help     key.Binding
	Quit     key.Binding
	Enter    key.Binding
	Escape   key.Binding
	SaveForm key.Binding
	Language key.Binding
}

var Keys KeyMap

func init() {
	UpdateKeybindings()
}

func UpdateKeybindings() {
	msg := i18n.Get()

	Keys = KeyMap{
		// Navigation
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("k/up", msg.KeyUp),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("j/down", msg.KeyDown),
		),
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("h/left", msg.KeyLeft),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("l/right", msg.KeyRight),
		),
		NextTab: key.NewBinding(
			key.WithKeys("tab", "l"),
			key.WithHelp("tab", msg.KeyNextTab),
		),
		PrevTab: key.NewBinding(
			key.WithKeys("shift+tab", "h"),
			key.WithHelp("shift+tab", msg.KeyPrevTab),
		),
		Projects: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", msg.KeyProjects),
		),

		// Task operations
		NewTask: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", msg.KeyNewTask),
		),
		NewTaskGeneral: key.NewBinding(
			key.WithKeys("A"),
			key.WithHelp("A", msg.KeyNewTaskGen),
		),
		EditTask: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", msg.KeyEditTask),
		),
		DeleteTask: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", msg.KeyDeleteTask),
		),
		CompleteTask: key.NewBinding(
			key.WithKeys("x", " "),
			key.WithHelp("x/space", msg.KeyCompleteTask),
		),
		DeleteDone: key.NewBinding(
			key.WithKeys("D"),
			key.WithHelp("D", msg.KeyDeleteDone),
		),
		AssocProjects: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", msg.KeyAssocProjects),
		),

		// Move task
		MoveToday: key.NewBinding(
			key.WithKeys("1"),
			key.WithHelp("1", msg.KeyMoveToday),
		),
		MoveWeek: key.NewBinding(
			key.WithKeys("2"),
			key.WithHelp("2", msg.KeyMoveWeek),
		),
		MoveNotUrgent: key.NewBinding(
			key.WithKeys("3"),
			key.WithHelp("3", msg.KeyMoveNotUrgent),
		),
		MoveGeneral: key.NewBinding(
			key.WithKeys("4"),
			key.WithHelp("4", msg.KeyMoveGeneral),
		),

		// Project operations
		NewProject: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", msg.KeyNewProject),
		),
		EditProject: key.NewBinding(
			key.WithKeys("e", "E"),
			key.WithHelp("e", msg.KeyEditProject),
		),
		DeleteProject: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", msg.KeyDeleteTask),
		),
		CompleteProject: key.NewBinding(
			key.WithKeys("X", " "),
			key.WithHelp("X/space", msg.KeyCompleteProj),
		),

		// General
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", msg.KeyHelp),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", msg.KeyQuit),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", msg.KeyEnter),
		),
		Escape: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", msg.KeyEscape),
		),
		SaveForm: key.NewBinding(
			key.WithKeys("shift+enter", "ctrl+s"),
			key.WithHelp("shift+enter/ctrl+s", msg.KeySaveForm),
		),
		Language: key.NewBinding(
			key.WithKeys("L"),
			key.WithHelp("L", msg.KeyLanguage),
		),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NewTask, k.CompleteTask, k.EditTask, k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.NextTab, k.PrevTab, k.Projects},
		{k.NewTask, k.NewTaskGeneral, k.EditTask, k.DeleteTask},
		{k.CompleteTask, k.DeleteDone, k.AssocProjects},
		{k.MoveToday, k.MoveWeek, k.MoveNotUrgent, k.MoveGeneral},
		{k.NewProject, k.EditProject, k.CompleteProject},
		{k.Help, k.Language, k.Quit, k.Escape},
	}
}
