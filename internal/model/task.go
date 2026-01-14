package model

import (
	"t7t/internal/i18n"
	"time"

	"github.com/google/uuid"
)

type Category string

const (
	CategoryToday     Category = "today"
	CategoryWeek      Category = "week"
	CategoryNotUrgent Category = "not_urgent"
	CategoryGeneral   Category = "general"
)

func (c Category) String() string {
	return CategoryString(c)
}

func CategoryString(c Category) string {
	m := i18n.Get()
	switch c {
	case CategoryToday:
		return m.CategoryToday
	case CategoryWeek:
		return m.CategoryWeek
	case CategoryNotUrgent:
		return m.CategoryNotUrgent
	case CategoryGeneral:
		return m.CategoryGeneral
	default:
		return string(c)
	}
}

type Task struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    Category  `json:"category"`
	Completed   bool      `json:"completed"`
	ProjectIDs  []string  `json:"project_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(name, description string, category Category) *Task {
	now := time.Now()
	return &Task{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Category:    category,
		Completed:   false,
		ProjectIDs:  []string{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Task) ToggleComplete() {
	t.Completed = !t.Completed
	t.UpdatedAt = time.Now()
}

func (t *Task) SetCategory(category Category) {
	t.Category = category
	t.UpdatedAt = time.Now()
}

func (t *Task) Update(name, description string) {
	t.Name = name
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *Task) SetProjects(projectIDs []string) {
	t.ProjectIDs = projectIDs
	t.UpdatedAt = time.Now()
}

func (t *Task) HasProject(projectID string) bool {
	for _, id := range t.ProjectIDs {
		if id == projectID {
			return true
		}
	}
	return false
}

func (t Task) FilterValue() string {
	return t.Name
}

func (t Task) Title() string {
	return t.Name
}

func (t Task) Description_() string {
	return t.Description
}
