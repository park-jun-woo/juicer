//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectControllers_BadFile 테스트
package nestjs

import "testing"

func TestCollectControllers_BadFile(t *testing.T) {
	result := collectControllers([]string{"/nonexistent/file.ts"}, "/tmp")
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
