//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestApplyPartialType 테스트
package nestjs

import "testing"

func TestApplyPartialType(t *testing.T) {
	got := applyPartialType(sampleFields())
	if len(got) != 3 {
		t.Fatalf("got %d", len(got))
	}
	for _, f := range got {
		if !f.optional {
			t.Fatalf("field %s should be optional", f.name)
		}
	}
}
