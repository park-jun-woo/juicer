//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what Request 구조체에 내용이 있는지 확인한다
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

func hasContent(r *scanner.Request) bool {
	return len(r.PathParams) > 0 || len(r.Query) > 0 || len(r.Headers) > 0 || r.Body != nil || len(r.Files) > 0 || len(r.FormFields) > 0
}
