//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what extractBodyKwargs Body(..., embed=True, alias="article") 추출 테스트
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

func TestExtractBodyKwargs_NoBodyCall(t *testing.T) {
	src := []byte("def f(x: int): pass\n")
	root, _ := parsePython(src)
	fn := findChildByType(root, "function_definition")
	params := findChildByType(fn, "parameters")
	tp := findChildByType(params, "typed_parameter")
	if tp == nil {
		t.Skip("no typed_parameter")
	}
	alias, embed := extractBodyKwargs(tp, src)
	if alias != "" || embed {
		t.Fatalf("alias=%q embed=%v", alias, embed)
	}
}

func TestExtractBodyKwargs_EmbedFalse(t *testing.T) {
	src := []byte(`def f(item: Item = Body(embed=False)): pass` + "\n")
	root, _ := parsePython(src)
	fn := findChildByType(root, "function_definition")
	params := findChildByType(fn, "parameters")
	for i := 0; i < int(params.ChildCount()); i++ {
		c := params.Child(i)
		if isParamNode(c) {
			_, embed := extractBodyKwargs(c, src)
			if embed {
				t.Fatalf("expected embed=false")
			}
			return
		}
	}
}
