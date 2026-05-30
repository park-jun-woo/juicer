//ff:func feature=scan type=test control=sequence topic=spring
//ff:what 다수 순수 헬퍼 함수 테스트
package spring

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"

	sitter "github.com/smacker/go-tree-sitter"
)

func parseS(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	return root, b
}

func TestIntPtr(t *testing.T) {
	if p := intPtr(9); p == nil || *p != 9 {
		t.Fatalf("got %v", p)
	}
}

func TestParseInt(t *testing.T) {
	if parseInt("204") != 204 || parseInt("x9y") != 9 || parseInt("") != 0 {
		t.Fatal("parseInt")
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

func TestIsCollectionType(t *testing.T) {
	if !isCollectionType("List<X>") || isCollectionType("X") {
		t.Fatal("collection")
	}
}

func TestIsJavaType(t *testing.T) {
	if !isJavaType("String") || !isJavaType("List<X>") || isJavaType("MyDto") {
		t.Fatal("javaType")
	}
}

func TestIsPrimitiveType(t *testing.T) {
	if !isPrimitiveType("int") || isPrimitiveType("MyDto") {
		t.Fatal("primitive")
	}
}

func TestExtractGenericInner(t *testing.T) {
	if extractGenericInner("List<UserDto>") != "UserDto" || extractGenericInner("String") != "" {
		t.Fatal("generic inner")
	}
}

func TestAngleBracketDelta(t *testing.T) {
	if angleBracketDelta('<') != 1 || angleBracketDelta('>') != -1 || angleBracketDelta('z') != 0 {
		t.Fatal("delta")
	}
}

func TestMergeRoles(t *testing.T) {
	if mergeRoles([]string{"a"}, []string{"b"})[0] != "b" {
		t.Fatal("method wins")
	}
	if mergeRoles([]string{"a"}, nil)[0] != "a" {
		t.Fatal("class fallback")
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

func TestNodeTextFindChild(t *testing.T) {
	root, src := parseS(t, `class C {}`)
	classes := findAllByType(root, "class_declaration")
	if findChildByType(classes[0], "class_body") == nil {
		t.Fatal("class_body")
	}
	if findChildByType(classes[0], "nope") != nil {
		t.Fatal("nil expected")
	}
	id := findChildByType(classes[0], "identifier")
	if id != nil && nodeText(id, src) != "C" {
		t.Fatalf("name %q", nodeText(id, src))
	}
}

func TestChildrenOfType(t *testing.T) {
	root, _ := parseS(t, `class C { int a; int b; }`)
	body := findAllByType(root, "class_body")[0]
	if len(childrenOfType(body, "field_declaration")) != 2 {
		t.Fatal("children")
	}
}

func TestUnquoteJava(t *testing.T) {
	if unquoteJava(`"x"`) != "x" || unquoteJava("y") != "y" {
		t.Fatal("unquote")
	}
}

func TestStripGenericPure(t *testing.T) {
	if stripGeneric("List<X>") != "List" || stripGeneric("X") != "X" {
		t.Fatal("strip")
	}
}

func TestSplitGenericArgs(t *testing.T) {
	got := splitGenericArgs("A, Map<B,C>, D")
	if len(got) != 3 || got[1] != "Map<B,C>" {
		t.Fatalf("got %v", got)
	}
}

func TestSubstituteType(t *testing.T) {
	m := map[string]string{"T": "string"}
	if substituteType("T", m) != "string" || substituteType("[]T", m) != "[]string" {
		t.Fatal("substitute")
	}
	if substituteType("List<T>", m) != "List<string>" {
		t.Fatal("generic substitute")
	}
}

func TestUnwrapReturnType(t *testing.T) {
	if _, ok := unwrapReturnType("void"); ok {
		t.Fatal("void")
	}
	if got, ok := unwrapReturnType("ResponseEntity<UserDto>"); ok || got != "UserDto" {
		t.Fatalf("RE single: %q %v", got, ok)
	}
	if got, ok := unwrapReturnType("List<UserDto>"); !ok || got != "UserDto" {
		t.Fatalf("List: %q %v", got, ok)
	}
	if _, ok := unwrapReturnType("ResponseEntity<Void>"); ok {
		t.Fatal("RE Void")
	}
}

func TestJavaTypeToOpenAPIString(t *testing.T) {
	if javaTypeToOpenAPIString("UUID") != "string:uuid" || javaTypeToOpenAPIString("String") != "string" {
		t.Fatal("oa string")
	}
}

func TestNormalizeRole(t *testing.T) {
	if normalizeRole("ROLE_ADMIN") != "ADMIN" || normalizeRole("USER") != "USER" {
		t.Fatal("normalize")
	}
}

func TestHasAnnotation(t *testing.T) {
	root, src := parseS(t, `class C { @NotNull String x; }`)
	field := findAllByType(root, "field_declaration")[0]
	if !hasAnnotation(field, src, "NotNull") || hasAnnotation(field, src, "Email") {
		t.Fatal("hasAnnotation")
	}
}

func TestFindModifiers(t *testing.T) {
	root, _ := parseS(t, `class C { public int x; }`)
	field := findAllByType(root, "field_declaration")[0]
	if findModifiers(field) == nil {
		t.Fatal("modifiers")
	}
}

func TestIsStaticField(t *testing.T) {
	root, _ := parseS(t, `class C { static int a; int b; }`)
	fields := findAllByType(root, "field_declaration")
	if !isStaticField(fields[0]) || isStaticField(fields[1]) {
		t.Fatal("static")
	}
}
