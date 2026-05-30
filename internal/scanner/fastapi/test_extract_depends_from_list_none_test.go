//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDependsFromList_None 테스트
package fastapi

import "testing"

func TestExtractDependsFromList_None(t *testing.T) {
	list, src := firstList(t, []byte("x = [a, b, c]\n"))
	deps := extractDependsFromList(list, src)
	if len(deps) != 0 {
		t.Fatalf("expected none, got %v", deps)
	}
}
