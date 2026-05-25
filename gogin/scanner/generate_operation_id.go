//ff:func feature=scan type=extract control=sequence
//ff:what 핸들러 이름에서 operationId를 생성한다
package scanner

import (
	"strings"
)

func generateOperationID(ep Endpoint) string {
	handler := ep.Handler
	// "file.go:h.ListBuildings" → "h.ListBuildings"
	if idx := strings.LastIndex(handler, ":"); idx >= 0 {
		handler = handler[idx+1:]
	}

	// "(inline)" → path+method 기반 생성
	if handler == "(inline)" || handler == "" {
		return pathMethodToOperationID(ep.Method, ep.Path)
	}

	// "h.ListBuildings" → "ListBuildings" → "listBuildings"
	if idx := strings.LastIndex(handler, "."); idx >= 0 {
		handler = handler[idx+1:]
	}

	// 후위 "()" 제거
	handler = strings.TrimSuffix(handler, "()")

	return lcFirst(handler)
}

