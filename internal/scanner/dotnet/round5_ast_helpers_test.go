//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what round5 AST 헬퍼 공용 유틸 + 단순 헬퍼 테스트
package dotnet

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

// firstOfType returns the first descendant node of the given type.
func firstOfType(t *testing.T, root *sitter.Node, typ string) *sitter.Node {
	t.Helper()
	nodes := findAllByType(root, typ)
	if len(nodes) == 0 {
		t.Fatalf("no %s node found", typ)
	}
	return nodes[0]
}

func TestWalkNodes_Round5(t *testing.T) {
	root, _ := parseCS(t, "class C { void M() {} }")
	count := 0
	walkNodes(root, func(n *sitter.Node) { count++ })
	if count < 3 {
		t.Fatalf("walkNodes visited too few nodes: %d", count)
	}
}

func TestFindAttributeInList_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { [HttpGet("/x")] public void M() {} }`)
	attrList := firstOfType(t, root, "attribute_list")
	if got := findAttributeInList(attrList, src, "HttpGet"); got == nil {
		t.Fatal("expected HttpGet attribute found")
	}
	if got := findAttributeInList(attrList, src, "Missing"); got != nil {
		t.Fatal("expected nil for missing attribute")
	}
}

func TestFindStringLiteralInArg_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { F("hello"); } }`)
	arg := firstOfType(t, root, "argument")
	if got := findStringLiteralInArg(arg, src); got != "hello" {
		t.Fatalf("got %q", got)
	}
	// arg with no string literal
	root2, src2 := parseCS(t, `class C { void M() { F(42); } }`)
	arg2 := firstOfType(t, root2, "argument")
	if got := findStringLiteralInArg(arg2, src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestFindIntLiteralInArg_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { F(42); } }`)
	arg := firstOfType(t, root, "argument")
	v, ok := findIntLiteralInArg(arg, src)
	if !ok || v != 42 {
		t.Fatalf("got %d %v", v, ok)
	}
	root2, src2 := parseCS(t, `class C { void M() { F("x"); } }`)
	arg2 := firstOfType(t, root2, "argument")
	if _, ok := findIntLiteralInArg(arg2, src2); ok {
		t.Fatal("expected no int literal")
	}
}

func TestExtractFirstStringFromArgs_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { F("a", "b"); } }`)
	args := firstOfType(t, root, "argument_list")
	if got := extractFirstStringFromArgs(args, src); got != "a" {
		t.Fatalf("got %q", got)
	}
	root2, src2 := parseCS(t, `class C { void M() { F(1, 2); } }`)
	args2 := firstOfType(t, root2, "argument_list")
	if got := extractFirstStringFromArgs(args2, src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestExtractUsingNamespace_Round5(t *testing.T) {
	root, src := parseCS(t, "using System.Text;\nclass C {}")
	using := firstOfType(t, root, "using_directive")
	if got := extractUsingNamespace(using, src); got != "System.Text" {
		t.Fatalf("got %q", got)
	}
}

func TestAttributeNamedArg_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { [Authorize(Roles = "Admin")] void M() {} }`)
	attr := firstOfType(t, root, "attribute")
	if got := attributeNamedArg(attr, src, "Roles"); got != "Admin" {
		t.Fatalf("got %q", got)
	}
	if got := attributeNamedArg(attr, src, "Policy"); got != "" {
		t.Fatalf("expected empty for missing named arg, got %q", got)
	}
}

func TestMatchHTTPAttribute_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { [HttpPost("/items")] void M() {} }`)
	attrList := firstOfType(t, root, "attribute_list")
	method, path, ok := matchHTTPAttribute(attrList, src)
	if !ok || method != "POST" || path != "/items" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}
	// non-http attribute list
	root2, src2 := parseCS(t, `class C { [Serializable] void M() {} }`)
	attrList2 := firstOfType(t, root2, "attribute_list")
	if _, _, ok := matchHTTPAttribute(attrList2, src2); ok {
		t.Fatal("expected no http attribute")
	}
}

func TestExtractClassRoles_Round5(t *testing.T) {
	root, src := parseCS(t, `[Authorize(Roles = "Admin,User")] class C {}`)
	cls := firstOfType(t, root, "class_declaration")
	roles := extractClassRoles(cls, src)
	if len(roles) != 2 {
		t.Fatalf("expected 2 roles, got %v", roles)
	}
}

func TestMethodLevelRoute_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { [Route("custom")] public void M() {} }`)
	m := firstOfType(t, root, "method_declaration")
	if got := methodLevelRoute(m, src); got != "custom" {
		t.Fatalf("got %q", got)
	}
	root2, src2 := parseCS(t, `class C { public void M() {} }`)
	m2 := firstOfType(t, root2, "method_declaration")
	if got := methodLevelRoute(m2, src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestExtractPropertyName_And_Type_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { public string Name { get; set; } }`)
	prop := firstOfType(t, root, "property_declaration")
	if got := extractPropertyName(prop, src); got != "Name" {
		t.Fatalf("name: got %q", got)
	}
	if got := extractPropertyTypeName(prop, src); got != "string" {
		t.Fatalf("type: got %q", got)
	}
}

func TestExtractRecordParam_Round5(t *testing.T) {
	root, src := parseCS(t, `public record R(string Name, int? Age);`)
	params := findAllByType(root, "parameter")
	if len(params) < 2 {
		t.Fatalf("expected 2 params, got %d", len(params))
	}
	f0 := extractRecordParam(params[0], src)
	if f0.Name != "Name" || f0.Type != "string" {
		t.Fatalf("param0: %+v", f0)
	}
	f1 := extractRecordParam(params[1], src)
	if f1.Name != "Age" || !f1.Nullable {
		t.Fatalf("param1 should be nullable int: %+v", f1)
	}
}

func TestExtractMethodParams_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { void M([FromQuery] string a, [FromBody] int b) {} }`)
	m := firstOfType(t, root, "method_declaration")
	var ep endpointInfo
	extractMethodParams(m, src, &ep)
	if len(ep.query) != 1 || ep.query[0].Name != "a" {
		t.Fatalf("query: %+v", ep.query)
	}
	if ep.bodyType == "" {
		t.Fatalf("expected body type set, got %+v", ep)
	}
}
