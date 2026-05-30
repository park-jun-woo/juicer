//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestApplyOmitType 테스트
package nestjs

import "testing"

func TestApplyOmitType(t *testing.T) {
	got := applyOmitType(sampleFields(), []string{"age"})
	if len(got) != 2 {
		t.Fatalf("expected 2, got %d", len(got))
	}
	for _, f := range got {
		if f.name == "age" {
			t.Fatal("age should be omitted")
		}
	}
}
