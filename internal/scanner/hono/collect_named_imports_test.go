//ff:func feature=scan type=test control=sequence topic=hono
//ff:what collectNamedImports 테스트
package hono

import "testing"

func TestCollectNamedImports(t *testing.T) {
	fi := mustParse(t, []byte(`import { users, items } from "./routes";`+"\n"))
	named := findAllByType(fi.Root, "named_imports")
	if len(named) == 0 {
		t.Fatal("no named_imports")
	}
	imports := map[string]string{}
	collectNamedImports(named[0], fi.Src, "routes.ts", imports)
	if imports["users"] != "routes.ts" || imports["items"] != "routes.ts" {
		t.Fatalf("imports = %v", imports)
	}
}
