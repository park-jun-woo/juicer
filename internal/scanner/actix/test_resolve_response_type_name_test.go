//ff:func feature=scan type=test topic=actix control=sequence
//ff:what resolveResponseTypeName struct_expression/identifier 분기 타입명 해석 테스트
package actix

import "testing"

func TestResolveResponseTypeName(t *testing.T) {
	// One source so identifier nodeText and block lookup share ctx.src.
	src := []byte(`fn f() { let resp = UserResponse { id: 1 }; resp }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	block := findAllByType(root, "block")[0]
	ctx := &responseCtx{block: block, src: src}

	// struct_expression branch
	se := findAllByType(root, "struct_expression")[0]
	if got := resolveResponseTypeName(se, ctx); got != "UserResponse" {
		t.Errorf("struct expr: got %q", got)
	}

	// identifier branch: the trailing `resp` identifier resolves via let binding.
	idents := findAllByType(root, "identifier")
	var respIdent = idents[len(idents)-1] // trailing expression-statement identifier
	if got := resolveResponseTypeName(respIdent, ctx); got != "UserResponse" {
		t.Errorf("identifier branch: got %q", got)
	}

	// neither struct nor identifier -> ""
	intLit := findAllByType(root, "integer_literal")[0]
	if got := resolveResponseTypeName(intLit, ctx); got != "" {
		t.Errorf("literal: got %q", got)
	}
}
