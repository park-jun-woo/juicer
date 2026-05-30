//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestApplyDataAnnotations 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyDataAnnotations(t *testing.T) {
	root, src := parseCS(t, `class C { [Required][StringLength(50)] public string Name { get; set; } }`)
	props := findAllByType(root, "property_declaration")
	f := &scanner.Field{}
	applyDataAnnotations(props[0], src, f)
	if f.Validate != "required" {
		t.Fatalf("validate: %q", f.Validate)
	}
	if f.MaxLength == nil || *f.MaxLength != 50 {
		t.Fatalf("maxlen: %v", f.MaxLength)
	}
}
