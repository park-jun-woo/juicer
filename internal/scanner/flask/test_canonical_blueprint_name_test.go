//ff:func feature=scan type=test control=sequence topic=flask
//ff:what canonicalBlueprintName이 alias를 원본명으로 역해석한다
package flask

import "testing"

func TestCanonicalBlueprintName(t *testing.T) {
	aliases := importAlias{"auth_blueprint": "auth"}
	if got := canonicalBlueprintName("auth_blueprint", aliases); got != "auth" {
		t.Errorf("aliased: expected auth, got %q", got)
	}
	if got := canonicalBlueprintName("main", aliases); got != "main" {
		t.Errorf("non-aliased: expected main unchanged, got %q", got)
	}
}
