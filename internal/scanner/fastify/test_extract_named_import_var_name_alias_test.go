//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractNamedImportVarName_Alias 테스트
package fastify

import "testing"

func TestExtractNamedImportVarName_Alias(t *testing.T) {
	c, src := importClause(t, `import { join as j } from "path";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "j" {
		t.Fatalf("got %q, want j (alias)", got)
	}
}
