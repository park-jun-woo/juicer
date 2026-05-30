//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestResolveBlueprintPrefixes_NoOverride 테스트
package flask

import "testing"

func TestResolveBlueprintPrefixes_NoOverride(t *testing.T) {
	def := flaskFile(t, `from flask import Blueprint
api = Blueprint("api", __name__, url_prefix="/api")
`)
	prefixes := resolveBlueprintPrefixes([]fileInfo{def})
	if prefixes["api"] != "/api" {
		t.Fatalf("expected /api, got %v", prefixes)
	}
}
