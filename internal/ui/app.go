package ui

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"t7t/internal/i18n"
	"t7t/internal/keys"
	"t7t/internal/model"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/maaslalani/confetty/confetti"
	"github.com/maaslalani/confetty/simulation"
)

type ViewMode int

const (
	ViewTasks ViewMode = iota
	ViewProjects
)

const confettiDuration = 1300 * time.Millisecond
const confettiFPS = 30.0

type confettiFrameMsg time.Time

func animateConfetti() tea.Cmd {
	return tea.Tick(time.Second/confettiFPS, func(t time.Time) tea.Msg {
		return confettiFrameMsg(t)
	})
}

type ModalType int

const (
	ModalNone ModalType = iota
	ModalNewTask
	ModalEditTask
	ModalNewProject
	ModalEditProject
	ModalAssociateProjects
	ModalHelp
	ModalConfirmDelete
	ModalLanguage
)

type App struct {
	store *model.Store

	viewMode   ViewMode
	activeTab  int
	tabs       []string
	categories []model.Category

	taskIndex    int
	projectIndex int

	taskListViewport         viewport.Model
	taskDetailViewport       viewport.Model
	projectListViewport      viewport.Model
	taskFormModalViewport    viewport.Model
	taskFormProjectsViewport viewport.Model
	projectsModalViewport    viewport.Model
	helpModalViewport        viewport.Model

	detailFocused bool

	lastModal ModalType

	modal         ModalType
	nameInput     textinput.Model
	descInput     textarea.Model
	selectedProjs map[string]bool

	help       help.Model
	showHelp   bool
	width      int
	height     int
	statusMsg  string
	statusErr  bool
	mdRenderer *glamour.TermRenderer

	editingTaskID    string
	editingProjectID string
	focusedInput     int

	deleteType string
	deleteID   string
	deleteName string

	confettiSystem  *simulation.System
	showConfetti    bool
	confettiEndTime time.Time

	languageIndex int
}

func NewApp(store *model.Store) *App {
	msg := i18n.Get()

	nameInput := textinput.New()
	nameInput.Placeholder = msg.PlaceholderName
	nameInput.CharLimit = 100
	nameInput.Width = 40
	nameInput.Blur()

	descInput := textarea.New()
	descInput.Placeholder = msg.PlaceholderDesc
	descInput.CharLimit = 2000
	descInput.SetWidth(40)
	descInput.SetHeight(6)
	descInput.Blur()

	h := help.New()
	h.ShowAll = false

	mdRenderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(60),
	)

	return &App{
		store:    store,
		viewMode: ViewTasks,
		tabs:     []string{msg.TabToday, msg.TabWeek, msg.TabNotUrgent, msg.TabGeneral},
		categories: []model.Category{
			model.CategoryToday,
			model.CategoryWeek,
			model.CategoryNotUrgent,
			model.CategoryGeneral,
		},
		activeTab:     0,
		taskIndex:     0,
		projectIndex:  0,
		modal:         ModalNone,
		nameInput:     nameInput,
		descInput:     descInput,
		selectedProjs: make(map[string]bool),
		help:          h,
		showHelp:      false,
		mdRenderer:    mdRenderer,
		focusedInput:  0,
		confettiSystem: &simulation.System{
			Particles: []*simulation.Particle{},
			Frame:     simulation.Frame{},
		},
		showConfetti:             false,
		taskListViewport:         viewport.New(0, 0),
		taskDetailViewport:       viewport.New(0, 0),
		projectListViewport:      viewport.New(0, 0),
		taskFormModalViewport:    viewport.New(0, 0),
		taskFormProjectsViewport: viewport.New(0, 0),
		projectsModalViewport:    viewport.New(0, 0),
		helpModalViewport:        viewport.New(0, 0),
		lastModal:                ModalNone,
		detailFocused:            false,
		languageIndex:            0,
	}
}

func (a *App) updateTexts() {
	msg := i18n.Get()
	a.tabs = []string{msg.TabToday, msg.TabWeek, msg.TabNotUrgent, msg.TabGeneral}
	a.nameInput.Placeholder = msg.PlaceholderName
	a.descInput.Placeholder = msg.PlaceholderDesc
	keys.UpdateKeybindings()
}

func (a *App) Init() tea.Cmd {
	return nil
}

func (a *App) startConfetti() tea.Cmd {
	a.showConfetti = true
	a.confettiEndTime = time.Now().Add(confettiDuration)
	a.confettiSystem.Particles = confetti.Spawn(a.width, a.height)
	return animateConfetti()
}

func (a *App) checkAllTodayTasksCompleted() bool {
	tasks := a.store.GetTasksByCategory(model.CategoryToday)
	if len(tasks) == 0 {
		return false
	}
	for _, t := range tasks {
		if !t.Completed {
			return false
		}
	}
	return true
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		a.help.Width = msg.Width
		a.confettiSystem.Frame.Width = msg.Width
		a.confettiSystem.Frame.Height = msg.Height

		// Calculate consistent header height:
		// logo (4) + gap (1) + tabs (1) + gap (1) + status (2) + panel border/padding (4)
		headerAndFooter := 13
		listWidth := a.width*60/100 - 4
		detailWidth := a.width*40/100 - 4
		contentHeight := a.height - headerAndFooter
		if contentHeight < 5 {
			contentHeight = 5
		}
		a.taskListViewport = viewport.New(listWidth, contentHeight)
		a.taskDetailViewport = viewport.New(detailWidth, contentHeight)
		a.projectListViewport = viewport.New(a.width-4, contentHeight)

		// Viewport para lista de projetos no modal de tarefa
		projectsListHeight := 6
		if a.height < 30 {
			projectsListHeight = 4
		}
		a.taskFormProjectsViewport = viewport.New(40, projectsListHeight)

		// Viewport para modais
		modalHeight := a.height - 10
		if modalHeight < 10 {
			modalHeight = 10
		}
		a.helpModalViewport = viewport.New(50, modalHeight)
		a.projectsModalViewport = viewport.New(40, modalHeight)

		return a, nil

	case confettiFrameMsg:
		if a.showConfetti {
			if time.Now().After(a.confettiEndTime) {
				a.showConfetti = false
				a.confettiSystem.Particles = []*simulation.Particle{}
				return a, nil
			}
			a.confettiSystem.Update()
			return a, animateConfetti()
		}
		return a, nil

	case tea.KeyMsg:
		if a.modal != ModalNone {
			return a.handleModalInput(msg)
		}

		switch {
		case key.Matches(msg, keys.Keys.Quit):
			return a, tea.Quit

		case key.Matches(msg, keys.Keys.Help):
			if a.modal == ModalNone {
				a.modal = ModalHelp
			}
			return a, nil

		case key.Matches(msg, keys.Keys.Language):
			if a.modal == ModalNone {
				a.modal = ModalLanguage
				// Set current language as selected
				langs := i18n.AvailableLanguages()
				currentLang := i18n.GetLanguage()
				for i, lang := range langs {
					if lang == currentLang {
						a.languageIndex = i
						break
					}
				}
			}
			return a, nil

		case msg.String() == "P":
			if a.viewMode == ViewTasks {
				a.viewMode = ViewProjects
				a.projectIndex = 0
			} else {
				a.viewMode = ViewTasks
			}
			return a, nil
		}

		if a.viewMode == ViewTasks {
			return a.handleTasksInput(msg)
		} else {
			return a.handleProjectsInput(msg)
		}
	}

	return a, tea.Batch(cmds...)
}

func (a *App) handleTasksInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	tasks := a.store.GetTasksByCategory(a.categories[a.activeTab])
	keyStr := msg.String()
	m := i18n.Get()

	// Quando o painel de detalhes esta focado
	if a.detailFocused {
		switch {
		case keyStr == "h" || keyStr == "left":
			a.detailFocused = false
			return a, nil
		case key.Matches(msg, keys.Keys.Down):
			a.taskDetailViewport.LineDown(1)
			return a, nil
		case key.Matches(msg, keys.Keys.Up):
			a.taskDetailViewport.LineUp(1)
			return a, nil
		case key.Matches(msg, keys.Keys.Escape):
			a.detailFocused = false
			return a, nil
		}
		// Permitir outras acoes mesmo com detalhe focado
	}

	// "l" ou "right" foca no painel de detalhes
	if !a.detailFocused && (keyStr == "l" || keyStr == "right") {
		a.detailFocused = true
		return a, nil
	}

	switch {
	case key.Matches(msg, keys.Keys.NextTab):
		a.activeTab = (a.activeTab + 1) % len(a.tabs)
		a.taskIndex = 0
		a.detailFocused = false
		a.taskDetailViewport.GotoTop()
		return a, nil

	case key.Matches(msg, keys.Keys.PrevTab):
		a.activeTab = (a.activeTab - 1 + len(a.tabs)) % len(a.tabs)
		a.taskIndex = 0
		a.detailFocused = false
		a.taskDetailViewport.GotoTop()
		return a, nil

	case key.Matches(msg, keys.Keys.Down):
		if len(tasks) > 0 {
			oldIndex := a.taskIndex
			a.taskIndex = (a.taskIndex + 1) % len(tasks)
			if oldIndex == len(tasks)-1 && a.taskIndex == 0 {
				a.taskListViewport.GotoTop()
			} else {
				a.taskListViewport.LineDown(1)
			}
			a.taskDetailViewport.GotoTop()
		}
		return a, nil

	case key.Matches(msg, keys.Keys.Up):
		if len(tasks) > 0 {
			oldIndex := a.taskIndex
			a.taskIndex = (a.taskIndex - 1 + len(tasks)) % len(tasks)
			if oldIndex == 0 && a.taskIndex == len(tasks)-1 {
				a.taskListViewport.GotoBottom()
			} else {
				a.taskListViewport.LineUp(1)
			}
			a.taskDetailViewport.GotoTop()
		}
		return a, nil

	case key.Matches(msg, keys.Keys.NewTask):
		a.modal = ModalNewTask
		a.nameInput.Reset()
		a.descInput.Reset()
		a.selectedProjs = make(map[string]bool)
		a.projectIndex = 0
		a.focusedInput = 0
		a.nameInput.Focus()
		a.descInput.Blur()
		return a, textinput.Blink

	case key.Matches(msg, keys.Keys.NewTaskGeneral):
		a.modal = ModalNewTask
		a.nameInput.Reset()
		a.descInput.Reset()
		a.selectedProjs = make(map[string]bool)
		a.projectIndex = 0
		a.focusedInput = 0
		a.nameInput.Focus()
		a.descInput.Blur()
		a.activeTab = 3
		return a, textinput.Blink

	case key.Matches(msg, keys.Keys.EditTask):
		if len(tasks) > 0 && a.taskIndex < len(tasks) {
			task := tasks[a.taskIndex]
			a.modal = ModalEditTask
			a.editingTaskID = task.ID
			a.nameInput.SetValue(task.Name)
			a.descInput.SetValue(task.Description)
			a.selectedProjs = make(map[string]bool)
			for _, pid := range task.ProjectIDs {
				a.selectedProjs[pid] = true
			}
			a.projectIndex = 0
			a.focusedInput = 0
			a.nameInput.Focus()
			a.descInput.Blur()
			return a, textinput.Blink
		}
		return a, nil

	case key.Matches(msg, keys.Keys.DeleteTask):
		if len(tasks) > 0 && a.taskIndex < len(tasks) {
			task := tasks[a.taskIndex]
			a.modal = ModalConfirmDelete
			a.deleteType = "task"
			a.deleteID = task.ID
			a.deleteName = task.Name
		}
		return a, nil

	case key.Matches(msg, keys.Keys.CompleteTask):
		if len(tasks) > 0 && a.taskIndex < len(tasks) {
			task := tasks[a.taskIndex]
			task.ToggleComplete()
			a.store.UpdateTask(task)
			if task.Completed {
				a.statusMsg = m.StatusTaskCompleted
				if a.activeTab == 0 && a.checkAllTodayTasksCompleted() {
					a.statusMsg = m.StatusAllTodayDone
					return a, a.startConfetti()
				}
			} else {
				a.statusMsg = m.StatusTaskReopened
			}
		}
		return a, nil

	case key.Matches(msg, keys.Keys.DeleteDone):
		category := a.categories[a.activeTab]
		var count int
		for _, t := range tasks {
			if t.Completed {
				count++
			}
		}
		if count > 0 {
			a.modal = ModalConfirmDelete
			a.deleteType = "completed"
			a.deleteID = string(category)
			a.deleteName = fmt.Sprintf(m.DeleteCountFormat, count)
		}
		return a, nil

	case key.Matches(msg, keys.Keys.AssocProjects):
		if len(tasks) > 0 && a.taskIndex < len(tasks) {
			task := tasks[a.taskIndex]
			a.modal = ModalAssociateProjects
			a.editingTaskID = task.ID
			a.selectedProjs = make(map[string]bool)
			for _, pid := range task.ProjectIDs {
				a.selectedProjs[pid] = true
			}
			a.projectIndex = 0
		}
		return a, nil

	case key.Matches(msg, keys.Keys.MoveToday):
		return a.moveTask(tasks, model.CategoryToday)
	case key.Matches(msg, keys.Keys.MoveWeek):
		return a.moveTask(tasks, model.CategoryWeek)
	case key.Matches(msg, keys.Keys.MoveNotUrgent):
		return a.moveTask(tasks, model.CategoryNotUrgent)
	case key.Matches(msg, keys.Keys.MoveGeneral):
		return a.moveTask(tasks, model.CategoryGeneral)
	}

	return a, nil
}

func (a *App) moveTask(tasks []*model.Task, category model.Category) (tea.Model, tea.Cmd) {
	m := i18n.Get()
	if len(tasks) > 0 && a.taskIndex < len(tasks) {
		task := tasks[a.taskIndex]
		task.SetCategory(category)
		a.store.UpdateTask(task)
		if a.taskIndex >= len(tasks)-1 && a.taskIndex > 0 {
			a.taskIndex--
		}
		a.statusMsg = fmt.Sprintf(m.StatusTaskMoved, model.CategoryString(category))
	}
	return a, nil
}

func (a *App) handleProjectsInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	projects := a.store.GetProjects()
	m := i18n.Get()

	switch {
	case key.Matches(msg, keys.Keys.Down):
		if len(projects) > 0 {
			oldIndex := a.projectIndex
			a.projectIndex = (a.projectIndex + 1) % len(projects)
			if oldIndex == len(projects)-1 && a.projectIndex == 0 {
				a.projectListViewport.GotoTop()
			} else {
				a.projectListViewport.LineDown(1)
			}
		}
		return a, nil

	case key.Matches(msg, keys.Keys.Up):
		if len(projects) > 0 {
			oldIndex := a.projectIndex
			a.projectIndex = (a.projectIndex - 1 + len(projects)) % len(projects)
			if oldIndex == 0 && a.projectIndex == len(projects)-1 {
				a.projectListViewport.GotoBottom()
			} else {
				a.projectListViewport.LineUp(1)
			}
		}
		return a, nil

	case key.Matches(msg, keys.Keys.NewProject):
		a.modal = ModalNewProject
		a.nameInput.Reset()
		a.descInput.Blur()
		a.nameInput.Focus()
		return a, textinput.Blink

	case key.Matches(msg, keys.Keys.EditProject):
		if len(projects) > 0 && a.projectIndex < len(projects) {
			proj := projects[a.projectIndex]
			a.modal = ModalEditProject
			a.editingProjectID = proj.ID
			a.nameInput.SetValue(proj.Name)
			a.descInput.Blur()
			a.nameInput.Focus()
			return a, textinput.Blink
		}
		return a, nil

	case key.Matches(msg, keys.Keys.DeleteProject):
		if len(projects) > 0 && a.projectIndex < len(projects) {
			proj := projects[a.projectIndex]
			a.modal = ModalConfirmDelete
			a.deleteType = "project"
			a.deleteID = proj.ID
			a.deleteName = proj.Name
		}
		return a, nil

	case key.Matches(msg, keys.Keys.CompleteProject):
		if len(projects) > 0 && a.projectIndex < len(projects) {
			proj := projects[a.projectIndex]
			proj.ToggleComplete()
			a.store.UpdateProject(proj)
			if proj.Completed {
				a.statusMsg = m.StatusProjectCompleted
			} else {
				a.statusMsg = m.StatusProjectReopened
			}
		}
		return a, nil
	}

	return a, nil
}

func (a *App) handleModalInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	keyStr := msg.String()
	m := i18n.Get()

	if key.Matches(msg, keys.Keys.Escape) {
		a.modal = ModalNone
		a.nameInput.Blur()
		a.descInput.Blur()
		return a, nil
	}

	// Shift+Enter ou Ctrl+S salva formularios
	if keyStr == "shift+enter" || keyStr == "ctrl+s" {
		if a.modal == ModalNewTask || a.modal == ModalEditTask ||
			a.modal == ModalNewProject || a.modal == ModalEditProject ||
			a.modal == ModalAssociateProjects {
			return a.confirmModal()
		}
	}

	if a.modal == ModalHelp {
		switch {
		case key.Matches(msg, keys.Keys.Down):
			a.helpModalViewport.LineDown(1)
		case key.Matches(msg, keys.Keys.Up):
			a.helpModalViewport.LineUp(1)
		default:
			a.modal = ModalNone
		}
		return a, nil
	}

	if a.modal == ModalLanguage {
		langs := i18n.AvailableLanguages()
		switch {
		case key.Matches(msg, keys.Keys.Down):
			a.languageIndex = (a.languageIndex + 1) % len(langs)
		case key.Matches(msg, keys.Keys.Up):
			a.languageIndex = (a.languageIndex - 1 + len(langs)) % len(langs)
		case key.Matches(msg, keys.Keys.Enter):
			i18n.SetLanguage(langs[a.languageIndex])
			a.updateTexts()
			a.statusMsg = i18n.Get().LanguageChanged
			a.modal = ModalNone
		}
		return a, nil
	}

	if a.modal == ModalConfirmDelete {
		switch msg.String() {
		case "y", "Y", "s", "S", "enter":
			if a.deleteType == "task" {
				a.store.DeleteTask(a.deleteID)
				tasks := a.store.GetTasksByCategory(a.categories[a.activeTab])
				if a.taskIndex >= len(tasks) && a.taskIndex > 0 {
					a.taskIndex--
				}
				a.statusMsg = m.StatusTaskDeleted
			} else if a.deleteType == "project" {
				a.store.DeleteProject(a.deleteID)
				projects := a.store.GetProjects()
				if a.projectIndex >= len(projects) && a.projectIndex > 0 {
					a.projectIndex--
				}
				a.statusMsg = m.StatusProjectDeleted
			} else if a.deleteType == "completed" {
				category := model.Category(a.deleteID)
				a.store.DeleteCompletedTasks(category)
				a.taskIndex = 0
				a.statusMsg = m.StatusCompletedDeleted
			}
			a.modal = ModalNone
			a.deleteType = ""
			a.deleteID = ""
			a.deleteName = ""
			return a, nil
		case "n", "N", "esc":
			a.modal = ModalNone
			a.deleteType = ""
			a.deleteID = ""
			a.deleteName = ""
			return a, nil
		}
		return a, nil
	}

	if a.modal == ModalAssociateProjects {
		projects := a.store.GetProjects()
		switch {
		case key.Matches(msg, keys.Keys.Down):
			if len(projects) > 0 {
				oldIndex := a.projectIndex
				a.projectIndex = (a.projectIndex + 1) % len(projects)
				if oldIndex == len(projects)-1 && a.projectIndex == 0 {
					a.projectsModalViewport.GotoTop()
				} else {
					a.projectsModalViewport.LineDown(1)
				}
			}
		case key.Matches(msg, keys.Keys.Up):
			if len(projects) > 0 {
				oldIndex := a.projectIndex
				a.projectIndex = (a.projectIndex - 1 + len(projects)) % len(projects)
				if oldIndex == 0 && a.projectIndex == len(projects)-1 {
					a.projectsModalViewport.GotoBottom()
				} else {
					a.projectsModalViewport.LineUp(1)
				}
			}
		case msg.String() == " ":
			if len(projects) > 0 && a.projectIndex < len(projects) {
				proj := projects[a.projectIndex]
				a.selectedProjs[proj.ID] = !a.selectedProjs[proj.ID]
			}
		case key.Matches(msg, keys.Keys.Enter):
			return a.confirmModal()
		}
		return a, nil
	}

	if a.modal == ModalNewProject || a.modal == ModalEditProject {
		if key.Matches(msg, keys.Keys.Enter) {
			return a.confirmModal()
		}
		a.nameInput, cmd = a.nameInput.Update(msg)
		return a, cmd
	}

	if a.modal == ModalNewTask || a.modal == ModalEditTask {
		projects := a.store.GetProjects()

		if a.focusedInput == 2 {
			switch keyStr := msg.String(); keyStr {
			case "tab":
				a.focusedInput = 0
				a.nameInput.Focus()
				return a, nil
			case "shift+tab":
				a.focusedInput = 1
				a.descInput.Focus()
				return a, nil
			case "j", "down":
				if len(projects) > 0 {
					oldIndex := a.projectIndex
					a.projectIndex = (a.projectIndex + 1) % len(projects)
					if oldIndex == len(projects)-1 && a.projectIndex == 0 {
						a.taskFormProjectsViewport.GotoTop()
					} else {
						a.taskFormProjectsViewport.LineDown(1)
					}
				}
				return a, nil
			case "k", "up":
				if len(projects) > 0 {
					oldIndex := a.projectIndex
					a.projectIndex = (a.projectIndex - 1 + len(projects)) % len(projects)
					if oldIndex == 0 && a.projectIndex == len(projects)-1 {
						a.taskFormProjectsViewport.GotoBottom()
					} else {
						a.taskFormProjectsViewport.LineUp(1)
					}
				}
				return a, nil
			case " ", "enter":
				if len(projects) > 0 && a.projectIndex < len(projects) {
					proj := projects[a.projectIndex]
					a.selectedProjs[proj.ID] = !a.selectedProjs[proj.ID]
				}
				return a, nil
			}
		}

		switch msg.String() {
		case "tab":
			if a.focusedInput == 0 {
				a.focusedInput = 1
				a.nameInput.Blur()
				a.descInput.Focus()
			} else if a.focusedInput == 1 {
				a.focusedInput = 2
				a.descInput.Blur()
			}
			return a, nil

		case "shift+tab":
			if a.focusedInput == 1 {
				a.focusedInput = 0
				a.descInput.Blur()
				a.nameInput.Focus()
			} else if a.focusedInput == 0 {
				a.focusedInput = 2
				a.nameInput.Blur()
			}
			return a, nil

		case "enter":
			if a.focusedInput == 0 {
				a.focusedInput = 1
				a.nameInput.Blur()
				a.descInput.Focus()
				return a, nil
			}
		}

		if a.focusedInput == 0 {
			a.nameInput, cmd = a.nameInput.Update(msg)
		} else if a.focusedInput == 1 {
			a.descInput, cmd = a.descInput.Update(msg)
		}
		return a, cmd
	}

	return a, nil
}

func (a *App) confirmModal() (tea.Model, tea.Cmd) {
	m := i18n.Get()

	getSelectedProjectIDs := func() []string {
		var ids []string
		for pid, selected := range a.selectedProjs {
			if selected {
				ids = append(ids, pid)
			}
		}
		return ids
	}

	switch a.modal {
	case ModalNewTask:
		name := strings.TrimSpace(a.nameInput.Value())
		if name != "" {
			task := model.NewTask(name, a.descInput.Value(), a.categories[a.activeTab])
			task.ProjectIDs = getSelectedProjectIDs()
			a.store.AddTask(task)
			a.statusMsg = m.StatusTaskCreated
		}

	case ModalEditTask:
		name := strings.TrimSpace(a.nameInput.Value())
		if name != "" && a.editingTaskID != "" {
			if task := a.store.GetTask(a.editingTaskID); task != nil {
				task.Update(name, a.descInput.Value())
				task.ProjectIDs = getSelectedProjectIDs()
				a.store.UpdateTask(task)
				a.statusMsg = m.StatusTaskUpdated
			}
		}

	case ModalNewProject:
		name := strings.TrimSpace(a.nameInput.Value())
		if name != "" {
			proj := model.NewProject(name)
			a.store.AddProject(proj)
			a.statusMsg = m.StatusProjectCreated
		}

	case ModalEditProject:
		name := strings.TrimSpace(a.nameInput.Value())
		if name != "" && a.editingProjectID != "" {
			if proj := a.store.GetProject(a.editingProjectID); proj != nil {
				proj.Update(name)
				a.store.UpdateProject(proj)
				a.statusMsg = m.StatusProjectUpdated
			}
		}

	case ModalAssociateProjects:
		if a.editingTaskID != "" {
			if task := a.store.GetTask(a.editingTaskID); task != nil {
				var projectIDs []string
				for pid, selected := range a.selectedProjs {
					if selected {
						projectIDs = append(projectIDs, pid)
					}
				}
				task.SetProjects(projectIDs)
				a.store.UpdateTask(task)
				a.statusMsg = m.StatusProjectsAssoc
			}
		}
	}

	a.modal = ModalNone
	return a, nil
}

func (a *App) View() string {
	if a.width == 0 {
		return i18n.Get().Loading
	}

	var content string

	if a.viewMode == ViewTasks {
		content = a.viewTasks()
	} else {
		content = a.viewProjects()
	}

	if a.modal != ModalNone {
		content = a.viewModal(content)
	}

	if a.showConfetti {
		content = a.renderConfettiOverlay(content)
	}

	return content
}

func (a *App) renderConfettiOverlay(base string) string {
	plane := make([][]string, a.height)
	for i := range plane {
		plane[i] = make([]string, a.width)
		for j := range plane[i] {
			plane[i][j] = " "
		}
	}

	for _, p := range a.confettiSystem.Particles {
		pos := p.Physics.Position()
		x := int(pos.X)
		y := int(pos.Y)

		if y >= 0 && y < a.height && x >= 0 && x < a.width {
			plane[y][x] = p.Char
		}
	}

	var confettiBuilder strings.Builder
	for i := range plane {
		for j := range plane[i] {
			confettiBuilder.WriteString(plane[i][j])
		}
		if i < len(plane)-1 {
			confettiBuilder.WriteString("\n")
		}
	}

	confettiStr := confettiBuilder.String()

	return lipgloss.JoinVertical(lipgloss.Top, base, confettiStr)
}

func (a *App) viewTasks() string {
	// Render header elements
	logo := LogoStyle.Render(LogoArt)
	tabs := a.renderTabs()

	// Calculate heights
	logoHeight := lipgloss.Height(logo)
	tabsHeight := lipgloss.Height(tabs)
	statusHeight := 2
	gaps := 2               // empty lines between elements
	panelBorderPadding := 4 // border (2) + padding (2) from ListPanelStyle

	headerHeight := logoHeight + tabsHeight + gaps
	contentHeight := a.height - headerHeight - statusHeight - panelBorderPadding

	if contentHeight < 5 {
		contentHeight = 5
	}

	// Render panels
	listWidth := a.width*60/100 - 4
	detailWidth := a.width*40/100 - 4

	tasks := a.store.GetTasksByCategory(a.categories[a.activeTab])

	listContent := a.renderTaskList(tasks, listWidth, contentHeight)
	a.taskListViewport.SetContent(listContent)

	listStyle := ListPanelStyle.Width(listWidth).Height(contentHeight)
	detailStyle := DetailPanelStyle.Width(detailWidth).Height(contentHeight)

	if a.detailFocused {
		listStyle = listStyle.BorderForeground(lipgloss.Color("240"))
		detailStyle = detailStyle.BorderForeground(lipgloss.Color("212"))
	}

	listPanel := listStyle.Render(a.taskListViewport.View())

	detailContent := a.renderTaskDetail(tasks, detailWidth, contentHeight)
	a.taskDetailViewport.SetContent(detailContent)
	detailPanel := detailStyle.Render(a.taskDetailViewport.View())

	panels := lipgloss.JoinHorizontal(lipgloss.Top, listPanel, detailPanel)
	statusBar := a.renderStatusBar()

	// Join all elements vertically
	return lipgloss.JoinVertical(lipgloss.Left,
		logo,
		"",
		tabs,
		"",
		panels,
		statusBar,
	)
}

func (a *App) renderTabs() string {
	var tabs []string
	m := i18n.Get()

	for i, tab := range a.tabs {
		var style lipgloss.Style
		if i == a.activeTab {
			style = ActiveTabStyle
		} else {
			style = InactiveTabStyle
		}
		tabs = append(tabs, style.Render(tab))
	}

	projStyle := InactiveTabStyle
	tabs = append(tabs, TabGapStyle.Render(" | "), projStyle.Render("[P] "+m.KeyProjects))

	return lipgloss.JoinHorizontal(lipgloss.Center, tabs...)
}

// contextRegex matches @word patterns (contexts)
var contextRegex = regexp.MustCompile(`@[\w-]+`)

// renderNameWithContexts renders a task name with context tags (@tag) highlighted
func renderNameWithContexts(name string, baseStyle lipgloss.Style) string {
	matches := contextRegex.FindAllStringIndex(name, -1)
	if len(matches) == 0 {
		return baseStyle.Render(name)
	}

	var result strings.Builder
	lastEnd := 0

	for _, match := range matches {
		start, end := match[0], match[1]
		// Render text before the context
		if start > lastEnd {
			result.WriteString(baseStyle.Render(name[lastEnd:start]))
		}
		// Render the context with ContextStyle
		result.WriteString(ContextStyle.Render(name[start:end]))
		lastEnd = end
	}

	// Render remaining text after last context
	if lastEnd < len(name) {
		result.WriteString(baseStyle.Render(name[lastEnd:]))
	}

	return result.String()
}

func (a *App) renderTaskList(tasks []*model.Task, width, height int) string {
	m := i18n.Get()

	if len(tasks) == 0 {
		return NormalItemStyle.Render(m.EmptyTaskList)
	}

	var lines []string
	for i, task := range tasks {
		var line string

		if i == a.taskIndex {
			line += CheckboxSelected
		} else {
			line += CheckboxNormal
		}

		if task.Completed {
			line += CheckboxChecked
		} else {
			line += CheckboxEmpty
		}

		prefixLen := 8
		availableWidth := width - prefixLen - 2

		projectNames := a.store.GetProjectNames(task.ProjectIDs)
		var projectsStr string
		if len(projectNames) > 0 {
			projectsStr = " [" + strings.Join(projectNames, ", ") + "]"
		}

		name := task.Name
		maxNameLen := availableWidth - len(projectsStr)
		if maxNameLen < 10 {
			maxNameLen = availableWidth
			projectsStr = ""
		}
		if len(name) > maxNameLen {
			name = name[:maxNameLen-3] + "..."
		}

		var style lipgloss.Style
		if i == a.taskIndex {
			style = SelectedItemStyle
		} else if task.Completed {
			style = CompletedItemStyle
		} else {
			style = NormalItemStyle
		}

		line += renderNameWithContexts(name, style)

		if projectsStr != "" {
			line += ProjectNamesStyle.Render(projectsStr)
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

func (a *App) renderTaskDetail(tasks []*model.Task, width, height int) string {
	m := i18n.Get()

	if len(tasks) == 0 || a.taskIndex >= len(tasks) {
		return NormalItemStyle.Render(m.EmptyTaskDetail)
	}

	task := tasks[a.taskIndex]
	var b strings.Builder

	b.WriteString(DetailTitleStyle.Width(width).Render(task.Name))
	b.WriteString("\n\n")

	b.WriteString(DetailLabelStyle.Render(m.LabelStatus))
	if task.Completed {
		b.WriteString(StatusMessageStyle.Render(m.LabelCompleted))
	} else {
		b.WriteString(DetailValueStyle.Render(m.LabelPending))
	}
	b.WriteString("\n\n")

	b.WriteString(DetailLabelStyle.Render(m.LabelDescription))
	b.WriteString("\n")
	if task.Description == "" {
		b.WriteString(DetailValueStyle.Render(m.LabelNoDesc))
	} else {
		if a.mdRenderer != nil {
			rendered, err := a.mdRenderer.Render(task.Description)
			if err == nil {
				b.WriteString(strings.TrimSpace(rendered))
			} else {
				b.WriteString(DetailValueStyle.Render(task.Description))
			}
		}
	}
	b.WriteString("\n\n")

	b.WriteString(DetailLabelStyle.Render(m.LabelProjects))
	b.WriteString("\n")
	projectNames := a.store.GetProjectNames(task.ProjectIDs)
	if len(projectNames) == 0 {
		b.WriteString(DetailValueStyle.Render(m.LabelNoProjects))
	} else {
		var projLines []string
		for _, name := range projectNames {
			projLines = append(projLines, "- "+name)
		}
		b.WriteString(DetailValueStyle.Render(strings.Join(projLines, "\n")))
	}

	return b.String()
}

func (a *App) viewProjects() string {
	m := i18n.Get()

	// Render header elements
	logo := LogoStyle.Render(LogoArt)
	tab := InactiveTabStyle.Render(m.ProjectsBackToTasks)

	// Calculate heights
	logoHeight := lipgloss.Height(logo)
	tabHeight := lipgloss.Height(tab)
	statusHeight := 2
	gaps := 2               // empty lines between elements
	panelBorderPadding := 4 // border (2) + padding (2) from ListPanelStyle

	headerHeight := logoHeight + tabHeight + gaps
	contentHeight := a.height - headerHeight - statusHeight - panelBorderPadding

	if contentHeight < 5 {
		contentHeight = 5
	}

	// Render content
	projects := a.store.GetProjects()
	var panel string

	if len(projects) == 0 {
		panel = NormalItemStyle.Render(m.EmptyProjectList)
	} else {
		listWidth := a.width - 4
		listContent := a.renderProjectList(projects, listWidth)
		a.projectListViewport.SetContent(listContent)
		panel = ListPanelStyle.Width(listWidth).Height(contentHeight).Render(a.projectListViewport.View())
	}

	// Render status bar
	var parts []string
	if a.statusMsg != "" {
		parts = append(parts, StatusMessageStyle.Render(a.statusMsg))
	}

	helpText := HelpKeyStyle.Render("a") + HelpDescStyle.Render(":"+m.HelpNew+" ") +
		HelpKeyStyle.Render("e") + HelpDescStyle.Render(":"+m.HelpEdit+" ") +
		HelpKeyStyle.Render("d") + HelpDescStyle.Render(":"+m.HelpDelete+" ") +
		HelpKeyStyle.Render("space") + HelpDescStyle.Render(":"+m.HelpComplete+" ") +
		HelpKeyStyle.Render("P") + HelpDescStyle.Render(":"+m.HelpBack+" ") +
		HelpKeyStyle.Render("q") + HelpDescStyle.Render(":"+m.HelpQuit)

	parts = append(parts, helpText)
	statusBar := StatusBarStyle.Render(strings.Join(parts, " | "))

	// Join all elements vertically
	return lipgloss.JoinVertical(lipgloss.Left,
		logo,
		"",
		tab,
		"",
		panel,
		statusBar,
	)
}

func (a *App) renderProjectList(projects []*model.Project, width int) string {
	m := i18n.Get()
	var lines []string

	for i, proj := range projects {
		var line string

		if i == a.projectIndex {
			line += CheckboxSelected
		} else {
			line += CheckboxNormal
		}

		if proj.Completed {
			line += CheckboxChecked
		} else {
			line += CheckboxEmpty
		}

		name := proj.Name

		var taskCount int
		for _, t := range a.store.Tasks {
			if t.HasProject(proj.ID) && !t.Completed {
				taskCount++
			}
		}

		var style lipgloss.Style
		if i == a.projectIndex {
			style = SelectedItemStyle
		} else if proj.Completed {
			style = CompletedItemStyle
		} else {
			style = NormalItemStyle
		}

		line += style.Render(name)

		if taskCount > 0 {
			line += " " + ProjectBadgeStyle.Render(fmt.Sprintf(m.ProjectsTaskCount, taskCount))
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

func (a *App) renderStatusBar() string {
	m := i18n.Get()
	var parts []string

	if a.statusMsg != "" {
		parts = append(parts, StatusMessageStyle.Render(a.statusMsg))
	}

	var helpText string
	if a.detailFocused {
		helpText = HelpKeyStyle.Render("j/k") + HelpDescStyle.Render(":"+m.HelpScroll+" ") +
			HelpKeyStyle.Render("h") + HelpDescStyle.Render(":"+m.HelpBack+" ") +
			HelpKeyStyle.Render("?") + HelpDescStyle.Render(":"+m.HelpHelp+" ") +
			HelpKeyStyle.Render("q") + HelpDescStyle.Render(":"+m.HelpQuit)
	} else {
		helpText = HelpKeyStyle.Render("a") + HelpDescStyle.Render(":"+m.HelpNew+" ") +
			HelpKeyStyle.Render("space") + HelpDescStyle.Render(":"+m.HelpComplete+" ") +
			HelpKeyStyle.Render("e") + HelpDescStyle.Render(":"+m.HelpEdit+" ") +
			HelpKeyStyle.Render("l") + HelpDescStyle.Render(":"+m.HelpDetails+" ") +
			HelpKeyStyle.Render("?") + HelpDescStyle.Render(":"+m.HelpHelp+" ") +
			HelpKeyStyle.Render("q") + HelpDescStyle.Render(":"+m.HelpQuit)
	}

	parts = append(parts, helpText)
	return StatusBarStyle.Render(strings.Join(parts, " | "))
}

func (a *App) viewModal(background string) string {
	var modalContent string

	switch a.modal {
	case ModalNewTask:
		modalContent = a.renderTaskForm(i18n.Get().ModalNewTask)
	case ModalEditTask:
		modalContent = a.renderTaskForm(i18n.Get().ModalEditTask)
	case ModalNewProject:
		modalContent = a.renderProjectForm(i18n.Get().ModalNewProject)
	case ModalEditProject:
		modalContent = a.renderProjectForm(i18n.Get().ModalEditProject)
	case ModalAssociateProjects:
		modalContent = a.renderAssociateProjectsModal()
	case ModalHelp:
		modalContent = a.renderHelpModal()
	case ModalConfirmDelete:
		modalContent = a.renderConfirmDeleteModal()
	case ModalLanguage:
		modalContent = a.renderLanguageModal()
	}

	modal := ModalStyle.Render(modalContent)

	modalWidth := lipgloss.Width(modal)
	modalHeight := lipgloss.Height(modal)

	x := (a.width - modalWidth) / 2
	y := (a.height - modalHeight) / 2

	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

	return lipgloss.Place(a.width, a.height, lipgloss.Center, lipgloss.Center, modal)
}

func (a *App) renderTaskForm(title string) string {
	m := i18n.Get()
	var b strings.Builder

	b.WriteString(ModalTitleStyle.Render(title))
	b.WriteString("\n\n")

	nameLabel := m.FormName
	if a.focusedInput == 0 {
		nameLabel = "> " + m.FormName
	}
	b.WriteString(InputLabelStyle.Render(nameLabel))
	b.WriteString("\n")
	b.WriteString(a.nameInput.View())
	b.WriteString("\n\n")

	descLabel := m.FormDescMarkdown
	if a.focusedInput == 1 {
		descLabel = "> " + m.FormDescMarkdown
	}
	b.WriteString(InputLabelStyle.Render(descLabel))
	b.WriteString("\n")
	b.WriteString(a.descInput.View())
	b.WriteString("\n\n")

	projLabel := m.LabelProjects
	if a.focusedInput == 2 {
		projLabel = "> " + m.LabelProjects
	}
	b.WriteString(InputLabelStyle.Render(projLabel))
	b.WriteString("\n")

	projects := a.store.GetProjects()

	if len(projects) == 0 {
		b.WriteString(HelpDescStyle.Render("  " + m.FormNoProjects))
	} else {
		var projectsList strings.Builder

		for i, proj := range projects {
			var line string

			if a.focusedInput == 2 && i == a.projectIndex {
				line = " >"
			} else {
				line = "  "
			}

			if a.selectedProjs[proj.ID] {
				line += " [x] "
			} else {
				line += " [ ] "
			}

			var style lipgloss.Style
			if a.focusedInput == 2 && i == a.projectIndex {
				style = SelectedItemStyle
			} else if a.selectedProjs[proj.ID] {
				style = DetailValueStyle
			} else {
				style = HelpDescStyle
			}

			line += style.Render(proj.Name)
			projectsList.WriteString(line + "\n")
		}

		a.taskFormProjectsViewport.SetContent(projectsList.String())
		b.WriteString(a.taskFormProjectsViewport.View())
	}

	b.WriteString("\n")

	if a.focusedInput == 2 {
		b.WriteString(HelpDescStyle.Render(m.HintNavProjects))
	} else {
		b.WriteString(HelpDescStyle.Render(m.HintFormFields))
	}

	return b.String()
}

func (a *App) renderProjectForm(title string) string {
	m := i18n.Get()
	var b strings.Builder

	b.WriteString(ModalTitleStyle.Render(title))
	b.WriteString("\n\n")

	b.WriteString(InputLabelStyle.Render(m.FormProjectName))
	b.WriteString("\n")
	b.WriteString(a.nameInput.View())
	b.WriteString("\n\n")

	b.WriteString(HelpDescStyle.Render(m.HintProjectForm))

	return b.String()
}

func (a *App) renderAssociateProjectsModal() string {
	m := i18n.Get()
	projects := a.store.GetProjects()

	if len(projects) == 0 {
		return NormalItemStyle.Render(m.HintNoProjectAvail)
	}

	var b strings.Builder
	b.WriteString(ModalTitleStyle.Render(m.ModalAssocProjects))
	b.WriteString("\n\n")

	for i, proj := range projects {
		var line string

		if i == a.projectIndex {
			line += CheckboxSelected
		} else {
			line += CheckboxNormal
		}

		if a.selectedProjs[proj.ID] {
			line += CheckboxChecked
		} else {
			line += CheckboxEmpty
		}

		var style lipgloss.Style
		if i == a.projectIndex {
			style = SelectedItemStyle
		} else {
			style = NormalItemStyle
		}

		line += style.Render(proj.Name)
		b.WriteString(line + "\n")
	}

	b.WriteString("\n")
	b.WriteString(HelpDescStyle.Render(m.HintAssocProjects))

	a.projectsModalViewport.SetContent(b.String())
	return a.projectsModalViewport.View()
}

func (a *App) renderHelpModal() string {
	m := i18n.Get()
	var b strings.Builder

	b.WriteString(ModalTitleStyle.Render(m.ModalHelp))
	b.WriteString("\n\n")

	helpItems := [][]string{
		{m.HelpNavSection, ""},
		{"j/k ou setas", m.HelpNavList},
		{"Tab/Shift+Tab", m.HelpNavTabs},
		{"P", m.HelpNavProjects},
		{"", ""},
		{m.HelpTaskSection, ""},
		{"a", m.HelpTaskNew},
		{"A", m.HelpTaskNewGen},
		{"e", m.HelpTaskEdit},
		{"d", m.HelpTaskDelete},
		{"space ou x", m.HelpTaskComplete},
		{"D", m.HelpTaskDeleteDone},
		{"p", m.HelpTaskAssoc},
		{"", ""},
		{m.HelpMoveSection, ""},
		{"1", m.HelpMoveToday},
		{"2", m.HelpMoveWeek},
		{"3", m.HelpMoveNotUrgent},
		{"4", m.HelpMoveGeneral},
		{"", ""},
		{m.HelpProjSection, ""},
		{"a", m.HelpProjNew},
		{"e", m.HelpProjEdit},
		{"d", m.HelpProjDelete},
		{"space ou X", m.HelpProjComplete},
		{"", ""},
		{m.HelpFormSection, ""},
		{"Tab", m.HelpFormTab},
		{"Ctrl+S", m.HelpFormSave},
		{"Esc", m.HelpFormCancel},
		{"", ""},
		{m.HelpGeneralSection, ""},
		{"?", m.HelpGeneralHelp},
		{"L", m.HelpGeneralLanguage},
		{"q ou Ctrl+C", m.HelpGeneralQuit},
	}

	for _, item := range helpItems {
		if item[0] == "" {
			b.WriteString("\n")
		} else if item[1] == "" {
			b.WriteString(DetailLabelStyle.Render(item[0]) + "\n")
		} else {
			b.WriteString(HelpKeyStyle.Render(fmt.Sprintf("%-15s", item[0])))
			b.WriteString(HelpDescStyle.Render(item[1]) + "\n")
		}
	}

	b.WriteString("\n")
	b.WriteString(HelpDescStyle.Render(m.HintCloseHelp))

	a.helpModalViewport.SetContent(b.String())
	return a.helpModalViewport.View()
}

func (a *App) renderConfirmDeleteModal() string {
	m := i18n.Get()
	var b strings.Builder

	b.WriteString(ModalTitleStyle.Render(m.ModalConfirmDelete))
	b.WriteString("\n\n")

	var message string
	switch a.deleteType {
	case "task":
		message = m.ConfirmDeleteTask
	case "project":
		message = m.ConfirmDeleteProject
	case "completed":
		message = m.ConfirmDeleteCompleted
	}

	b.WriteString(DetailValueStyle.Render(message))
	b.WriteString("\n\n")
	b.WriteString(SelectedItemStyle.Render(fmt.Sprintf("  %s", a.deleteName)))
	b.WriteString("\n\n")

	b.WriteString(HelpKeyStyle.Render("S/Enter"))
	b.WriteString(HelpDescStyle.Render(" " + m.ConfirmYes + "   "))
	b.WriteString(HelpKeyStyle.Render("N/Esc"))
	b.WriteString(HelpDescStyle.Render(" " + m.ConfirmNo))

	return b.String()
}

func (a *App) renderLanguageModal() string {
	m := i18n.Get()
	var b strings.Builder

	b.WriteString(ModalTitleStyle.Render(m.ModalLanguage))
	b.WriteString("\n\n")

	b.WriteString(DetailLabelStyle.Render(m.LanguageSelect))
	b.WriteString("\n\n")

	langs := i18n.AvailableLanguages()
	for i, lang := range langs {
		var line string

		if i == a.languageIndex {
			line += CheckboxSelected
		} else {
			line += CheckboxNormal
		}

		currentLang := i18n.GetLanguage()
		if lang == currentLang {
			line += CheckboxChecked
		} else {
			line += CheckboxEmpty
		}

		var style lipgloss.Style
		if i == a.languageIndex {
			style = SelectedItemStyle
		} else {
			style = NormalItemStyle
		}

		line += style.Render(i18n.LanguageDisplayName(lang))
		b.WriteString(line + "\n")
	}

	b.WriteString("\n")
	b.WriteString(HelpDescStyle.Render(m.LanguageHint))

	return b.String()
}
