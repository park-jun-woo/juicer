//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what resolveEndpointEnumPaths 엔드포인트 path/paths 내 enum 멤버식 인플레이스 해석 테스트
package nestjs

import "testing"

func TestResolveEndpointEnumPaths(t *testing.T) {
	src := []byte(`enum RouteKey { Asset = 'assets', Img = 'images' }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	pc := enumPathCtx{root: root, src: src, absFile: "f.ts"}
	eps := []endpointInfo{{
		path:  "RouteKey.Asset",
		paths: []string{"RouteKey.Img", "/literal"},
	}}
	resolveEndpointEnumPaths(eps, pc)
	if eps[0].path != "assets" {
		t.Errorf("path: %q", eps[0].path)
	}
	if eps[0].paths[0] != "images" || eps[0].paths[1] != "/literal" {
		t.Errorf("paths: %v", eps[0].paths)
	}
}
