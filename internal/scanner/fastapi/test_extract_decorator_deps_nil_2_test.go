//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDecoratorDeps_Nil 테스트
package fastapi

import "testing"

func TestExtractDecoratorDeps_Nil(t *testing.T) {
	if deps := extractDecoratorDeps(nil, nil); deps != nil {
		t.Fatalf("expected nil, got %v", deps)
	}
}
