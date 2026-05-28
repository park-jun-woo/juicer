//ff:type feature=scan type=model topic=hono
//ff:what 추출된 라우트 정보 구조체
package hono

import "github.com/park-jun-woo/codistill/internal/scanner/zod"

type routeInfo struct {
	Method        string
	Path          string
	Handler       string
	OwnerVar      string
	Middleware    []string
	Line          int
	ZodValidators []zod.ValidatorInfo
}
