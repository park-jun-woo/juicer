//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what autoloadSegments: index 파일/비-index 파일/중첩 디렉터리 세그먼트 변환을 검증
package fastify

import "testing"

func TestAutoloadSegments(t *testing.T) {
	if got := autoloadSegments("auth/index.ts"); len(got) != 1 || got[0] != "auth" {
		t.Fatalf("index in subdir: want [auth], got %v", got)
	}
	if got := autoloadSegments("index.ts"); len(got) != 0 {
		t.Fatalf("root index: want [], got %v", got)
	}
	got := autoloadSegments("api/users/profile.ts")
	if len(got) != 3 || got[0] != "api" || got[1] != "users" || got[2] != "profile" {
		t.Fatalf("want [api users profile], got %v", got)
	}
}
