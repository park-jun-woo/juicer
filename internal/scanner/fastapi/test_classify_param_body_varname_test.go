//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what classifyParam Body 파라미터 var_name, alias, embed 추출 테스트
package fastapi

import "testing"

func TestClassifyParam_BodyVarName(t *testing.T) {
	src := []byte(`def f(new_article: NewArticle = Body(..., embed=True, alias="article")): pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	pathNames := map[string]bool{}

	ri := &routeInfo{}
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if isParamNode(child) {
			classifyParam(child, src, ri, pathNames, nil)
		}
	}

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
}
