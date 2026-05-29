//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 체인 메서드 호출(.post(...) 등)의 arguments에서 validate(x.y) 크로스파일 Joi 참조를 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// chainJoiRefs — .route().post(auth(...), validate(x.y), handler) 형태의 체인 메서드
// 호출에서 validate(importName.member) Joi 참조 목록을 추출한다.
func chainJoiRefs(call *sitter.Node, src []byte) []joiValidatorRef {
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	return extractJoiRefsFromArgs(collectArgNodes(args), src)
}
