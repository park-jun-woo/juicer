//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_NoHandlerExprs 테스트
package gogin

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestAnalyzeHandlers_NoHandlerExprs(t *testing.T) {
	endpoints := []scanner.Endpoint{
		{Method: "GET", Path: "/test"},
	}
	analyzeHandlers(nil, endpoints, ".", nil)
	// Should not crash, handlerExprs is empty so loop skipped
}
