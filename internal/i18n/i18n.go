package i18n

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type Language string

const (
	Portuguese Language = "pt-BR"
	English    Language = "en"
)

type Messages struct {
	// Tab names
	TabToday     string `json:"tab_today"`
	TabWeek      string `json:"tab_week"`
	TabNotUrgent string `json:"tab_not_urgent"`
	TabGeneral   string `json:"tab_general"`

	// Category names (for status messages)
	CategoryToday     string `json:"category_today"`
	CategoryWeek      string `json:"category_week"`
	CategoryNotUrgent string `json:"category_not_urgent"`
	CategoryGeneral   string `json:"category_general"`

	// Status messages
	StatusTaskCompleted    string `json:"status_task_completed"`
	StatusAllTodayDone     string `json:"status_all_today_done"`
	StatusTaskReopened     string `json:"status_task_reopened"`
	StatusTaskMoved        string `json:"status_task_moved"`
	StatusTaskCreated      string `json:"status_task_created"`
	StatusTaskUpdated      string `json:"status_task_updated"`
	StatusProjectsAssoc    string `json:"status_projects_assoc"`
	StatusProjectCompleted string `json:"status_project_completed"`
	StatusProjectReopened  string `json:"status_project_reopened"`
	StatusProjectCreated   string `json:"status_project_created"`
	StatusProjectUpdated   string `json:"status_project_updated"`
	StatusTaskDeleted      string `json:"status_task_deleted"`
	StatusProjectDeleted   string `json:"status_project_deleted"`
	StatusCompletedDeleted string `json:"status_completed_deleted"`

	// Placeholders
	PlaceholderName string `json:"placeholder_name"`
	PlaceholderDesc string `json:"placeholder_desc"`

	// Empty states
	EmptyTaskList    string `json:"empty_task_list"`
	EmptyTaskDetail  string `json:"empty_task_detail"`
	EmptyProjectList string `json:"empty_project_list"`

	// Form labels
	LabelStatus      string `json:"label_status"`
	LabelCompleted   string `json:"label_completed"`
	LabelPending     string `json:"label_pending"`
	LabelDescription string `json:"label_description"`
	LabelNoDesc      string `json:"label_no_desc"`
	LabelProjects    string `json:"label_projects"`
	LabelNoProjects  string `json:"label_no_projects"`

	// Modal titles
	ModalNewTask       string `json:"modal_new_task"`
	ModalEditTask      string `json:"modal_edit_task"`
	ModalNewProject    string `json:"modal_new_project"`
	ModalEditProject   string `json:"modal_edit_project"`
	ModalAssocProjects string `json:"modal_assoc_projects"`
	ModalHelp          string `json:"modal_help"`
	ModalConfirmDelete string `json:"modal_confirm_delete"`
	ModalLanguage      string `json:"modal_language"`

	// Delete confirmation
	ConfirmDeleteTask      string `json:"confirm_delete_task"`
	ConfirmDeleteProject   string `json:"confirm_delete_project"`
	ConfirmDeleteCompleted string `json:"confirm_delete_completed"`
	ConfirmYes             string `json:"confirm_yes"`
	ConfirmNo              string `json:"confirm_no"`

	// Form fields
	FormName         string `json:"form_name"`
	FormDescMarkdown string `json:"form_desc_markdown"`
	FormProjectName  string `json:"form_project_name"`
	FormNoProjects   string `json:"form_no_projects"`

	// Form hints
	HintNavProjects    string `json:"hint_nav_projects"`
	HintFormFields     string `json:"hint_form_fields"`
	HintProjectForm    string `json:"hint_project_form"`
	HintAssocProjects  string `json:"hint_assoc_projects"`
	HintCloseHelp      string `json:"hint_close_help"`
	HintNoProjectAvail string `json:"hint_no_project_avail"`

	// Status bar help (tasks view)
	HelpNew      string `json:"help_new"`
	HelpComplete string `json:"help_complete"`
	HelpEdit     string `json:"help_edit"`
	HelpDetails  string `json:"help_details"`
	HelpHelp     string `json:"help_help"`
	HelpQuit     string `json:"help_quit"`
	HelpScroll   string `json:"help_scroll"`
	HelpBack     string `json:"help_back"`

	// Status bar help (projects view)
	HelpDelete string `json:"help_delete"`

	// Keybinding help text
	KeyUp            string `json:"key_up"`
	KeyDown          string `json:"key_down"`
	KeyLeft          string `json:"key_left"`
	KeyRight         string `json:"key_right"`
	KeyNextTab       string `json:"key_next_tab"`
	KeyPrevTab       string `json:"key_prev_tab"`
	KeyProjects      string `json:"key_projects"`
	KeyNewTask       string `json:"key_new_task"`
	KeyNewTaskGen    string `json:"key_new_task_gen"`
	KeyEditTask      string `json:"key_edit"`
	KeyDeleteTask    string `json:"key_delete"`
	KeyCompleteTask  string `json:"key_complete"`
	KeyDeleteDone    string `json:"key_delete_done"`
	KeyAssocProjects string `json:"key_assoc_projects"`
	KeyMoveToday     string `json:"key_move_today"`
	KeyMoveWeek      string `json:"key_move_week"`
	KeyMoveNotUrgent string `json:"key_move_not_urgent"`
	KeyMoveGeneral   string `json:"key_move_general"`
	KeyNewProject    string `json:"key_new_project"`
	KeyEditProject   string `json:"key_edit_project"`
	KeyCompleteProj  string `json:"key_complete_project"`
	KeyHelp          string `json:"key_help"`
	KeyQuit          string `json:"key_quit"`
	KeyEnter         string `json:"key_enter"`
	KeyEscape        string `json:"key_escape"`
	KeySaveForm      string `json:"key_save"`
	KeyLanguage      string `json:"key_language"`

	// Help modal keys (left column)
	HelpKeyNavList      string `json:"help_key_nav_list"`
	HelpKeyNavTabs      string `json:"help_key_nav_tabs"`
	HelpKeyComplete     string `json:"help_key_complete"`
	HelpKeyCompleteProj string `json:"help_key_complete_proj"`
	HelpKeyQuit         string `json:"help_key_quit"`

	// Help modal sections
	HelpNavSection      string `json:"help_nav_section"`
	HelpNavList         string `json:"help_nav_list"`
	HelpNavTabs         string `json:"help_nav_tabs"`
	HelpNavProjects     string `json:"help_nav_projects"`
	HelpTaskSection     string `json:"help_task_section"`
	HelpTaskNew         string `json:"help_task_new"`
	HelpTaskNewGen      string `json:"help_task_new_gen"`
	HelpTaskEdit        string `json:"help_task_edit"`
	HelpTaskDelete      string `json:"help_task_delete"`
	HelpTaskComplete    string `json:"help_task_complete"`
	HelpTaskDeleteDone  string `json:"help_task_delete_done"`
	HelpTaskAssoc       string `json:"help_task_assoc"`
	HelpMoveSection     string `json:"help_move_section"`
	HelpMoveToday       string `json:"help_move_today"`
	HelpMoveWeek        string `json:"help_move_week"`
	HelpMoveNotUrgent   string `json:"help_move_not_urgent"`
	HelpMoveGeneral     string `json:"help_move_general"`
	HelpProjSection     string `json:"help_proj_section"`
	HelpProjNew         string `json:"help_proj_new"`
	HelpProjEdit        string `json:"help_proj_edit"`
	HelpProjDelete      string `json:"help_proj_delete"`
	HelpProjComplete    string `json:"help_proj_complete"`
	HelpFormSection     string `json:"help_form_section"`
	HelpFormTab         string `json:"help_form_tab"`
	HelpFormSave        string `json:"help_form_save"`
	HelpFormCancel      string `json:"help_form_cancel"`
	HelpGeneralSection  string `json:"help_general_section"`
	HelpGeneralHelp     string `json:"help_general_help"`
	HelpGeneralQuit     string `json:"help_general_quit"`
	HelpGeneralLanguage string `json:"help_general_language"`

	// Projects view
	ProjectsBackToTasks string `json:"projects_back_to_tasks"`
	ProjectsTaskCount   string `json:"projects_task_count"`
	ProjectCompleted    string `json:"project_completed"`

	// Language selection
	LanguageSelect   string `json:"language_select"`
	LanguagePortBR   string `json:"language_port_br"`
	LanguageEnglish  string `json:"language_english"`
	LanguageHint     string `json:"language_hint"`
	LanguageChanged  string `json:"language_changed"`

	// Error messages
	ErrorInitStorage string `json:"error_init_storage"`
	ErrorRunApp      string `json:"error_run_app"`

	// Loading
	Loading string `json:"loading"`

	// Delete count format
	DeleteCountFormat string `json:"delete_count_format"`
}

var (
	currentLang Language = Portuguese
	messages    *Messages
	mu          sync.RWMutex
)

var ptBR = &Messages{
	// Tab names
	TabToday:     "Hoje",
	TabWeek:      "Essa Semana",
	TabNotUrgent: "Nao Urgente",
	TabGeneral:   "Lista Geral",

	// Category names
	CategoryToday:     "Hoje",
	CategoryWeek:      "Essa Semana",
	CategoryNotUrgent: "Nao Urgente",
	CategoryGeneral:   "Lista Geral",

	// Status messages
	StatusTaskCompleted:    "Tarefa concluida",
	StatusAllTodayDone:     "Parabens! Todas as tarefas de hoje concluidas!",
	StatusTaskReopened:     "Tarefa reaberta",
	StatusTaskMoved:        "Tarefa movida para %s",
	StatusTaskCreated:      "Tarefa criada",
	StatusTaskUpdated:      "Tarefa atualizada",
	StatusProjectsAssoc:    "Projetos associados",
	StatusProjectCompleted: "Projeto concluido",
	StatusProjectReopened:  "Projeto reaberto",
	StatusProjectCreated:   "Projeto criado",
	StatusProjectUpdated:   "Projeto atualizado",
	StatusTaskDeleted:      "Tarefa deletada",
	StatusProjectDeleted:   "Projeto deletado",
	StatusCompletedDeleted: "Tarefas concluidas deletadas",

	// Placeholders
	PlaceholderName: "Nome...",
	PlaceholderDesc: "Descricao (suporta Markdown)...",

	// Empty states
	EmptyTaskList:    "Nenhuma tarefa nesta lista.\n\nPressione 'a' para criar uma nova tarefa.",
	EmptyTaskDetail:  "Selecione uma tarefa para ver detalhes",
	EmptyProjectList: "Nenhum projeto cadastrado.\n\nPressione 'a' para criar um novo projeto.",

	// Form labels
	LabelStatus:      "Status: ",
	LabelCompleted:   "Concluida",
	LabelPending:     "Pendente",
	LabelDescription: "Descricao:",
	LabelNoDesc:      "(sem descricao)",
	LabelProjects:    "Projetos:",
	LabelNoProjects:  "(nenhum projeto)",

	// Modal titles
	ModalNewTask:       "Nova Tarefa",
	ModalEditTask:      "Editar Tarefa",
	ModalNewProject:    "Novo Projeto",
	ModalEditProject:   "Editar Projeto",
	ModalAssocProjects: "Associar Projetos",
	ModalHelp:          "Ajuda - Atalhos",
	ModalConfirmDelete: "Confirmar Exclusao",
	ModalLanguage:      "Selecionar Idioma",

	// Delete confirmation
	ConfirmDeleteTask:      "Deseja realmente deletar a tarefa?",
	ConfirmDeleteProject:   "Deseja realmente deletar o projeto?",
	ConfirmDeleteCompleted: "Deseja realmente deletar todas as tarefas concluidas?",
	ConfirmYes:             "confirmar",
	ConfirmNo:              "cancelar",

	// Form fields
	FormName:         "Nome:",
	FormDescMarkdown: "Descricao (Markdown):",
	FormProjectName:  "Nome do Projeto:",
	FormNoProjects:   "(nenhum projeto cadastrado)",

	// Form hints
	HintNavProjects:    "j/k: navegar | Space: selecionar | Tab: proximo | Ctrl+S: salvar",
	HintFormFields:     "Tab: alternar campos | Ctrl+S: salvar | Esc: cancelar",
	HintProjectForm:    "Enter: confirmar | Esc: cancelar",
	HintAssocProjects:  "Space: selecionar | Enter: confirmar | Esc: cancelar",
	HintCloseHelp:      "Pressione qualquer tecla para fechar",
	HintNoProjectAvail: "Nenhum projeto disponivel.\nCrie um projeto primeiro (P).",

	// Status bar help
	HelpNew:      "nova",
	HelpComplete: "concluir",
	HelpEdit:     "editar",
	HelpDetails:  "detalhes",
	HelpHelp:     "ajuda",
	HelpQuit:     "sair",
	HelpScroll:   "scroll",
	HelpBack:     "voltar",
	HelpDelete:   "deletar",

	// Keybinding help text
	KeyUp:            "cima",
	KeyDown:          "baixo",
	KeyLeft:          "esquerda",
	KeyRight:         "direita",
	KeyNextTab:       "proxima aba",
	KeyPrevTab:       "aba anterior",
	KeyProjects:      "projetos",
	KeyNewTask:       "nova tarefa",
	KeyNewTaskGen:    "nova na lista geral",
	KeyEditTask:      "editar",
	KeyDeleteTask:    "deletar",
	KeyCompleteTask:  "concluir",
	KeyDeleteDone:    "deletar concluidas",
	KeyAssocProjects: "associar projetos",
	KeyMoveToday:     "mover p/ Hoje",
	KeyMoveWeek:      "mover p/ Semana",
	KeyMoveNotUrgent: "mover p/ Nao Urgente",
	KeyMoveGeneral:   "mover p/ Lista Geral",
	KeyNewProject:    "novo projeto",
	KeyEditProject:   "editar projeto",
	KeyCompleteProj:  "concluir projeto",
	KeyHelp:          "ajuda",
	KeyQuit:          "sair",
	KeyEnter:         "confirmar",
	KeyEscape:        "cancelar",
	KeySaveForm:      "salvar",
	KeyLanguage:      "idioma",

	// Help modal keys (left column)
	HelpKeyNavList:      "j/k ou setas",
	HelpKeyNavTabs:      "Tab/Shift+Tab",
	HelpKeyComplete:     "space ou x",
	HelpKeyCompleteProj: "space ou X",
	HelpKeyQuit:         "q ou Ctrl+C",

	// Help modal sections
	HelpNavSection:      "Navegacao",
	HelpNavList:         "Mover na lista",
	HelpNavTabs:         "Alternar abas",
	HelpNavProjects:     "Tela de projetos",
	HelpTaskSection:     "Tarefas",
	HelpTaskNew:         "Nova tarefa na aba atual",
	HelpTaskNewGen:      "Nova tarefa na Lista Geral",
	HelpTaskEdit:        "Editar tarefa",
	HelpTaskDelete:      "Deletar tarefa",
	HelpTaskComplete:    "Concluir/reabrir tarefa",
	HelpTaskDeleteDone:  "Deletar concluidas",
	HelpTaskAssoc:       "Associar projetos a tarefa",
	HelpMoveSection:     "Mover tarefa",
	HelpMoveToday:       "Mover para Hoje",
	HelpMoveWeek:        "Mover para Essa Semana",
	HelpMoveNotUrgent:   "Mover para Nao Urgente",
	HelpMoveGeneral:     "Mover para Lista Geral",
	HelpProjSection:     "Projetos (tela P)",
	HelpProjNew:         "Novo projeto",
	HelpProjEdit:        "Editar projeto",
	HelpProjDelete:      "Deletar projeto",
	HelpProjComplete:    "Concluir projeto",
	HelpFormSection:     "Formularios",
	HelpFormTab:         "Proximo campo",
	HelpFormSave:        "Salvar",
	HelpFormCancel:      "Cancelar",
	HelpGeneralSection:  "Geral",
	HelpGeneralHelp:     "Mostrar/fechar ajuda",
	HelpGeneralQuit:     "Sair",
	HelpGeneralLanguage: "Trocar idioma",

	// Projects view
	ProjectsBackToTasks: "[P] Voltar para Tarefas",
	ProjectsTaskCount:   "%d tarefas",
	ProjectCompleted:    "Concluido",

	// Language selection
	LanguageSelect:  "Selecione o idioma:",
	LanguagePortBR:  "Portugues (Brasil)",
	LanguageEnglish: "English",
	LanguageHint:    "j/k: navegar | Enter: selecionar | Esc: cancelar",
	LanguageChanged: "Idioma alterado",

	// Error messages
	ErrorInitStorage: "Erro ao inicializar armazenamento: %v\n",
	ErrorRunApp:      "Erro ao executar aplicacao: %v\n",

	// Loading
	Loading: "Carregando...",

	// Delete count format
	DeleteCountFormat: "%d tarefa(s) concluida(s)",
}

var en = &Messages{
	// Tab names
	TabToday:     "Today",
	TabWeek:      "This Week",
	TabNotUrgent: "Not Urgent",
	TabGeneral:   "General List",

	// Category names
	CategoryToday:     "Today",
	CategoryWeek:      "This Week",
	CategoryNotUrgent: "Not Urgent",
	CategoryGeneral:   "General List",

	// Status messages
	StatusTaskCompleted:    "Task completed",
	StatusAllTodayDone:     "Congrats! All today's tasks completed!",
	StatusTaskReopened:     "Task reopened",
	StatusTaskMoved:        "Task moved to %s",
	StatusTaskCreated:      "Task created",
	StatusTaskUpdated:      "Task updated",
	StatusProjectsAssoc:    "Projects associated",
	StatusProjectCompleted: "Project completed",
	StatusProjectReopened:  "Project reopened",
	StatusProjectCreated:   "Project created",
	StatusProjectUpdated:   "Project updated",
	StatusTaskDeleted:      "Task deleted",
	StatusProjectDeleted:   "Project deleted",
	StatusCompletedDeleted: "Completed tasks deleted",

	// Placeholders
	PlaceholderName: "Name...",
	PlaceholderDesc: "Description (supports Markdown)...",

	// Empty states
	EmptyTaskList:    "No tasks in this list.\n\nPress 'a' to create a new task.",
	EmptyTaskDetail:  "Select a task to see details",
	EmptyProjectList: "No projects registered.\n\nPress 'a' to create a new project.",

	// Form labels
	LabelStatus:      "Status: ",
	LabelCompleted:   "Completed",
	LabelPending:     "Pending",
	LabelDescription: "Description:",
	LabelNoDesc:      "(no description)",
	LabelProjects:    "Projects:",
	LabelNoProjects:  "(no projects)",

	// Modal titles
	ModalNewTask:       "New Task",
	ModalEditTask:      "Edit Task",
	ModalNewProject:    "New Project",
	ModalEditProject:   "Edit Project",
	ModalAssocProjects: "Associate Projects",
	ModalHelp:          "Help - Shortcuts",
	ModalConfirmDelete: "Confirm Deletion",
	ModalLanguage:      "Select Language",

	// Delete confirmation
	ConfirmDeleteTask:      "Do you really want to delete this task?",
	ConfirmDeleteProject:   "Do you really want to delete this project?",
	ConfirmDeleteCompleted: "Do you really want to delete all completed tasks?",
	ConfirmYes:             "confirm",
	ConfirmNo:              "cancel",

	// Form fields
	FormName:         "Name:",
	FormDescMarkdown: "Description (Markdown):",
	FormProjectName:  "Project Name:",
	FormNoProjects:   "(no projects registered)",

	// Form hints
	HintNavProjects:    "j/k: navigate | Space: select | Tab: next | Ctrl+S: save",
	HintFormFields:     "Tab: switch fields | Ctrl+S: save | Esc: cancel",
	HintProjectForm:    "Enter: confirm | Esc: cancel",
	HintAssocProjects:  "Space: select | Enter: confirm | Esc: cancel",
	HintCloseHelp:      "Press any key to close",
	HintNoProjectAvail: "No projects available.\nCreate a project first (P).",

	// Status bar help
	HelpNew:      "new",
	HelpComplete: "complete",
	HelpEdit:     "edit",
	HelpDetails:  "details",
	HelpHelp:     "help",
	HelpQuit:     "quit",
	HelpScroll:   "scroll",
	HelpBack:     "back",
	HelpDelete:   "delete",

	// Keybinding help text
	KeyUp:            "up",
	KeyDown:          "down",
	KeyLeft:          "left",
	KeyRight:         "right",
	KeyNextTab:       "next tab",
	KeyPrevTab:       "prev tab",
	KeyProjects:      "projects",
	KeyNewTask:       "new task",
	KeyNewTaskGen:    "new in general list",
	KeyEditTask:      "edit",
	KeyDeleteTask:    "delete",
	KeyCompleteTask:  "complete",
	KeyDeleteDone:    "delete completed",
	KeyAssocProjects: "associate projects",
	KeyMoveToday:     "move to Today",
	KeyMoveWeek:      "move to Week",
	KeyMoveNotUrgent: "move to Not Urgent",
	KeyMoveGeneral:   "move to General List",
	KeyNewProject:    "new project",
	KeyEditProject:   "edit project",
	KeyCompleteProj:  "complete project",
	KeyHelp:          "help",
	KeyQuit:          "quit",
	KeyEnter:         "confirm",
	KeyEscape:        "cancel",
	KeySaveForm:      "save",
	KeyLanguage:      "language",

	// Help modal keys (left column)
	HelpKeyNavList:      "j/k or arrows",
	HelpKeyNavTabs:      "Tab/Shift+Tab",
	HelpKeyComplete:     "space or x",
	HelpKeyCompleteProj: "space or X",
	HelpKeyQuit:         "q or Ctrl+C",

	// Help modal sections
	HelpNavSection:      "Navigation",
	HelpNavList:         "Move in list",
	HelpNavTabs:         "Switch tabs",
	HelpNavProjects:     "Projects screen",
	HelpTaskSection:     "Tasks",
	HelpTaskNew:         "New task in current tab",
	HelpTaskNewGen:      "New task in General List",
	HelpTaskEdit:        "Edit task",
	HelpTaskDelete:      "Delete task",
	HelpTaskComplete:    "Complete/reopen task",
	HelpTaskDeleteDone:  "Delete completed",
	HelpTaskAssoc:       "Associate projects to task",
	HelpMoveSection:     "Move task",
	HelpMoveToday:       "Move to Today",
	HelpMoveWeek:        "Move to This Week",
	HelpMoveNotUrgent:   "Move to Not Urgent",
	HelpMoveGeneral:     "Move to General List",
	HelpProjSection:     "Projects (P screen)",
	HelpProjNew:         "New project",
	HelpProjEdit:        "Edit project",
	HelpProjDelete:      "Delete project",
	HelpProjComplete:    "Complete project",
	HelpFormSection:     "Forms",
	HelpFormTab:         "Next field",
	HelpFormSave:        "Save",
	HelpFormCancel:      "Cancel",
	HelpGeneralSection:  "General",
	HelpGeneralHelp:     "Show/close help",
	HelpGeneralQuit:     "Quit",
	HelpGeneralLanguage: "Change language",

	// Projects view
	ProjectsBackToTasks: "[P] Back to Tasks",
	ProjectsTaskCount:   "%d tasks",
	ProjectCompleted:    "Completed",

	// Language selection
	LanguageSelect:  "Select language:",
	LanguagePortBR:  "Portugues (Brasil)",
	LanguageEnglish: "English",
	LanguageHint:    "j/k: navigate | Enter: select | Esc: cancel",
	LanguageChanged: "Language changed",

	// Error messages
	ErrorInitStorage: "Error initializing storage: %v\n",
	ErrorRunApp:      "Error running application: %v\n",

	// Loading
	Loading: "Loading...",

	// Delete count format
	DeleteCountFormat: "%d completed task(s)",
}

func init() {
	messages = ptBR
	loadSavedLanguage()
}

func getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".t7t", "language.json")
}

func loadSavedLanguage() {
	configPath := getConfigPath()
	if configPath == "" {
		return
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	var config struct {
		Language Language `json:"language"`
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return
	}

	SetLanguage(config.Language)
}

func saveLanguage(lang Language) {
	configPath := getConfigPath()
	if configPath == "" {
		return
	}

	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return
	}

	config := struct {
		Language Language `json:"language"`
	}{
		Language: lang,
	}

	data, err := json.Marshal(config)
	if err != nil {
		return
	}

	os.WriteFile(configPath, data, 0644)
}

func SetLanguage(lang Language) {
	mu.Lock()
	defer mu.Unlock()

	currentLang = lang
	switch lang {
	case English:
		messages = en
	default:
		messages = ptBR
	}
	saveLanguage(lang)
}

func GetLanguage() Language {
	mu.RLock()
	defer mu.RUnlock()
	return currentLang
}

func Get() *Messages {
	mu.RLock()
	defer mu.RUnlock()
	return messages
}

func AvailableLanguages() []Language {
	return []Language{Portuguese, English}
}

func LanguageDisplayName(lang Language) string {
	switch lang {
	case English:
		return "English"
	case Portuguese:
		return "Portugues (Brasil)"
	default:
		return string(lang)
	}
}
