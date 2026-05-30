//ff:func feature=scan type=test control=sequence topic=spring
//ff:what 어노테이션/필드 추출 함수 테스트
package spring

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstField(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	root, src := parseS(t, javaSrc)
	fields := findAllByType(root, "field_declaration")
	if len(fields) == 0 {
		t.Fatal("no field")
	}
	return fields[0], src
}

func TestApplyEmailAnnotation(t *testing.T) {
	field, src := firstField(t, `class C { @Email private String email; }`)
	f := &scanner.Field{}
	applyEmailAnnotation(field, src, f)
	if f.Type != "string:email" {
		t.Fatalf("got %q", f.Type)
	}
}

func TestApplyMinMaxAnnotation(t *testing.T) {
	field, src := firstField(t, `class C { @Min(1) @Max(10) private int x; }`)
	f := &scanner.Field{}
	applyMinAnnotation(field, src, f)
	applyMaxAnnotation(field, src, f)
	if f.Minimum == nil || *f.Minimum != 1 || f.Maximum == nil || *f.Maximum != 10 {
		t.Fatalf("got min=%v max=%v", f.Minimum, f.Maximum)
	}
}

func TestApplySizeAnnotation(t *testing.T) {
	field, src := firstField(t, `class C { @Size(min = 2, max = 8) private String s; }`)
	f := &scanner.Field{}
	applySizeAnnotation(field, src, f)
	if f.MinLength == nil || *f.MinLength != 2 || f.MaxLength == nil || *f.MaxLength != 8 {
		t.Fatalf("got %v %v", f.MinLength, f.MaxLength)
	}
}

func TestApplyValidationAnnotations(t *testing.T) {
	field, src := firstField(t, `class C { @NotBlank @Min(1) private int x; }`)
	f := &scanner.Field{}
	applyValidationAnnotations(field, src, f)
	if f.Validate != "required" || f.Minimum == nil {
		t.Fatalf("got %+v", f)
	}
}

func TestApplyJsonProperty(t *testing.T) {
	field, src := firstField(t, `class C { @JsonProperty("user_name") private String userName; }`)
	f := &scanner.Field{}
	applyJsonProperty(field, src, f)
	if f.JSON != "user_name" {
		t.Fatalf("got %q", f.JSON)
	}
}

func TestFindAnnotation(t *testing.T) {
	field, src := firstField(t, `class C { @Size(min = 1) private String s; }`)
	if findAnnotation(field, src, "Size") == nil || findAnnotation(field, src, "Missing") != nil {
		t.Fatal("findAnnotation")
	}
}

func TestFirstStringArg(t *testing.T) {
	root, src := parseS(t, `class C { @RequestMapping("/x") void m() {} }`)
	anns := findAllByType(root, "annotation")
	if len(anns) == 0 {
		t.Skip("no annotation")
	}
	if got := firstStringArg(anns[0], src); got != "/x" {
		t.Fatalf("got %q", got)
	}
}

func TestAnnotationIntValue(t *testing.T) {
	root, src := parseS(t, `class C { @Size(min = 3) String s; }`)
	ann := findAllByType(root, "annotation")[0]
	if v, ok := annotationIntValue(ann, src, "min"); !ok || v != 3 {
		t.Fatalf("got %d %v", v, ok)
	}
}

func TestAnnotationElementValue(t *testing.T) {
	root, src := parseS(t, `class C { @RequestMapping(value = "/x") void m() {} }`)
	ann := findAllByType(root, "annotation")[0]
	if got := annotationElementValue(ann, src, "value"); got != "/x" {
		t.Fatalf("got %q", got)
	}
}

func TestSingleIntArg(t *testing.T) {
	root, src := parseS(t, `class C { @Min(7) int x; }`)
	ann := findAllByType(root, "annotation")[0]
	if v, ok := singleIntArg(ann, src); !ok || v != 7 {
		t.Fatalf("got %d %v", v, ok)
	}
}

func TestAnnotationName(t *testing.T) {
	root, src := parseS(t, `class C { @org.springframework.web.bind.annotation.GetMapping void m() {} }`)
	anns := findAllByType(root, "marker_annotation")
	if len(anns) == 0 {
		t.Skip("no marker annotation")
	}
	if got := annotationName(anns[0], src); got != "GetMapping" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractFieldNameAndType(t *testing.T) {
	field, src := firstField(t, `class D { private String name; }`)
	if extractFieldName(field, src) != "name" || extractFieldType(field, src) != "String" {
		t.Fatal("field name/type")
	}
}

func TestExtractOneField(t *testing.T) {
	field, src := firstField(t, `class D { @JsonProperty("n") @NotNull private String name; }`)
	f := extractOneField(field, src)
	if f.Name != "name" || f.JSON != "n" || f.Validate != "required" {
		t.Fatalf("got %+v", f)
	}
}

func TestExtractEnumValues(t *testing.T) {
	root, src := parseS(t, `enum Status { OPEN, CLOSED }`)
	en := findAllByType(root, "enum_declaration")[0]
	vals := extractEnumValues(en, src)
	if len(vals) != 2 || vals[0] != "OPEN" {
		t.Fatalf("got %v", vals)
	}
}
