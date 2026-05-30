//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplyJsonProperty 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyJsonProperty(t *testing.T) {
	root, _ := parseJava([]byte(`class D { @JsonProperty("user_name") private String userName; }`))
	src := []byte(`class D { @JsonProperty("user_name") private String userName; }`)
	field := findAllByType(root, "field_declaration")[0]
	f := &scanner.Field{Name: "userName"}
	applyJsonProperty(field, src, f)
	if f.JSON != "user_name" {
		t.Fatalf("got %q", f.JSON)
	}
}
