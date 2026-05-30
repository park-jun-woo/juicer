//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what applyOmitType / applyPickType / applyPartialType / applyDecoratorToField / parseObjectArg / collectPrefixCandidates / detectGlobalPrefixInFile / detectURIVersioning / resolveImportPath / resolveNonRelativeImport / tryResolveTS / findTSFiles 테스트
package nestjs

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func sampleFields() []scanner.Field {
	return []scanner.Field{
		{Name: "name", Type: "string", Validate: "required"},
		{Name: "email", Type: "string", Validate: "required,email"},
		{Name: "age", Type: "number", Validate: "required"},
	}
}

func TestApplyOmitType(t *testing.T) {
	got := applyOmitType(sampleFields(), []string{"age"})
	if len(got) != 2 {
		t.Fatalf("expected 2, got %d", len(got))
	}
	for _, f := range got {
		if f.name == "age" {
			t.Fatal("age should be omitted")
		}
	}
}

func TestApplyPickType(t *testing.T) {
	got := applyPickType(sampleFields(), []string{"email"})
	if len(got) != 1 || got[0].name != "email" {
		t.Fatalf("got %+v", got)
	}
}

func TestApplyPartialType(t *testing.T) {
	got := applyPartialType(sampleFields())
	if len(got) != 3 {
		t.Fatalf("got %d", len(got))
	}
	for _, f := range got {
		if !f.optional {
			t.Fatalf("field %s should be optional", f.name)
		}
	}
}

func TestApplyDecoratorToField(t *testing.T) {
	src := []byte(`class D { @IsOptional() name: string; }`)
	root, _ := parseTypeScript(src)
	props := findAllByType(root, "public_field_definition")
	var f dtoField
	applyDecoratorToField(decoratorInfo{name: "IsOptional"}, nil, src, &f)
	if !f.optional {
		t.Fatal("expected optional")
	}
	var f2 dtoField
	applyDecoratorToField(decoratorInfo{name: "MinLength", arg: "5"}, nil, src, &f2)
	if f2.minLength == nil || *f2.minLength != 5 {
		t.Fatalf("minLength: %v", f2.minLength)
	}
	var f3 dtoField
	applyDecoratorToField(decoratorInfo{name: "IsEnum", arg: "Status"}, nil, src, &f3)
	if f3.enumTypeName != "Status" {
		t.Fatalf("enumTypeName: %q", f3.enumTypeName)
	}
	_ = props
}

func TestParseObjectArg(t *testing.T) {
	src := []byte(`const o = { path: 'auth', version: '1' };`)
	root, _ := parseTypeScript(src)
	obj := findAllByType(root, "object")[0]
	d := &decoratorInfo{objectProps: map[string]string{}}
	parseObjectArg(obj, src, d)
	if d.arg != "auth" {
		t.Fatalf("arg: %q", d.arg)
	}
	if d.objectProps["version"] != "1" {
		t.Fatalf("props: %v", d.objectProps)
	}
}

func TestCollectPrefixCandidates(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.ts", "x")
	writeFile(t, dir, "src/app.module.ts", "x")
	cands := collectPrefixCandidates(dir)
	if len(cands) < 2 {
		t.Fatalf("expected main.ts + others, got %v", cands)
	}
	if filepath.Base(cands[0]) != "main.ts" {
		t.Fatalf("main.ts should be first: %v", cands)
	}
}

func TestCollectPrefixCandidates_NoSrc(t *testing.T) {
	dir := t.TempDir()
	cands := collectPrefixCandidates(dir)
	if len(cands) != 1 {
		t.Fatalf("expected just main path, got %v", cands)
	}
}

func TestDetectGlobalPrefixInFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "main.ts", `app.setGlobalPrefix('api');`)
	prefix, found := detectGlobalPrefixInFile(filepath.Join(dir, "main.ts"))
	if !found || prefix != "api" {
		t.Fatalf("got %q %v", prefix, found)
	}
}

func TestDetectGlobalPrefixInFile_NotFound(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "main.ts", `app.listen(3000);`)
	prefix, found := detectGlobalPrefixInFile(filepath.Join(dir, "main.ts"))
	if found || prefix != "" {
		t.Fatalf("got %q %v", prefix, found)
	}
}

func TestDetectGlobalPrefixInFile_Missing(t *testing.T) {
	_, found := detectGlobalPrefixInFile("/no/such.ts")
	if found {
		t.Fatal("expected not found for missing file")
	}
}

func TestDetectURIVersioning(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.ts", `app.enableVersioning({ type: VersioningType.URI });`)
	if !detectURIVersioning(dir) {
		t.Fatal("expected URI versioning detected")
	}
}

func TestDetectURIVersioning_NoMain(t *testing.T) {
	if detectURIVersioning(t.TempDir()) {
		t.Fatal("expected false when no main.ts")
	}
}

func TestTryResolveTS(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "mod.ts", "x")
	if got := tryResolveTS(filepath.Join(dir, "mod")); got != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %q", got)
	}
	writeFile(t, dir, "pkg/index.ts", "x")
	if got := tryResolveTS(filepath.Join(dir, "pkg")); got != filepath.Join(dir, "pkg/index.ts") {
		t.Fatalf("index: %q", got)
	}
	if got := tryResolveTS(filepath.Join(dir, "missing")); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestResolveImportPath_Relative(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto/create.dto.ts", "x")
	got := resolveImportPath(dir, "./dto/create.dto")
	if got != filepath.Join(dir, "dto/create.dto.ts") {
		t.Fatalf("got %q", got)
	}
}

func TestResolveImportPath_NonRelativeNoRoot(t *testing.T) {
	if got := resolveImportPath("/x", "@nestjs/common"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestResolveNonRelativeImport_AtAlias(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/decorators/field.ts", "x")
	got := resolveNonRelativeImport(dir, "@/decorators/field")
	if got != filepath.Join(dir, "src/decorators/field.ts") {
		t.Fatalf("got %q", got)
	}
}

func TestFindTSFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/app.ts", "x")
	writeFile(t, dir, "src/types.d.ts", "x")
	writeFile(t, dir, "src/node_modules/lib.ts", "x")
	files, err := findTSFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel, _ := filepath.Rel(dir, f)
		if rel != filepath.Join("src", "app.ts") {
			t.Errorf("unexpected file: %s", rel)
		}
	}
	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %v", files)
	}
}

func TestFindTSFiles_NoSrc(t *testing.T) {
	files, err := findTSFiles(t.TempDir())
	if err != nil || files != nil {
		t.Fatalf("expected nil, got %v %v", files, err)
	}
}
