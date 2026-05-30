//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterAliasesFromStmt_NoClause 테스트
package express

import "testing"

func TestCollectRouterAliasesFromStmt_NoClause(t *testing.T) {

	fi := mustParse(t, []byte(`import 'express';`))
	aliases := map[string]bool{}
	collectRouterAliasesFromStmt(firstImportStmt(t, fi), fi.Src, aliases)
	if len(aliases) != 0 {
		t.Fatalf("expected none, got %v", aliases)
	}
}
