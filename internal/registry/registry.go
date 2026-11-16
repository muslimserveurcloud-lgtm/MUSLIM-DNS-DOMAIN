package registry

import (
	"errors"
	"strings"
	"sync"
)

type Extension struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Domain struct {
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Owner     string `json:"owner"`
}

type Registry struct {
	mu         sync.RWMutex
	extensions map[string]Extension
	domains    map[string]Domain
}

func New() *Registry {
	return &Registry{
		extensions: make(map[string]Extension),
		domains:    make(map[string]Domain),
	}
}

func (r *Registry) AddExtension(name string) error {
	name = strings.TrimPrefix(strings.ToLower(strings.TrimSpace(name)), ".")
	if name == "" || strings.Contains(name, ".") {
		return errors.New("invalid extension")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.extensions[name] = Extension{
		Name:   name,
		Status: "development",
	}
	return nil
}

func (r *Registry) RegisterDomain(name, owner string) error {
	name = strings.ToLower(strings.TrimSpace(name))
	owner = strings.TrimSpace(owner)

	parts := strings.Split(name, ".")
	if len(parts) != 2 || parts[0] == "" || owner == "" {
		return errors.New("invalid domain or owner")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	ext := parts[1]
	if _, ok := r.extensions[ext]; !ok {
		return errors.New("extension not managed")
	}

	if _, exists := r.domains[name]; exists {
		return errors.New("domain already registered")
	}

	r.domains[name] = Domain{
		Name:      name,
		Extension: ext,
		Owner:     owner,
	}
	return nil
}

func (r *Registry) ListExtensions() []Extension {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]Extension, 0, len(r.extensions))
	for _, ext := range r.extensions {
		result = append(result, ext)
	}
	return result
}

func (r *Registry) GetDomain(name string) (Domain, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	d, ok := r.domains[strings.ToLower(strings.TrimSpace(name))]
	return d, ok
}
