//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeHandlers_EmptyExprs 테스트
package scanner

import "testing"

func TestAnalyzeHandlers_EmptyExprs(t *testing.T) {
	eps := []Endpoint{{Method: "GET", Path: "/test"}}
	analyzeHandlers(nil, eps, ".")
}
