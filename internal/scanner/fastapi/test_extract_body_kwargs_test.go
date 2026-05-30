//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestExtractBodyKwargs 테스트
package fastapi

import "testing"

func TestExtractBodyKwargs(t *testing.T) {
	src := []byte(`def f(new_article: NewArticle = Body(..., embed=True, alias="article")): pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")

	var found bool
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if !isParamNode(child) {
			continue
		}
		alias, embed := extractBodyKwargs(child, src)
		if alias != "article" {
			t.Errorf("alias: got %q, want %q", alias, "article")
		}
		if !embed {
			t.Error("embed: got false, want true")
		}
		found = true
		break
	}
	if !found {
		t.Fatal("no param node found")
	}
}
