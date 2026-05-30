//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what 다수 순수 헬퍼 함수 테스트
package quarkus

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestIntPtr(t *testing.T) {
	p := intPtr(7)
	if p == nil || *p != 7 {
		t.Fatalf("got %v", p)
	}
}

func TestParseInt(t *testing.T) {
	if parseInt("123") != 123 {
		t.Fatal("123")
	}
	if parseInt("a1b2") != 12 {
		t.Fatal("a1b2")
	}
	if parseInt("") != 0 {
		t.Fatal("empty")
	}
}

func TestJoinPath(t *testing.T) {
	if got := joinPath("/api", "users", "/{id}/"); got != "/api/users/{id}" {
		t.Fatalf("got %q", got)
	}
	if got := joinPath("", "", ""); got != "/" {
		t.Fatalf("got %q", got)
	}
	if got := joinPath("/users"); got != "/users" {
		t.Fatalf("got %q", got)
	}
}

func TestDefaultStatusForMethod(t *testing.T) {
	if defaultStatusForMethod("POST") != "201" {
		t.Fatal("POST")
	}
	if defaultStatusForMethod("GET") != "200" {
		t.Fatal("GET")
	}
}

func TestIsCollectionType(t *testing.T) {
	if !isCollectionType("List<String>") {
		t.Fatal("List")
	}
	if !isCollectionType("Set<Long>") {
		t.Fatal("Set")
	}
	if isCollectionType("String") {
		t.Fatal("String")
	}
}

func TestIsJavaType(t *testing.T) {
	if !isJavaType("String") {
		t.Fatal("String")
	}
	if !isJavaType("List<Foo>") {
		t.Fatal("List")
	}
	if !isJavaType("int[]") {
		t.Fatal("array")
	}
	if isJavaType("MyCustomDto") {
		t.Fatal("custom dto should be false")
	}
}

func TestIsPrimitiveType(t *testing.T) {
	if !isPrimitiveType("int") {
		t.Fatal("int")
	}
	if isPrimitiveType("MyDto") {
		t.Fatal("dto")
	}
}

func TestFieldTypeToScannerType(t *testing.T) {
	if got := fieldTypeToScannerType("String"); got != "string" {
		t.Fatalf("string: %q", got)
	}
	if got := fieldTypeToScannerType("UUID"); got != "string:uuid" {
		t.Fatalf("uuid: %q", got)
	}
	if got := fieldTypeToScannerType("List<String>"); got != "array:string" {
		t.Fatalf("array: %q", got)
	}
}

func TestJavaTypeToOpenAPI(t *testing.T) {
	if javaTypeToOpenAPI("String").Type != "string" {
		t.Fatal("string")
	}
	if javaTypeToOpenAPI("Long").Format != "int64" {
		t.Fatal("long")
	}
	if javaTypeToOpenAPI("UUID").Format != "uuid" {
		t.Fatal("uuid")
	}
}

func TestJavaTypeToOpenAPIString(t *testing.T) {
	if got := javaTypeToOpenAPIString("UUID"); got != "string:uuid" {
		t.Fatalf("got %q", got)
	}
	if got := javaTypeToOpenAPIString("String"); got != "string" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractGenericInner(t *testing.T) {
	if got := extractGenericInner("List<UserDto>"); got != "UserDto" {
		t.Fatalf("got %q", got)
	}
	if got := extractGenericInner("String"); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestAngleBracketDelta(t *testing.T) {
	if angleBracketDelta('<') != 1 || angleBracketDelta('>') != -1 || angleBracketDelta('x') != 0 {
		t.Fatal("delta")
	}
}

func TestMergeRoles(t *testing.T) {
	if got := mergeRoles([]string{"a"}, []string{"b"}); got[0] != "b" {
		t.Fatal("method wins")
	}
	if got := mergeRoles([]string{"a"}, nil); got[0] != "a" {
		t.Fatal("class fallback")
	}
}

func TestHasContent(t *testing.T) {
	if hasContent(&scanner.Request{}) {
		t.Fatal("empty")
	}
	if !hasContent(&scanner.Request{PathParams: []scanner.Param{{Name: "id"}}}) {
		t.Fatal("path params")
	}
	if !hasContent(&scanner.Request{Body: &scanner.Body{}}) {
		t.Fatal("body")
	}
}

func TestNodeTextAndFindChild(t *testing.T) {
	root, _ := parseJava([]byte(`class C {}`))
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no class")
	}
	if findChildByType(classes[0], "class_body") == nil {
		t.Fatal("expected class_body")
	}
	if findChildByType(classes[0], "nonexistent") != nil {
		t.Fatal("expected nil")
	}
	id := findChildByType(classes[0], "identifier")
	if id != nil && nodeText(id, []byte(`class C {}`)) != "C" {
		t.Fatalf("name %q", nodeText(id, []byte(`class C {}`)))
	}
}

func TestChildrenOfType(t *testing.T) {
	root, _ := parseJava([]byte(`class C { int a; int b; }`))
	body := findAllByType(root, "class_body")[0]
	fields := childrenOfType(body, "field_declaration")
	if len(fields) != 2 {
		t.Fatalf("got %d", len(fields))
	}
}

func TestExtractRoleStrings(t *testing.T) {
	root, _ := parseJava([]byte(`class C { @RolesAllowed({"admin", "user"}) void m() {} }`))
	args := findAllByType(root, "annotation_argument_list")
	if len(args) == 0 {
		t.Skip("no args")
	}
	roles := extractRoleStrings(args[0], []byte(`class C { @RolesAllowed({"admin", "user"}) void m() {} }`))
	if len(roles) != 2 || roles[0] != "admin" {
		t.Fatalf("got %v", roles)
	}
}

func TestIsStaticField(t *testing.T) {
	root, _ := parseJava([]byte(`class C { static int a; int b; }`))
	fields := findAllByType(root, "field_declaration")
	if !isStaticField(fields[0]) {
		t.Fatal("a is static")
	}
	if isStaticField(fields[1]) {
		t.Fatal("b is not static")
	}
}
