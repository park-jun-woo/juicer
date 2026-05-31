//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what TestExtractOneResponse_RenderRedirect: render→200/html, redirect→302/empty (체인 상태코드 포함) (Phase140)
package express

import "testing"

func TestExtractOneResponse_RenderRedirect(t *testing.T) {
	tests := []struct {
		src        string
		wantStatus string
		wantKind   string
	}{
		{`res.render('view');`, "200", "html"},
		{`res.status(201).render('view');`, "201", "html"},
		{`res.redirect('/home');`, "302", "empty"},
		{`res.status(301).redirect('/home');`, "301", "empty"},
	}
	for _, tt := range tests {
		fi := mustParse(t, []byte(tt.src))
		r := extractOneResponse(firstCallExpr(t, fi), fi.Src)
		if r == nil {
			t.Errorf("%s: expected response, got nil", tt.src)
			continue
		}
		if r.Status != tt.wantStatus || r.Kind != tt.wantKind {
			t.Errorf("%s: got %s/%s, want %s/%s", tt.src, r.Status, r.Kind, tt.wantStatus, tt.wantKind)
		}
	}
}
