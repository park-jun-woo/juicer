//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectControllers_Empty 테스트
package nestjs

import "testing"

func TestCollectControllers_Empty(t *testing.T) {
	result := collectControllers(nil, "/tmp")
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
