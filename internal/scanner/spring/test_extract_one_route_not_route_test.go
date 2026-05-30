//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractOneRoute_NotRoute 테스트
package spring

import "testing"

func TestExtractOneRoute_NotRoute(t *testing.T) {
	fi := sFileInfo(t, `class C { public void helper() {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	if _, ok := extractOneRoute(m, fi); ok {
		t.Fatal("expected false")
	}
}
