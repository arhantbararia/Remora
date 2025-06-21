package commands

import (
	"remora/pkg/resp"
	"strings"
	"sync"
)

type HandlerFunc func([]resp.Value) resp.Value



var (
	registry = make(map[string]HandlerFunc)
	registryMu sync.RWMutex

)


func Register(command string, handler HandlerFunc) {
	registryMu.Lock()
	defer registryMu.Unlock()

	cmd := strings.ToUpper(command)
	if _, exists := registry[cmd]; exists {
		panic("command already registered: " + cmd)
	}
	registry[cmd] = handler
}



func GetHandler(command string) (HandlerFunc, bool ) {
	registryMu.RLock()
	defer registryMu.RUnlock()

	handler, ok := registry[strings.ToUpper(command)]
	return handler, ok

}

