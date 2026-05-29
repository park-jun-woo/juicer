//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 체인 메서드 호출(.post(...) 등)의 arguments에서 validateRequest({body: z...}) zod 검증 스키마를 추출한다
package express

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner/zod"
)

// chainZodValidators — .route().post(validateRequest({body: z.object(...)}), handler) 형태의
// 체인 메서드 호출에서 zod 검증 스키마 목록을 추출한다.
func chainZodValidators(call *sitter.Node, src []byte) []zod.ValidatorInfo {
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	return extractZodValidatorsFromArgs(collectArgNodes(args), src, 0)
}
