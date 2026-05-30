//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractOneField 테스트
package spring

import "testing"

func TestExtractOneField(t *testing.T) {
	field, src := firstField(t, `class D { @JsonProperty("n") @NotNull private String name; }`)
	f := extractOneField(field, src)
	if f.Name != "name" || f.JSON != "n" || f.Validate != "required" {
		t.Fatalf("got %+v", f)
	}
}
