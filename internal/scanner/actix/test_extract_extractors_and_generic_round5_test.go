//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractExtractors_And_Generic_Round5 테스트
package actix

import "testing"

func TestExtractExtractors_And_Generic_Round5(t *testing.T) {
	fi := aFi(t, `async fn handler(body: web::Json<CreateReq>, id: web::Path<i64>) -> impl Responder { "" }`)
	fn := aFirst(t, fi.root, "function_item")
	exts := extractExtractors(fn, fi.src)
	if len(exts) == 0 {
		t.Fatalf("expected extractors, got %d", len(exts))
	}

	gt := aFirst(t, fi.root, "generic_type")
	_ = buildGenericExtractor(gt, fi.src)
}
