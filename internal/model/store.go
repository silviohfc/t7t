package model

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type Store struct {
	Tasks    []*Task    `json:"tasks"`
	Projects []*Project `json:"projects"`
	path     string
	mu       sync.RWMutex
}

func NewStore() (*Store, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dataDir := filepath.Join(homeDir, ".t7t")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	store := &Store{
		Tasks:    []*Task{},
		Projects: []*Project{},
		path:     filepath.Join(dataDir, "data.json"),
	}

	if err := store.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return store, nil
}

func (s *Store) Load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, s)
}

func (s *Store) Save() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

// Task operations

func (s *Store) AddTask(task *Task) error {
	s.Tasks = append(s.Tasks, task)
	return s.Save()
}

func (s *Store) UpdateTask(task *Task) error {
	return s.Save()
}

func (s *Store) DeleteTask(id string) error {
	for i, t := range s.Tasks {
		if t.ID == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			break
		}
	}
	return s.Save()
}

func (s *Store) DeleteCompletedTasks(category Category) error {
	var remaining []*Task
	for _, t := range s.Tasks {
		if !(t.Completed && t.Category == category) {
			remaining = append(remaining, t)
		}
	}
	s.Tasks = remaining
	return s.Save()
}

func (s *Store) GetTasksByCategory(category Category) []*Task {
	var tasks []*Task
	for _, t := range s.Tasks {
		if t.Category == category {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func (s *Store) GetTask(id string) *Task {
	for _, t := range s.Tasks {
		if t.ID == id {
			return t
		}
	}
	return nil
}

// Project operations

func (s *Store) AddProject(project *Project) error {
	s.Projects = append(s.Projects, project)
	return s.Save()
}

func (s *Store) UpdateProject(project *Project) error {
	return s.Save()
}

func (s *Store) DeleteProject(id string) error {
	for i, p := range s.Projects {
		if p.ID == id {
			s.Projects = append(s.Projects[:i], s.Projects[i+1:]...)
			break
		}
	}
	// Remove project from all tasks
	for _, t := range s.Tasks {
		var newProjectIDs []string
		for _, pid := range t.ProjectIDs {
			if pid != id {
				newProjectIDs = append(newProjectIDs, pid)
			}
		}
		t.ProjectIDs = newProjectIDs
	}
	return s.Save()
}

func (s *Store) GetProject(id string) *Project {
	for _, p := range s.Projects {
		if p.ID == id {
			return p
		}
	}
	return nil
}

func (s *Store) GetProjects() []*Project {
	return s.Projects
}

func (s *Store) GetProjectNames(ids []string) []string {
	var names []string
	for _, id := range ids {
		if p := s.GetProject(id); p != nil {
			names = append(names, p.Name)
		}
	}
	return names
}
