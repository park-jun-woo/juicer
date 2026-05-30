//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestCaptureScope_And_ApplyScope_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCaptureScope_And_ApplyScope_Round5(t *testing.T) {
	fi := aFi(t, `fn config(cfg: &mut web::ServiceConfig) { cfg.service(web::scope("/api").service(list_users)); }`)
	var scopes []scopeInfo
	walkNodes(fi.root, func(n *sitter.Node) {
		captureScope(n, fi.src, &scopes)
	})

	endpoints := []scanner.Endpoint{{Method: "GET", Path: "/users", Handler: "list_users"}}
	applyScopePrefixes([]*fileInfo{fi}, endpoints)
	applyScopeToEndpoints(scopeInfo{prefix: "/api", handlers: []string{"list_users"}}, endpoints)
	if endpoints[0].Path != "/api/users" {
		t.Fatalf("scope not applied: %q", endpoints[0].Path)
	}
}
