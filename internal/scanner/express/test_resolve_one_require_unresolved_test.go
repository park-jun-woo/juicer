//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestResolveOneRequire_Unresolved 테스트
package express

import "testing"

func TestResolveOneRequire_Unresolved(t *testing.T) {
	fi := mustParse(t, []byte(`const r = require('external-pkg');`))
	imports := map[string]string{}
	resolveOneRequire(firstLexDecl(t, fi), fi.Src, t.TempDir(), imports, t.TempDir(), nil)
	if len(imports) != 0 {
		t.Fatalf("got %v", imports)
	}
}
