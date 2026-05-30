//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestApplyLengthConstraint 테스트
package nestjs

import "testing"

func TestApplyLengthConstraint(t *testing.T) {
	var p *int
	applyLengthConstraint("5", &p)
	if p == nil || *p != 5 {
		t.Fatalf("got %v", p)
	}
	var q *int
	applyLengthConstraint("notanum", &q)
	if q != nil {
		t.Fatal("expected nil on parse error")
	}
}
