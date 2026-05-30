//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestApplyHandlerSignature_Round5 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyHandlerSignature_Round5(t *testing.T) {
	fi := aFi(t, `async fn handler(body: web::Json<CreateReq>) -> impl Responder { "" }`)
	fn := aFirst(t, fi.root, "function_item")
	ep := &scanner.Endpoint{}
	applyHandlerSignature(ep, fn, fi.src, structIndex{}, map[string][]scanner.Field{})
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatalf("expected body from handler signature: %+v", ep.Request)
	}
}
