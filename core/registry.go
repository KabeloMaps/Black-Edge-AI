package core

import "sync"

type Source interface {
	Name() string
	Fetch() ([]byte, error) // raw payload
}

var (
	registry = make(map[string]Source)
	lock     sync.RWMutex
)

func RegisterSource(src Source) {
	lock.Lock()
	defer lock.Unlock()
	registry[src.Name()] = src
}

func GetSource(name string) (Source, bool) {
	lock.RLock()
	defer lock.RUnlock()
	src, ok := registry[name]
	return src, ok
}

func GetAllSources() []Source {
	lock.RLock()
	defer lock.RUnlock()
	list := []Source{}
	for _, src := range registry {
		list = append(list, src)
	}
	return list
}