//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractNamedImportVarName_MultipleSpecsFirst 테스트
package fastify

import "testing"

func TestExtractNamedImportVarName_MultipleSpecsFirst(t *testing.T) {

	c, src := importClause(t, `import { readFile, writeFile } from "fs";`+"\n")
	if got := extractNamedImportVarName(c, src); got != "readFile" {
		t.Fatalf("expected readFile (first), got %q", got)
	}
}
