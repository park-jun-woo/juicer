//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractNamedImportVarName_EmptyBraces 테스트
package fastify

import "testing"

func TestExtractNamedImportVarName_EmptyBraces(t *testing.T) {

	c, src := importClause(t, `import {} from "path";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "" {
		t.Fatalf("expected empty for empty braces, got %q", got)
	}
}
