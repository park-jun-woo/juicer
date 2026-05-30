//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestConsumeFieldChild_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestConsumeFieldChild_Round5(t *testing.T) {
	fi := aFi(t, `struct S { #[serde(rename = "n")] name: String }`)
	var fields []scanner.Field
	var pending []serdeAttr
	field := aFirst(t, fi.root, "field_declaration")
	for i := 0; i < int(field.ChildCount()); i++ {
		fields, pending = consumeFieldChild(field.Child(i), fi.src, fields, pending)
	}
}
