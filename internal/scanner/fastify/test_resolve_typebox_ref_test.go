//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what resolveTypeBoxRef identifier → vars 매핑 노드 해석 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestResolveTypeBoxRef(t *testing.T) {
	fi := mustParse(t, []byte(`UserSchema`))
	id := findAllByType(fi.Root, "identifier")[0]
	target := mustParse(t, []byte(`{ id: Type.Integer() }`)).Root
	vars := map[string]*sitter.Node{"UserSchema": target}
	if got := resolveTypeBoxRef(id, fi.Src, vars); got != target {
		t.Errorf("resolved: %v", got)
	}
	// not an identifier
	objNode := mustParse(t, []byte(`({})`)).Root
	if resolveTypeBoxRef(objNode, fi.Src, vars) != nil {
		t.Error("non-identifier should be nil")
	}
	// nil vars
	if resolveTypeBoxRef(id, fi.Src, nil) != nil {
		t.Error("nil vars should be nil")
	}
}
