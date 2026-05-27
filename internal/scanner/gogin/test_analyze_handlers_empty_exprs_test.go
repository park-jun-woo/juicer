//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeHandlers_EmptyExprs 테스트
package gogin

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAnalyzeHandlers_EmptyExprs(t *testing.T) {
	eps := []scanner.Endpoint{{Method: "GET", Path: "/test"}}
	analyzeHandlers(nil, eps, ".", nil, nil)
}
