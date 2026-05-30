//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestApplyRescanResults_NoMatch_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyRescanResults_NoMatch_Round5(t *testing.T) {
	ctx := emptyGroupCtx()
	ctx.epIndex = map[struct {
		file string
		line int
	}]int{}

	applyRescanResults([]scanner.Endpoint{{File: "a.go", Line: 1}}, ctx)
}
