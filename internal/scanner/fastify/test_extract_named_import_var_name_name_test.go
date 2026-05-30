//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractNamedImportVarName_Name 테스트
package fastify

import "testing"

func TestExtractNamedImportVarName_Name(t *testing.T) {
	c, src := importClause(t, `import { join } from "path";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "join" {
		t.Fatalf("got %q, want join", got)
	}
}
