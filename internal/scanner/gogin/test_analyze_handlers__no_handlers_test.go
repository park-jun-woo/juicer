//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_NoHandlers 테스트
package gogin

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestAnalyzeHandlers_NoHandlers(t *testing.T) {
	eps := []scanner.Endpoint{{Method: "GET", Path: "/test"}}
	analyzeHandlers(nil, eps, ".", nil)
}
