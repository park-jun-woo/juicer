//ff:func feature=scan type=test topic=express control=sequence
//ff:what isCollectableSource .d.ts 제외 및 소스 확장자 수집 테스트
package express

import "testing"

func TestIsCollectableSource(t *testing.T) {
	if isCollectableSource("types.d.ts") {
		t.Error(".d.ts must be excluded")
	}
	if !isCollectableSource("routes.ts") {
		t.Error(".ts should be collectable")
	}
	if isCollectableSource("readme.md") {
		t.Error(".md not collectable")
	}
}
