//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what applyEmailAnnotation / applyMinAnnotation / applyMaxAnnotation / applySizeAnnotation / applyValidationAnnotations / hasAnnotation / findAnnotation 테스트
package quarkus

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func firstFieldDecl(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	src := []byte(javaSrc)
	root, err := parseJava(src)
	if err != nil {
		t.Fatal(err)
	}
	fields := findAllByType(root, "field_declaration")
	if len(fields) == 0 {
		t.Fatal("no field_declaration")
	}
	return fields[0], src
}

func TestApplyEmailAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Email private String email; }`)
	f := &scanner.Field{}
	applyEmailAnnotation(field, src, f)
	if f.Type != "string:email" {
		t.Fatalf("got %q", f.Type)
	}
}

func TestApplyEmailAnnotation_None(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { private String name; }`)
	f := &scanner.Field{}
	applyEmailAnnotation(field, src, f)
	if f.Type != "" {
		t.Fatalf("expected unchanged, got %q", f.Type)
	}
}

func TestApplyMinAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Min(5) private int x; }`)
	f := &scanner.Field{}
	applyMinAnnotation(field, src, f)
	if f.Minimum == nil || *f.Minimum != 5 {
		t.Fatalf("got %v", f.Minimum)
	}
}

func TestApplyMinAnnotation_ValueArg(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Min(value = 3) private int x; }`)
	f := &scanner.Field{}
	applyMinAnnotation(field, src, f)
	if f.Minimum == nil || *f.Minimum != 3 {
		t.Fatalf("got %v", f.Minimum)
	}
}

func TestApplyMinAnnotation_None(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { private int x; }`)
	f := &scanner.Field{}
	applyMinAnnotation(field, src, f)
	if f.Minimum != nil {
		t.Fatal("expected nil")
	}
}

func TestApplyMaxAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Max(100) private int x; }`)
	f := &scanner.Field{}
	applyMaxAnnotation(field, src, f)
	if f.Maximum == nil || *f.Maximum != 100 {
		t.Fatalf("got %v", f.Maximum)
	}
}

func TestApplySizeAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Size(min = 1, max = 10) private String s; }`)
	f := &scanner.Field{}
	applySizeAnnotation(field, src, f)
	if f.MinLength == nil || *f.MinLength != 1 {
		t.Fatalf("minLen: %v", f.MinLength)
	}
	if f.MaxLength == nil || *f.MaxLength != 10 {
		t.Fatalf("maxLen: %v", f.MaxLength)
	}
}

func TestApplyValidationAnnotations(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @NotNull @Min(1) private int x; }`)
	f := &scanner.Field{}
	applyValidationAnnotations(field, src, f)
	if f.Validate != "required" {
		t.Fatalf("validate: %q", f.Validate)
	}
	if f.Minimum == nil || *f.Minimum != 1 {
		t.Fatalf("min: %v", f.Minimum)
	}
}

func TestHasAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @NotNull private String name; }`)
	if !hasAnnotation(field, src, "NotNull") {
		t.Fatal("expected NotNull")
	}
	if hasAnnotation(field, src, "Email") {
		t.Fatal("unexpected Email")
	}
}

func TestHasAnnotation_NoModifiers(t *testing.T) {
	// a node with no modifiers
	root, _ := parseJava([]byte(`class C { void m() {} }`))
	classes := findAllByType(root, "class_body")
	if hasAnnotation(classes[0], []byte(`class C { void m() {} }`), "NotNull") {
		t.Fatal("expected false")
	}
}

func TestFindAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Size(min = 1) private String s; }`)
	if findAnnotation(field, src, "Size") == nil {
		t.Fatal("expected Size annotation")
	}
	if findAnnotation(field, src, "Missing") != nil {
		t.Fatal("expected nil")
	}
}
