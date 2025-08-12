package commands

import (
	"remora/pkg/resp"
	"remora/pkg/store"
	"strings"
	"sync"
)

type HandlerFunc func(*store.Store, []resp.Value) resp.Value

var (
	Registry   = make(map[string]HandlerFunc) 
	RegistryMu sync.RWMutex
)

func Register(command string, handler HandlerFunc) {
	RegistryMu.Lock()
	defer RegistryMu.Unlock()

	cmd := strings.ToUpper(command)
	if _, exists := Registry[cmd]; exists {
		panic("command already registered: " + cmd)
	}
	Registry[cmd] = handler

}
func GetHandler(command string) (HandlerFunc, bool) {
	RegistryMu.RLock()
	defer RegistryMu.RUnlock()

	handler, ok := Registry[strings.ToUpper(command)]

	return handler, ok

}
