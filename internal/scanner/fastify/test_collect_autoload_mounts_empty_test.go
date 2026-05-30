//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectAutoloadMounts_Empty 테스트
package fastify

import "testing"

func TestCollectAutoloadMounts_Empty(t *testing.T) {

	if m := collectAutoloadMounts(map[string]*fileInfo{}, "/root"); len(m) != 0 {
		t.Fatalf("expected no mounts, got %d", len(m))
	}
}
