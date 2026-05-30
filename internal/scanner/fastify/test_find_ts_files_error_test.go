//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestFindTSFiles_Error 테스트
package fastify

import "testing"

func TestFindTSFiles_Error(t *testing.T) {
	_, err := findTSFiles("/no/such/path/zzz")
	if err == nil {
		t.Fatal("expected error for missing root")
	}
}
