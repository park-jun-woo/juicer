//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterAliasesFromStmt_Alias 테스트
package express

import "testing"

func TestCollectRouterAliasesFromStmt_Alias(t *testing.T) {
	fi := mustParse(t, []byte(`import { Router as R } from 'express';`))
	aliases := map[string]bool{}
	collectRouterAliasesFromStmt(firstImportStmt(t, fi), fi.Src, aliases)
	if !aliases["R"] {
		t.Fatalf("expected alias R, got %v", aliases)
	}
}
