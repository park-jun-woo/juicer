//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestAnalyzeExpr_FuncLitNoCtx_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestAnalyzeExpr_FuncLitNoCtx_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}

	expr := parseExpr(t, `func(){ }`)
	analyzeExpr(ep, expr, nil, buildFuncIndex(nil))
}
