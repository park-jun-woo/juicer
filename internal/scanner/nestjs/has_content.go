//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what Request 구조체에 내용이 있는지 확인한다
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// hasContent returns true if the request has any non-empty field.
func hasContent(r *scanner.Request) bool {
	return len(r.PathParams) > 0 || len(r.Query) > 0 || r.Body != nil || len(r.Files) > 0
}
