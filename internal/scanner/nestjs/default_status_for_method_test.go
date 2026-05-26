//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDefaultStatusForMethod_Post 테스트
package nestjs

import "testing"

func TestDefaultStatusForMethod_Post(t *testing.T) {
	if got := defaultStatusForMethod("POST"); got != "201" {
		t.Fatalf("expected 201, got %s", got)
	}
}
