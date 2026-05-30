//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what 다수 nestjs 헬퍼 함수 테스트 (apply/remove/field/resolve/extract/has/is/object/read/find)
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyLengthConstraint(t *testing.T) {
	var p *int
	applyLengthConstraint("5", &p)
	if p == nil || *p != 5 {
		t.Fatalf("got %v", p)
	}
	var q *int
	applyLengthConstraint("notanum", &q)
	if q != nil {
		t.Fatal("expected nil on parse error")
	}
}

func TestRemoveRequired(t *testing.T) {
	if got := removeRequired("required,email"); got != "email" {
		t.Fatalf("got %q", got)
	}
	if got := removeRequired("required"); got != "" {
		t.Fatalf("got %q", got)
	}
	if got := removeRequired("email,min:1"); got != "email,min:1" {
		t.Fatalf("got %q", got)
	}
}

func TestFieldToDTOField(t *testing.T) {
	f := scanner.Field{Name: "email", Type: "string", Validate: "required,email"}
	df := fieldToDTOField(f)
	if df.name != "email" || df.tsType != "string" || df.optional {
		t.Fatalf("got %+v", df)
	}
	// no validate "required" -> optional
	df2 := fieldToDTOField(scanner.Field{Name: "x", Type: ""})
	if !df2.optional || df2.tsType != "string" {
		t.Fatalf("got %+v", df2)
	}
}

func TestConvertOneDtoField(t *testing.T) {
	df := dtoField{name: "age", tsType: "number", optional: false, validators: []string{"min:0"}}
	sf := convertOneDtoField(df)
	if sf.Name != "age" {
		t.Fatalf("got %+v", sf)
	}
	if sf.Validate == "" {
		t.Fatalf("expected validate built, got %+v", sf)
	}
}

func TestScannerFieldsToDTOFields(t *testing.T) {
	fields := []scanner.Field{{Name: "a", Type: "string"}, {Name: "b", Type: "number"}}
	got := scannerFieldsToDTOFields(fields)
	if len(got) != 2 || got[0].name != "a" {
		t.Fatalf("got %+v", got)
	}
}

func TestResolveParamName(t *testing.T) {
	// explicit arg wins
	if got := resolveParamName(DecParam, "userId", "id", "/users/:id"); got != "userId" {
		t.Fatalf("got %q", got)
	}
	// empty @Param() with single path param
	if got := resolveParamName(DecParam, "", "id", "/users/:id"); got != "id" {
		t.Fatalf("got %q", got)
	}
	// empty @Param() with multiple path params -> fallback to go param
	if got := resolveParamName(DecParam, "", "p", "/a/:x/:y"); got != "p" {
		t.Fatalf("got %q", got)
	}
	// non-param decorator -> go param name
	if got := resolveParamName("Query", "", "q", "/x"); got != "q" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractPathParamNames(t *testing.T) {
	if got := extractPathParamNames("/users/:id/posts/:postId"); len(got) != 2 || got[0] != "id" {
		t.Fatalf("got %v", got)
	}
	if got := extractPathParamNames("/users/{id}"); len(got) != 1 || got[0] != "id" {
		t.Fatalf("got %v", got)
	}
	if got := extractPathParamNames("/static/path"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestHasSetGlobalPrefix(t *testing.T) {
	src := []byte(`app.setGlobalPrefix('api');`)
	root, _ := parseTypeScript(src)
	calls := findAllByType(root, "call_expression")
	if !hasSetGlobalPrefix(calls[0], src) {
		t.Fatal("expected true")
	}
	src2 := []byte(`app.listen(3000);`)
	root2, _ := parseTypeScript(src2)
	calls2 := findAllByType(root2, "call_expression")
	if hasSetGlobalPrefix(calls2[0], src2) {
		t.Fatal("expected false")
	}
}

func TestIsEnableURIVersioning(t *testing.T) {
	src := []byte(`app.enableVersioning({ type: VersioningType.URI });`)
	root, _ := parseTypeScript(src)
	call := findAllByType(root, "call_expression")[0]
	if !isEnableURIVersioning(call, src) {
		t.Fatal("expected true")
	}
}

func TestIsEnableURIVersioning_NoArgs(t *testing.T) {
	src := []byte(`app.enableVersioning();`)
	root, _ := parseTypeScript(src)
	call := findAllByType(root, "call_expression")[0]
	if !isEnableURIVersioning(call, src) {
		t.Fatal("expected true for no-args default URI")
	}
}

func TestIsEnableURIVersioning_NotVersioning(t *testing.T) {
	src := []byte(`app.listen(3000);`)
	root, _ := parseTypeScript(src)
	call := findAllByType(root, "call_expression")[0]
	if isEnableURIVersioning(call, src) {
		t.Fatal("expected false")
	}
}

func TestObjectHasURIType(t *testing.T) {
	src := []byte(`const o = { type: VersioningType.URI };`)
	root, _ := parseTypeScript(src)
	obj := findAllByType(root, "object")[0]
	if !objectHasURIType(obj, src) {
		t.Fatal("expected true")
	}
	src2 := []byte(`const o = { type: VersioningType.HEADER };`)
	root2, _ := parseTypeScript(src2)
	obj2 := findAllByType(root2, "object")[0]
	if objectHasURIType(obj2, src2) {
		t.Fatal("expected false")
	}
}

func TestFindPairStringValue(t *testing.T) {
	src := []byte(`const o = { name: 'hello' };`)
	root, _ := parseTypeScript(src)
	pairs := findAllByType(root, "pair")
	if got := findPairStringValue(pairs[0], src); got != "hello" {
		t.Fatalf("got %q", got)
	}
	src2 := []byte(`const o = { count: 5 };`)
	root2, _ := parseTypeScript(src2)
	pairs2 := findAllByType(root2, "pair")
	if got := findPairStringValue(pairs2[0], src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestExtractFactoryStringArray(t *testing.T) {
	src := []byte(`const x = OmitType(Base, ['a', 'b']);`)
	root, _ := parseTypeScript(src)
	args := findAllByType(root, "arguments")[0]
	got := extractFactoryStringArray(args, src)
	if len(got) != 2 || got[0] != "a" {
		t.Fatalf("got %v", got)
	}
}

func TestExtractFactoryStringArray_None(t *testing.T) {
	src := []byte(`const x = PartialType(Base);`)
	root, _ := parseTypeScript(src)
	args := findAllByType(root, "arguments")[0]
	if got := extractFactoryStringArray(args, src); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestReadEnvPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, ".env", "# comment\nAPI_PREFIX=api/v1\nOTHER=x\n")
	if got := readEnvPrefix(dir, ".env"); got != "api/v1" {
		t.Fatalf("got %q", got)
	}
	if got := readEnvPrefix(dir, ".missing"); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestReadConfigDefault(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/config/app.config.ts", `export default { apiPrefix: process.env.API_PREFIX || 'api' };`)
	if got := readConfigDefault(dir); got != "api" {
		t.Fatalf("got %q", got)
	}
	if got := readConfigDefault(t.TempDir()); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestFallbackGlobalPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, ".env", "API_PREFIX=fromenv\n")
	if got := fallbackGlobalPrefix(dir); got != "fromenv" {
		t.Fatalf("got %q", got)
	}
	// empty project -> ""
	if got := fallbackGlobalPrefix(t.TempDir()); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
