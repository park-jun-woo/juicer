//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what dotnet 순수/AST 헬퍼 함수 테스트
package dotnet

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"

	sitter "github.com/smacker/go-tree-sitter"
)

func parseCS(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parseCSharp(b)
	if err != nil {
		t.Fatal(err)
	}
	return root, b
}

func TestIntPtr(t *testing.T) {
	if p := intPtr(8); p == nil || *p != 8 {
		t.Fatalf("got %v", p)
	}
}

func TestJoinPath(t *testing.T) {
	if got := joinPath("/api", "users", "/{id}/"); got != "/api/users/{id}" {
		t.Fatalf("got %q", got)
	}
	if got := joinPath("", ""); got != "/" {
		t.Fatalf("got %q", got)
	}
}

func TestDefaultStatusForMethod(t *testing.T) {
	if defaultStatusForMethod("POST") != "201" || defaultStatusForMethod("GET") != "200" {
		t.Fatal("status")
	}
}

func TestStripGeneric(t *testing.T) {
	if stripGeneric("List<X>") != "List" || stripGeneric("X") != "X" {
		t.Fatal("strip")
	}
}

func TestStripNullable(t *testing.T) {
	n, ok := stripNullable("int?")
	if n != "int" || !ok {
		t.Fatalf("nullable: %q %v", n, ok)
	}
	n2, ok2 := stripNullable("int")
	if n2 != "int" || ok2 {
		t.Fatalf("non-nullable: %q %v", n2, ok2)
	}
}

func TestExtractGenericInner(t *testing.T) {
	if extractGenericInner("List<UserDto>") != "UserDto" {
		t.Fatal("generic")
	}
	if extractGenericInner("String") != "String" {
		t.Fatal("plain")
	}
}

func TestMergeRoles(t *testing.T) {
	if mergeRoles([]string{"a"}, []string{"b"})[0] != "b" {
		t.Fatal("method")
	}
	if mergeRoles([]string{"a"}, nil)[0] != "a" {
		t.Fatal("class")
	}
}

func TestUnquoteCSharp(t *testing.T) {
	if unquoteCSharp(`"x"`) != "x" || unquoteCSharp("y") != "y" {
		t.Fatal("unquote")
	}
}

func TestIsPrimitiveType(t *testing.T) {
	if !isPrimitiveType("int") || isPrimitiveType("UserDto") {
		t.Fatal("primitive")
	}
}

func TestCsharpTypeToOpenAPIType(t *testing.T) {
	if csharpTypeToOpenAPIType("int") != "integer" {
		t.Fatal("int")
	}
	if csharpTypeToOpenAPIType("UserDto") != "object" {
		t.Fatal("unknown -> object")
	}
}

func TestCsharpTypeToOpenAPI(t *testing.T) {
	if csharpTypeToOpenAPI("string").Type != "string" {
		t.Fatal("string")
	}
	if csharpTypeToOpenAPI("Guid").Format != "uuid" {
		t.Fatal("guid")
	}
	arr := csharpTypeToOpenAPI("List<UserDto>")
	if arr.Type != "array" || arr.Items != "UserDto" {
		t.Fatalf("array: %+v", arr)
	}
	if csharpTypeToOpenAPI("int?").Type != "integer" {
		t.Fatal("nullable int")
	}
}

func TestHasContent(t *testing.T) {
	if hasContent(&scanner.Request{}) {
		t.Fatal("empty")
	}
	if !hasContent(&scanner.Request{Body: &scanner.Body{}}) {
		t.Fatal("body")
	}
}

func TestAppendValidate(t *testing.T) {
	if appendValidate("", "x") != "x" || appendValidate("a", "b") != "a,b" {
		t.Fatal("append")
	}
}

func TestIsPathParam(t *testing.T) {
	if !isPathParam("id", "/users/{id}") || isPathParam("id", "/users") {
		t.Fatal("path param")
	}
}

func TestUnwrapReturnType(t *testing.T) {
	if got, ok := unwrapReturnType("ActionResult<UserDto>"); ok || got != "UserDto" {
		t.Fatalf("ActionResult: %q %v", got, ok)
	}
	if got, ok := unwrapReturnType("Task<List<UserDto>>"); !ok || got != "UserDto" {
		t.Fatalf("Task<List>: %q %v", got, ok)
	}
	if got, ok := unwrapReturnType("UserDto[]"); !ok || got != "UserDto" {
		t.Fatalf("array: %q %v", got, ok)
	}
}

func TestNodeTextAndFindChild(t *testing.T) {
	root, src := parseCS(t, `class C {}`)
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no class")
	}
	id := findChildByType(classes[0], "identifier")
	if id != nil && nodeText(id, src) != "C" {
		t.Fatalf("name %q", nodeText(id, src))
	}
	if findChildByType(classes[0], "nonexistent") != nil {
		t.Fatal("nil expected")
	}
}

func TestChildrenOfType(t *testing.T) {
	root, _ := parseCS(t, `class C { int a; int b; }`)
	body := findAllByType(root, "declaration_list")[0]
	fields := childrenOfType(body, "field_declaration")
	if len(fields) != 2 {
		t.Fatalf("got %d", len(fields))
	}
}

func TestLastIdentifier(t *testing.T) {
	root, src := parseCS(t, `class C { [System.Foo] void m() {} }`)
	qn := findAllByType(root, "qualified_name")
	if len(qn) == 0 {
		t.Skip("no qualified name")
	}
	if got := lastIdentifier(qn[0], src); got != "Foo" {
		t.Fatalf("got %q", got)
	}
}

func TestAttributeName(t *testing.T) {
	root, src := parseCS(t, `class C { [HttpGet] void m() {} }`)
	attrs := findAllByType(root, "attribute")
	if len(attrs) == 0 {
		t.Skip("no attribute")
	}
	if got := attributeName(attrs[0], src); got != "HttpGet" {
		t.Fatalf("got %q", got)
	}
}
