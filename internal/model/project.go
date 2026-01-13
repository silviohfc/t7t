package model

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProject(name string) *Project {
	now := time.Now()
	return &Project{
		ID:        uuid.New().String(),
		Name:      name,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (p *Project) ToggleComplete() {
	p.Completed = !p.Completed
	p.UpdatedAt = time.Now()
}

func (p *Project) Update(name string) {
	p.Name = name
	p.UpdatedAt = time.Now()
}

func (p Project) FilterValue() string {
	return p.Name
}

func (p Project) Title() string {
	return p.Name
}

func (p Project) Description() string {
	if p.Completed {
		return "Concluido"
	}
	return ""
}
