//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Flask URL 규칙을 OpenAPI 경로 형식으로 변환한다
package flask

import "testing"

func TestFlaskPathToOpenAPI(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"/users/<int:id>", "/users/{id}"},
		{"/users/<user_id>", "/users/{user_id}"},
		{"/items/<float:price>", "/items/{price}"},
		{"/files/<path:subpath>", "/files/{subpath}"},
		{"/items/<uuid:item_id>", "/items/{item_id}"},
		{"/users/<int:uid>/posts/<int:pid>", "/users/{uid}/posts/{pid}"},
		{"/health", "/health"},
	}
	for _, tt := range tests {
		got := flaskPathToOpenAPI(tt.input)
		if got != tt.want {
			t.Errorf("flaskPathToOpenAPI(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
