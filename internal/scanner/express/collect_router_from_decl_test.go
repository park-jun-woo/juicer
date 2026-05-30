//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectRouterFromDecl: express()/express.Router()/Router별칭 + 미해당 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstLexDecl(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	decls := findAllByType(fi.Root, "lexical_declaration")
	if len(decls) == 0 {
		t.Fatal("no lexical_declaration")
	}
	return decls[0]
}

func TestCollectRouterFromDecl_ExpressRouterCall(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	routers := map[string]bool{}
	collectRouterFromDecl(firstLexDecl(t, fi), fi, routers, map[string]bool{})
	if !routers["r"] {
		t.Fatalf("expected r, got %v", routers)
	}
}

func TestCollectRouterFromDecl_AliasCall(t *testing.T) {
	fi := mustParse(t, []byte(`const r = Router();`))
	routers := map[string]bool{}
	collectRouterFromDecl(firstLexDecl(t, fi), fi, routers, map[string]bool{"Router": true})
	if !routers["r"] {
		t.Fatalf("expected r via alias, got %v", routers)
	}
}

func TestCollectRouterFromDecl_NoInitValue(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 5;`))
	routers := map[string]bool{}
	collectRouterFromDecl(firstLexDecl(t, fi), fi, routers, map[string]bool{})
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}

func TestCollectRouterFromDecl_NotRouterCall(t *testing.T) {
	fi := mustParse(t, []byte(`const x = foo();`))
	routers := map[string]bool{}
	collectRouterFromDecl(firstLexDecl(t, fi), fi, routers, map[string]bool{})
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}
