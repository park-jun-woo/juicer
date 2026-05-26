//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDefaultStatusForMethod_Get 테스트
package nestjs

import "testing"

func TestDefaultStatusForMethod_Get(t *testing.T) {
	if got := defaultStatusForMethod("GET"); got != "200" {
		t.Fatalf("expected 200, got %s", got)
	}
}
