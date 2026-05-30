//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractOneRoute_NotRoute 테스트
package quarkus

import "testing"

func TestExtractOneRoute_NotRoute(t *testing.T) {
	fi := qFileInfo(t, `class R { public void helper() {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	if _, ok := extractOneRoute(m, fi); ok {
		t.Fatal("expected false for non-route method")
	}
}
