//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what enumPathCtx.resolve 멤버식 해석/비해석 폴백 테스트
package nestjs

import "testing"

func TestEnumPathCtxResolve(t *testing.T) {
	src := []byte(`enum RouteKey { Asset = 'assets' }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	pc := enumPathCtx{root: root, src: src, absFile: "f.ts"}
	if got := pc.resolve("RouteKey.Asset"); got != "assets" {
		t.Errorf("resolved: %q", got)
	}
	// unresolvable -> unchanged
	if got := pc.resolve("/literal"); got != "/literal" {
		t.Errorf("fallback: %q", got)
	}
}
