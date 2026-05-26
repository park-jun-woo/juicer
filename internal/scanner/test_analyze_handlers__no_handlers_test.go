//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_NoHandlers 테스트
package scanner

import "testing"

func TestAnalyzeHandlers_NoHandlers(t *testing.T) {
	eps := []Endpoint{{Method: "GET", Path: "/test"}}
	analyzeHandlers(nil, eps, ".", nil)
}
