//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectAutoloadRequireNames 테스트
package fastify

import "testing"

func TestCollectAutoloadRequireNames(t *testing.T) {
	src := `
const autoload = require("@fastify/autoload");
const x = require("other");
const y = 5;
`
	fi := mustParse(t, []byte(src))
	names := map[string]bool{}
	collectAutoloadRequireNames(fi, names)
	if !names["autoload"] {
		t.Fatalf("expected 'autoload' collected, got %v", names)
	}
	if names["x"] || names["y"] {
		t.Fatalf("non-autoload requires should not be collected, got %v", names)
	}
}
