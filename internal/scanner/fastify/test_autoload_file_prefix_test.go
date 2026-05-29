//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what autoloadFilePrefix: 베이스 prefix + 디렉터리 세그먼트 합성과 비-ts 거부를 검증
package fastify

import "testing"

func TestAutoloadFilePrefix(t *testing.T) {
	got, ok := autoloadFilePrefix("/base", "/base/auth/index.ts", "/api")
	if !ok || got != "/api/auth" {
		t.Fatalf("want /api/auth ok, got %q ok=%v", got, ok)
	}
	if _, ok := autoloadFilePrefix("/base", "/base/auth/schema.d.ts", "/api"); ok {
		t.Error("expected .d.ts rejected")
	}
}
