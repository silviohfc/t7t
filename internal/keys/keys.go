package keys

import "github.com/charmbracelet/bubbles/key"

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
}

var Keys = KeyMap{
	// Navigation
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("k/up", "cima"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("j/down", "baixo"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("h/left", "esquerda"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("l/right", "direita"),
	),
	NextTab: key.NewBinding(
		key.WithKeys("tab", "l"),
		key.WithHelp("tab", "proxima aba"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("shift+tab", "h"),
		key.WithHelp("shift+tab", "aba anterior"),
	),
	Projects: key.NewBinding(
		key.WithKeys("P"),
		key.WithHelp("P", "projetos"),
	),

	// Task operations
	NewTask: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "nova tarefa"),
	),
	NewTaskGeneral: key.NewBinding(
		key.WithKeys("A"),
		key.WithHelp("A", "nova na lista geral"),
	),
	EditTask: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "editar"),
	),
	DeleteTask: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "deletar"),
	),
	CompleteTask: key.NewBinding(
		key.WithKeys("x", " "),
		key.WithHelp("x/space", "concluir"),
	),
	DeleteDone: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "deletar concluidas"),
	),
	AssocProjects: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "associar projetos"),
	),

	// Move task
	MoveToday: key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("1", "mover p/ Hoje"),
	),
	MoveWeek: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "mover p/ Semana"),
	),
	MoveNotUrgent: key.NewBinding(
		key.WithKeys("3"),
		key.WithHelp("3", "mover p/ Nao Urgente"),
	),
	MoveGeneral: key.NewBinding(
		key.WithKeys("4"),
		key.WithHelp("4", "mover p/ Lista Geral"),
	),

	// Project operations
	NewProject: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "novo projeto"),
	),
	EditProject: key.NewBinding(
		key.WithKeys("e", "E"),
		key.WithHelp("e", "editar projeto"),
	),
	DeleteProject: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "deletar projeto"),
	),
	CompleteProject: key.NewBinding(
		key.WithKeys("X", " "),
		key.WithHelp("X/space", "concluir projeto"),
	),

	// General
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "ajuda"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "sair"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirmar"),
	),
	Escape: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancelar"),
	),
	SaveForm: key.NewBinding(
		key.WithKeys("shift+enter", "ctrl+s"),
		key.WithHelp("shift+enter/ctrl+s", "salvar"),
	),
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
		{k.Help, k.Quit, k.Escape},
	}
}
