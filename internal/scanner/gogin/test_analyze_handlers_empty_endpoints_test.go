//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_EmptyEndpoints 테스트
package gogin

import (
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAnalyzeHandlers_EmptyEndpoints(t *testing.T) {
	analyzeHandlers(nil, nil, ".", nil, nil)
	analyzeHandlers(nil, []scanner.Endpoint{}, ".", nil, nil)
}
