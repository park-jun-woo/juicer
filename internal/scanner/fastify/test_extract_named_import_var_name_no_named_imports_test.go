//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractNamedImportVarName_NoNamedImports 테스트
package fastify

import "testing"

func TestExtractNamedImportVarName_NoNamedImports(t *testing.T) {

	c, src := importClause(t, `import Fastify from "fastify";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "" {
		t.Fatalf("expected empty for default import, got %q", got)
	}
}
