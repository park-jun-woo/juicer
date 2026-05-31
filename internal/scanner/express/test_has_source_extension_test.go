//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what hasSourceExtension 알려진 소스 확장자 매칭 테스트
package express

import "testing"

func TestHasSourceExtension(t *testing.T) {
	for _, name := range []string{"a.ts", "a.tsx", "a.js", "a.jsx", "a.mjs", "a.cjs"} {
		if !hasSourceExtension(name) {
			t.Errorf("%q should be source", name)
		}
	}
	for _, name := range []string{"a.go", "a.txt", "a"} {
		if hasSourceExtension(name) {
			t.Errorf("%q should not be source", name)
		}
	}
}
