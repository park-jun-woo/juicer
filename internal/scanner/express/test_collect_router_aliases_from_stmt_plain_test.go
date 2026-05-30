//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterAliasesFromStmt_Plain 테스트
package express

import "testing"

func TestCollectRouterAliasesFromStmt_Plain(t *testing.T) {
	fi := mustParse(t, []byte(`import { Router } from 'express';`))
	aliases := map[string]bool{}
	collectRouterAliasesFromStmt(firstImportStmt(t, fi), fi.Src, aliases)
	if !aliases["Router"] {
		t.Fatalf("expected Router, got %v", aliases)
	}
}
