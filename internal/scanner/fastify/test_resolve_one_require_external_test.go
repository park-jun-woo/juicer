//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveOneRequire_External 테스트
package fastify

import "testing"

func TestResolveOneRequire_External(t *testing.T) {
	d, fi := reqDeclarator(t, `const fastify = require("fastify");`+"\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, t.TempDir(), imports)
	if len(imports) != 0 {
		t.Fatalf("external require should not resolve, got %v", imports)
	}
}
