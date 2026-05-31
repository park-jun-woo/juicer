//ff:func feature=scan type=test topic=joi control=sequence
//ff:what ApplyMethod 타입/검증/enum 반영 테스트
package joi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyMethod(t *testing.T) {
	var f scanner.Field
	ApplyMethod(&f, ChainMethod{Name: "string"})
	ApplyMethod(&f, ChainMethod{Name: "required"})
	ApplyMethod(&f, ChainMethod{Name: "email"})
	if f.Type != "string" || f.Validate != "required,email" {
		t.Errorf("string/required/email: %+v", f)
	}

	var d scanner.Field
	ApplyMethod(&d, ChainMethod{Name: "date"})
	if d.Type != "string" || d.Validate != "date-time" {
		t.Errorf("date: %+v", d)
	}

	var v scanner.Field
	ApplyMethod(&v, ChainMethod{Name: "valid", Args: []string{"A", "B"}})
	if v.Type != "string" || !reflect.DeepEqual(v.Enum, []string{"A", "B"}) {
		t.Errorf("valid: %+v", v)
	}

	var n scanner.Field
	ApplyMethod(&n, ChainMethod{Name: "number"})
	ApplyMethod(&n, ChainMethod{Name: "uuid"})
	if n.Type != "number" || n.Validate != "uuid" {
		t.Errorf("number/uuid: %+v", n)
	}
}
