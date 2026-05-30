//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractZodValidatorsFromArgs: start초과 / call수집 / 비call스킵
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func argNodesOf(t *testing.T, fi *fileInfo) []*sitter.Node {
	t.Helper()
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	return collectArgNodes(args)
}

func TestExtractZodValidatorsFromArgs_Collect(t *testing.T) {
	// r.post('/x', validateRequest({body: s}), handler)
	fi := mustParse(t, []byte(`r.post('/x', validateRequest({ body: s }), handler);`))
	nodes := argNodesOf(t, fi)
	got := extractZodValidatorsFromArgs(nodes, fi.Src, 1)
	if len(got) != 1 || got[0].Target != "json" {
		t.Fatalf("got %+v", got)
	}
}

func TestExtractZodValidatorsFromArgs_StartBeyond(t *testing.T) {
	fi := mustParse(t, []byte(`r.post('/x');`))
	nodes := argNodesOf(t, fi)
	if got := extractZodValidatorsFromArgs(nodes, fi.Src, 1); got != nil {
		t.Fatalf("got %+v", got)
	}
}

func TestExtractZodValidatorsFromArgs_NonCallSkipped(t *testing.T) {
	fi := mustParse(t, []byte(`r.post('/x', handler);`))
	nodes := argNodesOf(t, fi)
	if got := extractZodValidatorsFromArgs(nodes, fi.Src, 1); got != nil {
		t.Fatalf("got %+v", got)
	}
}
