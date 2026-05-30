//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestApplyJsonProperty 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyJsonProperty(t *testing.T) {
	field, src := firstField(t, `class C { @JsonProperty("user_name") private String userName; }`)
	f := &scanner.Field{}
	applyJsonProperty(field, src, f)
	if f.JSON != "user_name" {
		t.Fatalf("got %q", f.JSON)
	}
}
