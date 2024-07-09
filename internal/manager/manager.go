package manager

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type Manager struct {
	passwords map[string]string
	filename  string
	mu        sync.Mutex
}

func New(filename string) (*Manager, error) {
	m := &Manager{
		passwords: make(map[string]string),
		filename:  filename,
	}
	err := m.load()
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	return m, nil
}

func (m *Manager) List() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	var names []string
	for name := range m.passwords {
		names = append(names, name)
	}
	return names
}

func (m *Manager) Get(name string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	password, ok := m.passwords[name]
	if !ok {
		return "", errors.New("password not found")
	}
	return password, nil
}

func (m *Manager) Set(name, password string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.passwords[name] = password
	return m.save()
}

func (m *Manager) load() error {
	file, err := os.ReadFile(m.filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &m.passwords)
}

func (m *Manager) save() error {
	data, err := json.MarshalIndent(m.passwords, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(m.filename, data, 0600)
}