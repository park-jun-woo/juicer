//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what Pass 3: web::scope의 prefix를 .service()로 등록된 핸들러 엔드포인트에 적용한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyScopePrefixes(files []*fileInfo, endpoints []scanner.Endpoint) {
	for _, fi := range files {
		for _, scope := range extractScopes(fi) {
			applyScopeToEndpoints(scope, endpoints)
		}
	}
}
