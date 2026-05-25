//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_Empty 테스트
package scanner

import "testing"

func TestAnalyzeHandlers_Empty(t *testing.T) {
	analyzeHandlers(nil, nil, ".")
	analyzeHandlers(nil, []Endpoint{}, ".")
}
