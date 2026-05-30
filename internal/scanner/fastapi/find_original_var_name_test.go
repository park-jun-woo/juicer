//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findOriginalVarName: prefix에 존재 / 미존재 분기
package fastapi

import "testing"

func TestFindOriginalVarName(t *testing.T) {
	fi := &fileInfo{prefixes: map[string]string{"router": "/api"}}
	if got := findOriginalVarName("router", fi); got != "router" {
		t.Fatalf("got %q", got)
	}
	if got := findOriginalVarName("missing", fi); got != "" {
		t.Fatalf("got %q", got)
	}
}
