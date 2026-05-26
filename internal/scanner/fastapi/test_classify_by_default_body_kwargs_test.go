//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what classifyByDefault Body 분기에서 var_name, alias, embed 추출 테스트
package fastapi

import "testing"

func TestClassifyByDefault_BodyKwargs(t *testing.T) {
	src := []byte(`def f(new_article: NewArticle = Body(..., embed=True, alias="article")): pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")

	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if !isParamNode(child) {
			continue
		}
		ri := &routeInfo{}
		classifyByDefault("Body", "new_article", "NewArticle", `Body(..., embed=True, alias="article")`, child, src, ri)
		if ri.bodyType != "NewArticle" {
			t.Fatalf("bodyType: got %q, want NewArticle", ri.bodyType)
		}
		if ri.bodyVarName != "new_article" {
			t.Fatalf("bodyVarName: got %q, want new_article", ri.bodyVarName)
		}
		if ri.bodyAlias != "article" {
			t.Fatalf("bodyAlias: got %q, want article", ri.bodyAlias)
		}
		if !ri.bodyEmbed {
			t.Fatal("bodyEmbed: got false, want true")
		}
		return
	}
	t.Fatal("no param node found")
}
