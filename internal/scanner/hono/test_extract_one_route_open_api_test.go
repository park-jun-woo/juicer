//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_OpenAPI 테스트
package hono

import "testing"

func TestExtractOneRoute_OpenAPI(t *testing.T) {
	r := oneRoute(t, `app.openapi(route, handler);`, map[string]bool{"app": true})

	_ = r
}
