//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractValidateRequest: 정상(ident/inline) + 각 nil/skip 분기
package express

import "testing"

func TestExtractValidateRequest_IdentAndInline(t *testing.T) {
	fi := mustParse(t, []byte(`validateRequest({ body: createSchema, query: z.object({}), unknown: x });`))
	got := extractValidateRequest(firstCallExpr(t, fi), fi.Src)
	if len(got) != 2 {
		t.Fatalf("expected 2 validators (unknown key skipped), got %+v", got)
	}
	// body -> json with SchemaName
	if got[0].Target != "json" || got[0].SchemaName != "createSchema" {
		t.Fatalf("unexpected body validator %+v", got[0])
	}
	// query -> inline schema node
	if got[1].Target != "query" || got[1].SchemaNode == nil {
		t.Fatalf("unexpected query validator %+v", got[1])
	}
}

func TestExtractValidateRequest_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`))
	if got := extractValidateRequest(fi.Root, fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

func TestExtractValidateRequest_NoFnIdentifier(t *testing.T) {
	// callee is a member_expression -> no direct identifier child
	fi := mustParse(t, []byte(`mw.validate({ body: s });`))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

func TestExtractValidateRequest_NotValidateFn(t *testing.T) {
	fi := mustParse(t, []byte(`foo({ body: s });`))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

func TestExtractValidateRequest_NoArgsNode(t *testing.T) {
	fi := mustParse(t, []byte("validate`x`;"))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

func TestExtractValidateRequest_EmptyArgs(t *testing.T) {
	fi := mustParse(t, []byte(`validate();`))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

func TestExtractValidateRequest_FirstArgNotObject(t *testing.T) {
	fi := mustParse(t, []byte(`validate('x');`))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}

func TestExtractValidateRequest_NoMatchingKeys(t *testing.T) {
	fi := mustParse(t, []byte(`validate({ foo: s });`))
	if got := extractValidateRequest(firstCallExpr(t, fi), fi.Src); got != nil {
		t.Fatalf("got %+v", got)
	}
}
