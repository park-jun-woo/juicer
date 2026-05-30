//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplySchemaToRequest_Empty 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestApplySchemaToRequest_Empty(t *testing.T) {

	si := &schemaInfo{}
	req := &scanner.Request{}
	hasReq := false
	applySchemaToRequest(si, []byte(""), req, &hasReq, map[string]*sitter.Node{})
	if hasReq || req.Body != nil || len(req.PathParams) != 0 || len(req.Query) != 0 {
		t.Fatalf("expected no changes, got %v hasReq=%v", req, hasReq)
	}
}
