//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what appendValidate / applyConstraintRule / applyFlagRule / applyFormatRule / applyMaxRule / applyMinRule / applyTypeRule 테스트
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAppendValidate(t *testing.T) {
	if got := appendValidate("", "required"); got != "required" {
		t.Fatalf("got %q", got)
	}
	if got := appendValidate("required", "email"); got != "required,email" {
		t.Fatalf("got %q", got)
	}
}

func TestApplyFlagRule(t *testing.T) {
	f := &scanner.Field{}
	if !applyFlagRule(f, "nullable") || !f.Nullable {
		t.Fatal("nullable not applied")
	}
	f2 := &scanner.Field{}
	if !applyFlagRule(f2, "required") || f2.Validate != "required" {
		t.Fatalf("required not applied: %+v", f2)
	}
	f3 := &scanner.Field{}
	if applyFlagRule(f3, "email") {
		t.Fatal("unexpected handled for email")
	}
}

func TestApplyFormatRule(t *testing.T) {
	f := &scanner.Field{}
	if !applyFormatRule(f, "email") {
		t.Fatal("email not handled")
	}
	if f.Type != "string" || f.Validate != "email" {
		t.Fatalf("got %+v", f)
	}
	f2 := &scanner.Field{}
	if applyFormatRule(f2, "notaformat") {
		t.Fatal("unexpected handled")
	}
}

func TestApplyTypeRule(t *testing.T) {
	f := &scanner.Field{}
	isNum := false
	if !applyTypeRule(f, "integer", &isNum) || f.Type != "integer" || !isNum {
		t.Fatalf("integer: %+v num=%v", f, isNum)
	}
	f2 := &scanner.Field{}
	isNum2 := false
	if !applyTypeRule(f2, "string", &isNum2) || f2.Type != "string" || isNum2 {
		t.Fatalf("string: %+v num=%v", f2, isNum2)
	}
	f3 := &scanner.Field{}
	if applyTypeRule(f3, "unknown", &isNum2) {
		t.Fatal("unexpected handled for unknown type")
	}
}

func TestApplyMaxRule(t *testing.T) {
	f := &scanner.Field{}
	applyMaxRule(f, "10", true)
	if f.Maximum == nil || *f.Maximum != 10 {
		t.Fatalf("number max: %+v", f)
	}
	f2 := &scanner.Field{}
	applyMaxRule(f2, "5", false)
	if f2.MaxLength == nil || *f2.MaxLength != 5 {
		t.Fatalf("string maxlen: %+v", f2)
	}
	f3 := &scanner.Field{}
	applyMaxRule(f3, "notanum", true)
	if f3.Maximum != nil {
		t.Fatal("expected no change on parse error")
	}
}

func TestApplyMinRule(t *testing.T) {
	f := &scanner.Field{}
	applyMinRule(f, "2", true)
	if f.Minimum == nil || *f.Minimum != 2 {
		t.Fatalf("number min: %+v", f)
	}
	f2 := &scanner.Field{}
	applyMinRule(f2, "3", false)
	if f2.MinLength == nil || *f2.MinLength != 3 {
		t.Fatalf("string minlen: %+v", f2)
	}
	f3 := &scanner.Field{}
	applyMinRule(f3, "x", false)
	if f3.MinLength != nil {
		t.Fatal("expected no change on parse error")
	}
}

func TestApplyConstraintRule(t *testing.T) {
	f := &scanner.Field{}
	applyConstraintRule(f, "max:7", false)
	if f.MaxLength == nil || *f.MaxLength != 7 {
		t.Fatalf("max: %+v", f)
	}
	f2 := &scanner.Field{}
	applyConstraintRule(f2, "min:1", true)
	if f2.Minimum == nil || *f2.Minimum != 1 {
		t.Fatalf("min: %+v", f2)
	}
	f3 := &scanner.Field{}
	applyConstraintRule(f3, "in:a,b,c", false)
	if len(f3.Enum) != 3 || f3.Enum[0] != "a" {
		t.Fatalf("in: %+v", f3.Enum)
	}
	f4 := &scanner.Field{}
	applyConstraintRule(f4, "unrelated:x", false)
	if f4.MaxLength != nil || f4.Minimum != nil || len(f4.Enum) != 0 {
		t.Fatalf("expected no change: %+v", f4)
	}
}
