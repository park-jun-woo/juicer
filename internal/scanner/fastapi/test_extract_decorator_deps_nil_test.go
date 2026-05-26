//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDecoratorDeps_NilCallNode nil 입력 테스트
package fastapi

import "testing"

func TestExtractDecoratorDeps_NilCallNode(t *testing.T) {
	deps := extractDecoratorDeps(nil, nil)
	if len(deps) != 0 {
		t.Fatalf("expected 0 deps for nil, got %d", len(deps))
	}
}
