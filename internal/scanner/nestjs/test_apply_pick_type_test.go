//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestApplyPickType 테스트
package nestjs

import "testing"

func TestApplyPickType(t *testing.T) {
	got := applyPickType(sampleFields(), []string{"email"})
	if len(got) != 1 || got[0].name != "email" {
		t.Fatalf("got %+v", got)
	}
}
