//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectDirTSFiles_Error 테스트
package fastify

import "testing"

func TestCollectDirTSFiles_Error(t *testing.T) {

	if files := collectDirTSFiles("/no/such/dir/xyz123"); files != nil {
		t.Fatalf("expected nil on error, got %v", files)
	}
}
