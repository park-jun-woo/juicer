//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestPrefixKey 테스트
package hono

import "testing"

func TestPrefixKey(t *testing.T) {
	if got := prefixKey("file.ts", "app"); got != "file.ts\x00app" {
		t.Fatalf("got %q", got)
	}
}
