//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_QuotedPathUnquoted 테스트
package hono

import "testing"

func TestExtractOneRoute_QuotedPathUnquoted(t *testing.T) {

	r := oneRoute(t, `app.delete('/items/:id', h);`, map[string]bool{"app": true})
	if r == nil || r.Path != "/items/:id" || r.Method != "DELETE" {
		t.Fatalf("got %+v", r)
	}
}
