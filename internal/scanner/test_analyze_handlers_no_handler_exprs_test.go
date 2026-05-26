//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_NoHandlerExprs 테스트
package scanner

import (
	"testing"
)

func TestAnalyzeHandlers_NoHandlerExprs(t *testing.T) {
	endpoints := []Endpoint{
		{Method: "GET", Path: "/test"},
	}
	analyzeHandlers(nil, endpoints, ".", nil)
	// Should not crash, handlerExprs is empty so loop skipped
}
