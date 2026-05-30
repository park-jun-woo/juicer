//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractOneZodValidator 테스트
package hono

import "testing"

func TestExtractOneZodValidator_SchemaName(t *testing.T) {
	call, src := firstCallExpr(t, `zValidator("json", createUserSchema);`)
	v := extractOneZodValidator(call, src)
	if v == nil || v.Target != "json" || v.SchemaName != "createUserSchema" {
		t.Fatalf("got %+v", v)
	}
}

func TestExtractOneZodValidator_SchemaNode(t *testing.T) {
	call, src := firstCallExpr(t, `zValidator("query", z.object({ a: z.string() }));`)
	v := extractOneZodValidator(call, src)
	if v == nil || v.Target != "query" || v.SchemaName != "" || v.SchemaNode == nil {
		t.Fatalf("got %+v", v)
	}
}

func TestExtractOneZodValidator_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const x = foo;`))
	id := findAllByType(fi.Root, "identifier")[0]
	if v := extractOneZodValidator(id, fi.Src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}

func TestExtractOneZodValidator_NotZValidator(t *testing.T) {
	call, src := firstCallExpr(t, `other("json", s);`)
	if v := extractOneZodValidator(call, src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}

func TestExtractOneZodValidator_TooFewArgs(t *testing.T) {
	call, src := firstCallExpr(t, `zValidator("json");`)
	if v := extractOneZodValidator(call, src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}

func TestExtractOneZodValidator_NoIdentifierFn(t *testing.T) {
	// member-call: function part is member_expression, no direct identifier child
	call, src := firstCallExpr(t, `a.b("json", s);`)
	if v := extractOneZodValidator(call, src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}

func TestExtractOneZodValidator_TargetNotString(t *testing.T) {
	call, src := firstCallExpr(t, `zValidator(targetVar, schema);`)
	if v := extractOneZodValidator(call, src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}
