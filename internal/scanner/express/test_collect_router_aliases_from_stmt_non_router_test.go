//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterAliasesFromStmt_NonRouter 테스트
package express

import "testing"

func TestCollectRouterAliasesFromStmt_NonRouter(t *testing.T) {
	fi := mustParse(t, []byte(`import { Foo } from 'express';`))
	aliases := map[string]bool{}
	collectRouterAliasesFromStmt(firstImportStmt(t, fi), fi.Src, aliases)
	if len(aliases) != 0 {
		t.Fatalf("expected none, got %v", aliases)
	}
}
