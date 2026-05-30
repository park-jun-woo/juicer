//ff:func feature=scan type=test control=sequence topic=django
//ff:what findFuncView — 이름으로 함수 뷰 검색 분기를 검증
package django

import "testing"

func TestFindFuncView(t *testing.T) {
	views := []funcViewInfo{{name: "health"}, {name: "ping"}}
	if fv := findFuncView(views, "ping"); fv == nil || fv.name != "ping" {
		t.Fatalf("expected to find ping, got %v", fv)
	}
	if fv := findFuncView(views, "missing"); fv != nil {
		t.Fatalf("expected nil, got %v", fv)
	}
}
