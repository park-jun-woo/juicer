//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_EmptyEndpoints 테스트
package scanner

import (
	"testing"
)

func TestAnalyzeHandlers_EmptyEndpoints(t *testing.T) {
	analyzeHandlers(nil, nil, ".", nil)
	analyzeHandlers(nil, []Endpoint{}, ".", nil)
}
